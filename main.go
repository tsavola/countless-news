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
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/kljensen/snowball/english"
	"github.com/mmcdole/gofeed"
)

const (
	indexName   = "www/index.html"
	indexNameGz = "www/index.html.gz"
	stageName   = "tmp/index.html"
	stageNameGz = "tmp/index.html.gz"
)

type Nation struct {
	Name      string
	NameToken string // Initialized by InitNations()
	Flag      string
	Sources   []Source
}

type Source interface {
	URL() string
	Headline(*gofeed.Item) string
	Match(*gofeed.Item) bool
}

// S is the simplest Source implementation.
type S struct {
	url string
}

func (x S) URL() string                       { return x.url }
func (x S) Headline(item *gofeed.Item) string { return item.Title }
func (x S) Match(*gofeed.Item) bool           { return true }

// SSubstring matches only titles which contain a substring.  It also strips
// CDATA wrapper from the title.
type SSubstring struct {
	S
	substring string
}

var (
	cdataRegexp = regexp.MustCompile(`^<!\[CDATA\[(.*)\]\]>.*`)
)

func (x SSubstring) Headline(item *gofeed.Item) string {
	matches := cdataRegexp.FindStringSubmatch(item.Title)
	if len(matches) >= 2 {
		return matches[1]
	} else {
		return item.Title
	}
}

func (x SSubstring) Match(item *gofeed.Item) bool {
	return strings.Contains(strings.ToLower(item.Title), x.substring)
}

type IndexItem struct {
	Nation    *Nation
	Headline  string
	URL       *url.URL
	Timestamp string
	Score     float64
}

type Index []*IndexItem

func (a Index) Len() int           { return len(a) }
func (a Index) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Index) Less(i, j int) bool { return a[i].Score > a[j].Score }

func InitNations() {
	for i := range NationalSources {
		NationalSources[i].NameToken = english.Stem(NationalSources[i].Name, false)
	}
}

func Update() {
	var (
		index    Index
		maxScore float64
	)

	for i := range NationalSources {
		nation := &NationalSources[i]

		log.Printf("Processing %s", nation.Name)

		source, item, score := ChooseFrom(nation)
		if item == nil {
			log.Printf("No news for %s", nation.Name)
			continue
		}

		ts := item.Published
		if ts == "" {
			ts = item.Updated
		}

		u, _ := url.Parse(item.Link) // Already validated by Algorithm()

		index = append(index, &IndexItem{
			Nation:    nation,
			Headline:  strings.TrimSpace(source.Headline(item)),
			URL:       u,
			Timestamp: ts,
			Score:     score,
		})

		if score > maxScore {
			maxScore = score
		}
	}

	sort.Sort(index)

	html := Render(index, maxScore)

	if !func() (ok bool) {
		f, err := os.Create(stageName)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()

		if _, err := f.Write(html); err != nil {
			log.Print(err)
			return
		}

		ok = true
		return
	}() {
		return
	}

	cmd := exec.Command("zopfli", stageName)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Print(err)
		return
	}

	if err := os.Rename(stageNameGz, indexNameGz); err == nil {
		if err := os.Rename(stageName, indexName); err != nil {
			log.Print(err)
		}
	} else {
		log.Print(err)
		return
	}

	log.Print("Index updated")
}

func main() {
	if err := InitStopWords(); err != nil {
		log.Print(err)
		return
	}

	InitNations()

	for {
		Update()

		now := time.Now()

		t := now.Add(time.Hour)
		t = t.Add(-time.Minute * time.Duration(t.Minute()))       // Truncate
		t = t.Add(time.Duration(rand.Int63n(int64(time.Minute)))) // Jitter

		d := t.Sub(now)
		log.Printf("Sleeping for %v", d)
		time.Sleep(d)
	}
}
