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

type scanLineIter struct {
	scanner   *scanner.Scanner
	buffer    runeBuffer
	lineWidth int
	lboPos    int
	indent    string
}

func newScanLineIter(text string, lineWidth int) scanLineIter {
	sc := new(scanner.Scanner)
	sc.Init(strings.NewReader(text))

	iter := scanLineIter{}
	iter.scanner = sc
	iter.buffer = newRuneBuffer(lineWidth)
	iter.lineWidth = lineWidth
	return iter
}

func (iter *scanLineIter) setIndent(indent int) {
	iter.indent = strings.Repeat(" ", indent)
}

func (iter *scanLineIter) resetText(text string) {
	iter.scanner.Init(strings.NewReader(text))
	iter.buffer.length = 0
}

func (iter *scanLineIter) Next() (string, IterStatus) {
	lineWidth := iter.lineWidth - len(iter.indent)

	for r := iter.scanner.Next(); r != scanner.EOF; r = iter.scanner.Next() {
		lboTyp := lineBreakOppotunity(r)

		var line string
		if lboTyp == lbo_break {
			line = string(iter.buffer.slice())
			iter.lboPos = 0
			iter.buffer.length = 0
			return iter.indent + line, ITER_HAS_MORE
		}

		if iter.buffer.length == 0 && lboTyp == lbo_space {
			continue
		}

		length := iter.buffer.length + runeWidth(r)
		lboPos := iter.lboPos

		if length > lineWidth {
			switch lboTyp {
			case lbo_before, lbo_both, lbo_space:
				lboPos = iter.buffer.length
			}
			if lboPos == 0 {
				lboPos = lineWidth
			}

			line := string(iter.buffer.runes[0:lboPos])
			iter.buffer.cr(lboPos)

			switch lboTyp {
			case lbo_space:
				iter.lboPos = 0
			case lbo_after, lbo_both:
				iter.buffer.add(r)
				iter.lboPos = iter.buffer.length
			default:
				iter.buffer.add(r)
				iter.lboPos = 0
			}

			return iter.indent + line, ITER_HAS_MORE
		}

		iter.buffer.add(r)
		switch lboTyp {
		case lbo_before:
			iter.lboPos = iter.buffer.length - 1
		case lbo_after, lbo_both, lbo_space:
			iter.lboPos = iter.buffer.length
		}
	}

	return iter.indent + string(iter.buffer.slice()), ITER_NO_MORE
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
