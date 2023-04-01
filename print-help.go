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

// MarginsAndIndentExceedLineWidth is an error which indicates that the sum of
// left margin, right margin, and indent is greater than line width.
// It means that there are no width to print texts.
type MarginsAndIndentExceedLineWidth struct {
	LineWidth, MarginLeft, MarginRight, Indent int
}

func (e MarginsAndIndentExceedLineWidth) Error() string {
	return "MarginsAndIndentExceedLineWidth"
}

// WrapOpts is a struct type which holds options for wrapping texts.
// This struct type has the following field for wrap options: MarginLeft,
// MarginRight, and Indent.
// MarginLeft and MarginRight is space widths on both sides, and these are
// applied to all lines.
// Indent is a space width on left side, and this is applied to second line or
// later of each option.
type WrapOpts struct {
	MarginLeft  int
	MarginRight int
	Indent      int
}

// HelpIter is a struct type to iterate lines of a help text.
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

// Next is a method which returns a line of a help text and a status which
// indicates this HelpIter has more texts or not.
// If there are more lines, the returned IterStatus value is ITER_HAS_MORE,
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

// MakeHelp is a function to make a line iterator of a help text.
// This function makes a help text from a usage text, option configurations
// ([]OptCfg), and wrap options (WrapOpts).
//
// A help text consists of an usage section and options section, and options
// section consists of title parts and description parts.
// On a title part, a option name, aliases, and a .AtParam field of OptCfg are
// enumerated.
// On a description part, a .Desc field of OptCfg is put and it follows a title
// part with an indent, specified in WrapOpts,
//
// On the both sides of a help text, left margin and right margin of which size
// are specified in WrapOpts can be put.
// These margins are applied to all lines of a help text.
//
// The sum of left margin, right margin, and indent have to be less than the
// line width, because if not, there is no width to output texts.
// The line width is obtained from the terminal width.
func MakeHelp(
	usage string, optCfgs []OptCfg, wrapOpts WrapOpts,
) (HelpIter, error) {
	lineWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
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
		return HelpIter{}, MarginsAndIndentExceedLineWidth{
			LineWidth:   lineWidth,
			MarginLeft:  wrapOpts.MarginLeft,
			MarginRight: wrapOpts.MarginRight,
			Indent:      indent,
		}
	}
	lineWidth -= wrapOpts.MarginLeft + wrapOpts.MarginRight

	texts[0] = usage
	for i, cfg := range optCfgs {
		texts[i+1] = makeOptHelp(texts[i+1], cfg, indent)
	}

	return newHelpIter(texts, lineWidth, indent, wrapOpts.MarginLeft), nil
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

	if cfg.HasParam && len(cfg.AtParam) > 0 {
		title += " " + cfg.AtParam
	}

	return title
}

func makeOptHelp(title string, cfg OptCfg, indent int) string {
	w := titleWidth(title)
	if w+2 > indent {
		title += "\n" + strings.Repeat(" ", indent) + cfg.Desc
	} else {
		title += strings.Repeat(" ", indent-w) + cfg.Desc
	}
	return title
}

func titleWidth(title string) int {
	w := 0
	for _, r := range title[:] {
		w += runeWidth(r)
	}
	return w
}

// PrintHelp is a function which output a help text to stdout.
// This function calls MakeHelp function to make a help text inside itself.
func PrintHelp(usage string, optCfgs []OptCfg, wrapOpts WrapOpts) error {
	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	if err != nil {
		return err
	}

	for {
		line, status := iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
	return nil
}
