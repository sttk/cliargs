// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

type Help struct {
	marginLeft, marginRight int
	blocks                  []block
}

type block struct {
	indent, marginLeft, marginRight int
	texts                           []string
}

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

func (help Help) Iter() HelpIter {
	if len(help.blocks) == 0 {
		return HelpIter{}
	}

	lineWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		lineWidth = 80
	}

	return HelpIter{
		lineWidth: lineWidth,
		blocks:    help.blocks,
		blockIter: newBlockIter(
			help.blocks[0].texts,
			lineWidth-help.blocks[0].marginLeft-help.blocks[0].marginRight,
			help.blocks[0].indent,
			help.blocks[0].marginLeft,
		),
	}
}

type HelpIter struct {
	lineWidth int
	blocks    []block
	blockIter blockIter
}

func (iter *HelpIter) Next() (string, IterStatus) {
	line, status := iter.blockIter.next()
	if status == ITER_NO_MORE {
		if len(iter.blocks) <= 1 {
			return line, ITER_NO_MORE
		}
		iter.blocks = iter.blocks[1:]
		iter.blockIter = newBlockIter(
			iter.blocks[0].texts,
			iter.lineWidth-iter.blocks[0].marginLeft-iter.blocks[0].marginRight,
			iter.blocks[0].indent,
			iter.blocks[0].marginLeft,
		)
	}
	return line, ITER_HAS_MORE
}

type blockIter struct {
	texts    []string
	index    int
	indent   int
	margin   string
	lineIter lineIter
}

func newBlockIter(texts []string, lineWidth, indent, margin int) blockIter {
	if len(texts) == 0 {
		return blockIter{}
	}
	return blockIter{
		texts:    texts,
		indent:   indent,
		margin:   strings.Repeat(" ", margin),
		lineIter: newLineIter(texts[0], lineWidth),
	}
}

func (iter *blockIter) next() (string, IterStatus) {
	if len(iter.texts) == 0 {
		return "", ITER_NO_MORE
	}

	line, status := iter.lineIter.Next()
	if len(line) > 0 {
		line = iter.margin + line
	}
	if status == ITER_NO_MORE {
		iter.index++
		if iter.index >= len(iter.texts) {
			return line, status
		}
		iter.lineIter.resetText(iter.texts[iter.index])
		iter.lineIter.setIndent(0)
		return line, ITER_HAS_MORE
	}

	iter.lineIter.setIndent(iter.indent)
	return line, status
}

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
			if cfg.Name == anyOption {
				continue
			}
			texts[i] = makeOptTitle(cfg)
			width := textWidth(texts[i])
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
			if cfg.Name == anyOption {
				continue
			}
			texts[i] = makeOptTitle(cfg)
			widths[i] = textWidth(texts[i])
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
			if cfg.Name == anyOption {
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
	title := cfg.Name
	switch len(title) {
	case 0:
	case 1:
		title = "-" + title
	default:
		title = "--" + title
	}

	for _, alias := range cfg.Aliases {
		switch len(alias) {
		case 0:
		case 1:
			title += ", -" + alias
		default:
			title += ", --" + alias
		}
	}

	if cfg.HasArg && len(cfg.ArgHelp) > 0 {
		title += " " + cfg.ArgHelp
	}

	return title
}

func textWidth(text string) int {
	w := 0
	for _, r := range text {
		w += runeWidth(r)
	}
	return w
}

func (help Help) Print() {
	iter := help.Iter()

	for {
		line, status := iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}
