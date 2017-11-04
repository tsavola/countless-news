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

type Result struct {
	Nation *Nation
	Source Source
	Item   *gofeed.Item
	Score  float64
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
	log.Print("Updating index")

	results := make(chan Result)

	for i := range NationalSources {
		go func(nation *Nation) {
			r := Result{Nation: nation}
			defer func() {
				results <- r
			}()
			r.Source, r.Item, r.Score = ChooseFrom(nation)
		}(&NationalSources[i])
	}

	var (
		index    Index
		maxScore float64
	)

	for range NationalSources {
		r := <-results
		if r.Item == nil {
			log.Printf("No news for %s", r.Nation.Name)
			continue
		}

		ts := r.Item.Published
		if ts == "" {
			ts = r.Item.Updated
		}

		link := r.Item.Link
		if strings.HasPrefix(link, "/") { // Relative
			sourceURL, _ := url.Parse(r.Source.URL())
			link = sourceURL.Scheme + "://" + sourceURL.Host + link
		}

		itemURL, _ := url.Parse(link)

		index = append(index, &IndexItem{
			Nation:    r.Nation,
			Headline:  strings.TrimSpace(r.Source.Headline(r.Item)),
			URL:       itemURL,
			Timestamp: ts,
			Score:     r.Score,
		})

		if r.Score > maxScore {
			maxScore = r.Score
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
