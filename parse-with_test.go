package cliargs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/cliargs"
	"github.com/sttk/cliargs/errors"
	"github.com/sttk/cliargs/validators"
)

func TestParseWith_zeroCfgAndZeroArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
}

func TestParseWith_zeroCfgAndOneCommandArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"/path/to/app", "foo-bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"foo-bar"})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
}

func TestParseWith_zeroCfgAndOneLongOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"path/to/app", "--foo-bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:foo-bar}")

	switch e := err.(type) {
	case errors.UnconfiguredOption:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_zeroCfgAndOneShortOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"path/to/app", "-f"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:f}")

	switch e := err.(type) {
	case errors.UnconfiguredOption:
		assert.Equal(t, e.Option, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgAndZeroOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	os.Args = []string{"app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgAndOneCmdArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	os.Args = []string{"path/to/app", "foo-bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Args, []string{"foo-bar"})
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgAndOneLongOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	os.Args = []string{"path/to/app", "--foo-bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgAndOneShortOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"f"}},
	}

	os.Args = []string{"app", "-f"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgAndOneDifferentLongOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	os.Args = []string{"app", "--boo-far"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:boo-far}")
	switch e := err.(type) {
	case errors.UnconfiguredOption:
		assert.Equal(t, e.Option, "boo-far")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "boo-far")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgAndOneDifferentShortOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"f"}},
	}

	os.Args = []string{"app", "-b"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:b}")
	switch e := err.(type) {
	case errors.UnconfiguredOption:
		assert.Equal(t, e.Option, "b")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "b")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_anyOptCfgAndOneDifferentLongOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
		cliargs.OptCfg{Names: []string{"*"}},
	}

	os.Args = []string{"app", "--boo-far"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.True(t, cmd.HasOpt("boo-far"))
	assert.Equal(t, cmd.OptArg("boo-far"), "")
	assert.Equal(t, cmd.OptArgs("boo-far"), []string(nil))
}

func TestParseWith_anyOptCfgAndOneDifferentShortOpt(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"f"}},
		cliargs.OptCfg{Names: []string{"*"}},
	}

	os.Args = []string{"app", "-b"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, cmd.Name, "app")
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.True(t, cmd.HasOpt("b"))
	assert.Equal(t, cmd.OptArg("b"), "")
	assert.Equal(t, cmd.OptArgs("b"), []string(nil))
}

func TestParseWith_oneCfgHasArgAndOneLongOptHasArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar=ABC"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar", ""}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{""})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar="}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{""})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgHasArgAndOneShortOptHasArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"f"}, HasArg: true},
	}

	os.Args = []string{"app", "-f", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC"})

	os.Args = []string{"app", "-f=ABC"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC"})

	os.Args = []string{"app", "-f", ""}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{""})

	os.Args = []string{"app", "-f="}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string{""})
}

func TestParseWith_oneCfgHasArgButOneLongOptHasNoArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true},
	}

	os.Args = []string{"app", "--foo-bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionNeedsArg{Option:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case errors.OptionNeedsArg:
		assert.Equal(t, e.Option, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgHasArgAndOneShortOptHasNoArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"f"}, HasArg: true},
	}

	os.Args = []string{"app", "-f"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionNeedsArg{Option:f,StoreKey:f}")
	switch e := err.(type) {
	case errors.OptionNeedsArg:
		assert.Equal(t, e.Option, "f")
		assert.Equal(t, e.StoreKey, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgHasNoArgAndOneLongOptHasArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"ABC"})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar=ABC"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case errors.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar", ""}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{""})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar="}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case errors.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgHasNoArgAndOneShortOptHasArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"f"}},
	}

	os.Args = []string{"app", "-f", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"ABC"})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "-f=ABC"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:f,StoreKey:f}")
	switch e := err.(type) {
	case errors.OptionTakesNoArg:
		assert.Equal(t, e.Option, "f")
		assert.Equal(t, e.StoreKey, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "-f", ""}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{""})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "-f="}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionTakesNoArg{Option:f,StoreKey:f}")
	switch e := err.(type) {
	case errors.OptionTakesNoArg:
		assert.Equal(t, e.Option, "f")
		assert.Equal(t, e.StoreKey, "f")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgHasNoArgButIsArray(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: false, IsArray: true},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ConfigIsArrayButHasNoArg{StoreKey:foo-bar,Name:foo-bar}")
	switch e := err.(type) {
	case errors.ConfigIsArrayButHasNoArg:
		assert.Equal(t, e.StoreKey, "foo-bar")
		assert.Equal(t, e.Name, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgIsArrayAndOptHasOneArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true, IsArray: true},
		cliargs.OptCfg{Names: []string{"f"}, HasArg: true, IsArray: true},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar", "ABC", "--foo-bar=DEF"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "-f", "ABC"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC"})

	os.Args = []string{"app", "-f", "ABC", "-f=DEF"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.True(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "ABC")
	assert.Equal(t, cmd.OptArgs("f"), []string{"ABC", "DEF"})
}

func TestParseWith_oneCfgHasAliasesAndArgMatchesName(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Names:  []string{"foo-bar", "f", "b"},
			HasArg: true,
		},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar=ABC"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_hasOptsOfBothNameAndAliase(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Names:   []string{"foo-bar", "f"},
			HasArg:  true,
			IsArray: true,
		},
	}

	os.Args = []string{"app", "-f", "ABC", "--foo-bar=DEF"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "-f=ABC", "--foo-bar", "DEF"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgIsNotArrayButOptsAreMultiple(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Names:   []string{"foo-bar", "f"},
			HasArg:  true,
			IsArray: false,
		},
	}

	os.Args = []string{"app", "-f", "ABC", "--foo-bar=DEF"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionIsNotArray{Option:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case errors.OptionIsNotArray:
		assert.Equal(t, e.Option, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "-f=1", "ABC", "--foo-bar", "2"}

	cmd = cliargs.NewCmd()
	err = cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionIsNotArray{Option:foo-bar,StoreKey:foo-bar}")
	switch e := err.(type) {
	case errors.OptionIsNotArray:
		assert.Equal(t, e.Option, "foo-bar")
		assert.Equal(t, e.StoreKey, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"ABC"})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "1")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"1"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_specifyDefault(t *testing.T) {
	defer reset()

	os.Args = []string{"app"}

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"bar"}, HasArg: true, Defaults: []string{"A"}},
		cliargs.OptCfg{Names: []string{"baz"}, HasArg: true, IsArray: true, Defaults: []string{"A"}},
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
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
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: false, Defaults: []string{"A"}},
	}

	os.Args = []string{"app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "ConfigHasDefaultsButHasNoArg{StoreKey:foo-bar,Name:foo-bar}")
	switch e := err.(type) {
	case errors.ConfigHasDefaultsButHasNoArg:
		assert.Equal(t, e.StoreKey, "foo-bar")
		assert.Equal(t, e.Name, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_multipleArgs(t *testing.T) {
	defer reset()

	os.Args = []string{"app", "--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux"}

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

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"qux", "quux"})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.True(t, cmd.HasOpt("baz"))
	assert.True(t, cmd.HasOpt("X"))
	assert.True(t, cmd.HasOpt("corge"))
	assert.Equal(t, cmd.OptArg("baz"), "1")
	assert.Equal(t, cmd.OptArgs("baz"), []string{"1", "2"})
	assert.Equal(t, cmd.OptArg("corge"), "99")
	assert.Equal(t, cmd.OptArgs("corge"), []string{"99"})
}

func TestParseWith_parseAllArgsEvenIfError(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo", "f"}},
	}

	os.Args = []string{"app", "-ef", "bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"bar"})
	assert.True(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.False(t, cmd.HasOpt("e"))

	switch e := err.(type) {
	case errors.UnconfiguredOption:
		assert.Equal(t, e.Option, "e")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "e")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_parseAllArgsEvenIfShortOptionValueIsError(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"e"}},
		cliargs.OptCfg{Names: []string{"foo", "f"}},
	}

	os.Args = []string{"app", "-ef=123", "bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"bar"})
	assert.False(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.True(t, cmd.HasOpt("e"))

	switch e := err.(type) {
	case errors.OptionTakesNoArg:
		assert.Equal(t, e.Option, "f")
		assert.Equal(t, e.StoreKey, "foo")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_parseAllArgsEvenIfLongOptionValueIsError(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"e"}},
		cliargs.OptCfg{Names: []string{"foo", "f"}},
	}

	os.Args = []string{"app", "--foo=123", "-e", "bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"bar"})
	assert.False(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
	assert.True(t, cmd.HasOpt("e"))

	switch e := err.(type) {
	case errors.OptionTakesNoArg:
		assert.Equal(t, e.Option, "foo")
		assert.Equal(t, e.StoreKey, "foo")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_ignoreCfgifNamesIsEmpty(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{}},
		cliargs.OptCfg{
			Names: []string{"foo"},
		},
	}

	os.Args = []string{"app", "--foo", "bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"bar"})
	assert.True(t, cmd.HasOpt("foo"))
}

func TestParseWith_optionNameIsDuplicated(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo", "f"}},
		cliargs.OptCfg{Names: []string{"bar", "f"}},
	}

	os.Args = []string{"app", "--foo", "--bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.Equal(t, err.Error(), "OptionNameIsDuplicated{StoreKey:bar,Name:f}")
	switch e := err.(type) {
	case errors.OptionNameIsDuplicated:
		assert.Equal(t, e.Name, "f")
		assert.Equal(t, e.StoreKey, "bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_useStoreKey(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey: "FooBar",
			Names:    []string{"f", "foo"},
		},
	}

	os.Args = []string{"app", "--foo", "bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"bar"})

	assert.True(t, cmd.HasOpt("FooBar"))
	assert.False(t, cmd.HasOpt("foo"))
	assert.False(t, cmd.HasOpt("f"))
}

func TestParseWith_StoreKeyIsDuplicated(t *testing.T) {
	defer reset()

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

	os.Args = []string{"app", "--foo", "bar"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.Equal(t, err.Error(), "StoreKeyIsDuplicated{StoreKey:FooBar,Name:b}")
	switch e := err.(type) {
	case errors.StoreKeyIsDuplicated:
		assert.Equal(t, e.StoreKey, "FooBar")
		assert.Equal(t, e.Name, "b")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "b")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseWith_acceptAllOptionsIfStoreKeyIsAterisk(t *testing.T) {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{StoreKey: "*"},
	}

	os.Args = []string{"app", "--foo", "--bar", "baz"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"baz"})

	assert.True(t, cmd.HasOpt("foo"))
	assert.True(t, cmd.HasOpt("bar"))
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

	os.Args = []string{"app", "--foo", "1", "-f=2", "--Bar=3"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.True(t, cmd.HasOpt("Bar"))
	assert.Equal(t, cmd.OptArg("Bar"), "1")
	assert.Equal(t, cmd.OptArgs("Bar"), []string{"1", "2", "3"})
}

func TestParseWith_oneCfgUsingValidator(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey:  "foo",
			HasArg:    true,
			Validator: &validators.ValidateInt16,
		},
	}

	os.Args = []string{"app", "--foo", "123", "def", "ghi"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"def", "ghi"})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "123")
	assert.Equal(t, cmd.OptArgs("foo"), []string{"123"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))
}

func TestParseWith_oneCfgUsingValidator_validationError(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey:  "Foo",
			Names:     []string{"foo", "f"},
			HasArg:    true,
			Validator: &validators.ValidateInt8,
		},
	}

	os.Args = []string{"app", "-f", "128", "def", "ghi"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	switch e := err.(type) {
	case errors.OptionArgIsInvalid:
		assert.Equal(t, e.StoreKey, "Foo")
		assert.Equal(t, e.Option, "f")
		assert.Equal(t, e.OptArg, "128")
		assert.Equal(t, e.Cause.Error(), "strconv.ParseInt: parsing \"128\": value out of range")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"def", "ghi"})
	assert.False(t, cmd.HasOpt("Foo"))
	assert.Equal(t, cmd.OptArg("Foo"), "")
	assert.Equal(t, cmd.OptArgs("Foo"), []string(nil))
}

func TestParseWith_oneCfgBeingArrayAndUsingValidator(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey:  "Foo",
			Names:     []string{"foo", "f"},
			HasArg:    true,
			IsArray:   true,
			Validator: &validators.ValidateInt32,
		},
	}

	os.Args = []string{"app", "--foo", "123", "def", "-f", "456", "ghi"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"def", "ghi"})
	assert.True(t, cmd.HasOpt("Foo"))
	assert.Equal(t, cmd.OptArg("Foo"), "123")
	assert.Equal(t, cmd.OptArgs("Foo"), []string{"123", "456"})
}

func TestParseWith_oneCfgBeingArrayAndUsingValidator_validationError(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey:  "Foo",
			Names:     []string{"foo", "f"},
			HasArg:    true,
			IsArray:   true,
			Validator: &validators.ValidateInt32,
		},
	}

	os.Args = []string{"app", "-f=123", "--foo", "ABC", "def", "-f", "456", "ghi"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	switch e := err.(type) {
	case errors.OptionArgIsInvalid:
		assert.Equal(t, e.StoreKey, "Foo")
		assert.Equal(t, e.Option, "foo")
		assert.Equal(t, e.OptArg, "ABC")
		assert.Equal(t, e.Cause.Error(), "strconv.ParseInt: parsing \"ABC\": invalid syntax")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"def", "ghi"})
	assert.True(t, cmd.HasOpt("Foo"))
	assert.Equal(t, cmd.OptArg("Foo"), "123", "456")
	assert.Equal(t, cmd.OptArgs("Foo"), []string{"123", "456"})
}

///

func TestParseUntilSubCmdWith_zeroCfg_thereIsNoArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"app"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})

	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))

	assert.False(t, subCmd.HasOpt("foo-bar"))
	assert.Equal(t, subCmd.OptArg("foo-bar"), "")
	assert.Equal(t, subCmd.OptArgs("foo-bar"), []string(nil))
}

func TestParseUntilSubCmdWith_zeroCfg_thereIsSubCmd(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"/path/to/app", "foo-bar", "baz"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))

	assert.Equal(t, subCmd.Name, "foo-bar")
	assert.Equal(t, subCmd.Args, []string{})
	assert.False(t, subCmd.HasOpt("foo-bar"))
	assert.Equal(t, subCmd.OptArg("foo-bar"), "")
	assert.Equal(t, subCmd.OptArgs("foo-bar"), []string(nil))

	err = subCmd.Parse()
	assert.Nil(t, err)

	assert.Equal(t, subCmd.Name, "foo-bar")
	assert.Equal(t, subCmd.Args, []string{"baz"})
	assert.False(t, subCmd.HasOpt("foo-bar"))
	assert.Equal(t, subCmd.OptArg("foo-bar"), "")
	assert.Equal(t, subCmd.OptArgs("foo-bar"), []string(nil))
}

func TestParseUntilSubCmdWith_zeroCfg_oneOptWithNoArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{}

	os.Args = []string{"path/to/app", "--foo-bar"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "UnconfiguredOption{Option:foo-bar}")

	switch e := err.(type) {
	case errors.UnconfiguredOption:
		assert.Equal(t, e.Option, "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	switch e := err.(type) {
	case errors.InvalidOption:
		assert.Equal(t, e.GetOption(), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))
	assert.False(t, subCmd.HasOpt("foo-bar"))
	assert.Equal(t, subCmd.OptArg("foo-bar"), "")
	assert.Equal(t, subCmd.OptArgs("foo-bar"), []string(nil))

	err = subCmd.Parse()
	assert.Nil(t, err)

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))
	assert.False(t, subCmd.HasOpt("foo-bar"))
	assert.Equal(t, subCmd.OptArg("foo-bar"), "")
	assert.Equal(t, subCmd.OptArgs("foo-bar"), []string(nil))
}

func TestParseUntilSubCmdWith_zeroCfg_oneOptWithArg(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	os.Args = []string{"app", "--foo-bar=ABC"}

	cmd = cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))

	os.Args = []string{"app", "--foo-bar", ""}

	cmd = cliargs.NewCmd()
	subCmd, err = cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{""})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))

	os.Args = []string{"app", "--foo-bar="}

	cmd = cliargs.NewCmd()
	subCmd, err = cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{""})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))
}

func TestParseUntilSubCmdWith_oneCfgAndSubCmd(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}},
	}

	os.Args = []string{"app", "--foo-bar", "ABC"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string(nil))
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "ABC")
	assert.Equal(t, subCmd.Args, []string{})

	err = subCmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, subCmd.Name, "ABC")
	assert.Equal(t, subCmd.Args, []string{})
}

func TestParseUntilSubCmdWith_oneCfgHavingOptArgAndSubCmd(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{Names: []string{"foo-bar"}, HasArg: true},
	}

	os.Args = []string{"app", "--foo-bar", "ABC", "def", "ghi"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo-bar"))
	assert.Equal(t, cmd.OptArg("foo-bar"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "def")
	assert.Equal(t, subCmd.Args, []string{})

	err = subCmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, subCmd.Name, "def")
	assert.Equal(t, subCmd.Args, []string{"ghi"})
}

func TestParseUntilSubCmdWith_oneCfgUsingStoreKeyAndSubCmd(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{StoreKey: "foo", HasArg: true},
	}

	os.Args = []string{"app", "--foo", "ABC", "def", "ghi"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "ABC")
	assert.Equal(t, cmd.OptArgs("foo"), []string{"ABC"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "def")
	assert.Equal(t, subCmd.Args, []string{})

	err = subCmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, subCmd.Name, "def")
	assert.Equal(t, subCmd.Args, []string{"ghi"})
}

func TestParseUntilSubCmdWith_oneCfgUsingValidatorAndSubCmd(t *testing.T) {
	defer reset()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			StoreKey:  "foo",
			HasArg:    true,
			Validator: &validators.ValidateInt32,
		},
	}

	os.Args = []string{"app", "--foo", "123", "def", "ghi"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "123")
	assert.Equal(t, cmd.OptArgs("foo"), []string{"123"})
	assert.False(t, cmd.HasOpt("f"))
	assert.Equal(t, cmd.OptArg("f"), "")
	assert.Equal(t, cmd.OptArgs("f"), []string(nil))

	assert.Equal(t, subCmd.Name, "def")
	assert.Equal(t, subCmd.Args, []string{})

	err = subCmd.Parse()

	assert.Nil(t, err)
	assert.Equal(t, subCmd.Name, "def")
	assert.Equal(t, subCmd.Args, []string{"ghi"})
}
