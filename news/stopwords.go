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
	_ "embed"
	"strings"
)

//go:embed stopwords.txt
var stopWordData string

var stopWords = getStopWords()

func getStopWords() map[string]struct{} {
	words := make(map[string]struct{})

	for _, word := range strings.Fields(stopWordData) {
		word = strings.TrimSpace(word)
		if word != "" {
			words[word] = struct{}{}
		}
	}

	words["n"] = struct{}{} // Trimmed "‘n’"
	words["news"] = struct{}{}
	words["weather"] = struct{}{}

	return words
}
