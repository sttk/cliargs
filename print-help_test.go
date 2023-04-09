package cliargs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeHelp_emptyUsage_noOptCfg_emptyWrapOpts(t *testing.T) {
	optCfgs := []OptCfg{}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp("", optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)
}

func TestMakeHelp_shortUsage_noOptCfg_emptyWrapOpts(t *testing.T) {
	usage := "abcdefghijklmnopqrstuvwxyz"
	optCfgs := []OptCfg{}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage)
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)
}

// This text is quoted from https://go.dev/doc/
const longUsage string = "The Go programming language is an open source project to make programmers more productive."

func TestMakeHelp_longUsage_noOptCfg_emptyWrapOpts(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage[0:79])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, usage[79:90])
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)
}

func TestMakeHelp_longUsage_oneShortOptCfg_emptyWrapOpts(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
	}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage[0:79])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, usage[79:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--foo  This is the description of --foo option.")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestMakeHelp_longUsage_twoShortAndLongOptCfg_emptyWrapOpts(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "bar-baz",
			Aliases:  []string{"b"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage[0:79])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, usage[79:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--foo          This is the description of --foo option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--bar-baz, -b  This is the description of --bar-baz option. This option takes ")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "               one parameter.")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestMakeHelp_longUsage_twoShortAndLongOptCfg_largeIndent(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "bar-baz",
			Aliases:  []string{"b"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{Indent: 20}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage[0:79])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, usage[79:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--foo               This is the description of --foo option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--bar-baz, -b       This is the description of --bar-baz option. This option ")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                    takes one parameter.")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestMakeHelp_longUsage_twoShortAndLongOptCfg_shortIndent(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "bar-baz",
			Aliases:  []string{"b"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{Indent: 10}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage[0:79])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, usage[79:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--foo     This is the description of --foo option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "--bar-baz, -b")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "          This is the description of --bar-baz option. This option takes one ")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "          parameter.")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestMakeHelp_longUsage_twoShortAndLongOptCfg_margins(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "bar-baz",
			Aliases:  []string{"b"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{MarginLeft: 5, MarginRight: 5}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, "     "+usage[0:62])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     "+usage[62:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     --foo          This is the description of --foo option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     --bar-baz, -b  This is the description of --bar-baz option. This ")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                    option takes one parameter.")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestMakeHelp_optNameIsShortAndOptAliasIsLong(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "b",
			Aliases:  []string{"bar-baz"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{MarginLeft: 5, MarginRight: 5}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, "     "+usage[0:62])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     "+usage[62:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     --foo          This is the description of --foo option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "     -b, --bar-baz  This is the description of --bar-baz option. This ")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "                    option takes one parameter.")
	assert.Equal(t, status, ITER_NO_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestMakeHelp_marginsAndIndentExceedLineWidth(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "b",
			Aliases:  []string{"bar-baz"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{MarginLeft: 50, MarginRight: 50, Indent: 10}

	_, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "MarginsAndIndentExceedLineWidth{"+
		"LineWidth:80,MarginLeft:50,MarginRight:50,Indent:10}")
	switch err.(type) {
	case MarginsAndIndentExceedLineWidth:
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).LineWidth, 80)
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).MarginLeft, 50)
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).MarginRight, 50)
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).Indent, 10)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestMakeHelp_useAtParam(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name:     "d",
			Aliases:  []string{"data"},
			HasParam: true,
			Desc:     "This is the description of --data option.",
			AtParam:  "<data>",
		},
		OptCfg{
			Name:     "f",
			Aliases:  []string{"fail"},
			HasParam: false,
			Desc:     "This is the description of --fail option.",
		},
		OptCfg{
			Name:     "o",
			Aliases:  []string{"output"},
			HasParam: true,
			Desc:     "This is the description of --output option.",
			AtParam:  "<file>",
		},
		OptCfg{
			Name:     "s",
			Aliases:  []string{"silent"},
			HasParam: false,
			Desc:     "This is the description of --silent option.",
		},
	}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)

	line, status := iter.Next()
	assert.Equal(t, line, usage[0:79])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, usage[79:90])
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "-d, --data <data>    This is the description of --data option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "-f, --fail           This is the description of --fail option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "-o, --output <file>  This is the description of --output option.")
	assert.Equal(t, status, ITER_HAS_MORE)

	line, status = iter.Next()
	assert.Equal(t, line, "-s, --silent         This is the description of --silent option.")
	assert.Equal(t, status, ITER_NO_MORE)

	iter, _ = MakeHelp(usage, optCfgs, wrapOpts)
	for {
		line, status = iter.Next()
		fmt.Println(line)
		if status == ITER_NO_MORE {
			break
		}
	}
}

func TestHelpIter_textsIsEmpty(t *testing.T) {
	iter := newHelpIter([]string{}, 0, 0, 0)
	line, status := iter.Next()
	assert.Equal(t, line, "")
	assert.Equal(t, status, ITER_NO_MORE)
}

func TestPrintHelp(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "b",
			Aliases:  []string{"bar-baz"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{MarginLeft: 5, MarginRight: 5, Indent: 10}

	err := PrintHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)
}

func TestPrintHelp_error(t *testing.T) {
	usage := longUsage
	optCfgs := []OptCfg{
		OptCfg{
			Name: "foo",
			Desc: "This is the description of --foo option.",
		},
		OptCfg{
			Name:     "b",
			Aliases:  []string{"bar-baz"},
			HasParam: true,
			Desc:     "This is the description of --bar-baz option. This option takes one parameter.",
		},
	}
	wrapOpts := WrapOpts{MarginLeft: 50, MarginRight: 50, Indent: 10}

	err := PrintHelp(usage, optCfgs, wrapOpts)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "MarginsAndIndentExceedLineWidth{"+
		"LineWidth:80,MarginLeft:50,MarginRight:50,Indent:10}")
	switch err.(type) {
	case MarginsAndIndentExceedLineWidth:
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).LineWidth, 80)
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).MarginLeft, 50)
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).MarginRight, 50)
		assert.Equal(t, err.(MarginsAndIndentExceedLineWidth).Indent, 10)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestMakeHelp_optCfgsIncludesAnyOption(t *testing.T) {
	usage := "abcdefg"
	optCfgs := []OptCfg{
		OptCfg{Name: "foo", Desc: "description of foo."},
		OptCfg{Name: "bar-baz", Desc: "description of bar-baz."},
		OptCfg{Name: "*", Desc: "description of *."},
		OptCfg{Name: "qux", Desc: "description of qux."},
	}
	wrapOpts := WrapOpts{}

	iter, err := MakeHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)
	text, status := iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, text, usage)
	text, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, text, "--foo      description of foo.")
	text, status = iter.Next()
	assert.Equal(t, status, ITER_HAS_MORE)
	assert.Equal(t, text, "--bar-baz  description of bar-baz.")
	text, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, text, "--qux      description of qux.")
	text, status = iter.Next()
	assert.Equal(t, status, ITER_NO_MORE)
	assert.Equal(t, text, "")
}

func TestPrintHelp_optCfgsIncludesAnyOption(t *testing.T) {
	usage := "abcdefg"
	optCfgs := []OptCfg{
		OptCfg{Name: "foo", Desc: "description of foo."},
		OptCfg{Name: "bar-baz", Desc: "description of bar-baz."},
		OptCfg{Name: "*", Desc: "description of *."},
		OptCfg{Name: "qux", Desc: "description of qux."},
	}
	wrapOpts := WrapOpts{MarginLeft: 5}

	err := PrintHelp(usage, optCfgs, wrapOpts)
	assert.Nil(t, err)
}
