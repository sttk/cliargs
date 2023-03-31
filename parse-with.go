// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

// ConfigIsArrayButHasNoParam is an error which indicates that an option
// configuration contradicts that the option must be an array
// (.IsArray = true) but must have no option parameter (.HasParam = false).
type ConfigIsArrayButHasNoParam struct{ Option string }

func (e ConfigIsArrayButHasNoParam) Error() string {
	return "ConfigIsArrayButHasNoParam"
}

// ConfigHasDefaultButHasNoParam is an error which indicates that an option
// configuration contradicts that the option has default value
// (.Default != nil) but must have no option parameter (.HasParam = false).
type ConfigHasDefaultButHasNoParam struct{ Option string }

func (e ConfigHasDefaultButHasNoParam) Error() string {
	return "ConfigHasDefaultButHasNoParam"
}

// UnconfiguredOption is an error which indicates that there is no
// configuration about the input option.
type UnconfiguredOption struct{ Option string }

func (e UnconfiguredOption) Error() string {
	return "UnconfiguredOption"
}

// OptionNeedsParam is an error which indicates that an option is input with
// no option parameter though its option configuration requires option
// parameters (.HasParam = true).
type OptionNeedsParam struct{ Option string }

func (e OptionNeedsParam) Error() string {
	return "OptionNeedsParam"
}

// OptionTakesNoParam is an error which indicates that an option isinput with
// an option parameter though its option configuration does not accept option
// parameters (.HasParam = false).
type OptionTakesNoParam struct{ Option string }

func (e OptionTakesNoParam) Error() string {
	return "OptionTakesNoParam"
}

// OptionIsNotArray is an error which indicates that an option is input with
// an option parameter multiple times though its option configuration specifies
// the option is not an array (.IsArray = false).
type OptionIsNotArray struct{ Option string }

func (e OptionIsNotArray) Error() string {
	return "OptionIsNotArray"
}

const anyOption = "*"

// OptCfg is a structure that represents an option configuration.
// An option configuration consists of fields: Name, Aliases, HasParam,
// IsArray, Default, OnParsed, Desc, and AtParam.

// Name is the option name and Aliases are the another names.
// Options given by those names in command line arguments are all registered to
// Args with the Name.
//
// HasParam and IsArray are flags which allows the option to take option
// parameters.
// If both HasParam and IsArray are true, the option can take one or multiple
// option parameters.
// If HasParam is true and IsArray is false, the option can take only one
// option parameter.
// If both HasParam and IsArray are false, the option can take no option
// parameter.
//
// Default is the field to specify the default value for when the option is not
// given in command line arguments.
//
// OnParsed is the field for the event handler which is called when the option
// has been parsed.
// This handler receives a string array which is the option parameter(s) as its
// argument.
// If this field is nil, nothing is done after parsing.
//
// Desc is the field to set the description of the option.
//
// AtParam is a display at a parameter position of this option in a help text.
// This string is for a display like: -o, --option <value>.
type OptCfg struct {
	Name     string
	Aliases  []string
	HasParam bool
	IsArray  bool
	Default  []string
	OnParsed *func([]string) error
	Desc     string
	AtParam  string
}

// ParseWith is a function which parses command line arguments with option
// configurations.
// This function divides command line arguments to command parameters and
// options, and an option consists of a name and option parameters.
// Options are divided to long format options and short format options.
// About long/short format options, since they are same with Parse function,
// see the comment of the function.
//
// This function allows only options declared in option configurations.
// A option configuration has fields: Name, Aliases, HasParam, IsArray, and
// Default.
// When an option matches Name or includes in Aliases in an option
// configuration, the option is registered in Args with the Name.
// If both HasParam and IsArray are true, the option can has one or multiple
// option parameters, and if HasParam is true and IsArray is false, the option
// can has only one option parameter, otherwise the option cannot have option
// parameter.
// If Default is specified and the option is not given in command line
// arguments, the value of Default is set to the option parameter.
//
// If options not declared in option configurations are given in command line
// arguments, this function basically returns UnconfiguredOption error.
// If you want to allow other options, add an option configuration of which
// Name is "*" (but HasParam and IsArray of this configuration is ignored).
func ParseWith(args []string, optCfgs []OptCfg) (Args, error) {
	hasAnyOpt := false
	cfgMap := make(map[string]int)
	for i, cfg := range optCfgs {
		if !cfg.HasParam {
			if cfg.IsArray {
				err := ConfigIsArrayButHasNoParam{Option: cfg.Name}
				return Args{cmdParams: empty}, err
			}
			if cfg.Default != nil {
				err := ConfigHasDefaultButHasNoParam{Option: cfg.Name}
				return Args{cmdParams: empty}, err
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

	var takeParam = func(opt string) bool {
		i, exists := cfgMap[opt]
		if exists {
			return optCfgs[i].HasParam
		}
		return false
	}

	var cmdParams = make([]string, 0)
	var optParams = make(map[string][]string)

	var collCmdParams = func(params ...string) error {
		cmdParams = append(cmdParams, params...)
		return nil
	}
	var collOptParams = func(opt string, params ...string) error {
		i, exists := cfgMap[opt]
		if !exists {
			if !hasAnyOpt {
				return UnconfiguredOption{Option: opt}
			}

			arr := optParams[opt]
			if arr == nil {
				arr = empty
			}
			optParams[opt] = append(arr, params...)
			return nil
		}

		cfg := optCfgs[i]
		if !cfg.HasParam {
			if len(params) > 0 {
				return OptionTakesNoParam{Option: cfg.Name}
			}
		} else {
			if len(params) == 0 {
				return OptionNeedsParam{Option: cfg.Name}
			}
		}

		arr := optParams[cfg.Name]
		if arr == nil {
			arr = empty
		}
		arr = append(arr, params...)

		if !cfg.IsArray {
			if len(arr) > 1 {
				return OptionIsNotArray{Option: cfg.Name}
			}
		}

		optParams[cfg.Name] = arr
		return nil
	}

	err := parseArgs(args, collCmdParams, collOptParams, takeParam)
	if err != nil {
		return Args{cmdParams: empty}, err
	}

	for _, cfg := range optCfgs {
		arr, exists := optParams[cfg.Name]
		if !exists && cfg.Default != nil {
			arr = cfg.Default
			optParams[cfg.Name] = arr
		}
		if cfg.OnParsed != nil {
			err = (*cfg.OnParsed)(arr)
			if err != nil {
				return Args{cmdParams: empty}, err
			}
		}
	}

	return Args{cmdParams: cmdParams, optParams: optParams}, err
}
