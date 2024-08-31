// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"github.com/sttk/cliargs/errors"
)

const anyOption = "*"

// ParseWith is the method which parses command line arguments with option configurations.
// This function divides command line arguments to command arguments and options.
// And an option consists of a name and an option argument.
// Options are separated to long format options and short format options.
// About long/short format options, since they are same with Parse function, see the comment of
// that function.
//
// This function allows only options declared in option configurations.
// An option configuration has fields: StoreKey, Names, HasArg, IsArray, Defaults, Desc and
// ArgInHelp.
// When an option matches one of the Names in an option configuration, the option is registered
// into Cmd with StoreKey.
// If both HasArg and IsArray are true, the option can have one or multiple option argumentsr, and
// if HasArg is true and IsArray is false, the option can have only one option argument, otherwise
// the option cannot have option arguments.
// If Defaults field is specified and no option value is given in command line arguments, the value
// of Defaults is set as the option arguments.
//
// If options not declared in option configurationsi are given in command line arguments, this
// function basically returns UnconfiguradOption error.
// However, if you want to allow other options, add an option configuration of which StoreKey or
// the first element of Names is "*".
func (cmd *Cmd) ParseWith(optCfgs []OptCfg) error {
	_, err := cmd.parseArgsWith(optCfgs, false)
	cmd.OptCfgs = optCfgs
	return err
}

// ParseUntilSubCmdWith is the method which parses command line arguments with option
// configurations but stops parsing when encountering first command argument.
//
// This method creates and returns a new Cmd instance that holds the command line arguments
// starting from the first command argument.
//
// This method parses command line arguments in the same way as the Cmd#parse_with method,
// except that it only parses the command line arguments before the first command argument.
func (cmd *Cmd) ParseUntilSubCmdWith(optCfgs []OptCfg) (Cmd, error) {
	idx, err := cmd.parseArgsWith(optCfgs, true)
	cmd.OptCfgs = optCfgs
	if idx < 0 {
		return Cmd{}, err
	}
	return cmd.subCmd(idx), err
}

func (cmd *Cmd) parseArgsWith(
	optCfgs []OptCfg,
	untilFirstArg bool,
) (int, error) {

	const ANY_OPT = "*"
	hasAnyOpt := false

	var EMPTY_STRUCT struct{}

	optMap := make(map[string]struct{})
	cfgMap := make(map[string]int)

	for i, cfg := range optCfgs {
		var names []string
		for _, nm := range cfg.Names {
			if len(nm) > 0 {
				names = append(names, nm)
				break
			}
		}

		storeKey := cfg.StoreKey
		if len(storeKey) == 0 && len(names) > 0 {
			storeKey = names[0]
		}

		if len(storeKey) == 0 {
			continue
		}

		if storeKey == ANY_OPT {
			hasAnyOpt = true
			continue
		}

		var firstName string
		if len(names) > 0 {
			firstName = names[0]
		} else {
			firstName = storeKey
		}

		_, exists := optMap[storeKey]
		if exists {
			e := errors.StoreKeyIsDuplicated{StoreKey: storeKey, Name: firstName}
			return -1, e
		}
		optMap[storeKey] = EMPTY_STRUCT

		if !cfg.HasArg {
			if cfg.IsArray {
				e := errors.ConfigIsArrayButHasNoArg{StoreKey: storeKey, Name: firstName}
				return -1, e
			}
			if cfg.Defaults != nil {
				e := errors.ConfigHasDefaultsButHasNoArg{StoreKey: storeKey, Name: firstName}
				return -1, e
			}
		}

		if len(names) == 0 {
			cfgMap[firstName] = i
		} else {
			for _, nm := range cfg.Names {
				_, exists := cfgMap[nm]
				if exists {
					e := errors.OptionNameIsDuplicated{StoreKey: storeKey, Name: nm}
					return -1, e
				}
				cfgMap[nm] = i
			}
		}
	}

	var takeOptArgs = func(opt string) bool {
		i, exists := cfgMap[opt]
		if exists {
			return optCfgs[i].HasArg
		}
		return false
	}

	var collectArgs = func(arg string) {
		cmd.Args = append(cmd.Args, arg)
	}

	var collectOpts = func(name string, a ...string) error {
		i, exists := cfgMap[name]
		if exists {
			cfg := optCfgs[i]

			var storeKey string
			if len(cfg.StoreKey) > 0 {
				storeKey = cfg.StoreKey
			} else {
				for _, nm := range cfg.Names {
					if len(nm) > 0 {
						storeKey = nm
						break
					}
				}
			}

			if len(a) > 0 {
				if !cfg.HasArg {
					return errors.OptionTakesNoArg{
						Option:   name,
						StoreKey: storeKey,
					}
				}

				arr, _ := cmd.opts[storeKey]
				if len(arr) > 0 {
					if !cfg.IsArray {
						return errors.OptionIsNotArray{
							StoreKey: storeKey,
							Option:   name,
						}
					}
					if cfg.Validator != nil {
						err := (*cfg.Validator)(storeKey, name, a[0])
						if err != nil {
							return err
						}
					}
					cmd.opts[storeKey] = append(arr, a[0])
				} else {
					if cfg.Validator != nil {
						err := (*cfg.Validator)(storeKey, name, a[0])
						if err != nil {
							return err
						}
					}
					cmd.opts[storeKey] = append(arr, a[0])
				}
			} else {
				if cfg.HasArg {
					return errors.OptionNeedsArg{
						Option:   name,
						StoreKey: storeKey,
					}
				}

				_, exists := cmd.opts[storeKey]
				if !exists {
					cmd.opts[storeKey] = nil
				}
			}

			return nil
		} else {
			if !hasAnyOpt {
				return errors.UnconfiguredOption{
					Option: name,
				}
			}

			if len(a) > 0 {
				cmd.opts[name] = append(cmd.opts[name], a[0])
			} else {
				cmd.opts[name] = nil
			}

			return nil
		}
	}

	idx, err := parseArgs(
		cmd._args,
		collectArgs,
		collectOpts,
		takeOptArgs,
		untilFirstArg,
	)

	for _, cfg := range optCfgs {
		storeKey := cfg.StoreKey
		if len(storeKey) == 0 && len(cfg.Names) > 0 {
			for _, nm := range cfg.Names {
				if len(nm) > 0 {
					storeKey = nm
					break
				}
			}
		}

		if len(storeKey) == 0 {
			continue
		}

		if storeKey == ANY_OPT {
			continue
		}

		arr, exists := cmd.opts[storeKey]
		if !exists && cfg.Defaults != nil {
			arr = cfg.Defaults
			cmd.opts[storeKey] = arr
			exists = true
		}

		if exists && cfg.onParsed != nil {
			e := (*cfg.onParsed)(arr)
			if e != nil && err == nil {
				err = e
			}
		}
	}

	return idx, err
}
