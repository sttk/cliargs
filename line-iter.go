// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"strings"
	"text/scanner"
	"unicode"
)

// IterStatus is a type which indicates an iterator status.
// If this value is true, a iterator has more items.
// If this value is false, a iterator has no more items.
type IterStatus bool

const (
	ITER_HAS_MORE IterStatus = true  // A iterator has more items.
	ITER_NO_MORE  IterStatus = false // A iterator has no more items.
)

type lboType int

const (
	lbo_before lboType = iota
	lbo_after
	lbo_both
	lbo_never
	lbo_break
	lbo_space
)

type lineIter struct {
	scanner *scanner.Scanner
	buffer  runeBuffer
	width   [2]int /* 0: width before lbo, 1: width after lbo */
	lboPos  int
	limit   int
	indent  string
}

func newLineIter(text string, lineWidth int) lineIter {
	sc := new(scanner.Scanner)
	sc.Init(strings.NewReader(text))

	iter := lineIter{}
	iter.scanner = sc
	iter.buffer = newRuneBuffer(lineWidth)
	iter.limit = lineWidth
	return iter
}

func (iter *lineIter) setIndent(indent int) {
	iter.indent = strings.Repeat(" ", indent)
}

func (iter *lineIter) resetText(text string) {
	iter.scanner.Init(strings.NewReader(text))
	iter.buffer.length = 0
	iter.width[0] = 0
	iter.width[1] = 0
	iter.lboPos = 0
}

func (iter *lineIter) Next() (string, IterStatus) {
	limit := iter.limit - len(iter.indent)

	var line string

	for r := iter.scanner.Next(); r != scanner.EOF; r = iter.scanner.Next() {
		lboTyp := lineBreakOppotunity(r)

		if lboTyp == lbo_break {
			line = string(iter.buffer.slice())
			iter.buffer.length = 0
			iter.width[0] = 0
			iter.width[1] = 0
			iter.lboPos = 0
			if len(line) > 0 {
				line = iter.indent + line
			}
			return line, ITER_HAS_MORE
		}

		if iter.buffer.length == 0 && lboTyp == lbo_space {
			continue
		}

		runeW := runeWidth(r)
		lboPos := iter.lboPos

		if (iter.width[0] + iter.width[1] + runeW) > limit {
			switch lboTyp {
			case lbo_before, lbo_both, lbo_space:
				lboPos = iter.buffer.length
			}
			if lboPos == 0 {
				//iter.width[0] += iter.width[1]
				iter.width[1] = 0
				lboPos = iter.buffer.length
			}

			line := string(iter.buffer.runes[0:lboPos])
			iter.buffer.cr(lboPos)

			switch lboTyp {
			case lbo_space:
				iter.width[0] = 0
				iter.width[1] = 0
				iter.lboPos = 0
			//case lbo_before:
			//	iter.buffer.add(r)
			//	iter.width[0] = runeW
			//	iter.width[1] = 0
			//	iter.lboPos = 0
			case lbo_after, lbo_both:
				iter.buffer.add(r)
				iter.width[0] = iter.width[1] + runeW
				iter.width[1] = 0
				iter.lboPos = iter.buffer.length
			default:
				iter.buffer.add(r)
				iter.width[0] = iter.width[1] + runeW
				iter.width[1] = 0
				iter.lboPos = 0
			}

			if len(line) > 0 {
				line = iter.indent + line
			}
			return line, ITER_HAS_MORE
		}

		if runeW > 0 {
			iter.buffer.add(r)
		}
		switch lboTyp {
		//case lbo_before:
		//	iter.lboPos = iter.buffer.length - 1
		//	iter.width[0] += iter.width[1]
		//	iter.width[1] = runeW
		case lbo_after, lbo_both, lbo_space:
			iter.lboPos = iter.buffer.length
			iter.width[0] += iter.width[1] + runeW
			iter.width[1] = 0
		default:
			iter.width[1] += runeW
		}
	}

	line = string(iter.buffer.slice())
	iter.buffer.length = 0

	if len(line) > 0 {
		line = iter.indent + line
	}
	return line, ITER_NO_MORE
}

func lineBreakOppotunity(r rune) lboType {
	if r == 0x0a || r == 0x0d {
		return lbo_break
	}
	if unicode.IsSpace(r) {
		return lbo_space
	}
	if unicode.IsPunct(r) {
		return lbo_after
	}
	return lbo_never
}

func runeWidth(r rune) int {
	if !unicode.IsPrint(r) {
		return 0
	}
	return 1
}
