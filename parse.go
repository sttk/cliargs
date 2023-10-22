// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

// OptionHasInvalidChar is an error which indicates that an invalid character
// is found in the option.
type OptionHasInvalidChar struct{ Option string }

func (e OptionHasInvalidChar) Error() string {
	return fmt.Sprintf("OptionHasInvalidChar{Option:%s}", e.Option)
}

var (
	empty            = make([]string, 0)
	rangeOfAlphabets = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0041, 0x005a, 1}, // A-Z
			{0x0061, 0x007a, 1}, // a-z
		},
	}
	rangeOfAlNumMarks = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x002d, 0x002d, 1}, // -
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005a, 1}, // A-Z
			{0x0061, 0x007a, 1}, // a-z
		},
	}
)

// Cmd is a structure which contains a command name, command arguments, and
// option arguments that are parsed from command line arguments without
// configurations.
// And this provides methods to check if they are specified or to obtain them.
type Cmd struct {
	Name string
	args []string
	opts map[string][]string
}

// HasOpt is a method which checks if the option is specified in command line
// arguments.
func (cmd Cmd) HasOpt(name string) bool {
	_, exists := cmd.opts[name]
	return exists
}

// OptArg is a method to get a option argument which is firstly specified
// with opt in command line arguments.
func (cmd Cmd) OptArg(name string) string {
	arr := cmd.opts[name]
	// If no entry, map returns a nil slice.
	// If a value of a found entry is an empty slice.
	// Both returned values are zero length in common.
	if len(arr) == 0 {
		return ""
	} else {
		return arr[0]
	}
}

// OptArgs is a method to get option arguments which are all specified with
// name in command line arguments.
func (cmd Cmd) OptArgs(name string) []string {
	return cmd.opts[name]
}

// Args is a method to get command arguments which are specified in command
// line arguments and are not associated with any options.
func (cmd Cmd) Args() []string {
	return cmd.args
}

// Parse is a function to parse command line arguments without configurations.
// This function divides command line arguments to command arguments, which
// are not associated with any options, and options, of which each has a name
// and option arguments.
// If an option appears multiple times in command line arguments, the option
// has multiple option arguments.
// Options are divided to long format options and short format options.
//
// A long format option starts with "--" and follows multiple characters which
// consists of alphabets, numbers, and '-'.
// (A character immediately after the heading "--" allows only an alphabet.)
// A long format option can be followed by "=" and its option argument.
//
// A short format option starts with "-" and follows single character which is
// an alphabet.
// Multiple short options can be combined into one argument.
// (For example -a -b -c can be combined into -abc.)
// Moreover, a short option can be followed by "=" and its option argument.
// In case of combined short options, only the last short option can take an
// option argument.
// (For example, -abc=3 is equal to -a -b -c=3.)
func Parse() (Cmd, error) {
	var args = make([]string, 0)
	var opts = make(map[string][]string)

	var collectArgs = func(a ...string) error {
		args = append(args, a...)
		return nil
	}
	var collectOpts = func(name string, a ...string) error {
		arr, exists := opts[name]
		if !exists {
			arr = empty
		}
		opts[name] = append(arr, a...)
		return nil
	}

	var cmdName string
	if len(os.Args) > 0 {
		cmdName = path.Base(os.Args[0])
	}

	var osArgs1 []string
	if len(os.Args) > 1 {
		osArgs1 = os.Args[1:]
	}

	err := parseArgs(osArgs1, collectArgs, collectOpts, _false)

	return Cmd{Name: cmdName, args: args, opts: opts}, err
}

func _false(_ string) bool {
	return false
}

func parseArgs(
	osArgs []string,
	collectArgs func(...string) error,
	collectOpts func(string, ...string) error,
	takeArgs func(string) bool,
) error {

	isNonOpt := false
	prevOptTakingArgs := ""
	var firstErr error = nil

L0:
	for iArg, arg := range osArgs {
		if isNonOpt {
			err := collectArgs(arg)
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue L0
			}

		} else if len(prevOptTakingArgs) > 0 {
			err := collectOpts(prevOptTakingArgs, arg)
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue L0
			}
			prevOptTakingArgs = ""

		} else if strings.HasPrefix(arg, "--") {
			if len(arg) == 2 {
				isNonOpt = true
				continue L0
			}

			arg = arg[2:]
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						err := collectOpts(arg[0:i], arg[i+1:])
						if err != nil {
							if firstErr == nil {
								firstErr = err
							}
							continue L0
						}
						break
					}
					if !unicode.Is(rangeOfAlNumMarks, r) {
						if firstErr == nil {
							firstErr = OptionHasInvalidChar{Option: arg}
						}
						continue L0
					}
				} else {
					if !unicode.Is(rangeOfAlphabets, r) {
						if firstErr == nil {
							firstErr = OptionHasInvalidChar{Option: arg}
						}
						continue L0
					}
				}
				i++
			}

			if i == len(arg) {
				if takeArgs(arg) && iArg < len(osArgs)-1 {
					prevOptTakingArgs = arg
					continue L0
				}
				err := collectOpts(arg)
				if err != nil {
					if firstErr == nil {
						firstErr = err
					}
					continue L0
				}
			}

		} else if strings.HasPrefix(arg, "-") {
			if len(arg) == 1 {
				err := collectArgs(arg)
				if err != nil {
					if firstErr == nil {
						firstErr = err
					}
				}
				continue L0
			}

			arg := arg[1:]
			var name string
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						if len(name) > 0 {
							err := collectOpts(name, arg[i+1:])
							if err != nil {
								if firstErr == nil {
									firstErr = err
								}
							}
						}
						continue L0
					}
					if len(name) > 0 {
						err := collectOpts(name)
						if err != nil {
							if firstErr == nil {
								firstErr = err
							}
						}
					}
				}
				if !unicode.Is(rangeOfAlphabets, r) {
					if firstErr == nil {
						firstErr = OptionHasInvalidChar{Option: string(r)}
					}
					name = ""
				} else {
					name = string(r)
				}
				i++
			}

			if i == len(arg) && len(name) > 0 {
				if takeArgs(name) && iArg < len(osArgs)-1 {
					prevOptTakingArgs = name
				} else {
					err := collectOpts(name)
					if err != nil {
						if firstErr == nil {
							firstErr = err
						}
						continue L0
					}
				}
			}

		} else {
			err := collectArgs(arg)
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue L0
			}
		}
	}

	return firstErr
}
