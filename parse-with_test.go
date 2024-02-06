package cliargs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/cliargs"
)

func TestParseWith_zeroCfgAndZeroArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{}

	osArgs := []string{"app"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_zeroCfgAndOneCommandArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{}

	osArgs := []string{"/path/to/app", "foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{"foo-bar"})
}

func testParseWith_zeroCfgAndOneLongOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{}

	osArgs := []string{"path/to/app", "--foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_zeroCfgAndOneShortOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{}

	osArgs := []string{"path/to/app", "-f"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:f}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Option, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgAndZeroOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"app"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgAndOneCmdArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"path/to/app", "foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{"foo-bar"})
}

func TestParseWith_oneCfgAndOneLongOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"path/to/app", "--foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgAndOneShortOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "f"},
	}

	osArgs := []string{"app", "-f"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{})
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgAndOneDifferentLongOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"app", "--boo-far"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:boo-far}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Option, "boo-far")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "boo-far")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgAndOneDifferentShortOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "f"},
	}

	osArgs := []string{"app", "-b"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:b}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Option, "b")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "b")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_anyOptCfgAndOneDifferentLongOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
		cliargs.OptCfg{Name: "*"},
	}

	osArgs := []string{"app", "--boo-far"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	assert.True(t, cmd.HasOpt("boo-far"))
	assert.Equal(t, cmd.OptArg("boo-far"), "")
	assert.Equal(t, cmd.OptArgs("boo-far"), []string{})
}

func TestParseWith_anyOptCfgAndOneDifferentShortOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "f"},
		cliargs.OptCfg{Name: "*"},
	}

	osArgs := []string{"app", "-b"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	assert.True(t, cmd.HasOpt("b"))
	assert.Equal(t, cmd.OptArg("b"), "")
	assert.Equal(t, cmd.OptArgs("b"), []string{})
}

func TestParseWith_oneCfgHasArgAndOneLongOptHasArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar", HasArg: true},
	}

	osArgs := []string{"app", "--foo-bar", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "--foo-bar=ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "--foo-bar", ""}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{""})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "--foo-bar="}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{""})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasArgAndOneShortOptHasArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "f", HasArg: true},
	}

	osArgs := []string{"app", "-f", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC"})
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f=ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC"})
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f", ""}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{""})
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f="}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{""})
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasArgButOneLongOptHasNoArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar", HasArg: true},
	}

	osArgs := []string{"app", "--foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionNeedsArg{Option:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionNeedsArg:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasArgAndOneShortOptHasNoArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "f", HasArg: true},
	}

	osArgs := []string{"app", "-f"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionNeedsArg{Option:f}")
	switch e := err.(type) {
	case cliargs.OptionNeedsArg:
		assert.Equal(t, e.Option, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasNoArgAndOneLongOptHasArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"app", "--foo-bar", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{"ABC"})

	osArgs = []string{"app", "--foo-bar=ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "--foo-bar", ""}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{""})

	osArgs = []string{"app", "--foo-bar="}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasNoArgAndOneShortOptHasArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "f"},
	}

	osArgs := []string{"app", "-f", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{})
	assert.Equal(t, cmd.Args(), []string{"ABC"})

	osArgs = []string{"app", "-f=ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:f}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Option, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f", ""}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{})
	assert.Equal(t, cmd.Args(), []string{""})

	osArgs = []string{"app", "-f="}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:f}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Option, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasNoArgButIsArray(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar", HasArg: false, IsArray: true},
	}

	osArgs := []string{"app", "--foo-bar", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ConfigIsArrayButHasNoArg{Option:foo-bar}")
	switch e := err.(type) {
	case cliargs.ConfigIsArrayButHasNoArg:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgIsArrayAndOptHasOneArg(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar", HasArg: true, IsArray: true},
		cliargs.OptCfg{Name: "f", HasArg: true, IsArray: true},
	}

	osArgs := []string{"app", "--foo-bar", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "--foo-bar", "ABC", "--foo-bar=DEF"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f", "ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC"})
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f", "ABC", "-f=DEF"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC", "DEF"})
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasAliasesAndArgMatchesName(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f", "b"},
			HasArg:  true,
		},
	}

	osArgs := []string{"app", "--foo-bar", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "--foo-bar=ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgHasAliasesAndArgMatchesAliases(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
		},
	}

	osArgs := []string{"app", "-f", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f=ABC"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_combineOptsByNameAndAliases(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			IsArray: true,
		},
	}

	osArgs := []string{"app", "-f", "ABC", "--foo-bar=DEF"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})

	osArgs = []string{"app", "-f=ABC", "--foo-bar", "DEF"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_oneCfgIsNotArrayButOptsAreMultiple(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo-bar",
			Aliases: []string{"f"},
			HasArg:  true,
			IsArray: false,
		},
	}

	osArgs := []string{"app", "-f", "ABC", "--foo-bar=DEF"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionIsNotArray{Option:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionIsNotArray:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_specifyDefault(t *testing.T) {
	osArgs := []string{"app"}
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "bar", HasArg: true, Defaults: []string{"A"}},
		cliargs.OptCfg{Name: "baz", HasArg: true, IsArray: true, Defaults: []string{"A"}},
	}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo"))
	assert.True(t, cmd.HasOpt("bar"))
	assert.True(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArg("bar"), "A")
	assert.Equal(t, cmd.OptArg("baz"), "A")
	assert.Equal(t, cmd.OptArgs("foo"), []string(nil))
	assert.Equal(t, cmd.OptArgs("bar"), []string{"A"})
	assert.Equal(t, cmd.OptArgs("baz"), []string{"A"})
}

func TestParseWith_oneCfgHasNoArgButHasDefault(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar", HasArg: false, Defaults: []string{"A"}},
	}

	osArgs := []string{"app"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ConfigHasDefaultsButHasNoArg{Option:foo-bar}")
	switch e := err.(type) {
	case cliargs.ConfigHasDefaultsButHasNoArg:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{})
}

func TestParseWith_multipleArgs(t *testing.T) {
	osArgs := []string{"app", "--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux"}

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo-bar"},
		cliargs.OptCfg{
			Name:    "baz",
			Aliases: []string{"z"},
			HasArg:  true,
			IsArray: true,
		},
		cliargs.OptCfg{Name: "corge", HasArg: true, Defaults: []string{"99"}},
		cliargs.OptCfg{Name: "*"},
	}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.True(t, cmd.HasOpt("baz"))
	assert.True(t, cmd.HasOpt("X"))
	assert.True(t, cmd.HasOpt("corge"))
	assert.Equal(t, cmd.OptArg("baz"), "1")
	assert.Equal(t, cmd.OptArgs("baz"), []string{"1", "2"})
	assert.Equal(t, cmd.OptArg("corge"), "99")
	assert.Equal(t, cmd.OptArgs("corge"), []string{"99"})
	assert.Equal(t, cmd.Args(), []string{"qux", "quux"})
}

func TestParseWith_parseAllArgsEvenIfError(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "foo", Aliases: []string{"f"}},
	}

	osArgs := []string{"app", "-ef", "bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")

	assert.True(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.False(t, cmd.HasOpt("e"))
	assert.Equal(t, cmd.Args(), []string{"bar"})

	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Option, "e")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "e")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_parseAllArgsEvenIfShortOptionValueIsError(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "e"},
		cliargs.OptCfg{Name: "foo", Aliases: []string{"f"}},
	}

	osArgs := []string{"app", "-ef=123", "bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")

	assert.False(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.True(t, cmd.HasOpt("e"))
	assert.Equal(t, cmd.Args(), []string{"bar"})

	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_parseAllArgsEvenIfLongOptionValueIsError(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Name: "e"},
		cliargs.OptCfg{Name: "foo", Aliases: []string{"f"}},
	}

	osArgs := []string{"app", "--foo=123", "-e", "bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")

	assert.False(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.True(t, cmd.HasOpt("e"))
	assert.Equal(t, cmd.Args(), []string{"bar"})

	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "foo")
	default:
		assert.Fail(t, err.Error())
	}
}
