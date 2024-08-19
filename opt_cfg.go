// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

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
	Validator *func(string, string, string) error
	Desc      string
	ArgInHelp string
	onParsed  *func([]string) error
}
