// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"path"
)

// StoreKeyIsDuplicated is the error which indicates that a store key in an
// option configuration is duplicated another among all option configurations.
type StoreKeyIsDuplicated struct{ StoreKey string }

// Error is the method to retrieve the message of this error.
func (e StoreKeyIsDuplicated) Error() string {
	return fmt.Sprintf("StoreKeyIsDuplicated{StoreKey:%s}", e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e StoreKeyIsDuplicated) GetOpt() string {
	return e.StoreKey
}

// ConfigIsArrayButHasNoArg is the error which indicates that an option
// configuration contradicts that the option must be an array
// (.IsArray = true) but must have no option argument (.HasArg = false).
type ConfigIsArrayButHasNoArg struct{ StoreKey string }

// Error is the method to retrieve the message of this error.
func (e ConfigIsArrayButHasNoArg) Error() string {
	return fmt.Sprintf("ConfigIsArrayButHasNoArg{StoreKey:%s}", e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e ConfigIsArrayButHasNoArg) GetOpt() string {
	return e.StoreKey
}

// ConfigHasDefaultsButHasNoArg is the error which indicates that an option
// configuration contradicts that the option has default value
// (.Defaults != nil) but must have no option argument (.HasArg = false).
type ConfigHasDefaultsButHasNoArg struct{ StoreKey string }

// Error is the method to retrieve the message of this error.
func (e ConfigHasDefaultsButHasNoArg) Error() string {
	return fmt.Sprintf("ConfigHasDefaultsButHasNoArg{StoreKey:%s}", e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e ConfigHasDefaultsButHasNoArg) GetOpt() string {
	return e.StoreKey
}

// OptionNameIsDuplicated is the error which indicates that an option name
// in Names field is duplicated another among all option configurations.
type OptionNameIsDuplicated struct{ Name, StoreKey string }

// Error is the method to retrieve the message of this error.
func (e OptionNameIsDuplicated) Error() string {
	return fmt.Sprintf("OptionNameIsDuplicated{Name:%s,StoreKey:%s}",
		e.Name, e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e OptionNameIsDuplicated) GetOpt() string {
	return e.Name
}

// UnconfiguredOption
type UnconfiguredOption struct{ Name string }

// Error is the method to retrieve the message of this error.
func (e UnconfiguredOption) Error() string {
	return fmt.Sprintf("UnconfiguredOption{Name:%s}", e.Name)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e UnconfiguredOption) GetOpt() string {
	return e.Name
}

// OptionTakesNoArg
type OptionTakesNoArg struct{ Name, StoreKey string }

// Error is the method to retrieve the message of this error.
func (e OptionTakesNoArg) Error() string {
	return fmt.Sprintf("OptionTakesNoArg{Name:%s,StoreKey:%s}",
		e.Name, e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e OptionTakesNoArg) GetOpt() string {
	return e.Name
}

// OptionNeedsArg
type OptionNeedsArg struct{ Name, StoreKey string }

// Error is the method to retrieve the message of this error.
func (e OptionNeedsArg) Error() string {
	return fmt.Sprintf("OptionNeedsArg{Name:%s,StoreKey:%s}",
		e.Name, e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e OptionNeedsArg) GetOpt() string {
	return e.Name
}

// OptionIsNotArray
type OptionIsNotArray struct{ Name, StoreKey string }

// Error is the method to retrieve the message of this error.
func (e OptionIsNotArray) Error() string {
	return fmt.Sprintf("OptionIsNotArray{Name:%s,StoreKey:%s}",
		e.Name, e.StoreKey)
}

// GetOpt is the method to retrieve the store key that caused this error.
func (e OptionIsNotArray) GetOpt() string {
	return e.Name
}

const anyOption = "*"

// OptCfg is the struct that represents an option configuration.
// An option configuration consists of fields: StoreKey, Names, HasArg,
// IsArray, Defaults, Desc, and ArgInHelp.
//
// The StoreKey field is the key to store a option value(s) in the option map.
// If this key is not specified or empty, the first element of Names field is
// used instead.
//
// The Names field is the array for specifing the option name and the aliases.
// The order of the names in this array are used in a help text.
//
// HasArg and IsArray are flags which allow the option to take option
// arguments.
// If both HasArg and IsArray are true, the option can take one or multiple
// option arguments.
// If HasArg is true and IsArray is false, the option can take only one option
// arguments.
// If both HasArg and IsArray are false, the option can take no option
// argument.
//
// Defaults is the field to specified the default value for when the option is
// not given in command line arguments.
//
// OnParsed is the field for a function which is called when the option has
// been parsed.
//
// Desc is the field to set the description of the option.
//
// ArgInHelp is a display of the argument of this option in a help text.
// The example of the display is like: -o, --option <value>.
type OptCfg struct {
	StoreKey  string
	Names     []string
	HasArg    bool
	IsArray   bool
	Defaults  []string
	OnParsed  *func([]string) error
	Desc      string
	ArgInHelp string
}

// ParseWith is the function which parses command line arguments with option
// configurations.
// This function divides command line arguments to command arguments and
// options. And an option consists of a name and an option argument.
// Options are divided to long format options and short format options.
// About long/short format options, since they are same with Parse function,
// see the comment of that function.
//
// This function allows only options declared in option configurations.
// An option configuration has fields: StoreKey, Names, HasArg, IsArray,
// Defaults, Desc and ArgInHelp.
// When an option matches one of the Names in an option configuration, the
// option is registered into Cmd with StoreKey.
// If both HasArg and IsArray are true, the option can have one or multiple
// option argumentsr, and if HasArg is true and IsArray is false, the option
// can have only one option argument, otherwise the option cannot have option
// arguments.
// If Defaults field is specified and no option value is given in command line
// arguments, the value of Defaults is set as the option arguments.
//
// If options not declared in option configurationsi are given in command line
// arguments, this function basically returns UnconfiguradOption error.
// However, if you want to allow other options, add an option configuration of
// which StoreKey or the first element of Names is "*".
func ParseWith(osArgs []string, optCfgs []OptCfg) (Cmd, error) {
	var cmdName string
	if len(osArgs) > 0 {
		cmdName = path.Base(osArgs[0])
	}

	var opts = make(map[string][]string)
	var err error = nil

	cfgMap := make(map[string]int)
	hasAnyOpt := false

	for i, cfg := range optCfgs {
		storeKey := cfg.StoreKey
		if len(storeKey) == 0 && len(cfg.Names) > 0 {
			storeKey = cfg.Names[0]
		}

		if len(storeKey) == 0 {
			continue
		}

		if storeKey == anyOption {
			hasAnyOpt = true
			continue
		}

		_, exists := opts[storeKey]
		if exists {
			err = StoreKeyIsDuplicated{StoreKey: storeKey}
			return Cmd{Name: cmdName, args: empty}, err
		}
		opts[storeKey] = nil

		if !cfg.HasArg {
			if cfg.IsArray {
				err = ConfigIsArrayButHasNoArg{StoreKey: storeKey}
				return Cmd{Name: cmdName, args: empty}, err
			}
			if cfg.Defaults != nil {
				err = ConfigHasDefaultsButHasNoArg{StoreKey: storeKey}
				return Cmd{Name: cmdName, args: empty}, err
			}
		}

		for _, nm := range cfg.Names {
			_, exists := cfgMap[nm]
			if exists {
				err = OptionNameIsDuplicated{Name: nm, StoreKey: storeKey}
				return Cmd{Name: cmdName, args: empty}, err
			}
			cfgMap[nm] = i
		}
	}

	var takeArgs = func(opt string) bool {
		i, exists := cfgMap[opt]
		if exists {
			return optCfgs[i].HasArg
		}
		return false
	}

	args := make([]string, 0)
	opts = make(map[string][]string)

	var collectArgs = func(a ...string) {
		args = append(args, a...)
	}
	var collectOpts = func(name string, a ...string) error {
		i, exists := cfgMap[name]
		if !exists {
			if !hasAnyOpt {
				return UnconfiguredOption{Name: name}
			}

			arr := opts[name]
			if arr == nil {
				arr = empty
			}
			opts[name] = append(arr, a...)
			return nil
		}

		cfg := optCfgs[i]

		storeKey := cfg.StoreKey
		if len(storeKey) == 0 {
			storeKey = cfg.Names[0]
		}

		if !cfg.HasArg {
			if len(a) > 0 {
				return OptionTakesNoArg{Name: name, StoreKey: storeKey}
			}
		} else {
			if len(a) == 0 {
				return OptionNeedsArg{Name: name, StoreKey: storeKey}
			}
		}

		arr := opts[storeKey]
		if arr == nil {
			arr = empty
		}
		arr = append(arr, a...)

		if !cfg.IsArray {
			if len(arr) > 1 {
				return OptionIsNotArray{Name: name, StoreKey: storeKey}
			}
		}

		opts[storeKey] = arr
		return nil
	}

	var osArgs1 []string
	if len(osArgs) > 1 {
		osArgs1 = osArgs[1:]
	}

	err = parseArgs(osArgs1, collectArgs, collectOpts, takeArgs)

	for _, cfg := range optCfgs {
		if len(cfg.Names) == 0 {
			continue
		}

		storeKey := cfg.StoreKey
		if len(storeKey) == 0 {
			storeKey = cfg.Names[0]
		}

		arr, exists := opts[storeKey]
		if !exists && cfg.Defaults != nil {
			arr = cfg.Defaults
			opts[storeKey] = arr
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
