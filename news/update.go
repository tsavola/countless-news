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

package news

import (
	"bytes"
	"compress/gzip"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/kljensen/snowball/english"
	"github.com/mmcdole/gofeed"
)

var zopfli = "zopfli"

const (
	indexName   = "index.html"
	indexNameGz = "index.html.gz"
	stageName   = ".tmp_index.html"
	stageNameGz = ".tmp_index.html.gz"
)

type Notification struct {
	GzipIndexFilename string
}

type Nation struct {
	Name      string
	NameToken string // Initialized by init()
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

func init() {
	for i := range NationalSources {
		NationalSources[i].NameToken = english.Stem(NationalSources[i].Name, false)
	}
}

func update(dir string, infoLog Logger) error {
	infoLog.Printf("updating index")

	results := make(chan Result)

	for i := range NationalSources {
		go func(nation *Nation) {
			r := Result{Nation: nation}
			defer func() {
				results <- r
			}()
			r.Source, r.Item, r.Score = ChooseFrom(nation, infoLog)
		}(&NationalSources[i])
	}

	var (
		index    Index
		maxScore float64
	)

	for range NationalSources {
		r := <-results
		if r.Item == nil {
			infoLog.Printf("no news from %s", r.Nation.Name)
			continue
		}

		ts := r.Item.Published
		if ts == "" {
			ts = r.Item.Updated
		}

		sourceURL, _ := url.Parse(r.Source.URL())

		link := r.Item.Link
		if strings.HasPrefix(link, "/") { // Relative
			link = sourceURL.Scheme + "://" + sourceURL.Host + link
		}

		itemURL, _ := url.Parse(link)

		if sourceURL.Scheme == "https" && itemURL.Host == sourceURL.Host {
			itemURL.Scheme = "https"
		}

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

	if err := func() error {
		f, err := os.Create(path.Join(dir, stageName))
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(html)
		return err
	}(); err != nil {
		return err
	}

	err := compressWithZopfli(path.Join(dir, stageName))
	if err != nil {
		log.Printf("%v", err)
		err = compressWithGzip(path.Join(dir, stageNameGz), html)
	}
	if err != nil {
		return err
	}

	if err := os.Rename(path.Join(dir, stageNameGz), path.Join(dir, indexNameGz)); err == nil {
		if err := os.Rename(path.Join(dir, stageName), path.Join(dir, indexName)); err != nil {
			return err
		}
	} else {
		return err
	}

	infoLog.Printf("index updated")
	return nil
}

func compressWithZopfli(filename string) error {
	cmd := exec.Command(zopfli, filename)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func compressWithGzip(gzipFilename string, data []byte) error {
	buf := bytes.NewBuffer(make([]byte, 0, len(data)))

	w, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return os.WriteFile(gzipFilename, buf.Bytes(), 0666)
}

func UpdateLoop(dir string, notify chan<- Notification, infoLog, errorLog Logger) {
	if notify != nil {
		filename := path.Join(dir, indexNameGz)
		if syscall.Access(filename, syscall.F_OK) == nil {
			notify <- Notification{filename}
		}
	}

	for {
		if err := update(dir, infoLog); err == nil {
			if notify != nil {
				notify <- Notification{path.Join(dir, indexNameGz)}
			}
		} else {
			errorLog.Printf("%v", err)
		}

		now := time.Now()

		t := now.Add(time.Hour)
		t = t.Add(-time.Minute * time.Duration(t.Minute()))       // Truncate
		t = t.Add(time.Duration(rand.Int63n(int64(time.Minute)))) // Jitter

		d := t.Sub(now)
		infoLog.Printf("sleeping for %v", d)
		time.Sleep(d)
	}
}
