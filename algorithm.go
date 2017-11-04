// Copyright (c) 2017 Timo Savola. All Rights Reserved.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package main

import (
	"html"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/kljensen/snowball/english"
	"github.com/mmcdole/gofeed"
)

const (
	// Limiting the number of news items to consider gives more balanced scores
	// as some sources provide fewer news.
	maxItemCount = 25
)

var topCategories = map[string]struct{}{
	"front page":   struct{}{},
	"top articles": struct{}{},
}

var topInitialWords = map[string]struct{}{
	"breaking:": struct{}{},
}

var skipItemWords = map[string]struct{}{
	"roundup": struct{}{},
}

var (
	nonWordChar = regexp.MustCompile("[&”–-]")
	nonAlphabet = regexp.MustCompile("[^a-z]")
)

func isExtraneousWord(word string) (skipWord bool) {
	// "[WATCH]" and such tags.
	if word[0] == '[' && word[len(word)-1] == ']' {
		return true
	}

	switch word {
	case "UPDATE", "UPDATED", "VIDEO", "WATCH":
		return true
	}

	return
}

func tokenizeWord(word string, nation *Nation) (result string) {
	word = html.UnescapeString(word)

	if isExtraneousWord(word) {
		return
	}

	word = nonWordChar.ReplaceAllString(word, "")
	if word == "" {
		return
	}

	simpleWord := nonAlphabet.ReplaceAllString(strings.ToLower(word), "")
	if _, ignore := stopWords[simpleWord]; ignore {
		return
	}

	token := english.Stem(word, false)
	if token == nation.NameToken {
		return
	}

	result = token
	return
}

// ChooseFrom a nation's news sources.  The algorithm looks at each source
// separately.
func ChooseFrom(nation *Nation) (topSource Source, topItem *gofeed.Item, topScore float64) {
	for _, source := range nation.Sources {
		item, score := chooseFrom(source, nation)
		if topItem == nil || score > topScore {
			topSource = source
			topItem = item
			topScore = score
		}
	}
	return
}

// chooseFrom a news source.
func chooseFrom(source Source, nation *Nation) (topItem *gofeed.Item, topScore float64) {
	feed, err := gofeed.NewParser().ParseURL(source.URL())
	if err != nil {
		log.Printf("%s: %v", source.URL(), err)
		return
	}

	// Limit considered items to latest publications with valid-looking URLs.
	// They also have to match the Source implementation's criteria.

	var (
		items     = feed.Items
		itemOrder = make([]int64, len(items))
		itemCount = 0
	)

	for i, item := range items {
		itemURL, err := url.Parse(item.Link)
		if err != nil {
			log.Print(err)
			continue
		}

		if strings.HasPrefix(item.Link, "/") {
			// Relative is ok
		} else {
			switch itemURL.Scheme {
			case "https", "http":
				// ok
			default:
				log.Printf("Bad URL scheme: %s", item.Link)
				continue
			}
		}

		if !source.Match(item) {
			continue
		}

		if item.PublishedParsed != nil {
			itemOrder[i] = item.PublishedParsed.UnixNano()
		} else if item.UpdatedParsed != nil {
			itemOrder[i] = item.UpdatedParsed.UnixNano()
		} else {
			// Assume reverse publication order if no timestamp found.
			itemOrder[i] = int64(len(items) - i)
		}

		itemCount++
	}

	sort.Sort(ReverseItemOrdering{items, itemOrder})

	if itemCount > maxItemCount {
		itemCount = maxItemCount
	}
	if len(items) > itemCount {
		items = items[:itemCount]
	}

	// Parse unique words for each item, by first converting them to word stems
	// (tokens).  Flag a token as important if it appears as such in at least
	// one context.  If a news source has a convention for indicating top news,
	// collect them too.

	var (
		topItems        = make(map[int]struct{})
		itemsTokens     = make([]map[string]struct{}, len(items))
		tokenCounts     = make(map[string]int)
		importantTokens = make(map[string]struct{})
	)

	for i, item := range items {
		tokensOfItem := make(map[string]struct{})

		addToken := func(word, token string) {
			if char, _ := utf8.DecodeRuneInString(word); char != utf8.RuneError {
				if (unicode.IsLetter(char) && unicode.IsUpper(char)) || unicode.IsNumber(char) {
					importantTokens[token] = struct{}{}
				}
			}

			tokensOfItem[token] = struct{}{}
			tokenCounts[token]++
		}

		headline := source.Headline(item)

		// Remove "Updated | " prefix or " | John F. Editor" suffix.
		if prefix := strings.Index(headline, " | "); prefix >= 0 {
			suffix := len(headline) - (prefix + 3)
			if prefix >= suffix {
				headline = headline[:prefix]
			} else {
				headline = headline[len(headline)-suffix:]
			}
		}

		for j, word := range strings.Fields(headline) {
			lowerWord := strings.ToLower(word)

			if j == 0 {
				if _, top := topInitialWords[lowerWord]; top {
					topItems[i] = struct{}{}
					continue
				}
			}

			if _, skip := skipItemWords[lowerWord]; skip {
				continue
			}

			token := tokenizeWord(word, nation)
			if token == "" {
				continue
			}

			addToken(word, token)
		}

		for _, keyword := range item.Categories {
			if _, top := topCategories[strings.ToLower(keyword)]; top {
				topItems[i] = struct{}{}
			}
		}

		itemsTokens[i] = tokensOfItem
	}

	// Word count with unique words removed.  It will be used to calculate word
	// weights later.

	var (
		significantTokenCount int
	)

	for _, count := range tokenCounts {
		if count > 1 {
			significantTokenCount += count
		}
	}

	// Rare, important words carry the most weight.  Unique words carry no
	// weight, as we will be looking for similarities later.

	var (
		tokenWeights = make(map[string]float64)
	)

	for token, count := range tokenCounts {
		if count > 1 {
			weight := float64(significantTokenCount-count) / float64(significantTokenCount)

			if _, important := importantTokens[token]; !important {
				weight /= 2
			}

			tokenWeights[token] = weight
		}
	}

	// Score each item based on how many of the item's words appear in other
	// items.

	for i, thisTokens := range itemsTokens {
		if len(thisTokens) > 0 {
			var score float64

			for j, thatTokens := range itemsTokens {
				if i != j {
					var sum float64

					for token, _ := range thisTokens {
						if _, match := thatTokens[token]; match {
							sum += tokenWeights[token]
						}
					}

					score += sum / float64(len(thisTokens))
				}
			}

			if len(topItems) > 0 {
				if _, top := topItems[i]; !top {
					continue
				}
			}

			if score > topScore {
				topItem = items[i]
				topScore = score
			}
		}
	}

	if topItem == nil && len(items) > 0 {
		i := 0
		if len(topItems) > 0 {
			for i = range topItems {
				break
			}
		}
		topItem = items[i]
	}

	return
}

type ReverseItemOrdering struct {
	Items []*gofeed.Item
	Order []int64
}

func (x ReverseItemOrdering) Len() int {
	return len(x.Items)
}

func (x ReverseItemOrdering) Swap(i, j int) {
	x.Items[i], x.Items[j] = x.Items[j], x.Items[i]
	x.Order[i], x.Order[j] = x.Order[j], x.Order[i]
}

func (x ReverseItemOrdering) Less(i, j int) bool {
	return x.Order[i] > x.Order[j]
}
