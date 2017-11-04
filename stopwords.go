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
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	packageName   = "github.com/tsavola/countless-news"
	stopWordsFile = "postgres/src/backend/snowball/stopwords/english.stop"
)

var stopWords map[string]struct{}

func InitStopWords() (err error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = path.Join(os.Getenv("HOME"), "go")
	}

	filename := path.Join(gopath, "src", packageName, stopWordsFile)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	stopWords = make(map[string]struct{})

	for _, word := range strings.Fields(string(data)) {
		word = strings.TrimSpace(word)
		if word != "" {
			stopWords[word] = struct{}{}
		}
	}

	stopWords["n"] = struct{}{} // Trimmed "‘n’"
	stopWords["news"] = struct{}{}
	return
}
