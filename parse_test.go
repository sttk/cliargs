package cliargs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/cliargs"
)

var osArgs []string = os.Args

func resetOsArgs() {
	os.Args = osArgs
}

func TestParse_zeroArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 1)
	os.Args[0] = "/path/to/app"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_oneNonOptArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "abcd"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{"abcd"})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_oneLongOpt(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "--silent"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.True(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string{})
}

func TestParse_oneLongOptWithArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "--alphabet=ABC"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.True(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "ABC")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_oneShortOpt(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-s"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.Name, "app")
	assert.Equal(t, args.Args(), []string{})
	assert.False(t, args.HasOpt("a"))
	assert.Equal(t, args.OptArg("a"), "")
	assert.Equal(t, args.OptArgs("a"), []string(nil))
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptArg("alphabet"), "")
	assert.Equal(t, args.OptArgs("alphabet"), []string(nil))
	assert.True(t, args.HasOpt("s"))
	assert.Equal(t, args.OptArg("s"), "")
	assert.Equal(t, args.OptArgs("s"), []string{})
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptArg("silent"), "")
	assert.Equal(t, args.OptArgs("silent"), []string(nil))
}

func TestParse_oneShortOptWithArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-a=123"

	args, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, args.Name, "app")
	assert.Equal(t, args.Args(), []string{})
	assert.True(t, args.HasOpt("a"))
	assert.Equal(t, args.OptArg("a"), "123")
	assert.Equal(t, args.OptArgs("a"), []string{"123"})
	assert.False(t, args.HasOpt("alphabet"))
	assert.Equal(t, args.OptArg("alphabet"), "")
	assert.Equal(t, args.OptArgs("alphabet"), []string(nil))
	assert.False(t, args.HasOpt("s"))
	assert.Equal(t, args.OptArg("s"), "")
	assert.Equal(t, args.OptArgs("s"), []string(nil))
	assert.False(t, args.HasOpt("silent"))
	assert.Equal(t, args.OptArg("silent"), "")
	assert.Equal(t, args.OptArgs("silent"), []string(nil))
}

func TestParse_oneArgByMultipleShortOpts(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "-sa"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, len(cmd.Args()), 0)
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string{})
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string{})
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_oneArgByMultipleShortOptsWithArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "/path/to/app"
	os.Args[1] = "-sa=123"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "123")
	assert.Equal(t, cmd.OptArgs("a"), []string{"123"})
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_longOptNameIncludesHyphenMark(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "--aaa-bbb-ccc=123"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, len(cmd.Args()), 0)
	assert.True(t, cmd.HasOpt("aaa-bbb-ccc"))
	assert.Equal(t, cmd.OptArg("aaa-bbb-ccc"), "123")
	assert.Equal(t, cmd.OptArgs("aaa-bbb-ccc"), []string{"123"})
}

func TestParse_optsIncludesEqualMark(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "-sa=b=c"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "b=c")
	assert.Equal(t, cmd.OptArgs("a"), []string{"b=c"})
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_optsIncludesMarks(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-sa=1,2-3"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "1,2-3")
	assert.Equal(t, cmd.OptArgs("a"), []string{"1,2-3"})
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_illegalLongOptIfIncludingInvalidChar(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 4)
	os.Args[0] = "app"
	os.Args[1] = "-s"
	os.Args[2] = "--abc%def"
	os.Args[3] = "-a"

	cmd, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:abc%def}")
	switch e := err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, e.Option, "abc%def")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "abc%def")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string{})
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_illegalLongOptIfFirstCharIsNumber(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "--1abc"

	cmd, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:1abc}")
	switch e := err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, e.Option, "1abc")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "1abc")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_illegalLongOptIfFirstCharIsHyphen(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "---aaa=123"

	cmd, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:-aaa=123}")
	switch e := err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, e.Option, "-aaa=123")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "-aaa=123")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_IllegalCharInShortOpt(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 4)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-s"
	os.Args[2] = "--alphabet"
	os.Args[3] = "-s@"

	cmd, err := cliargs.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionHasInvalidChar{Option:@}")
	switch e := err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, e.Option, "@")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "@")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.True(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string{})
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_useEndOptMark(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 7)
	os.Args[0] = "app"
	os.Args[1] = "-s"
	os.Args[2] = "--"
	os.Args[3] = "-s"
	os.Args[4] = "--"
	os.Args[5] = "-s@"
	os.Args[6] = "xxx"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{"-s", "--", "-s@", "xxx"})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.True(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string{})
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_singleHyphen(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{"-"})
	assert.False(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string(nil))
	assert.False(t, cmd.HasOpt("alphabet"))
	assert.Equal(t, cmd.OptArg("alphabet"), "")
	assert.Equal(t, cmd.OptArgs("alphabet"), []string(nil))
	assert.False(t, cmd.HasOpt("s"))
	assert.Equal(t, cmd.OptArg("s"), "")
	assert.Equal(t, cmd.OptArgs("s"), []string(nil))
	assert.False(t, cmd.HasOpt("silent"))
	assert.Equal(t, cmd.OptArg("silent"), "")
	assert.Equal(t, cmd.OptArgs("silent"), []string(nil))
}

func TestParse_multipleArgs(t *testing.T) {
	defer resetOsArgs()

	os.Args = make([]string, 8)
	os.Args[0] = "app"
	os.Args[1] = "--foo-bar"
	os.Args[2] = "-a"
	os.Args[3] = "--baz"
	os.Args[4] = "-bc=3"
	os.Args[5] = "qux"
	os.Args[6] = "-c=4"
	os.Args[7] = "quux"

	cmd, err := cliargs.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "")
	assert.Equal(t, cmd.OptArgs("a"), []string{})
	assert.True(t, cmd.HasOpt("b"))
	assert.Equal(t, cmd.OptArg("b"), "")
	assert.Equal(t, cmd.OptArgs("b"), []string{})
	assert.True(t, cmd.HasOpt("c"))
	assert.Equal(t, cmd.OptArg("c"), "3")
	assert.Equal(t, cmd.OptArgs("c"), []string{"3", "4"})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{})
	assert.True(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("baz"), "")
	assert.Equal(t, cmd.OptArgs("baz"), []string{})
	assert.Equal(t, cmd.Args(), []string{"qux", "quux"})
}

func TestParse_parseAllArgsEvenIfError(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app", "--foo", "--1", "-b2ar", "--3", "baz"}

	cmd, err := cliargs.Parse()

	switch e := err.(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, e.Option, "1")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "1")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo"))
	assert.True(t, cmd.HasOpt("b"))
	assert.True(t, cmd.HasOpt("a"))
	assert.True(t, cmd.HasOpt("r"))
	assert.False(t, cmd.HasOpt("1"))
	assert.False(t, cmd.HasOpt("2"))
	assert.False(t, cmd.HasOpt("3"))
	assert.Equal(t, cmd.Args(), []string{"baz"})
}
