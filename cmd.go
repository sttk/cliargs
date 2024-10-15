// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"os"
	"path"
)

// Cmd is the structure that parses command line arguments and stores them.
// The results of parsing are stored by separating into command name, command arguments, options,
// and option arguments.
// And this provides methods to check if they are specified and to retrieve them.
type Cmd struct {
	Name    string
	Args    []string
	OptCfgs []OptCfg

	opts          map[string][]string
	isAfterEndOpt bool

	_args []string
}

// NewCmd is the function that creates a Cmd instance iwth command line arguments obtained from
// os.Args.
func NewCmd() Cmd {
	var name string
	if len(os.Args) > 0 {
		name = path.Base(os.Args[0])
	}

	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	return Cmd{Name: name, Args: []string{}, opts: make(map[string][]string), _args: args}
}

func (cmd Cmd) subCmd(fromIndex int, isAfterEndOpt bool) Cmd {
	var name string
	if len(cmd._args) > fromIndex {
		name = cmd._args[fromIndex]
	}

	var args []string
	if len(cmd._args) > fromIndex+1 {
		args = cmd._args[fromIndex+1:]
	}

	return Cmd{
		Name:          name,
		Args:          []string{},
		opts:          make(map[string][]string),
		isAfterEndOpt: isAfterEndOpt,
		_args:         args,
	}
}

// HasOpt is the method that checks whether an option with the specified name exists.
func (cmd Cmd) HasOpt(name string) bool {
	_, exists := cmd.opts[name]
	return exists
}

// OptArg is the method that returns the option argument with the specified name.
// If the option has multiple arguments, this method returns the first argument.
// If the option is a boolean flag, the method returns an empty string.
// If the option is not specified in the command line arguments, the return value
// of this method is an empty string.
func (cmd Cmd) OptArg(name string) string {
	arr := cmd.opts[name]
	// If no entry, map returns a nil slice.
	// len() methods of both a nil slice and a empty slice return zero in common.
	if len(arr) == 0 {
		return ""
	} else {
		return arr[0]
	}
}

// OptArgs is the method that returns the option arguments with the specified name.
// If the option has one or multiple arguments, this method returns an array of the arguments.
// If the option is a boolean flag, the method returns an empty slice.
// If the option is not specified in the command line arguments, the return value
// of this method is a nil slice.
func (cmd Cmd) OptArgs(name string) []string {
	return cmd.opts[name]
}

// String is the method that returns the string which represents the content of this instance.
func (cmd Cmd) String() string {
	return fmt.Sprintf("Cmd { Name: %s, Args: %v, Opts: %v }", cmd.Name, cmd.Args, cmd.opts)
}
