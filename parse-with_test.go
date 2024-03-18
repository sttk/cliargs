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

func TestParseWith_zeroCfgAndOneLongOpt(t *testing.T) {
	optCfgs := []cliargs.OptCfg{}

	osArgs := []string{"path/to/app", "--foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Name:foo-bar}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Name, "foo-bar")
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
	assert.Equal(t, err.Error(), "UnconfiguredOption{Name:f}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Name, "f")
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
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
		cliargs.OptCfg{Names: []string{"f"}},
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	osArgs := []string{"app", "--boo-far"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Name:boo-far}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Name, "boo-far")
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
		cliargs.OptCfg{Names: []string{"f"}},
	}

	osArgs := []string{"app", "-b"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Name:b}")
	switch e := err.(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, e.Name, "b")
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
		cliargs.OptCfg{Names: []string{"*"}},
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
		cliargs.OptCfg{Names: []string{"f"}},
		cliargs.OptCfg{Names: []string{"*"}},
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
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true},
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
		cliargs.OptCfg{Names: []string{"f"}, HasArg: true},
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
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true},
	}

	osArgs := []string{"app", "--foo-bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionNeedsArg{Name:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionNeedsArg:
		assert.Equal(t, e.Name, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
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
		cliargs.OptCfg{Names: []string{"f"}, HasArg: true},
	}

	osArgs := []string{"app", "-f"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionNeedsArg{Name:f,StoreKey:f}")
	switch e := err.(type) {
	case cliargs.OptionNeedsArg:
		assert.Equal(t, e.Name, "f")
		assert.Equal(t, e.StoreKey, "f")
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
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
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Name:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Name, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
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
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Name:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Name, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
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
		cliargs.OptCfg{Names: []string{"f"}},
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
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Name:f,StoreKey:f}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Name, "f")
		assert.Equal(t, e.StoreKey, "f")
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
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Name:f,StoreKey:f}")
	switch e := err.(type) {
	case cliargs.OptionTakesNoArg:
		assert.Equal(t, e.Name, "f")
		assert.Equal(t, e.StoreKey, "f")
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
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: false, IsArray: true},
	}

	osArgs := []string{"app", "--foo-bar", "ABC"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ConfigIsArrayButHasNoArg{StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.ConfigIsArrayButHasNoArg:
		assert.Equal(t, e.StoreKey, "foo-bar")
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
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true, IsArray: true},
		cliargs.OptCfg{Names: []string{"f"}, HasArg: true, IsArray: true},
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
			Names:  []string{"foo-bar", "f", "b"},
			HasArg: true,
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

func TestParseWith_hasOptsOfBothNameAndAliase(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Names:   []string{"foo-bar", "f"},
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
			Names:   []string{"foo-bar", "f"},
			HasArg:  true,
			IsArray: false,
		},
	}

	osArgs := []string{"app", "-f", "ABC", "--foo-bar=DEF"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionIsNotArray{Name:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionIsNotArray:
		assert.Equal(t, e.Name, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
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

	osArgs = []string{"app", "-f=1", "ABC", "--foo-bar", "2"}

	cmd, err = cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionIsNotArray{Name:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.OptionIsNotArray:
		assert.Equal(t, e.Name, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
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
	assert.Equal(t, cmd.OptArg("foo-bar"), "1")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"1"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
	assert.Equal(t, cmd.Args(), []string{"ABC"})
}

func TestParseWith_specifyDefault(t *testing.T) {
	osArgs := []string{"app"}
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"bar"}, HasArg: true, Defaults: []string{"A"}},
		cliargs.OptCfg{Names: []string{"baz"}, HasArg: true, IsArray: true, Defaults: []string{"A"}},
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
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: false, Defaults: []string{"A"}},
	}

	osArgs := []string{"app"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ConfigHasDefaultsButHasNoArg{StoreKey:foo-bar}")
	switch e := err.(type) {
	case cliargs.ConfigHasDefaultsButHasNoArg:
		assert.Equal(t, e.StoreKey, "foo-bar")
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
		cliargs.OptCfg{Names: []string{"foo-bar"}},
		cliargs.OptCfg{
			Names:   []string{"baz", "z"},
			HasArg:  true,
			IsArray: true,
		},
		cliargs.OptCfg{Names: []string{"corge"}, HasArg: true, Defaults: []string{"99"}},
		cliargs.OptCfg{Names: []string{"*"}},
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
		cliargs.OptCfg{Names: []string{"foo", "f"}},
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
		assert.Equal(t, e.Name, "e")
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
		cliargs.OptCfg{Names: []string{"e"}},
		cliargs.OptCfg{Names: []string{"foo", "f"}},
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
		assert.Equal(t, e.Name, "f")
		assert.Equal(t, e.StoreKey, "foo")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "f")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_parseAllArgsEvenIfLongOptionValueIsError(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"e"}},
		cliargs.OptCfg{Names: []string{"foo", "f"}},
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
		assert.Equal(t, e.Name, "foo")
		assert.Equal(t, e.StoreKey, "foo")
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

func TestParseWith_ignoreCfgifNamesIsEmpty(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{}},
		cliargs.OptCfg{
			Names: []string{"foo"},
		},
	}

	osArgs := []string{"app", "--foo", "bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")
	assert.Nil(t, err)

	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.Args(), []string{"bar"})
}

func TestParseWith_optionNameIsDuplicated(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo", "f"}},
		cliargs.OptCfg{Names: []string{"bar", "f"}},
	}

	osArgs := []string{"app", "--foo", "--bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")

	assert.Equal(t, err.Error(), "OptionNameIsDuplicated{Name:f,StoreKey:bar}")
	switch e := err.(type) {
	case cliargs.OptionNameIsDuplicated:
		assert.Equal(t, e.Name, "f")
		assert.Equal(t, e.StoreKey, "bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "f")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_useStoreKey(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "FooBar",
			Names:    []string{"f", "foo"},
		},
	}

	osArgs := []string{"app", "--foo", "bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")
	assert.Nil(t, err)

	assert.True(t, cmd.HasOpt("FooBar"))
	assert.False(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.Args(), []string{"bar"})
}

func TestParseWith_StoreKeyIsDuplicated(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "FooBar",
			Names:    []string{"f", "foo"},
		},
		cliargs.OptCfg{
			StoreKey: "FooBar",
			Names:    []string{"b", "bar"},
		},
	}

	osArgs := []string{"app", "--foo", "bar"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, err.Error(), "StoreKeyIsDuplicated{StoreKey:FooBar}")
	switch e := err.(type) {
	case cliargs.StoreKeyIsDuplicated:
		assert.Equal(t, e.StoreKey, "FooBar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case cliargs.InvalidOption:
		assert.Equal(t, e.GetOpt(), "FooBar")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_acceptAllOptionsIfStoreKeyIsAterisk(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{StoreKey: "*"},
	}

	osArgs := []string{"app", "--foo", "--bar", "baz"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")
	assert.Nil(t, err)

	assert.True(t, cmd.HasOpt("foo"))
	assert.True(t, cmd.HasOpt("bar"))
	assert.Equal(t, cmd.Args(), []string{"baz"})
}

func TestParseWith_acceptUnconfiguredOptionEvenIfItMatchesStoreKey(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "Bar",
			Names:    []string{"foo", "f"},
			HasArg:   true,
			IsArray:  true,
		},
		cliargs.OptCfg{
			StoreKey: "*",
		},
	}

	osArgs := []string{"app", "--foo", "1", "-f=2", "--Bar=3"}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	assert.Equal(t, cmd.Name, "app")
	assert.Nil(t, err)

	assert.True(t, cmd.HasOpt("Bar"))
	assert.Equal(t, cmd.OptArg("Bar"), "1")
	assert.Equal(t, cmd.OptArgs("Bar"), []string{"1", "2", "3"})
	assert.Equal(t, cmd.Args(), []string{})
}
