/*
 *    Copyright (C) 2017-2019 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

// Package templatehelper provides a func-map of common template functions
package templatehelper

import (
	"encoding/json"
	htmlTemplate "html/template"
	"regexp"
	"strings"
	"text/template"
	"time"
)

// FuncMap contains a few convenient template helpers
var (
	FuncMap = template.FuncMap{
		"JSON": func(values ...interface{}) htmlTemplate.JS {
			json, _ := json.Marshal(values)
			return htmlTemplate.JS(json)
		},
		"Left": func(values ...interface{}) string {
			s := values[0].(string)
			n := values[1].(int)
			if n > len(s) {
				n = len(s)
			}

			return s[:n]
		},
		"Matches": func(values ...interface{}) bool {
			ok, _ := regexp.MatchString(values[1].(string), values[0].(string))
			return ok
		},
		"Mid": func(values ...interface{}) string {
			s := values[0].(string)
			l := values[1].(int)
			if l > len(s) {
				l = len(s)
			}

			if len(values) > 2 {
				r := values[2].(int)
				if r > len(s) {
					r = len(s)
				}
				return s[l:r]
			}
			return s[l:]
		},
		"Right": func(values ...interface{}) string {
			s := values[0].(string)
			n := len(s) - values[1].(int)
			if n < 0 {
				n = 0
			}

			return s[n:]
		},
		"Last": func(values ...interface{}) string {
			return values[0].([]string)[len(values[0].([]string))-1]
		},
		// strings functions
		"Compare":      strings.Compare, // 1.5+ only
		"Contains":     strings.Contains,
		"ContainsAny":  strings.ContainsAny,
		"Count":        strings.Count,
		"EqualFold":    strings.EqualFold,
		"HasPrefix":    strings.HasPrefix,
		"HasSuffix":    strings.HasSuffix,
		"Index":        strings.Index,
		"IndexAny":     strings.IndexAny,
		"Join":         strings.Join,
		"LastIndex":    strings.LastIndex,
		"LastIndexAny": strings.LastIndexAny,
		"Repeat":       strings.Repeat,
		"Replace":      strings.Replace,
		"Split":        strings.Split,
		"SplitAfter":   strings.SplitAfter,
		"SplitAfterN":  strings.SplitAfterN,
		"SplitN":       strings.SplitN,
		"TimeNow":      time.Now,
		"Title":        strings.Title,
		"ToLower":      strings.ToLower,
		"ToTitle":      strings.ToTitle,
		"ToUpper":      strings.ToUpper,
		"Trim":         strings.Trim,
		"TrimLeft":     strings.TrimLeft,
		"TrimPrefix":   strings.TrimPrefix,
		"TrimRight":    strings.TrimRight,
		"TrimSpace":    strings.TrimSpace,
		"TrimSuffix":   strings.TrimSuffix,
	}
)
