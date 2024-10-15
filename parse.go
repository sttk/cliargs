// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"strings"
	"unicode"

	"github.com/sttk/cliargs/errors"
)

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

// Parse is the function to parse command line arguments without configurations.
// This function divides command line arguments to command arguments, which are not associated with
// any options, and options, of which each has a name and option arguments.
// If an option appears multiple times in command line arguments, the option has multiple option
// arguments.
// Options are divided to long format options and short format options.
//
// A long format option starts with "--" and follows multiple characters which consists of
// alphabets, numbers, and '-'.
// (A character immediately after the heading "--" allows only an alphabet.)
// A long format option can be followed by "=" and its option argument.
//
// A short format option starts with "-" and follows single character which is an alphabet.
// Multiple short options can be combined into one argument.
// (For example -a -b -c can be combined into -abc.)
// Moreover, a short option can be followed by "=" and its option argument.
// In case of combined short options, only the last short option can take an option argument.
// (For example, -abc=3 is equal to -a -b -c=3.)
func (cmd *Cmd) Parse() error {
	var collectArgs = func(a string) {
		cmd.Args = append(cmd.Args, a)
	}
	var collectOpts = func(name string, a ...string) error {
		arr, exists := cmd.opts[name]
		if !exists {
			arr = empty
		}
		cmd.opts[name] = append(arr, a...)
		return nil
	}

	_, _, err := parseArgs(cmd._args, collectArgs, collectOpts, takeOptArgs, false, cmd.isAfterEndOpt)
	return err
}

// ParseUntilSubCmd is the method that parses command line arguments without configurations but
// stops parsing when encountering first command argument.
//
// This method creates and returns a new Cmd instance that holds the command line arguments
// starting from the first command argument.
//
// This method parses command line arguments in the same way as the Cmd#parse method, except that
// it only parses the command line arguments before the first command argument.
func (cmd *Cmd) ParseUntilSubCmd() (Cmd, error) {
	var collectArgs = func(_arg string) {}

	var collectOpts = func(name string, a ...string) error {
		arr, exists := cmd.opts[name]
		if !exists {
			arr = empty
		}
		cmd.opts[name] = append(arr, a...)
		return nil
	}

	idx, isAfterEndOpt, err := parseArgs(
		cmd._args, collectArgs, collectOpts, takeOptArgs, true, cmd.isAfterEndOpt)
	if idx < 0 {
		return Cmd{}, err
	}
	return cmd.subCmd(idx, isAfterEndOpt), err
}

func takeOptArgs(_opt string) bool {
	return false
}

func parseArgs(
	osArgs []string,
	collectArgs func(string),
	collectOpts func(string, ...string) error,
	takeOptArgs func(string) bool,
	untilFirstArg bool,
	isAfterEndOpt bool,
) (int, bool, error) {

	prevOptTakingArgs := ""
	var firstErr error = nil

L0:
	for iArg, arg := range osArgs {
		if isAfterEndOpt {
			if untilFirstArg {
				return iArg, isAfterEndOpt, firstErr
			}
			collectArgs(arg)

		} else if len(prevOptTakingArgs) > 0 {
			err := collectOpts(prevOptTakingArgs, arg)
			prevOptTakingArgs = ""
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue L0
			}
		} else if strings.HasPrefix(arg, "--") {
			if len(arg) == 2 {
				isAfterEndOpt = true
				continue L0
			}

			arg = arg[2:]
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						rr := []rune(arg)
						err := collectOpts(string(rr[0:i]), string(rr[i+1:]))
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
							firstErr = errors.OptionContainsInvalidChar{Option: arg}
						}
						continue L0
					}
				} else {
					if !unicode.Is(rangeOfAlphabets, r) {
						if firstErr == nil {
							firstErr = errors.OptionContainsInvalidChar{Option: arg}
						}
						continue L0
					}
				}
				i++
			}

			if i == len(arg) {
				if takeOptArgs(arg) && iArg < len(osArgs)-1 {
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
				if untilFirstArg {
					return iArg, isAfterEndOpt, firstErr
				}
				collectArgs(arg)
				continue L0
			}

			arg := arg[1:]
			var name string
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						if len(name) > 0 {
							rr := []rune(arg)
							err := collectOpts(name, string(rr[i+1:]))
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
						firstErr = errors.OptionContainsInvalidChar{Option: string(r)}
					}
					name = ""
				} else {
					name = string(r)
				}
				i++
			}

			if i == len(arg) && len(name) > 0 {
				if takeOptArgs(name) && iArg < len(osArgs)-1 {
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
			if untilFirstArg {
				return iArg, isAfterEndOpt, firstErr
			}
			collectArgs(arg)
		}
	}

	return -1, isAfterEndOpt, firstErr
}
