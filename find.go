// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"strings"
)

// FindFirstArg is a function which returns an index, a name, a existent flag
// of first non option-format element in a specified string array.
// If non option-format element is found, a existent flag is true, but if
// the element is not found, the flag is false.
func FindFirstArg(osArgs []string) (index int, arg string, exists bool) {
	isNonOpt := false
	if len(osArgs) > 0 {
		for i, a := range osArgs[1:] {
			if isNonOpt {
				return i + 1, a, true
			} else if a == "--" {
				isNonOpt = true
				continue
			} else if !strings.HasPrefix(a, "-") {
				return i + 1, a, true
			}
		}
	}
	return -1, "", false
}
