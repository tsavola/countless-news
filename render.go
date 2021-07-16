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
	"bytes"
	"fmt"
	"html"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const (
	scoreColor    = "#f80"
	maxScoreWidth = 50 // pt
)

func Render(index Index, maxScore float64) []byte {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, `<!DOCTYPE html>`)
	fmt.Fprintf(buf, `<html lang="en">`)
	fmt.Fprintf(buf, `<head>`)
	fmt.Fprintf(buf, `<meta charset="utf-8">`)
	fmt.Fprintf(buf, `<title>Countless News</title>`)
	fmt.Fprintf(buf, `<meta name="description" content="Top news by country">`)
	fmt.Fprintf(buf, `<meta name="viewport" content="width=device-width, initial-scale=1.0">`)
	fmt.Fprintf(buf, `<link rel="stylesheet" type="text/css" href="style.css">`)
	fmt.Fprintf(buf, `</head>`)
	fmt.Fprintf(buf, `<body>`)
	fmt.Fprintf(buf, `<h1>Countless News</h1>`)
	fmt.Fprintf(buf, `<div id="root">`)
	fmt.Fprintf(buf, `<div id="news">`)

	for _, entry := range index {
		headline := html.UnescapeString(entry.Headline)

		domain := entry.URL.Host

		suffix, _ := publicsuffix.PublicSuffix(domain)
		num := strings.Count(suffix, ".") + 2

		parts := strings.Split(domain, ".")
		domain = strings.Join(parts[len(parts)-num:], ".")

		scoreWidth := entry.Score * maxScoreWidth / maxScore

		fmt.Fprintf(buf, `<div class="item">`)
		fmt.Fprintf(buf, `<div>`)
		fmt.Fprintf(buf, `<div class="nation" style="background: linear-gradient(-90deg, %s, transparent %.3fpt)">`, scoreColor, scoreWidth)
		fmt.Fprintf(buf, `<span class="name">%s</span>`, html.EscapeString(entry.Nation.Name))
		fmt.Fprintf(buf, `</div>`) // nation
		fmt.Fprintf(buf, `<div class="headline"><a href="%s" title="%s">%s</a></div>`, entry.URL, html.EscapeString(entry.Timestamp), html.EscapeString(headline))
		fmt.Fprintf(buf, `<div class="domain"><span class="%s">%s</span></div>`, entry.URL.Scheme, domain)
		fmt.Fprintf(buf, `</div>`) // anonymous
		fmt.Fprintf(buf, `</div>`) // item
	}

	fmt.Fprintf(buf, `</div>`) // news
	fmt.Fprintf(buf, `<div id="footer">`)
	fmt.Fprintf(buf, `<div>Local news for nonlocal people</div>`)
	fmt.Fprintf(buf, `<div>`)
	fmt.Fprintf(buf, `<div><a href="https://github.com/tsavola/countless-news/blob/master/algorithm.go">Algorithm</a></div>`)
	fmt.Fprintf(buf, `<div><a href="https://github.com/tsavola/countless-news/blob/master/sources.go">Sources</a></div>`)
	fmt.Fprintf(buf, `</div>`) // anonymous
	fmt.Fprintf(buf, `</div>`) // footer
	fmt.Fprintf(buf, `</div>`) // root
	fmt.Fprintf(buf, `</body>`)
	fmt.Fprintf(buf, `</html>`)

	return buf.Bytes()
}
