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
	marginLeft  int
	marginRight int
	blocks      []block
}

type block struct {
	indent      int
	marginLeft  int
	marginRight int
	bodies      []blockBody
}

type blockBody struct {
	firstIndent int
	text        string
}

// NewHelp is a function to construct a Help instance.
func NewHelp() Help {
	var help Help
	help.blocks = make([]block, 0, 2)
	return help
}

// NewHelpWithMargins is a function to construct a Help instance with setting left and right
// margins.
func NewHelpWithMargins(marginLeft, marginRight int) Help {
	var help Help
	help.marginLeft = marginLeft
	help.marginRight = marginRight
	help.blocks = make([]block, 0, 2)
	return help
}

// AddText is a method which adds a text to this Help instance.
func (help *Help) AddText(text string) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.bodies = []blockBody{
		blockBody{firstIndent: 0, text: text},
	}
	help.blocks = append(help.blocks, b)
}

// AddTextWithIndent is a method which adds a text with indent size to this Help instance.
func (help *Help) AddTextWithIndent(text string, indent int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.indent = indent
	b.bodies = []blockBody{
		blockBody{firstIndent: 0, text: text},
	}
	help.blocks = append(help.blocks, b)
}

// AddTextWithMargins is a method which adds a text with left and right mergins to this Help
// instance.
func (help *Help) AddTextWithMargins(text string, marginLeft, marginRight int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.marginLeft += marginLeft
	b.marginRight += marginRight
	b.bodies = []blockBody{
		blockBody{firstIndent: 0, text: text},
	}
	help.blocks = append(help.blocks, b)
}

// AddTextWithIndnetAndMargins is a method which adds a text with  indent size and left and right
// mergins to this Help instance.
func (help *Help) AddTextWithIndentAndMargins(text string, indent, marginLeft, marginRight int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.indent = indent
	b.marginLeft += marginLeft
	b.marginRight += marginRight
	b.bodies = []blockBody{
		blockBody{firstIndent: 0, text: text},
	}
	help.blocks = append(help.blocks, b)
}

// AddTexts is a method which adds an array of texts to this Help instance.
func (help *Help) AddTexts(texts []string) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	bodies := make([]blockBody, 0, len(texts))
	for _, text := range texts {
		bodies = append(bodies, blockBody{firstIndent: 0, text: text})
	}
	b.bodies = bodies
	help.blocks = append(help.blocks, b)
}

// AddTextsWithIndent is a method which adds an array of texts with indent size to this Help
// instance.
func (help *Help) AddTextsWithIndent(texts []string, indent int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.indent = indent
	bodies := make([]blockBody, 0, len(texts))
	for _, text := range texts {
		bodies = append(bodies, blockBody{firstIndent: 0, text: text})
	}
	b.bodies = bodies
	help.blocks = append(help.blocks, b)
}

// AddTextsWithMargins is a method which adds an array of texts with left and right mergins to
// this Help instance.
func (help *Help) AddTextsWithMargins(texts []string, marginLeft, marginRight int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.marginLeft += marginLeft
	b.marginRight += marginRight
	bodies := make([]blockBody, 0, len(texts))
	for _, text := range texts {
		bodies = append(bodies, blockBody{firstIndent: 0, text: text})
	}
	b.bodies = bodies
	help.blocks = append(help.blocks, b)
}

// AddTextsWithIndnetAndMargins is a method which adds an array of texts with  indent size and left
// and right mergins to this Help instance.
func (help *Help) AddTextsWithIndentAndMargins(
	texts []string, indent, marginLeft, marginRight int,
) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.indent = indent
	b.marginLeft += marginLeft
	b.marginRight += marginRight
	bodies := make([]blockBody, 0, len(texts))
	for _, text := range texts {
		bodies = append(bodies, blockBody{firstIndent: 0, text: text})
	}
	b.bodies = bodies
	help.blocks = append(help.blocks, b)
}

// AddOpts is a method which adds OptCfg(s) to this Help instance.
func (help *Help) AddOpts(optCfgs []OptCfg) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.bodies = createOptsHelp(optCfgs, &b.indent)
	help.blocks = append(help.blocks, b)
}

// AddOptsWithIndent is a method which adds OptCfg(s) with indent size to this Help instance.
func (help *Help) AddOptsWithIndent(optCfgs []OptCfg, indent int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.indent = indent
	b.bodies = createOptsHelp(optCfgs, &b.indent)
	help.blocks = append(help.blocks, b)
}

// AddOptsWithMargins is a method which adds OptCfg(s) with left and right margins to this Help
// instance.
func (help *Help) AddOptsWithMargins(optCfgs []OptCfg, marginLeft, marginRight int) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.marginLeft += marginLeft
	b.marginRight += marginRight
	b.bodies = createOptsHelp(optCfgs, &b.indent)
	help.blocks = append(help.blocks, b)
}

// AddOptsWithIndentAndMargins is a method which adds OptCfg(s) with indent size, left and right
// margins to this Help instance.
func (help *Help) AddOptsWithIndentAndMargins(
	optCfgs []OptCfg, indent, marginLeft, marginRight int,
) {
	b := block{
		marginLeft:  help.marginLeft,
		marginRight: help.marginRight,
	}
	b.indent = indent
	b.marginLeft += marginLeft
	b.marginRight += marginRight
	b.bodies = createOptsHelp(optCfgs, &b.indent)
	help.blocks = append(help.blocks, b)
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

// Print is a method which prints help texts to standard output.
func (help Help) Print() {
	iter := help.Iter()

	for {
		line, exists := iter.Next()
		if !exists {
			break
		}
		fmt.Println(line)
	}
}

func createOptsHelp(optCfgs []OptCfg, indent *int) []blockBody {
	bodies := make([]blockBody, 0, len(optCfgs))
	const ANY_OPT string = "*"

	if *indent > 0 {
		for _, cfg := range optCfgs {
			var names []string
			for _, nm := range cfg.Names {
				if len(nm) > 0 {
					names = append(names, nm)
					break
				}
			}

			storeKey := cfg.StoreKey
			if len(storeKey) == 0 && len(names) > 0 {
				storeKey = names[0]
			}

			if len(storeKey) == 0 {
				continue
			}

			if storeKey == ANY_OPT {
				continue
			}

			firstIndent, text := makeOptTitle(cfg)

			width := firstIndent + linebreak.TextWidth(text)

			if len(cfg.Desc) > 0 {
				if width+2 > *indent {
					text += "\n" + strings.Repeat(" ", *indent) + cfg.Desc
				} else {
					text += strings.Repeat(" ", *indent-width) + cfg.Desc
				}
			}

			bodies = append(bodies, blockBody{firstIndent, text})
		}
	} else {
		widths := make([]int, 0, len(bodies))
		maxIndent := 0

		for _, cfg := range optCfgs {
			var names []string
			for _, nm := range cfg.Names {
				if len(nm) > 0 {
					names = append(names, nm)
					break
				}
			}

			storeKey := cfg.StoreKey
			if len(storeKey) == 0 && len(names) > 0 {
				storeKey = names[0]
			}

			if len(storeKey) == 0 {
				continue
			}

			if storeKey == ANY_OPT {
				continue
			}

			firstIndent, text := makeOptTitle(cfg)

			width := firstIndent + linebreak.TextWidth(text)
			if maxIndent < width {
				maxIndent = width
			}

			bodies = append(bodies, blockBody{firstIndent, text})
			widths = append(widths, width)
		}

		maxIndent += 2
		*indent = maxIndent

		i := 0
		for _, cfg := range optCfgs {
			var names []string
			for _, nm := range cfg.Names {
				if len(nm) > 0 {
					names = append(names, nm)
					break
				}
			}

			storeKey := cfg.StoreKey
			if len(storeKey) == 0 && len(names) > 0 {
				storeKey = names[0]
			}

			if len(storeKey) == 0 {
				continue
			}

			if storeKey == ANY_OPT {
				continue
			}

			if len(cfg.Desc) > 0 {
				bodies[i].text += strings.Repeat(" ", maxIndent-widths[i]) + cfg.Desc
			}

			i += 1
		}
	}

	return bodies
}

func makeOptTitle(cfg OptCfg) (int, string) {
	headSpaces := 0
	lastSpaces := 0
	title := ""
	useStoreKey := true

	n := len(cfg.Names)

	for i, name := range cfg.Names {
		switch len(name) {
		case 0:
			if len(title) == 0 {
				headSpaces += 4
			} else if i != n-1 {
				lastSpaces += 4
			} else {
				lastSpaces += 2
			}
		case 1:
			if lastSpaces > 0 {
				title += "," + strings.Repeat(" ", lastSpaces-1)
			}
			lastSpaces = 0
			title += "-" + name
			if i != n-1 {
				lastSpaces += 2
			}
			useStoreKey = false
		default:
			if lastSpaces > 0 {
				title += "," + strings.Repeat(" ", lastSpaces-1)
			}
			lastSpaces = 0
			title += "--" + name
			if i != n-1 {
				lastSpaces += 2
			}
			useStoreKey = false
		}
	}

	if useStoreKey {
		switch len(cfg.StoreKey) {
		case 0:
		case 1:
			if lastSpaces > 0 {
				title += "," + strings.Repeat(" ", lastSpaces-1)
			}
			title += "-" + cfg.StoreKey
		default:
			if lastSpaces > 0 {
				title += "," + strings.Repeat(" ", lastSpaces-1)
			}
			title += "--" + cfg.StoreKey
		}
	}

	if len(cfg.ArgInHelp) > 0 {
		title += " " + cfg.ArgInHelp
	}

	return headSpaces, title
}

// HelpIter is a struct type to iterate lines of help texts.
type HelpIter struct {
	lineWidth int
	blocks    []block
	blockIter blockIter
}

type blockIter struct {
	bodies   []blockBody
	index    int
	indent   string
	margin   string
	lineIter linebreak.LineIter
}

// Next is a method which returns a line of a help text and a status which indicates this HelpIter
// has more texts or not.
// If there are more lines, the returned IterStatus value is true, otherwise the value is false.
func (iter *HelpIter) Next() (string, bool) {
	for {
		line, exists := iter.blockIter.next()
		if exists {
			return line, true
		}
		if len(iter.blocks) <= 1 {
			return "", false
		}
		iter.blocks = iter.blocks[1:]
		iter.blockIter = newBlockIter(iter.blocks[0], iter.lineWidth)
	}
}

func newBlockIter(block block, lineWidth int) blockIter {
	printWidth := lineWidth - block.marginLeft - block.marginRight
	if len(block.bodies) == 0 || printWidth <= block.indent {
		return blockIter{}
	}

	body := block.bodies[0]
	lineIter := linebreak.New(body.text, printWidth)
	lineIter.SetIndent(strings.Repeat(" ", body.firstIndent))

	return blockIter{
		bodies:   block.bodies,
		index:    0,
		indent:   strings.Repeat(" ", block.indent),
		margin:   strings.Repeat(" ", block.marginLeft),
		lineIter: lineIter,
	}
}

func (iter *blockIter) next() (string, bool) {
	if len(iter.bodies) == 0 {
		return "", false
	}

	for {
		line, exists := iter.lineIter.Next()
		iter.lineIter.SetIndent(iter.indent)
		if len(line) > 0 {
			line = iter.margin + line
		}
		if exists {
			return line, true
		}
		iter.index++
		if iter.index >= len(iter.bodies) {
			return "", false
		}
		body := iter.bodies[iter.index]
		iter.lineIter.Init(body.text)
		iter.lineIter.SetIndent(strings.Repeat(" ", body.firstIndent))
	}
}
