package cliargs_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/cliargs"
	"os"
	"testing"
)

var osArgs []string = os.Args

func resetOsArgs() {
	os.Args = osArgs
}

func TestParse_zeroArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 1)
	os.Args[0] = osArgs[0]

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_oneNonOptArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "abcd"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{"abcd"})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_oneLongOpt(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "--silent"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.True(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string{})
}

func TestParse_oneLongOptWithParam(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "--alphabet=ABC"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.True(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "ABC")
	assert.Equal(t, args.OptParams("alphabet"), []string{"ABC"})
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_oneShortOpt(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-s"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_oneShortOptWithParam(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-a=123"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "123")
	assert.Equal(t, args.OptParams("a"), []string{"123"})
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_oneArgByMultipleShortOpts(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-sa"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, len(args.CmdParams()), 0)
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string{})
	assert.False(t, args.HasOpt("alphabet"))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string{})
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_oneArgByMultipleShortOptsWithParam(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-sa=123"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "123")
	assert.Equal(t, args.OptParams("a"), []string{"123"})
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_longOptNameIncludesHyphenMark(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "--aaa-bbb-ccc=123"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, len(args.CmdParams()), 0)
	assert.True(t, args.HasOpt("aaa-bbb-ccc"))
	assert.Equal(t, args.OptParam("aaa-bbb-ccc"), "123")
	assert.Equal(t, args.OptParams("aaa-bbb-ccc"), []string{"123"})
}

func TestParse_optParamsIncludesEqualMark(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-sa=b=c"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "b=c")
	assert.Equal(t, args.OptParams("a"), []string{"b=c"})
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_optParamsIncludesMarks(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-sa=1,2-3"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{})
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "1,2-3")
	assert.Equal(t, args.OptParams("a"), []string{"1,2-3"})
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_illegalLongOptIfIncludingInvalidChar(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 4)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-s"
	os.Args[2] = "--abc%def"
	os.Args[3] = "-a"

	args, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:abc%def}")
	switch err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, err.(cliargs.OptionHasInvalidChar).Option, "abc%def")
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_illegalLongOptIfFirstCharIsNumber(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "--1abc"

	args, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:1abc}")
	switch err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, err.(cliargs.OptionHasInvalidChar).Option, "1abc")
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_illegalLongOptIfFirstCharIsHyphen(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "---aaa=123"

	args, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:-aaa=123}")
	switch err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, err.(cliargs.OptionHasInvalidChar).Option, "-aaa=123")
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_IllegalCharInShortOpt(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 4)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-s"
	os.Args[2] = "--alphabet"
	os.Args[3] = "-s@"

	args, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:@}")
	switch err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, err.(cliargs.OptionHasInvalidChar).Option, "@")
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, args.CmdParams(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_useEndOptMark(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 7)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-s"
	os.Args[2] = "--"
	os.Args[3] = "-s"
	os.Args[4] = "--"
	os.Args[5] = "-s@"
	os.Args[6] = "xxx"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{"-s", "--", "-s@", "xxx"})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_singleHyphen(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = osArgs[0]
	os.Args[1] = "-"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.CmdParams(), []string{"-"})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptParam("alphabet"), "")
	assert.Equal(t, args.OptParams("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptParam("s"), "")
	assert.Equal(t, args.OptParams("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptParam("silent"), "")
	assert.Equal(t, args.OptParams("silent"), []string(nil))
}

func TestParse_multipleArgs(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 8)
	os.Args[0] = osArgs[0]
	os.Args[1] = "--foo-bar"
	os.Args[2] = "-a"
	os.Args[3] = "--baz"
	os.Args[4] = "-bc=3"
	os.Args[5] = "qux"
	os.Args[6] = "-c=4"
	os.Args[7] = "quux"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptParam("a"), "")
	assert.Equal(t, args.OptParams("a"), []string{})
	assert.True(t, args.HasOpt("b"))
	assert.Equal(t, args.OptParam("b"), "")
	assert.Equal(t, args.OptParams("b"), []string{})
	assert.True(t, args.HasOpt("c"))
	assert.Equal(t, args.OptParam("c"), "3")
	assert.Equal(t, args.OptParams("c"), []string{"3", "4"})
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string{})
	assert.True(t, args.HasOpt("baz"))
	assert.Equal(t, args.OptParam("baz"), "")
	assert.Equal(t, args.OptParams("baz"), []string{})
	assert.Equal(t, args.CmdParams(), []string{"qux", "quux"})
}
