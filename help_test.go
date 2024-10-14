package cliargs_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/cliargs"
	"github.com/sttk/linebreak"
)

func TestHelp_NewHelp_empty(t *testing.T) {
	help := cliargs.NewHelp()
	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_NewHelpAndAddText_ifOneLineWithZeroWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("abc")

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "abc")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_NewHelpAndAddText_ifOneLineWithWrapping(t *testing.T) {
	termCols := linebreak.TermCols()
	text := strings.Repeat("a", termCols)
	text += "123456"

	help := cliargs.NewHelp()
	help.AddText(text)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, strings.Repeat("a", termCols))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "123456")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_NewHelpAndAddText_ifMultiLinesWithWrapping(t *testing.T) {
	termCols := linebreak.TermCols()
	text := strings.Repeat("a", termCols)
	text += "123456\n"
	text += strings.Repeat("b", termCols)
	text += "789"

	help := cliargs.NewHelp()
	help.AddText(text)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, strings.Repeat("a", termCols))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "123456")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, strings.Repeat("b", termCols))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "789")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_NewHelpWithMergins_andAddText(t *testing.T) {
	termCols := linebreak.TermCols()

	text := strings.Repeat("a", termCols-5-3)
	text += "12345\n"
	text += strings.Repeat("b", termCols-5-3)
	text += "6789"

	help := cliargs.NewHelpWithMargins(5, 3)
	help.AddText(text)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "     "+strings.Repeat("a", termCols-5-3))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "     12345")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "     "+strings.Repeat("b", termCols-5-3))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "     6789")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddTextWithMargings(t *testing.T) {
	termCols := linebreak.TermCols()

	text := strings.Repeat("a", termCols-5-3)
	text += "12345\n"
	text += strings.Repeat("b", termCols-5-3)
	text += "6789"

	help := cliargs.NewHelp()
	help.AddTextWithMargins(text, 5, 3)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "     "+strings.Repeat("a", termCols-5-3))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "     12345")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "     "+strings.Repeat("b", termCols-5-3))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "     6789")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddTextWithIndent(t *testing.T) {
	termCols := linebreak.TermCols()

	text := strings.Repeat("a", termCols)
	text += "12345\n"
	text += strings.Repeat("b", termCols-8)
	text += "6789"

	help := cliargs.NewHelp()
	help.AddTextWithIndent(text, 8)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, strings.Repeat("a", termCols))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "        12345")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "        "+strings.Repeat("b", termCols-8))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "        6789")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddTextWithIndentAndMargins(t *testing.T) {
	termCols := linebreak.TermCols()

	text := strings.Repeat("a", termCols-1-2)
	text += "12345\n"
	text += strings.Repeat("b", termCols-8-1-2)
	text += "6789"

	help := cliargs.NewHelp()
	help.AddTextWithIndentAndMargins(text, 8, 1, 2)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, " "+strings.Repeat("a", termCols-1-2))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "         12345")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "         "+strings.Repeat("b", termCols-8-1-2))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "         6789")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AdddText_ifTextIsEmpty(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("")

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddText_multipleTimes(t *testing.T) {
	termCols := linebreak.TermCols()

	text := strings.Repeat("a", termCols-4-3) +
		strings.Repeat("b", termCols-4-5-3) +
		strings.Repeat("c", termCols-4-5-3)

	help := cliargs.NewHelpWithMargins(1, 1)
	help.AddTextWithIndentAndMargins(text, 5, 3, 2)

	text = strings.Repeat("d", termCols-2-2) +
		strings.Repeat("e", termCols-2-5-2) +
		strings.Repeat("f", termCols-2-5-2)

	help.AddTextWithIndentAndMargins(text, 5, 1, 1)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "    "+strings.Repeat("a", termCols-7))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "         "+strings.Repeat("b", termCols-12))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "         "+strings.Repeat("c", termCols-12))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "  "+strings.Repeat("d", termCols-4))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "       "+strings.Repeat("e", termCols-9))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "       "+strings.Repeat("f", termCols-9))
}

func TestHelp_AddTexts_ifArrayIsEmpty(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddTexts([]string{})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddTexts_ifArrayHasAText(t *testing.T) {
	termCols := linebreak.TermCols()

	texts := []string{
		strings.Repeat("a", termCols) +
			strings.Repeat("b", termCols) +
			strings.Repeat("c", termCols),
	}

	help := cliargs.NewHelp()
	help.AddTexts(texts)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("a", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("b", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("c", termCols))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddTexts_ifArrayHasMultipleTexts(t *testing.T) {
	termCols := linebreak.TermCols()

	texts := []string{
		strings.Repeat("a", termCols) +
			strings.Repeat("b", termCols) +
			strings.Repeat("c", termCols),

		strings.Repeat("d", termCols) +
			strings.Repeat("e", termCols) +
			strings.Repeat("f", termCols),
	}

	help := cliargs.NewHelp()
	help.AddTexts(texts)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("a", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("b", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("c", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("d", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("e", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("f", termCols))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddTextsWithIndent(t *testing.T) {
	termCols := linebreak.TermCols()

	texts := []string{
		strings.Repeat("a", termCols) +
			strings.Repeat("b", termCols-5) +
			strings.Repeat("c", termCols-5),

		strings.Repeat("d", termCols) +
			strings.Repeat("e", termCols-5) +
			strings.Repeat("f", termCols-5),
	}

	help := cliargs.NewHelp()
	help.AddTextsWithIndent(texts, 5)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("a", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "     "+strings.Repeat("b", termCols-5))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "     "+strings.Repeat("c", termCols-5))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat("d", termCols))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "     "+strings.Repeat("e", termCols-5))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "     "+strings.Repeat("f", termCols-5))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddTextsWithMargins(t *testing.T) {
	termCols := linebreak.TermCols()

	texts := []string{
		strings.Repeat("a", termCols-3-3) +
			strings.Repeat("b", termCols-3-3) +
			strings.Repeat("c", termCols-3-3),

		strings.Repeat("d", termCols-3-3) +
			strings.Repeat("e", termCols-3-3) +
			strings.Repeat("f", termCols-3-3),
	}

	help := cliargs.NewHelpWithMargins(1, 1)
	help.AddTextsWithMargins(texts, 2, 2)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("a", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("b", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("c", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("d", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("e", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("f", termCols-6))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddTextsWithIndentAndMargins(t *testing.T) {
	termCols := linebreak.TermCols()

	texts := []string{
		strings.Repeat("a", termCols-3-3) +
			strings.Repeat("b", termCols-3-5-3) +
			strings.Repeat("c", termCols-3-5-3),

		strings.Repeat("d", termCols-3-3) +
			strings.Repeat("e", termCols-3-5-3) +
			strings.Repeat("f", termCols-3-5-3),
	}

	help := cliargs.NewHelpWithMargins(1, 1)
	help.AddTextsWithIndentAndMargins(texts, 5, 2, 2)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("a", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 8)+strings.Repeat("b", termCols-11))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 8)+strings.Repeat("c", termCols-11))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 3)+strings.Repeat("d", termCols-6))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 8)+strings.Repeat("e", termCols-11))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 8)+strings.Repeat("f", termCols-11))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOpts_withNoWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc:  "This is a description of option.",
		},
	})
	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "--foo-bar  This is a description of option.")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddOpts_withWrapping(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc:  strings.Repeat("a", termCols-11) + " bcdef",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.Equal(t, line, "--foo-bar  "+strings.Repeat("a", termCols-11))
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "           "+"bcdef")
	assert.True(t, exists)

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_NewHelpWithMargins_and_AddOpts(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelpWithMargins(4, 2)

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc: strings.Repeat("a", termCols-11-4-2) + " " + strings.Repeat("b", termCols-11-4-2) +
				"ccc",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "    --foo-bar  "+strings.Repeat("a", termCols-11-4-2))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 11+4)+strings.Repeat("b", termCols-11-4-2))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 11+4)+"ccc")

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddOptsWithMargins(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOptsWithMargins([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc: strings.Repeat("a", termCols-11-5-4) + " " + strings.Repeat("b", termCols-11-5-4) +
				"ccc",
		},
	}, 5, 4)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "     --foo-bar  "+strings.Repeat("a", termCols-11-5-4))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 11+5)+strings.Repeat("b", termCols-11-5-4))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 11+5)+"ccc")

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddOptsWithMargins_bothOfNewMethod_andAddTextWithMargins(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelpWithMargins(4, 2)

	help.AddOptsWithMargins([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc: strings.Repeat("a", termCols-11-5-4) + " " + strings.Repeat("b", termCols-11-5-4) +
				"ccc",
		},
	}, 1, 2)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "     --foo-bar  "+strings.Repeat("a", termCols-11-5-4))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 11+5)+strings.Repeat("b", termCols-11-5-4))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 11+5)+"ccc")

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddOptsWithIndent_ifIndentIsLongerThanTitle(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOptsWithIndent([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc:  strings.Repeat("a", termCols-12) + " " + strings.Repeat("b", termCols-12) + "ccc",
		},
	}, 12)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--foo-bar   "+strings.Repeat("a", termCols-12))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 12)+strings.Repeat("b", termCols-12))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 12)+"ccc")

	line, exists = iter.Next()
	assert.Equal(t, line, "")
	assert.False(t, exists)
}

func TestHelp_AddOptsWithIndent_ifIndentIsShorterThanTitle(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOptsWithIndent([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc:  strings.Repeat("a", termCols),
		},
	}, 10)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--foo-bar")

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 10)+strings.Repeat("a", termCols-10))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 10)+strings.Repeat("a", 10))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOptsWithIndentAndMargins(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOptsWithIndentAndMargins([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc:  strings.Repeat("a", termCols),
		},
	}, 6, 4, 2)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "    --foo-bar")

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 10)+strings.Repeat("a", termCols-6-4-2))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 10)+strings.Repeat("a", 6+4+2))

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOpts_ifOptsAreMultiple(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names:     []string{"foo-bar", "f"},
			HasArg:    true,
			Desc:      strings.Repeat("a", termCols-22) + " " + strings.Repeat("b", termCols-22) + "ccc",
			ArgInHelp: "<text>",
		},
		cliargs.OptCfg{
			Names: []string{"baz", "b"},
			Desc:  strings.Repeat("d", termCols-22) + " " + strings.Repeat("e", termCols-22) + "fff",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--foo-bar, -f <text>  "+strings.Repeat("a", termCols-22))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 22)+strings.Repeat("b", termCols-22))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 22)+"ccc")

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--baz, -b             "+strings.Repeat("d", termCols-22))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 22)+strings.Repeat("e", termCols-22))

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, strings.Repeat(" ", 22)+"fff")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOpts_ifNamesAreEmptyAndStoreKeyIsSpecified(t *testing.T) {
	help := cliargs.NewHelp()

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
			Desc:  "description",
		},
		cliargs.OptCfg{
			Names: []string{"*"},
			Desc:  "any option",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--foo-bar  description")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOpts_ifStoreKeyIsAnyOption(t *testing.T) {
	help := cliargs.NewHelp()

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "foo",
			Desc:     "description",
		},
		cliargs.OptCfg{
			StoreKey: "*",
			Desc:     "any option",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--foo  description")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOpts_ifFirstElementOfNamesIsAnyOption(t *testing.T) {
	help := cliargs.NewHelp()

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "foo-bar",
			Desc:     "description",
		},
		cliargs.OptCfg{
			StoreKey: "*",
			Desc:     "any option",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "--foo-bar  description")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOptsWithIndent_ifIndentIsLongerThanLineWidth(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOptsWithIndent([]cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "foo-bar",
			Desc:     "description",
		},
		cliargs.OptCfg{
			StoreKey: "baz",
			Desc:     "description",
		},
	}, termCols+1)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOptsWithMargins_ifSumOfMarginsAreEqualToLineWidth(t *testing.T) {
	termCols := linebreak.TermCols()

	help := cliargs.NewHelp()

	help.AddOptsWithMargins([]cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "foo-bar",
			Desc:     "description",
		},
		cliargs.OptCfg{
			StoreKey: "baz",
			Desc:     "description",
		},
	}, termCols-1, 1)

	iter := help.Iter()

	line, exists := iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}

func TestHelp_AddOpts_ifNamesContainsEmptyStrings(t *testing.T) {
	help := cliargs.NewHelp()

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"", "f", "foo-bar", "", ""},
			Desc:  "description",
		},
		cliargs.OptCfg{
			Names: []string{"b", "", "z", "baz"},
			Desc:  "description",
		},
	})

	iter := help.Iter()

	line, exists := iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "    -f, --foo-bar  description")

	line, exists = iter.Next()
	assert.True(t, exists)
	assert.Equal(t, line, "-b,     -z, --baz  description")

	line, exists = iter.Next()
	assert.False(t, exists)
	assert.Equal(t, line, "")
}
