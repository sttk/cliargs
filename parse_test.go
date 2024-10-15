package cliargs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/cliargs"
	"github.com/sttk/cliargs/errors"
)

var origOsArgs []string = os.Args

func reset() {
	os.Args = origOsArgs
}

func TestParse_zeroArg(t *testing.T) {
	defer reset()

	os.Args = make([]string, 1)
	os.Args[0] = "/path/to/app"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "abcd"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abcd"})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "--silent"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "--alphabet=ABC"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-s"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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

func TestParse_oneShortOptWithArg(t *testing.T) {
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-a=123"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("a"))
	assert.Equal(t, cmd.OptArg("a"), "123")
	assert.Equal(t, cmd.OptArgs("a"), []string{"123"})
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

func TestParse_oneArgByMultipleShortOpts(t *testing.T) {
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "-sa"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, len(cmd.Args), 0)
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
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "/path/to/app"
	os.Args[1] = "-sa=123"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "--aaa-bbb-ccc=123"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, len(cmd.Args), 0)
	assert.True(t, cmd.HasOpt("aaa-bbb-ccc"))
	assert.Equal(t, cmd.OptArg("aaa-bbb-ccc"), "123")
	assert.Equal(t, cmd.OptArgs("aaa-bbb-ccc"), []string{"123"})
}

func TestParse_optsIncludesEqualMark(t *testing.T) {
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "-sa=b=c"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-sa=1,2-3"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 4)
	os.Args[0] = "app"
	os.Args[1] = "-s"
	os.Args[2] = "--abc%def"
	os.Args[3] = "-a"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionContainsInvalidChar{Option:abc%def}")
	switch e := err.(type) {
	case errors.OptionContainsInvalidChar:
		assert.Equal(t, e.Option, "abc%def")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "abc%def")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "--1abc"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionContainsInvalidChar{Option:1abc}")
	switch e := err.(type) {
	case errors.OptionContainsInvalidChar:
		assert.Equal(t, e.Option, "1abc")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "1abc")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "app"
	os.Args[1] = "---aaa=123"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionContainsInvalidChar{Option:-aaa=123}")
	switch e := err.(type) {
	case errors.OptionContainsInvalidChar:
		assert.Equal(t, e.Option, "-aaa=123")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "-aaa=123")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 4)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-s"
	os.Args[2] = "--alphabet"
	os.Args[3] = "-s@"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionContainsInvalidChar{Option:@}")
	switch e := err.(type) {
	case errors.OptionContainsInvalidChar:
		assert.Equal(t, e.Option, "@")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "@")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	os.Args = make([]string, 7)
	os.Args[0] = "app"
	os.Args[1] = "-s"
	os.Args[2] = "--"
	os.Args[3] = "-s"
	os.Args[4] = "--"
	os.Args[5] = "-s@"
	os.Args[6] = "xxx"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"-s", "--", "-s@", "xxx"})
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
	defer reset()

	os.Args = make([]string, 2)
	os.Args[0] = "path/to/app"
	os.Args[1] = "-"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"-"})
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
	defer reset()

	os.Args = make([]string, 8)
	os.Args[0] = "app"
	os.Args[1] = "--foo-bar"
	os.Args[2] = "-a"
	os.Args[3] = "--baz"
	os.Args[4] = "-bc=3"
	os.Args[5] = "qux"
	os.Args[6] = "-c=4"
	os.Args[7] = "quux"

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"qux", "quux"})
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
}

func TestParse_parseAllArgsEvenIfError(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "--1", "-b2ar", "--3", "baz"}

	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	switch e := err.(type) {
	case errors.OptionContainsInvalidChar:
		assert.Equal(t, e.Option, "1")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "1")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"baz"})
	assert.True(t, cmd.HasOpt("foo"))
	assert.True(t, cmd.HasOpt("b"))
	assert.True(t, cmd.HasOpt("a"))
	assert.True(t, cmd.HasOpt("r"))
	assert.False(t, cmd.HasOpt("1"))
	assert.False(t, cmd.HasOpt("2"))
	assert.False(t, cmd.HasOpt("3"))
}

func TestParseUntilSubCmd_zeroArg(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "sub"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})

	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{})

	err = subCmd.Parse()
	assert.Equal(t, err, nil)

	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{})
}

func TestParseUntilSubCmd_oneOptArg(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "sub", "--bar=123"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.False(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.OptArg("bar"), "")
	assert.Equal(t, cmd.OptArgs("bar"), []string(nil))

	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))

	err = subCmd.Parse()

	assert.Equal(t, err, nil)
	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.True(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "123")
	assert.Equal(t, subCmd.OptArgs("bar"), []string{"123"})
}

func TestParseUntilSubCmd_oneCmdArg(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "sub", "bar"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.False(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.OptArg("bar"), "")
	assert.Equal(t, cmd.OptArgs("bar"), []string(nil))

	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))

	err = subCmd.Parse()

	assert.Equal(t, err, nil)
	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{"bar"})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))
}

func TestParseUntilSubCmd_noSubCmd(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "--bar=123"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.True(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.OptArg("bar"), "123")
	assert.Equal(t, cmd.OptArgs("bar"), []string{"123"})

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))

	err = subCmd.Parse()

	assert.Equal(t, err, nil)
	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))
}

func TestParseUntilSubCmd_subCmdIssingleHyphen(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "-", "--bar=123"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.False(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.OptArg("bar"), "")
	assert.Equal(t, cmd.OptArgs("bar"), []string(nil))

	assert.Equal(t, subCmd.Name, "-")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))

	err = subCmd.Parse()

	assert.Equal(t, err, nil)
	assert.Equal(t, subCmd.Name, "-")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.True(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "123")
	assert.Equal(t, subCmd.OptArgs("bar"), []string{"123"})
}

func TestParseUntilSubCmd_withEndOptMark(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "sub", "--", "--bar", "--baz=123", "-@"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.False(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.OptArg("bar"), "")
	assert.Equal(t, cmd.OptArgs("bar"), []string(nil))
	assert.False(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("baz"), "")
	assert.Equal(t, cmd.OptArgs("baz"), []string(nil))

	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))
	assert.False(t, subCmd.HasOpt("baz"))
	assert.Equal(t, subCmd.OptArg("baz"), "")
	assert.Equal(t, subCmd.OptArgs("baz"), []string(nil))

	err = subCmd.Parse()

	assert.Equal(t, err, nil)
	assert.Equal(t, subCmd.Name, "sub")
	assert.Equal(t, subCmd.Args, []string{"--bar", "--baz=123", "-@"})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))
	assert.False(t, subCmd.HasOpt("baz"))
	assert.Equal(t, subCmd.OptArg("baz"), "")
	assert.Equal(t, subCmd.OptArgs("baz"), []string(nil))
}

func TestParseUntilSubCmd_afterEndOptMark(t *testing.T) {
	defer reset()

	os.Args = []string{"/path/to/app", "--foo", "--", "--bar", "--baz=123", "-@"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()

	assert.Equal(t, err, nil)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.False(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.OptArg("bar"), "")
	assert.Equal(t, cmd.OptArgs("bar"), []string(nil))
	assert.False(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("baz"), "")
	assert.Equal(t, cmd.OptArgs("baz"), []string(nil))

	assert.Equal(t, subCmd.Name, "--bar")
	assert.Equal(t, subCmd.Args, []string{})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))
	assert.False(t, subCmd.HasOpt("baz"))
	assert.Equal(t, subCmd.OptArg("baz"), "")
	assert.Equal(t, subCmd.OptArgs("baz"), []string(nil))

	err = subCmd.Parse()

	assert.Equal(t, err, nil)
	assert.Equal(t, subCmd.Name, "--bar")
	assert.Equal(t, subCmd.Args, []string{"--baz=123", "-@"})

	assert.False(t, subCmd.HasOpt("foo"))
	assert.Equal(t, subCmd.OptArg("foo"), "")
	assert.Equal(t, subCmd.OptArgs("foo"), []string(nil))
	assert.False(t, subCmd.HasOpt("bar"))
	assert.Equal(t, subCmd.OptArg("bar"), "")
	assert.Equal(t, subCmd.OptArgs("bar"), []string(nil))
	assert.False(t, subCmd.HasOpt("baz"))
	assert.Equal(t, subCmd.OptArg("baz"), "")
	assert.Equal(t, subCmd.OptArgs("baz"), []string(nil))
}
