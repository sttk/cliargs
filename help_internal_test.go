package cliargs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeOptTitle_whenCfgHasOnlyOneLongName(t *testing.T) {
	cfg := OptCfg{Names: []string{"foo-bar"}}
	indent, title := makeOptTitle(cfg)

	assert.Equal(t, indent, 0)
	assert.Equal(t, title, "--foo-bar")
}

func TestMakeOptTitle_whenCfgHasOnlyOneShortName(t *testing.T) {
	cfg := OptCfg{Names: []string{"f"}}
	indent, title := makeOptTitle(cfg)

	assert.Equal(t, indent, 0)
	assert.Equal(t, title, "-f")
}

func TestMakeOptTitle_whenCfgHasMultipleName(t *testing.T) {
	cfg := OptCfg{Names: []string{"f", "b", "foo-bar"}}
	indent, title := makeOptTitle(cfg)

	assert.Equal(t, indent, 0)
	assert.Equal(t, title, "-f, -b, --foo-bar")
}

func TestMakeOptTitle_whenCfgHasNoNameButStoreKey(t *testing.T) {
	cfg := OptCfg{StoreKey: "Foo_Bar"}
	indent, title := makeOptTitle(cfg)

	assert.Equal(t, indent, 0)
	assert.Equal(t, title, "--Foo_Bar")
}

func TestMakeOptTitle_whenCfgNamesContainEmptyName(t *testing.T) {
	cfg := OptCfg{Names: []string{"", "f"}}
	indent, title := makeOptTitle(cfg)

	assert.Equal(t, indent, 4)
	assert.Equal(t, title, "-f")

	cfg = OptCfg{Names: []string{"", "f", "", "b", ""}}
	indent, title = makeOptTitle(cfg)

	assert.Equal(t, indent, 4)
	assert.Equal(t, title, "-f,     -b")
}

func TestMakeOptTitle_whenCfgNamesAreAllEmptyAndHasStoreKey(t *testing.T) {
	cfg := OptCfg{StoreKey: "FooBar", Names: []string{"", ""}}
	indent, title := makeOptTitle(cfg)

	assert.Equal(t, indent, 8)
	assert.Equal(t, title, "--FooBar")
}

func TestCreateOptsHelp_whenCfgHasOnlyOneLongName(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{Names: []string{"foo-bar"}},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "--foo-bar")
	assert.Equal(t, indent, 11)
}

func TestCreateOptsHelp_whenCfgHasOnlyOneLongNameAndDesc(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names: []string{"foo-bar"},
			Desc:  "The description of foo-bar.",
		},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "--foo-bar  The description of foo-bar.")
	assert.Equal(t, indent, 11)
}

func TestCreateOptsHelp_whenCfgHasOnlyOneLongNameAndDescAndArgInHelp(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names:     []string{"foo-bar"},
			Desc:      "The description of foo-bar.",
			ArgInHelp: "<num>",
		},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "--foo-bar <num>  The description of foo-bar.")
	assert.Equal(t, indent, 17)
}

func TestCreateOptsHelp_whenCfgHasOnlyOneShortName(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{Names: []string{"f"}},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "-f")
	assert.Equal(t, indent, 4)
}

func TestCreateOptsHelp_whenCfgHasOnlyOneShortNameAndDesc(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names: []string{"f"},
			Desc:  "The description of f.",
		},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "-f  The description of f.")
	assert.Equal(t, indent, 4)
}

func TestCreateOptsHelp_whenCfgHasOnlyOneShortNameAndDescAndArgInHelp(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names:     []string{"f"},
			Desc:      "The description of f.",
			ArgInHelp: "<n>",
		},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "-f <n>  The description of f.")
	assert.Equal(t, indent, 8)
}

func TestCreateOptsHelp_whenIndentIsPositiveAndLongThanTitle(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names:     []string{"foo-bar"},
			Desc:      "The description of foo-bar.",
			ArgInHelp: "<num>",
		},
	}
	indent := 19
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "--foo-bar <num>    The description of foo-bar.")
	assert.Equal(t, indent, 19)
}

func TestCreateOptsHelp_whenIndentIsPositiveAndShorterThanTitle(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names:     []string{"foo-bar"},
			Desc:      "The description of foo-bar.",
			ArgInHelp: "<num>",
		},
	}
	indent := 16
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "--foo-bar <num>\n                The description of foo-bar.")
	assert.Equal(t, indent, 16)

	indent = 10
	blockBodies = createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body = blockBodies[0]
	assert.Equal(t, body.firstIndent, 0)
	assert.Equal(t, body.text, "--foo-bar <num>\n          The description of foo-bar.")
	assert.Equal(t, indent, 10)
}

func TestCreateOptsHelp_whenNamesContainsEmptyStrings(t *testing.T) {
	cfgs := []OptCfg{
		OptCfg{
			Names:     []string{"", "", "f", "", "foo-bar", ""},
			Desc:      "The description of foo-bar.",
			ArgInHelp: "<num>",
		},
	}
	indent := 0
	blockBodies := createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body := blockBodies[0]
	assert.Equal(t, body.firstIndent, 8)
	assert.Equal(t, body.text, "-f,     --foo-bar <num>  The description of foo-bar.")
	assert.Equal(t, indent, 8+25)

	indent = 35 // longer than title width
	blockBodies = createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body = blockBodies[0]
	assert.Equal(t, body.firstIndent, 8)
	assert.Equal(t, body.text, "-f,     --foo-bar <num>    The description of foo-bar.")
	assert.Equal(t, indent, 35)

	indent = 33 // equal to title width
	blockBodies = createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body = blockBodies[0]
	assert.Equal(t, body.firstIndent, 8)
	assert.Equal(t, body.text, "-f,     --foo-bar <num>  The description of foo-bar.")
	assert.Equal(t, indent, 33)

	indent = 32 // shorter than title width
	blockBodies = createOptsHelp(cfgs, &indent)

	assert.Equal(t, len(blockBodies), 1)
	body = blockBodies[0]
	assert.Equal(t, body.firstIndent, 8)
	assert.Equal(t,
		body.text,
		"-f,     --foo-bar <num>\n                                The description of foo-bar.",
	)
	assert.Equal(t, indent, 32)
}
