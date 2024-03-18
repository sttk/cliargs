// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"strings"

	"github.com/sttk/linebreak"
)

// Help is a struct type which holds help text blocks and help options block.
type Help struct {
	marginLeft, marginRight int
	blocks                  []block
}

type block struct {
	indent, marginLeft, marginRight int
	texts                           []string
}

// NewHelp is a function to create a Help instance.
// This function can optionally take left margin and right margin as variadic
// arguments.
func NewHelp(wrapOpts ...int) Help {
	var help Help
	if len(wrapOpts) > 0 {
		help.marginLeft = wrapOpts[0]
	}
	if len(wrapOpts) > 1 {
		help.marginRight = wrapOpts[1]
	}
	help.blocks = make([]block, 0, 2)
	return help
}

// Iter is a method which creates a HelpIter instance.
func (help Help) Iter() HelpIter {
	if len(help.blocks) == 0 {
		return HelpIter{}
	}

	lineWidth := linebreak.TermCols()

	return HelpIter{
		lineWidth: lineWidth,
		blocks:    help.blocks,
		blockIter: newBlockIter(help.blocks[0], lineWidth),
	}
}

// HelpIter is a struct type to iterate lines of help texts.
type HelpIter struct {
	lineWidth int
	blocks    []block
	blockIter blockIter
}

// Next is a method which returns a line of a help text and a status which
// indicates this HelpIter has more texts or not.
// If there are more lines, the returned IterStatus value is ITER_HAS_MORE,
// otherwise the value is ITER_NO_MORE.
func (iter *HelpIter) Next() (string, bool) {
	line, more := iter.blockIter.next()
	if !more {
		if len(iter.blocks) <= 1 {
			return line, false
		}
		iter.blocks = iter.blocks[1:]
		iter.blockIter = newBlockIter(iter.blocks[0], iter.lineWidth)
	}
	return line, true
}

type blockIter struct {
	texts    []string
	index    int
	indent   int
	margin   string
	lineIter linebreak.LineIter
}

func newBlockIter(b block, lineWidth int) blockIter {
	if len(b.texts) == 0 {
		return blockIter{}
	}
	printWidth := lineWidth - b.marginLeft - b.marginRight
	if printWidth <= b.indent {
		return blockIter{}
	}
	return blockIter{
		texts:    b.texts,
		indent:   b.indent,
		margin:   strings.Repeat(" ", b.marginLeft),
		lineIter: linebreak.New(b.texts[0], printWidth),
	}
}

func (iter *blockIter) next() (string, bool) {
	if len(iter.texts) == 0 {
		return "", false
	}

	line, more := iter.lineIter.Next()
	if len(line) > 0 {
		line = iter.margin + line
	}
	if !more {
		iter.index++
		if iter.index >= len(iter.texts) {
			return line, more
		}
		iter.lineIter.Init(iter.texts[iter.index])
		iter.lineIter.SetIndent("")
		return line, true
	}

	iter.lineIter.SetIndent(strings.Repeat(" ", iter.indent))
	return line, more
}

// AddText is a method which adds a text to this Help instance.
// And this method can optionally set indent, left margin, and right margin as
// variadic arguments, too.
func (help *Help) AddText(text string, wrapOpts ...int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	if len(wrapOpts) > 0 {
		b.indent = wrapOpts[0]
	}
	if len(wrapOpts) > 1 {
		b.marginLeft += wrapOpts[1]
	}
	if len(wrapOpts) > 2 {
		b.marginRight += wrapOpts[2]
	}
	b.texts = []string{text}
	help.blocks = append(help.blocks, b)
}

// AddTexts is a method which adds an array of texts to this Help instance.
// And this method can optionally set indent, left margin, and right margin as
// variadic arguments, too.
func (help *Help) AddTexts(texts []string, wrapOpts ...int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	if len(wrapOpts) > 0 {
		b.indent = wrapOpts[0]
	}
	if len(wrapOpts) > 1 {
		b.marginLeft += wrapOpts[1]
	}
	if len(wrapOpts) > 2 {
		b.marginRight += wrapOpts[2]
	}
	b.texts = texts
	help.blocks = append(help.blocks, b)
}

// AddOpts is a method which adds OptCfg(s) to this Help instance.
// And this method can optionally set indent, left margin, and right margin as
// variadic arguments, too.
func (help *Help) AddOpts(optCfgs []OptCfg, wrapOpts ...int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	if len(wrapOpts) > 0 {
		b.indent = wrapOpts[0]
	}
	if len(wrapOpts) > 1 {
		b.marginLeft += wrapOpts[1]
	}
	if len(wrapOpts) > 2 {
		b.marginRight += wrapOpts[2]
	}

	texts := make([]string, len(optCfgs))

	if b.indent > 0 {
		i := 0
		for _, cfg := range optCfgs {
			storeKey := cfg.StoreKey
			if len(storeKey) == 0 {
				if len(cfg.Names) > 0 && cfg.Names[0] == anyOption {
					storeKey = cfg.Names[0]
				}
			}
			if storeKey == anyOption {
				continue
			}

			texts[i] = makeOptTitle(cfg)
			width := linebreak.TextWidth(texts[i])
			if width+2 > b.indent {
				texts[i] += "\n" + strings.Repeat(" ", b.indent) + cfg.Desc
			} else {
				texts[i] += strings.Repeat(" ", b.indent-width) + cfg.Desc
			}
			i++
		}
		texts = texts[0:i]

	} else {
		widths := make([]int, len(texts))
		indent := 0

		i := 0
		for _, cfg := range optCfgs {
			storeKey := cfg.StoreKey
			if len(storeKey) == 0 {
				if len(cfg.Names) > 0 && cfg.Names[0] == anyOption {
					storeKey = cfg.Names[0]
				}
			}
			if storeKey == anyOption {
				continue
			}

			texts[i] = makeOptTitle(cfg)
			widths[i] = linebreak.TextWidth(texts[i])
			if indent < widths[i] {
				indent = widths[i]
			}
			i++
		}
		texts = texts[0:i]
		indent += 2

		b.indent = indent

		i = 0
		for _, cfg := range optCfgs {
			storeKey := cfg.StoreKey
			if len(storeKey) == 0 {
				if len(cfg.Names) > 0 && cfg.Names[0] == anyOption {
					storeKey = cfg.Names[0]
				}
			}
			if storeKey == anyOption {
				continue
			}

			texts[i] += strings.Repeat(" ", indent-widths[i]) + cfg.Desc
			i++
		}
	}

	b.texts = texts
	help.blocks = append(help.blocks, b)
}

func makeOptTitle(cfg OptCfg) string {
	title := ""
	for _, name := range cfg.Names {
		switch len(name) {
		case 0:
		case 1:
			title += ", -" + name
		default:
			title += ", --" + name
		}
	}

	if cfg.HasArg && len(cfg.ArgInHelp) > 0 {
		title += " " + cfg.ArgInHelp
	}

	return title[2:]
}

// Print is a method which prints help texts to standard output.
func (help Help) Print() {
	iter := help.Iter()

	for {
		line, more := iter.Next()
		fmt.Println(line)
		if !more {
			break
		}
	}
}
