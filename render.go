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
	scoreWidth = 20 // pt
)

func Render(index Index, maxScore float64) []byte {
	buf := new(bytes.Buffer)

	fmt.Fprint(buf, `<!DOCTYPE html>`)
	fmt.Fprint(buf, `<html lang="en">`)
	fmt.Fprint(buf, `<head>`)
	fmt.Fprint(buf, `<meta charset="utf-8">`)
	fmt.Fprint(buf, `<title>Countless News</title>`)
	fmt.Fprint(buf, `<meta name="description" content="Top news by country">`)
	fmt.Fprint(buf, `<style>`)
	fmt.Fprintln(buf, `a { text-decoration: none }`)
	fmt.Fprintln(buf, `table { margin-left: auto; margin-right: auto }`)
	fmt.Fprintln(buf, `tr:nth-child(odd) { background-color: white }`)
	fmt.Fprintln(buf, `tr:nth-child(even) { background-color: #f9f9f9 }`)
	fmt.Fprintln(buf, `td { white-space: nowrap }`)
	fmt.Fprintln(buf, `.header { display: none }`)
	fmt.Fprintln(buf, `.nation { padding-left: 4pt; padding-right: 4pt }`)
	fmt.Fprintln(buf, `.score { vertical-align: bottom; padding-right: 2pt }`)
	fmt.Fprintf(buf, ".score div { width: %dpt; height: 1em }\n", scoreWidth)
	fmt.Fprintln(buf, `.score div div { height: 100%; background-color: lightgrey; margin-left: auto }`)
	fmt.Fprintln(buf, `.domain { opacity: 0.5; text-align: right; color: #444; padding-left: 4pt }`)
	fmt.Fprintln(buf, `.https { color: green }`)
	fmt.Fprintln(buf, `#footer { background-color: inherit }`)
	fmt.Fprintln(buf, `#footer td { text-align: right; padding-top: 1em; opacity: 0.5 }`)
	fmt.Fprint(buf, `</style>`)
	fmt.Fprint(buf, `</head>`)
	fmt.Fprint(buf, `<body>`)
	fmt.Fprint(buf, `<table>`)
	fmt.Fprint(buf, `<tr class="header">`)
	fmt.Fprint(buf, `<th>Flag</th>`)
	fmt.Fprint(buf, `<th>Nation</th>`)
	fmt.Fprint(buf, `<th>Score</th>`)
	fmt.Fprint(buf, `<th>Headline</th>`)
	fmt.Fprint(buf, `<th>Source</th>`)
	fmt.Fprint(buf, `</tr>`)

	for _, entry := range index {
		headline := entry.Headline
		if len(headline) > 100 {
			headline = headline[:99] + "…"
		}

		domain := entry.URL.Host

		suffix, _ := publicsuffix.PublicSuffix(domain)
		num := strings.Count(suffix, ".") + 2

		parts := strings.Split(domain, ".")
		domain = strings.Join(parts[len(parts)-num:len(parts)], ".")

		// TODO: something better
		var proxyHTML string
		if domain == "google.com" {
			entry.URL.Scheme = "https"
			proxyHTML = fmt.Sprintf(`<span class="%s">%s</span>, `, entry.URL.Scheme, domain)
			domain = "baltictimes.com"
		}

		fmt.Fprint(buf, `<tr>`)
		fmt.Fprintf(buf, `<td>%s</td>`, entry.Nation.Flag)
		fmt.Fprintf(buf, `<td class="nation">%s</td>`, html.EscapeString(entry.Nation.Name))
		fmt.Fprintf(buf, `<td class="score"><div title="%f"><div style="width: %.2fpt"></div></div></td>`, entry.Score, entry.Score*scoreWidth/maxScore)
		fmt.Fprintf(buf, `<td class="headline"><a href="%s" title="%s">%s</a></td>`, entry.URL, html.EscapeString(entry.Timestamp), html.EscapeString(headline))
		fmt.Fprintf(buf, `<td class="domain">%s<span class="%s">%s</span></td>`, proxyHTML, entry.URL.Scheme, domain)
		fmt.Fprint(buf, `</tr>`)
	}

	fmt.Fprint(buf, `<tr id="footer">`)
	fmt.Fprint(buf, `<td colspan="5">`)
	fmt.Fprint(buf, `<a href="https://github.com/tsavola/countless-news/blob/master/algorithm.go">Algorithm</a>`)
	fmt.Fprint(buf, ` · `)
	fmt.Fprint(buf, `<a href="https://github.com/tsavola/countless-news/blob/master/sources.go">Sources</a>`)
	fmt.Fprint(buf, `</td>`)
	fmt.Fprint(buf, `</tr>`)
	fmt.Fprint(buf, `</table>`)
	fmt.Fprint(buf, `</body>`)
	fmt.Fprint(buf, `</html>`)

	return buf.Bytes()
}
