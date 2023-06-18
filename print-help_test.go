package cliargs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/cliargs"
)

func TestNewHelp_empty(t *testing.T) {
	help := cliargs.NewHelp()
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_oneLine_withNoWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("abc")
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "abc")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_multiLines_withNoWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("abc\ndef")
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "abc")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "def")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_oneLine_withWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789i12345")
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "i12345")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_multiLines_withWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789i12345\n6789j123456789k123456789l123456789m123456789n123456789o123456789p123456789q12345678")
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "i12345")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "6789j123456789k123456789l123456789m123456789n123456789o123456789p123456789q12345")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_marginLeftByNewArgHelp(t *testing.T) {
	help := cliargs.NewHelp(5)
	help.AddText("abc\ndef")
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     abc")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     def")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_marginLeftAndMarginRightByNewArgHelps(t *testing.T) {
	help := cliargs.NewHelp(5, 3)
	help.AddText("a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789i12345\n6789")
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     a123456789b123456789c123456789d123456789e123456789f123456789g123456789h1")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     23456789i12345")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     6789")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_marginLeftByiAddHelpTextArg(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("abc\ndef", 0, 5)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     abc")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     def")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_marginLeftAndMarginRightByddHelpTextArgs(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789i12345\n6789", 0, 5, 3)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     a123456789b123456789c123456789d123456789e123456789f123456789g123456789h1")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     23456789i12345")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     6789")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_Indent(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddText("a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789i12345\n6789", 8)
	help.AddText("a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789i12345\n6789", 5)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "        i12345")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "        6789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "a123456789b123456789c123456789d123456789e123456789f123456789g123456789h123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     i12345")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     6789")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddTexts_zeroText(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddTexts([]string{})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	help = cliargs.NewHelp()
	help.AddTexts([]string(nil))
	iter = help.Iter()

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddTexts_oneText(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddTexts([]string{
		"a12345678 b12345678 c12345678 d12345678 " +
			"e12345678 f12345678 g12345678 h12345678 i123"})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "a12345678 b12345678 c12345678 d12345678 "+
		"e12345678 f12345678 g12345678 h12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "i123")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddTexts_multipleTexts(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddTexts([]string{
		"a12345678 b12345678 c12345678 d12345678 " +
			"e12345678 f12345678 g12345678 h12345678 i123",
		"j12345678 k12345678 l12345678 m12345678 " +
			"n12345678 o12345678 p12345678 q12345678 r123",
	})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "a12345678 b12345678 c12345678 d12345678 "+
		"e12345678 f12345678 g12345678 h12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "i123")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "j12345678 k12345678 l12345678 m12345678 "+
		"n12345678 o12345678 p12345678 q12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "r123")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddTexts_withIndent(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddTexts([]string{
		"a12345678  123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
		"b1234      123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
	}, 11)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "a12345678  123456789 123456789 123456789 123456789 123456789 123456789 123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "           123456789 123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "b1234      123456789 123456789 123456789 123456789 123456789 123456789 123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "           123456789 123456789")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddTexts_withMargins(t *testing.T) {
	help := cliargs.NewHelp(2, 2)
	help.AddTexts([]string{
		"a12345678  123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
		"b1234      123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789",
	}, 11, 5, 3)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "       a12345678  123456789 123456789 123456789 123456789 123456789 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                  123456789 123456789 123456789 123456789")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "       b1234      123456789 123456789 123456789 123456789 123456789 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                  123456789 123456789 123456789 123456789")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_zeroOpts(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withNoWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo-bar",
			Desc: "This is a description of option.",
		},
	})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar  This is a description of option.")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withWrapping(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f", "foo", "b", "bar"},
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678",
		},
	})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar, -f, --foo, -b, --bar  a12345678 b12345678 c12345678 d12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                                 e12345678 f12345678 g12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withMarginsByNewHelpArg(t *testing.T) {
	help := cliargs.NewHelp(5, 3)
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 h12345678",
			ArgHelp: "<text>",
		},
	})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     --foo-bar, -f <text>  a12345678 b12345678 c12345678 d12345678 e12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                           f12345678 g12345678 h12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withMarginsByAddOptsArg(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 h12345678",
			ArgHelp: "<text>",
		},
	}, 0, 5, 3)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     --foo-bar, -f <text>  a12345678 b12345678 c12345678 d12345678 e12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                           f12345678 g12345678 h12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withMarginsByNewHelpArgAndAddOptsArg(t *testing.T) {
	help := cliargs.NewHelp(2, 2)
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 h12345678",
			ArgHelp: "<text>",
		},
	}, 0, 3, 1)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "     --foo-bar, -f <text>  a12345678 b12345678 c12345678 d12345678 e12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                           f12345678 g12345678 h12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withIndentLongerThanTitle(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 h12345678",
			ArgHelp: "<text>",
		},
	}, 25)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar, -f <text>     a12345678 b12345678 c12345678 d12345678 e12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                         f12345678 g12345678 h12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_oneOpts_withIndentShorterThanTitle(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 h12345678",
			ArgHelp: "<text>",
		},
	}, 10)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar, -f <text>")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "          a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "          h12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_multipleOpts(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			Desc:    "a12345678 b12345678 c12345678 d12345678 e12345678 f12345678 g12345678 h12345678",
			ArgHelp: "<text>",
		},
		cliargs.OptCfg{
			Name:    "baz",
			Aliases: []string{"b"},
			Desc:    "i12345678 j12345678 k12345678 l12345678 m12345678 n12345678 o12345678 p12345678",
		},
	})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar, -f <text>  a12345678 b12345678 c12345678 d12345678 e12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                      f12345678 g12345678 h12345678")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--baz, -b             i12345678 j12345678 k12345678 l12345678 m12345678 ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                      n12345678 o12345678 p12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_hasAnyOption(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo-bar",
			Desc: "a12345678 b12345678",
		},
		cliargs.OptCfg{
			Name: "*",
			Desc: "c12345678 d12345678",
		},
	})
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar  a12345678 b12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_hasAnyOption_withIndent(t *testing.T) {
	help := cliargs.NewHelp()
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "*",
			Desc: "c12345678 d12345678",
		},
		cliargs.OptCfg{
			Name: "foo-bar",
			Desc: "a12345678 b12345678",
		},
	}, 5)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "--foo-bar")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     a12345678 b12345678")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestNewHelp_ifLineWidthLessThanSumOfMargins(t *testing.T) {
	help := cliargs.NewHelp(71, 10)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddText_ifLineWidthLessThanSumOfMargins(t *testing.T) {
	help := cliargs.NewHelp(10, 40)
	help.AddText("abcdefg", 10, 10, 10)
	help.AddText("hijklmn", 10, 10)
	help.AddText("opqrstu", 10, 10, 11)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                    hijklmn")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestAddOpts_ifLineWidthLessThanSunOfMargins(t *testing.T) {
	help := cliargs.NewHelp(10, 30)
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo-bar",
			Desc: "This is a description of option.",
		},
	}, 10, 10, 20)
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "baz",
			Desc: "This is a description of option.",
		},
	}, 10, 10)
	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "qux",
			Desc: "This is a description of option.",
		},
	}, 10, 10, 21)
	iter := help.Iter()

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                    --baz     This is a ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                              description of ")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                              option.")
	assert.Equal(t, status, cliargs.ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, cliargs.ITER_NO_MORE)
}

func TestPrint_curl(t *testing.T) {
	// The source of the following text is the output of `curl --help` in
	// curl 7.87.0. (https://curl.se/docs/copyright.html)

	help := cliargs.NewHelp()

	help.AddText("Usage: curl [options...] <url>")

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "d",
			Aliases: []string{"data"},
			Desc:    "HTTP POST data",
			HasArg:  true,
			ArgHelp: "<data>",
		},
		cliargs.OptCfg{
			Name:    "f",
			Aliases: []string{"fail"},
			Desc:    "Fail fast with no output on HTTP errors",
		},
		cliargs.OptCfg{
			Name:    "h",
			Aliases: []string{"help"},
			Desc:    "Get help for commands",
			HasArg:  true,
			ArgHelp: "<category>",
		},
		cliargs.OptCfg{
			Name:    "i",
			Aliases: []string{"include"},
			Desc:    "Include protocol response headers in the output",
		},
		cliargs.OptCfg{
			Name:    "o",
			Aliases: []string{"output"},
			Desc:    "Write to file instead of stdout",
			HasArg:  true,
			ArgHelp: "<file>",
		},
		cliargs.OptCfg{
			Name:    "O",
			Aliases: []string{"remote-name"},
			Desc:    "Write output to a file named as the remote file",
		},
		cliargs.OptCfg{
			Name:    "s",
			Aliases: []string{"silent"},
			Desc:    "Silent mode",
		},
		cliargs.OptCfg{
			Name:    "T",
			Aliases: []string{"upload-file"},
			Desc:    "Transfer local FILE to destination",
			HasArg:  true,
			ArgHelp: "<file>",
		},
		cliargs.OptCfg{
			Name:    "u",
			Aliases: []string{"user"},
			Desc:    "Server user and password",
			HasArg:  true,
			ArgHelp: "<user:password>",
		},
		cliargs.OptCfg{
			Name:    "A",
			Aliases: []string{"user-agent"},
			Desc:    "Send User-Agent <name> to server",
			HasArg:  true,
			ArgHelp: "<name>",
		},
		cliargs.OptCfg{
			Name:    "v",
			Aliases: []string{"verbose"},
			Desc:    "Make the operation more talkative",
		},
		cliargs.OptCfg{
			Name:    "V",
			Aliases: []string{"version"},
			Desc:    "Show version number and quit",
		},
	}, 0, 1)

	help.AddText(`
This is not the full help, this menu is stripped into categories.
Use "--help category" to get an overview of all categories.
For all options use the manual or "--help all".`)

	help.Print()
}
