// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"path"
)

// ConfigIsArrayButHasNoArg is the error which indicates that an option
// configuration contradicts that the option must be an array
// (.IsArray = true) but must have no option argument (.HasArg = false).
type ConfigIsArrayButHasNoArg struct{ Option string }

// Error is the method to retrieve the message of this error.
func (e ConfigIsArrayButHasNoArg) Error() string {
	return fmt.Sprintf("ConfigIsArrayButHasNoArg{Option:%s}", e.Option)
}

// GetOpt is the method to retrieve the option that caused this error.
func (e ConfigIsArrayButHasNoArg) GetOpt() string {
	return e.Option
}

// ConfigHasDefaultsButHasNoArg is the error which indicates that an option
// configuration contradicts that the option has default value
// (.Defaults != nil) but must have no option argument (.HasArg = false).
type ConfigHasDefaultsButHasNoArg struct{ Option string }

// Error is the method to retrieve the message of this error.
func (e ConfigHasDefaultsButHasNoArg) Error() string {
	return fmt.Sprintf("ConfigHasDefaultsButHasNoArg{Option:%s}", e.Option)
}

// GetOpt is the method to retrieve the option that caused this error.
func (e ConfigHasDefaultsButHasNoArg) GetOpt() string {
	return e.Option
}

// UnconfiguredOption is the error which indicates that there is no
// configuration about the input option.
type UnconfiguredOption struct{ Option string }

// Error is the method to retrieve the message of this error.
func (e UnconfiguredOption) Error() string {
	return fmt.Sprintf("UnconfiguredOption{Option:%s}", e.Option)
}

// GetOpt is the method to retrieve the option that caused this error.
func (e UnconfiguredOption) GetOpt() string {
	return e.Option
}

// OptionNeedsArg is the error which indicates that an option is input with
// no option argument though its option configuration requires option
// argument (.HasArg = true).
type OptionNeedsArg struct{ Option string }

// Error is the method to retrieve the message of this error.
func (e OptionNeedsArg) Error() string {
	return fmt.Sprintf("OptionNeedsArg{Option:%s}", e.Option)
}

// GetOpt is the method to retrieve the option that caused this error.
func (e OptionNeedsArg) GetOpt() string {
	return e.Option
}

// OptionTakesNoArg is the error which indicates that an option is input with
// an option argument though its option configuration does not accept option
// arguments (.HasArg = false).
type OptionTakesNoArg struct{ Option string }

// Error is the method to retrieve the message of this error.
func (e OptionTakesNoArg) Error() string {
	return fmt.Sprintf("OptionTakesNoArg{Option:%s}", e.Option)
}

// GetOpt is the method to retrieve the option that caused this error.
func (e OptionTakesNoArg) GetOpt() string {
	return e.Option
}

// OptionIsNotArray is the error which indicates that an option is input with
// an option argument multiple times though its option configuration specifies
// the option is not an array (.IsArray = false).
type OptionIsNotArray struct{ Option string }

// Error is the method to retrieve the message of this error.
func (e OptionIsNotArray) Error() string {
	return fmt.Sprintf("OptionIsNotArray{Option:%s}", e.Option)
}

// GetOpt is the method to retrieve the option that caused this error.
func (e OptionIsNotArray) GetOpt() string {
	return e.Option
}

const anyOption = "*"

// OptCfg is the struct that represents an option configuration.
// An option configuration consists of fields: Name, Aliases, HasArg,
// IsArray, Defaults, OnParsed, Desc, and ArgHelp.
//
// Name is the option name and Aliases are the another names.
// Options given by those names in command line arguments are all registered to
// Args with the Name.
//
// HasArg and IsArray are flags which allows the option to take option
// arguments.
// If both HasArg and IsArray are true, the option can take one or multiple
// option arguments.
// If HasArg is true and IsArray is false, the option can take only one
// option arguments.
// If both HasArg and IsArray are false, the option can take no option
// argument.
//
// Defaults is the field to specify the default value for when the option is
// not given in command line arguments.
//
// OnParsed is the field for the event handler which is called when the option
// has been parsed.
// This handler receives a string array which is the option argument(s) as its
// argument.
// If this field is nil, nothing is done after parsing.
//
// Desc is the field to set the description of the option.
//
// ArgHelp is a display at a argument position of this option in a help text.
// This string is for a display like: -o, --option <value>.
type OptCfg struct {
	Name     string
	Aliases  []string
	HasArg   bool
	IsArray  bool
	Defaults []string
	OnParsed *func([]string) error
	Desc     string
	ArgHelp  string
}

// ParseWith is the function which parses command line arguments with option
// configurations.
// This function divides command line arguments to command arguments and
// options, and an option consists of a name and option arguments.
// Options are divided to long format options and short format options.
// About long/short format options, since they are same with Parse function,
// see the comment of the function.
//
// This function allows only options declared in option configurations.
// A option configuration has fields: Name, Aliases, HasArg, IsArray, Defaults,
// Desc and ArgHelp.
// When an option matches Name or includes in Aliases in an option
// configuration, the option is registered into Cmd with the Name.
// If both HasArg and IsArray are true, the option can have one or multiple
// option arguments, and if HasArg is true and IsArray is false, the option
// can have only one option argument, otherwise the option cannot have option
// argument.
// If Defaults is specified and the option is not given in command line
// arguments, the value of Defaults is set to the option argument.
//
// If options not declared in option configurations are given in command line
// arguments, this function basically returns UnconfiguredOption error.
// If you want to allow other options, add an option configuration of which
// Name is "*" (but HasArg and IsArray of this configuration is ignored).
func ParseWith(osArgs []string, optCfgs []OptCfg) (Cmd, error) {
	var cmdName string
	if len(osArgs) > 0 {
		cmdName = path.Base(osArgs[0])
	}

	hasAnyOpt := false
	cfgMap := make(map[string]int)
	for i, cfg := range optCfgs {
		if !cfg.HasArg {
			if cfg.IsArray {
				err := ConfigIsArrayButHasNoArg{Option: cfg.Name}
				return Cmd{Name: cmdName, args: empty}, err
			}
			if cfg.Defaults != nil {
				err := ConfigHasDefaultsButHasNoArg{Option: cfg.Name}
				return Cmd{Name: cmdName, args: empty}, err
			}
		}
		if cfg.Name == anyOption {
			hasAnyOpt = true
			continue
		}
		cfgMap[cfg.Name] = i
		for _, a := range cfg.Aliases {
			cfgMap[a] = i
		}
	}

	var takeArgs = func(opt string) bool {
		i, exists := cfgMap[opt]
		if exists {
			return optCfgs[i].HasArg
		}
		return false
	}

	var args = make([]string, 0)
	var opts = make(map[string][]string)

	var collectArgs = func(a ...string) {
		args = append(args, a...)
	}
	var collectOpts = func(name string, a ...string) error {
		i, exists := cfgMap[name]
		if !exists {
			if !hasAnyOpt {
				return UnconfiguredOption{Option: name}
			}

			arr := opts[name]
			if arr == nil {
				arr = empty
			}
			opts[name] = append(arr, a...)
			return nil
		}

		cfg := optCfgs[i]
		if !cfg.HasArg {
			if len(a) > 0 {
				return OptionTakesNoArg{Option: cfg.Name}
			}
		} else {
			if len(a) == 0 {
				return OptionNeedsArg{Option: cfg.Name}
			}
		}

		arr := opts[cfg.Name]
		if arr == nil {
			arr = empty
		}
		arr = append(arr, a...)

		if !cfg.IsArray {
			if len(arr) > 1 {
				return OptionIsNotArray{Option: cfg.Name}
			}
		}

		opts[cfg.Name] = arr
		return nil
	}

	var osArgs1 []string
	if len(osArgs) > 1 {
		osArgs1 = osArgs[1:]
	}

	err := parseArgs(osArgs1, collectArgs, collectOpts, takeArgs)

	for _, cfg := range optCfgs {
		arr, exists := opts[cfg.Name]
		if !exists && cfg.Defaults != nil {
			arr = cfg.Defaults
			opts[cfg.Name] = arr
		}
		if cfg.OnParsed != nil {
			e := (*cfg.OnParsed)(arr)
			if e != nil && err == nil {
				err = e
			}
		}
	}

	return Cmd{Name: cmdName, args: args, opts: opts}, err
}
