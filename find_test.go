package cliargs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/cliargs"
)

func TestFindFirstArg_findAtFirstIndex(t *testing.T) {
	osArgs := []string{"app", "foo-bar", "-a", "--baz", "--bcd"}

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	assert.Equal(t, i, 1)
	assert.Equal(t, arg, "foo-bar")
	assert.True(t, exists)
}

func TestFindFirstArg_findAtMiddleIndex(t *testing.T) {
	osArgs := []string{"app", "--corge", "foo-bar", "-a", "--baz", "--bcd"}

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	assert.Equal(t, i, 2)
	assert.Equal(t, arg, "foo-bar")
	assert.True(t, exists)
}

func TestFindFirstArg_findAtLastIndex(t *testing.T) {
	osArgs := []string{"app", "--corge", "--foo-bar", "-a", "baz"}

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	assert.Equal(t, i, 4)
	assert.Equal(t, arg, "baz")
	assert.True(t, exists)
}

func TestFindFirstArg_returnFoundFirst(t *testing.T) {
	osArgs := []string{"app", "--corge", "foo-bar", "-a", "baz", "--bcd"}

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	assert.Equal(t, i, 2)
	assert.Equal(t, arg, "foo-bar")
	assert.True(t, exists)
}

func TestFindFirstArg_notFound(t *testing.T) {
	osArgs := []string{"app", "--corge", "--foo-bar", "-a", "-baz", "--bcd"}

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	assert.Equal(t, i, -1)
	assert.Equal(t, arg, "")
	assert.False(t, exists)
}

func TestFindFirstArg_supportDoubleHyphens(t *testing.T) {
	osArgs := []string{"app", "--", "--foo-bar", "-a", "-baz", "--bcd"}

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	assert.Equal(t, i, 2)
	assert.Equal(t, arg, "--foo-bar")
	assert.True(t, exists)
}
