// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"github.com/sttk-go/sabi"
	"golang.org/x/term"
	"os"
	"strings"
)

type /* error reason */ (
	// MarginsAndIndentExceedLineWidth is an error reason which indicates that
	// the sum of left margin, right margin, and indent is greater than line
	// width.
	// It means that there are no width to print texts.
	MarginsAndIndentExceedLineWidth struct {
		LineWidth, MarginLeft, MarginRight, Indent int
	}
)

// WrapOpts is a struct type which holds options for wrapping texts.
// This struct type has the following field for wrap options: MarginLeft,
// MarginRight, and Indent.
// MarginLeft and MarginRight is space widths on both sides, and these are
// output on all lines.
// Indent is space width on left side, and this is output on second line or
// later of each option.
type WrapOpts struct {
	MarginLeft  int
	MarginRight int
	Indent      int
}

// HelpIter is a struct type to iterate lines of help texts.
type HelpIter struct {
	texts    []string
	index    int
	indent   int
	margin   string
	lineIter lineIter
}

func newHelpIter(texts []string, lineWidth, indent, margin int) HelpIter {
	if len(texts) == 0 {
		return HelpIter{}
	}
	return HelpIter{
		texts:    texts,
		indent:   indent,
		margin:   strings.Repeat(" ", margin),
		lineIter: newLineIter(texts[0], lineWidth),
	}
}

// Next is a method which returns a line of help texts and a status which
// indicates this HelpIter has more texts or not.
// If there are more lines, the IterStatus value returned is ITER_HAS_MORE,
// otherwise the value is ITER_NO_MORE.
func (iter *HelpIter) Next() (string, IterStatus) {
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

	if iter.index > 0 {
		iter.lineIter.setIndent(iter.indent)
	}
	return line, status
}

// MakeHelp is a function to make a line iterator of help text.
// A help text consists of an usage section and options section, and options
// section consists of title parts and description parts.
func MakeHelp(
	usage string, optCfgs []OptCfg, wrapOpts WrapOpts,
) (HelpIter, sabi.Err) {
	lineWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	if lineWidth <= 0 {
		lineWidth = 80
	}

	texts := make([]string, len(optCfgs)+1)
	indent := 0

	for i, cfg := range optCfgs {
		t := makeOptTitle(cfg)
		if indent < len(t) {
			indent = len(t)
		}
		texts[i+1] = t
	}
	indent += 2

	if wrapOpts.Indent > 0 {
		indent = wrapOpts.Indent
	}

	if (wrapOpts.MarginLeft + wrapOpts.MarginRight + indent) >= lineWidth {
		return HelpIter{}, sabi.NewErr(MarginsAndIndentExceedLineWidth{
			LineWidth:   lineWidth,
			MarginLeft:  wrapOpts.MarginLeft,
			MarginRight: wrapOpts.MarginRight,
			Indent:      indent,
		})
	}
	lineWidth -= wrapOpts.MarginLeft + wrapOpts.MarginRight

	texts[0] = usage
	for i, cfg := range optCfgs {
		texts[i] += "\n"
		texts[i+1] = makeOptHelp(texts[i+1], cfg, indent)
	}

	return newHelpIter(texts, lineWidth, indent, wrapOpts.MarginLeft), sabi.Ok()
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

	return title
}

func makeOptHelp(title string, cfg OptCfg, indent int) string {
	if len(title)+2 > indent {
		title += "\n" + strings.Repeat(" ", indent) + cfg.Desc
	} else {
		title += strings.Repeat(" ", indent-len(title)) + cfg.Desc
	}
	return title
}

func PrintHelp(usage string, optCfgs []OptCfg, wrapOpts WrapOpts) sabi.Err {
	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	if !err.IsOk() {
		return err
	}

	for {
		line, status := iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
	return sabi.Ok()
}
