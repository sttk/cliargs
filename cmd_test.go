package cliargs

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var origOsArgs = os.Args

func reset() {
	os.Args = origOsArgs
}

func TestCmd_NewCmd_withOsArgs(t *testing.T) {
	defer reset()
	os.Args = []string{"/path/to/app", "--foo", "bar"}

	cmd := NewCmd()
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, cmd.String(), "Cmd { Name: app, Args: [], Opts: map[] }")
	assert.Equal(t, fmt.Sprintf("%v", cmd), "Cmd { Name: app, Args: [], Opts: map[] }")

	assert.Equal(t, cmd._args, []string{"--foo", "bar"})

	assert.Equal(t, cmd.HasOpt("foo"), false)
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string(nil))
}

func TestCmd_NewCmd_withNoArgs(t *testing.T) {
	defer reset()
	os.Args = []string{"/path/to/app"}

	cmd := NewCmd()
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, cmd.String(), "Cmd { Name: app, Args: [], Opts: map[] }")
	assert.Equal(t, fmt.Sprintf("%v", cmd), "Cmd { Name: app, Args: [], Opts: map[] }")

	assert.Equal(t, cmd._args, []string(nil))

	assert.Equal(t, cmd.HasOpt("foo"), false)
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string(nil))
}

func TestCmd_NewCmd_empty(t *testing.T) {
	defer reset()
	os.Args = []string{}

	cmd := NewCmd()
	assert.Equal(t, cmd.Name, "")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, cmd.String(), "Cmd { Name: , Args: [], Opts: map[] }")
	assert.Equal(t, fmt.Sprintf("%v", cmd), "Cmd { Name: , Args: [], Opts: map[] }")

	assert.Equal(t, cmd._args, []string(nil))

	assert.Equal(t, cmd.HasOpt("foo"), false)
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string(nil))
}

func TestCmd_subCmd(t *testing.T) {
	cmd := NewCmd()
	cmd._args = []string{"--foo", "-b", "qux", "--corge"}

	subCmd := cmd.subCmd(2, false)
	assert.Equal(t, subCmd.Name, "qux")
	assert.Equal(t, subCmd._args, []string{"--corge"})
}

func TestCmd_subCmd_withNoArg(t *testing.T) {
	cmd := NewCmd()
	cmd._args = []string{"--foo", "-b", "qux"}

	subCmd := cmd.subCmd(2, false)
	assert.Equal(t, subCmd.Name, "qux")
	assert.Equal(t, subCmd._args, []string(nil))
}

func TestCmd_subCmd_empty(t *testing.T) {
	cmd := NewCmd()
	cmd._args = []string{"--foo", "-b"}

	subCmd := cmd.subCmd(2, false)
	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd._args, []string(nil))
}

func TestCmd_NewCmd_HasOpt(t *testing.T) {
	defer reset()

	cmd := NewCmd()
	cmd.opts["foo-bar"] = []string(nil)
	cmd.opts["baz"] = []string{"123"}
	cmd.opts["qux"] = []string{"A", "B"}

	assert.Equal(t, cmd.HasOpt("foo-bar"), true)
	assert.Equal(t, cmd.HasOpt("baz"), true)
	assert.Equal(t, cmd.HasOpt("qux"), true)
	assert.Equal(t, cmd.HasOpt("quux"), false)
}

func TestCmd_NewCmd_OptArg(t *testing.T) {
	defer reset()

	cmd := NewCmd()
	cmd.opts["foo-bar"] = []string{}
	cmd.opts["baz"] = []string{"123"}
	cmd.opts["qux"] = []string{"A", "B"}

	assert.Equal(t, cmd.OptArg("foo-bar"), "")
	assert.Equal(t, cmd.OptArg("baz"), "123")
	assert.Equal(t, cmd.OptArg("qux"), "A")
	assert.Equal(t, cmd.OptArg("quux"), "")
}

func TestCmd_NewCmd_OptArgs(t *testing.T) {
	defer reset()

	cmd := NewCmd()
	cmd.opts["foo-bar"] = []string{}
	cmd.opts["baz"] = []string{"123"}
	cmd.opts["qux"] = []string{"A", "B"}

	assert.Equal(t, cmd.OptArgs("foo-bar"), []string{})
	assert.Equal(t, cmd.OptArgs("baz"), []string{"123"})
	assert.Equal(t, cmd.OptArgs("qux"), []string{"A", "B"})
	assert.Equal(t, cmd.OptArgs("quux"), []string(nil))
}
