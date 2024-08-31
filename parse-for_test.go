package cliargs_test

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/cliargs"
	errs "github.com/sttk/cliargs/errors"
)

func TestParseFor_emptyOptionStoreAndNoArgs(t *testing.T) {
	defer reset()

	type MyOptions struct{}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, cmd.OptCfgs, []cliargs.OptCfg{})
}

func TestParseFor_nonEmptyOptionStoreAndNoArgs(t *testing.T) {
	defer reset()

	type MyOptions struct {
		BoolVal    bool
		IntVal     int
		Int8Val    int8
		Int16Val   int16
		Int32Val   int32
		Int64Val   int64
		UintVal    uint
		Uint8Val   uint8
		Uint16Val  uint16
		Uint32Val  uint32
		Uint64Val  uint64
		Float32Val float32
		Float64Val float64
		StringVal  string
		IntArr     []int
		Int8Arr    []int8
		Int16Arr   []int16
		Int32Arr   []int32
		Int64Arr   []int64
		UintArr    []uint
		Uint8Arr   []uint8
		Uint16Arr  []uint16
		Uint32Arr  []uint32
		Uint64Arr  []uint64
		Float32Arr []float32
		Float64Arr []float64
		StringArr  []string
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 27)
	assert.False(t, options.BoolVal)
	assert.Equal(t, options.IntVal, 0)
	assert.Equal(t, options.Int8Val, int8(0))
	assert.Equal(t, options.Int16Val, int16(0))
	assert.Equal(t, options.Int32Val, int32(0))
	assert.Equal(t, options.Int64Val, int64(0))
	assert.Equal(t, options.UintVal, uint(0))
	assert.Equal(t, options.Uint8Val, uint8(0))
	assert.Equal(t, options.Uint16Val, uint16(0))
	assert.Equal(t, options.Uint32Val, uint32(0))
	assert.Equal(t, options.Uint64Val, uint64(0))
	assert.Equal(t, options.Float32Val, float32(0.0))
	assert.Equal(t, options.Float64Val, 0.0)
	assert.Equal(t, options.StringVal, "")
	assert.Equal(t, options.IntArr, []int(nil))
	assert.Equal(t, options.Int8Arr, []int8(nil))
	assert.Equal(t, options.Int16Arr, []int16(nil))
	assert.Equal(t, options.Int32Arr, []int32(nil))
	assert.Equal(t, options.Int64Arr, []int64(nil))
	assert.Equal(t, options.UintArr, []uint(nil))
	assert.Equal(t, options.Uint8Arr, []uint8(nil))
	assert.Equal(t, options.Uint16Arr, []uint16(nil))
	assert.Equal(t, options.Uint32Arr, []uint32(nil))
	assert.Equal(t, options.Uint64Arr, []uint64(nil))
	assert.Equal(t, options.Float32Arr, []float32(nil))
	assert.Equal(t, options.Float64Arr, []float64(nil))
	assert.Equal(t, options.StringArr, []string(nil))
}

func TestParseFor_dontOverwriteOptionsIfNoArgs(t *testing.T) {
	defer reset()

	type MyOptions struct {
		BoolVal    bool
		IntVal     int
		Int8Val    int8
		Int16Val   int16
		Int32Val   int32
		Int64Val   int64
		UintVal    uint
		Uint8Val   uint8
		Uint16Val  uint16
		Uint32Val  uint32
		Uint64Val  uint64
		Float32Val float32
		Float64Val float64
		StringVal  string
		IntArr     []int
		Int8Arr    []int8
		Int16Arr   []int16
		Int32Arr   []int32
		Int64Arr   []int64
		UintArr    []uint
		Uint8Arr   []uint8
		Uint16Arr  []uint16
		Uint32Arr  []uint32
		Uint64Arr  []uint64
		Float32Arr []float32
		Float64Arr []float64
		StringArr  []string
	}
	options := MyOptions{
		BoolVal:    true,
		IntVal:     111,
		Int8Val:    22,
		Int16Val:   333,
		Int32Val:   444,
		Int64Val:   555,
		UintVal:    666,
		Uint8Val:   77,
		Uint16Val:  888,
		Uint32Val:  999,
		Uint64Val:  1111,
		Float32Val: 0.123,
		Float64Val: 0.456789,
		StringVal:  "abcdefg",
		IntArr:     []int{1, 1, 1},
		Int8Arr:    []int8{2, 2},
		Int16Arr:   []int16{3, 3, 3},
		Int32Arr:   []int32{4, 4, 4},
		Int64Arr:   []int64{5, 5, 5},
		UintArr:    []uint{6, 6, 6},
		Uint8Arr:   []uint8{7, 7},
		Uint16Arr:  []uint16{8, 8, 8},
		Uint32Arr:  []uint32{9, 9, 9},
		Uint64Arr:  []uint64{1, 1, 1, 1},
		Float32Arr: []float32{0.1, 2.3},
		Float64Arr: []float64{0.45, 6.789},
		StringArr:  []string{"ab", "cd", "efg"},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 27)
	assert.True(t, options.BoolVal)
	assert.Equal(t, options.IntVal, 111)
	assert.Equal(t, options.Int8Val, int8(22))
	assert.Equal(t, options.Int16Val, int16(333))
	assert.Equal(t, options.Int32Val, int32(444))
	assert.Equal(t, options.Int64Val, int64(555))
	assert.Equal(t, options.UintVal, uint(666))
	assert.Equal(t, options.Uint8Val, uint8(77))
	assert.Equal(t, options.Uint16Val, uint16(888))
	assert.Equal(t, options.Uint32Val, uint32(999))
	assert.Equal(t, options.Uint64Val, uint64(1111))
	assert.Equal(t, options.Float32Val, float32(0.123))
	assert.Equal(t, options.Float64Val, 0.456789)
	assert.Equal(t, options.StringVal, "abcdefg")
	assert.Equal(t, options.IntArr, []int{1, 1, 1})
	assert.Equal(t, options.Int8Arr, []int8{2, 2})
	assert.Equal(t, options.Int16Arr, []int16{3, 3, 3})
	assert.Equal(t, options.Int32Arr, []int32{4, 4, 4})
	assert.Equal(t, options.Int64Arr, []int64{5, 5, 5})
	assert.Equal(t, options.UintArr, []uint{6, 6, 6})
	assert.Equal(t, options.Uint8Arr, []uint8{7, 7})
	assert.Equal(t, options.Uint16Arr, []uint16{8, 8, 8})
	assert.Equal(t, options.Uint32Arr, []uint32{9, 9, 9})
	assert.Equal(t, options.Uint64Arr, []uint64{1, 1, 1, 1})
	assert.Equal(t, options.Float32Arr, []float32{0.1, 2.3})
	assert.Equal(t, options.Float64Arr, []float64{0.45, 6.789})
	assert.Equal(t, options.StringArr, []string{"ab", "cd", "efg"})
}

func TestParseFor_optionIsBoolAndArgIsName(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Flag bool `optcfg:"flag,f"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app", "--flag", "abc"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.True(t, cmd.HasOpt("Flag"))
	assert.Equal(t, cmd.OptArg("Flag"), "")
	assert.Equal(t, cmd.OptArgs("Flag"), []string(nil))
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.True(t, options.Flag)
}

func TestParseFor_optionIsBoolAndArgIsAlias(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Flag bool `optcfg:"flag,f"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app", "-f", "abc"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.True(t, options.Flag)
}

func TestParseFor_optionsAreIntAndArgsAreNames(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntVal   int   `optcfg:"int-val,i"`
		Int8Val  int8  `optcfg:"int8-val,j"`
		Int16Val int16 `optcfg:"int16-val,k"`
		Int32Val int32 `optcfg:"int32-val,m"`
		Int64Val int64 `optcfg:"int64-val,n"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"--int-val", "1",
		"--int8-val", "2",
		"--int16-val", "3",
		"--int32-val", "4",
		"--int64-val", "5",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntVal, 1)
	assert.Equal(t, options.Int8Val, int8(2))
	assert.Equal(t, options.Int16Val, int16(3))
	assert.Equal(t, options.Int32Val, int32(4))
	assert.Equal(t, options.Int64Val, int64(5))
}

func TestParseFor_optionsAreIntAndArgsAreAliases(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntVal   int   `optcfg:"int-val,i"`
		Int8Val  int8  `optcfg:"int8-val,j"`
		Int16Val int16 `optcfg:"int16-val,k"`
		Int32Val int32 `optcfg:"int32-val,m"`
		Int64Val int64 `optcfg:"int64-val,n"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"-i", "1",
		"-j", "2",
		"-k", "3",
		"-m", "4",
		"-n", "5",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntVal, 1)
	assert.Equal(t, options.Int8Val, int8(2))
	assert.Equal(t, options.Int16Val, int16(3))
	assert.Equal(t, options.Int32Val, int32(4))
	assert.Equal(t, options.Int64Val, int64(5))
}

func TestParseFor_optionsAreUintAndArgsAreNames(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintVal   uint   `optcfg:"uint-val,i"`
		Uint8Val  uint8  `optcfg:"uint8-val,j"`
		Uint16Val uint16 `optcfg:"uint16-val,k"`
		Uint32Val uint32 `optcfg:"uint32-val,m"`
		Uint64Val uint64 `optcfg:"uint64-val,n"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"--uint-val", "1",
		"--uint8-val", "2",
		"--uint16-val", "3",
		"--uint32-val", "4",
		"--uint64-val", "5",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintVal, uint(1))
	assert.Equal(t, options.Uint8Val, uint8(2))
	assert.Equal(t, options.Uint16Val, uint16(3))
	assert.Equal(t, options.Uint32Val, uint32(4))
	assert.Equal(t, options.Uint64Val, uint64(5))
}

func TestParseFor_optionsAreUintAndArgsAreAliases(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintVal   uint   `optcfg:"uint-val,i"`
		Uint8Val  uint8  `optcfg:"uint8-val,j"`
		Uint16Val uint16 `optcfg:"uint16-val,k"`
		Uint32Val uint32 `optcfg:"uint32-val,m"`
		Uint64Val uint64 `optcfg:"uint64-val,n"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"-i", "1",
		"-j", "2",
		"-k", "3",
		"-m", "4",
		"-n", "5",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintVal, uint(1))
	assert.Equal(t, options.Uint8Val, uint8(2))
	assert.Equal(t, options.Uint16Val, uint16(3))
	assert.Equal(t, options.Uint32Val, uint32(4))
	assert.Equal(t, options.Uint64Val, uint64(5))
}

func TestParseFor_optionsAreFloatAndArgsAreNames(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Val float32 `optcfg:"float32-val,m"`
		Float64Val float64 `optcfg:"float64-val,n"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"--float32-val", "0.1234",
		"--float64-val", "0.5678",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Val, float32(0.1234))
	assert.Equal(t, options.Float64Val, 0.5678)
}

func TestParseFor_optionsAreFloatAndArgsAreAliases(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Val float32 `optcfg:"float32-val,m"`
		Float64Val float64 `optcfg:"float64-val,n"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"-m", "0.1234",
		"-n", "0.5678",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Val, float32(0.1234))
	assert.Equal(t, options.Float64Val, 0.5678)
}

func TestParseFor_optionsAreStringAndArgsAreNames(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringVal string `optcfg:"string-val,s"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"--string-val", "def",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringVal, "def")
}

func TestParseFor_optionsAreStringAndArgsAreAliases(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringVal string `optcfg:"string-val,s"`
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"-s", "def",
		"abc",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"abc"})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringVal, "def")
}

func TestParseFor_defaultValueIsInt(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntVal   int   `optcfg:"=11"`
		Int8Val  int8  `optcfg:"=22"`
		Int16Val int16 `optcfg:"=33"`
		Int32Val int32 `optcfg:"=44"`
		Int64Val int64 `optcfg:"=55"`
	}
	options := MyOptions{}

	os.Args = []string{"./app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntVal, 11)
	assert.Equal(t, options.Int8Val, int8(22))
	assert.Equal(t, options.Int16Val, int16(33))
	assert.Equal(t, options.Int32Val, int32(44))
	assert.Equal(t, options.Int64Val, int64(55))
}

func TestParseFor_defaultValueIsNegativeInt(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntVal   int   `optcfg:"=-11"`
		Int8Val  int8  `optcfg:"=-22"`
		Int16Val int16 `optcfg:"=-33"`
		Int32Val int32 `optcfg:"=-44"`
		Int64Val int64 `optcfg:"=-55"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntVal, -11)
	assert.Equal(t, options.Int8Val, int8(-22))
	assert.Equal(t, options.Int16Val, int16(-33))
	assert.Equal(t, options.Int32Val, int32(-44))
	assert.Equal(t, options.Int64Val, int64(-55))
}

func TestParseFor_defaultValueIsUint(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintVal   uint   `optcfg:"=11"`
		Uint8Val  uint8  `optcfg:"=22"`
		Uint16Val uint16 `optcfg:"=33"`
		Uint32Val uint32 `optcfg:"=44"`
		Uint64Val uint64 `optcfg:"=55"`
	}
	options := MyOptions{}

	os.Args = []string{"./app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintVal, uint(11))
	assert.Equal(t, options.Uint8Val, uint8(22))
	assert.Equal(t, options.Uint16Val, uint16(33))
	assert.Equal(t, options.Uint32Val, uint32(44))
	assert.Equal(t, options.Uint64Val, uint64(55))
}

func TestParseFor_defaultValueIsFloat(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Val float32 `optcfg:"=0.123"`
		Float64Val float64 `optcfg:"=0.456789"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Val, float32(0.123))
	assert.Equal(t, options.Float64Val, float64(0.456789))
}

func TestParseFor_defaultValueIsNegativeFloat(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Val float32 `optcfg:"=-0.123"`
		Float64Val float64 `optcfg:"=-0.456789"`
	}
	options := MyOptions{}

	os.Args = []string{"app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Val, float32(-0.123))
	assert.Equal(t, options.Float64Val, float64(-0.456789))
}

func TestParseFor_defaultValueIsIntArrayAndSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[]"`
		Int8Arr  []int8  `optcfg:"=[]"`
		Int16Arr []int16 `optcfg:"=[]"`
		Int32Arr []int32 `optcfg:"=[]"`
		Int64Arr []int64 `optcfg:"=[]"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{})
	assert.Equal(t, options.Int8Arr, []int8{})
	assert.Equal(t, options.Int16Arr, []int16{})
	assert.Equal(t, options.Int32Arr, []int32{})
	assert.Equal(t, options.Int64Arr, []int64{})
}

func TestParseFor_overwriteIntArrayWithDefaultValueIfSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[]"`
		Int8Arr  []int8  `optcfg:"=[]"`
		Int16Arr []int16 `optcfg:"=[]"`
		Int32Arr []int32 `optcfg:"=[]"`
		Int64Arr []int64 `optcfg:"=[]"`
	}
	options := MyOptions{
		IntArr:   []int{1},
		Int8Arr:  []int8{2},
		Int16Arr: []int16{3},
		Int32Arr: []int32{4},
		Int64Arr: []int64{5},
	}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{})
	assert.Equal(t, options.Int8Arr, []int8{})
	assert.Equal(t, options.Int16Arr, []int16{})
	assert.Equal(t, options.Int32Arr, []int32{})
	assert.Equal(t, options.Int64Arr, []int64{})
}

func TestParseFor_defaultValueIsIntArrayAndSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[1]"`
		Int8Arr  []int8  `optcfg:"=[2]"`
		Int16Arr []int16 `optcfg:"=[3]"`
		Int32Arr []int32 `optcfg:"=[4]"`
		Int64Arr []int64 `optcfg:"=[5]"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{1})
	assert.Equal(t, options.Int8Arr, []int8{2})
	assert.Equal(t, options.Int16Arr, []int16{3})
	assert.Equal(t, options.Int32Arr, []int32{4})
	assert.Equal(t, options.Int64Arr, []int64{5})
}

func TestParseFor_overwriteIntArrayWithDefaultValueIfSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[1]"`
		Int8Arr  []int8  `optcfg:"=[2]"`
		Int16Arr []int16 `optcfg:"=[3]"`
		Int32Arr []int32 `optcfg:"=[4]"`
		Int64Arr []int64 `optcfg:"=[5]"`
	}
	options := MyOptions{
		IntArr:   []int{11},
		Int8Arr:  []int8{22},
		Int16Arr: []int16{33},
		Int32Arr: []int32{44},
		Int64Arr: []int64{55},
	}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{1})
	assert.Equal(t, options.Int8Arr, []int8{2})
	assert.Equal(t, options.Int16Arr, []int16{3})
	assert.Equal(t, options.Int32Arr, []int32{4})
	assert.Equal(t, options.Int64Arr, []int64{5})
}

func TestParseFor_defaultValueIsIntArrayAndSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[1,2]"`
		Int8Arr  []int8  `optcfg:"=[2,3]"`
		Int16Arr []int16 `optcfg:"=[3,4]"`
		Int32Arr []int32 `optcfg:"=[4,5]"`
		Int64Arr []int64 `optcfg:"=[5,6]"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{1, 2})
	assert.Equal(t, options.Int8Arr, []int8{2, 3})
	assert.Equal(t, options.Int16Arr, []int16{3, 4})
	assert.Equal(t, options.Int32Arr, []int32{4, 5})
	assert.Equal(t, options.Int64Arr, []int64{5, 6})
}

func TestParseFor_overwriteIntArrayWithDefaultValueIfSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[1,2]"`
		Int8Arr  []int8  `optcfg:"=[2,3]"`
		Int16Arr []int16 `optcfg:"=[3,4]"`
		Int32Arr []int32 `optcfg:"=[4,5]"`
		Int64Arr []int64 `optcfg:"=[5,6]"`
	}
	options := MyOptions{
		IntArr:   []int{11},
		Int8Arr:  []int8{22},
		Int16Arr: []int16{33},
		Int32Arr: []int32{44},
		Int64Arr: []int64{55},
	}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{1, 2})
	assert.Equal(t, options.Int8Arr, []int8{2, 3})
	assert.Equal(t, options.Int16Arr, []int16{3, 4})
	assert.Equal(t, options.Int32Arr, []int32{4, 5})
	assert.Equal(t, options.Int64Arr, []int64{5, 6})
}

func TestParseFor_defaultValueIsNegativeIntArray(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=[-1,-2]"`
		Int8Arr  []int8  `optcfg:"=[-2,-3]"`
		Int16Arr []int16 `optcfg:"=[-3,-4]"`
		Int32Arr []int32 `optcfg:"=[-4,-5]"`
		Int64Arr []int64 `optcfg:"=[-5,-6]"`
	}
	options := MyOptions{}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{-1, -2})
	assert.Equal(t, options.Int8Arr, []int8{-2, -3})
	assert.Equal(t, options.Int16Arr, []int16{-3, -4})
	assert.Equal(t, options.Int32Arr, []int32{-4, -5})
	assert.Equal(t, options.Int64Arr, []int64{-5, -6})
}

func TestParseFor_defaultValueIsIntArraySeparatedByColons(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr   []int   `optcfg:"=:[-1:-2]"`
		Int8Arr  []int8  `optcfg:"=/[-2/-3]"`
		Int16Arr []int16 `optcfg:"=![-3!-4]"`
		Int32Arr []int32 `optcfg:"=|[-4|-5]"`
		Int64Arr []int64 `optcfg:"='[-5'-6]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.IntArr, []int{-1, -2})
	assert.Equal(t, options.Int8Arr, []int8{-2, -3})
	assert.Equal(t, options.Int16Arr, []int16{-3, -4})
	assert.Equal(t, options.Int32Arr, []int32{-4, -5})
	assert.Equal(t, options.Int64Arr, []int64{-5, -6})
}

func TestParseFor_defaultValueIsUintArrayAndSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=[]"`
		Uint8Arr  []uint8  `optcfg:"=[]"`
		Uint16Arr []uint16 `optcfg:"=[]"`
		Uint32Arr []uint32 `optcfg:"=[]"`
		Uint64Arr []uint64 `optcfg:"=[]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{})
	assert.Equal(t, options.Uint8Arr, []uint8{})
	assert.Equal(t, options.Uint16Arr, []uint16{})
	assert.Equal(t, options.Uint32Arr, []uint32{})
	assert.Equal(t, options.Uint64Arr, []uint64{})
}

func TestParseFor_overwriteUintArrayWithDefaultValueIfSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=[]"`
		Uint8Arr  []uint8  `optcfg:"=[]"`
		Uint16Arr []uint16 `optcfg:"=[]"`
		Uint32Arr []uint32 `optcfg:"=[]"`
		Uint64Arr []uint64 `optcfg:"=[]"`
	}
	options := MyOptions{
		UintArr:   []uint{1},
		Uint8Arr:  []uint8{2},
		Uint16Arr: []uint16{3},
		Uint32Arr: []uint32{4},
		Uint64Arr: []uint64{5},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{})
	assert.Equal(t, options.Uint8Arr, []uint8{})
	assert.Equal(t, options.Uint16Arr, []uint16{})
	assert.Equal(t, options.Uint32Arr, []uint32{})
	assert.Equal(t, options.Uint64Arr, []uint64{})
}

func TestParseFor_defaultValueIsUintArrayAndSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=[1]"`
		Uint8Arr  []uint8  `optcfg:"=[2]"`
		Uint16Arr []uint16 `optcfg:"=[3]"`
		Uint32Arr []uint32 `optcfg:"=[4]"`
		Uint64Arr []uint64 `optcfg:"=[5]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{1})
	assert.Equal(t, options.Uint8Arr, []uint8{2})
	assert.Equal(t, options.Uint16Arr, []uint16{3})
	assert.Equal(t, options.Uint32Arr, []uint32{4})
	assert.Equal(t, options.Uint64Arr, []uint64{5})
}

func TestParseFor_overwriteUintArrayWithDefaultValueIfSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=[1]"`
		Uint8Arr  []uint8  `optcfg:"=[2]"`
		Uint16Arr []uint16 `optcfg:"=[3]"`
		Uint32Arr []uint32 `optcfg:"=[4]"`
		Uint64Arr []uint64 `optcfg:"=[5]"`
	}
	options := MyOptions{
		UintArr:   []uint{11},
		Uint8Arr:  []uint8{22},
		Uint16Arr: []uint16{33},
		Uint32Arr: []uint32{44},
		Uint64Arr: []uint64{55},
	}

	os.Args = []string{"path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{1})
	assert.Equal(t, options.Uint8Arr, []uint8{2})
	assert.Equal(t, options.Uint16Arr, []uint16{3})
	assert.Equal(t, options.Uint32Arr, []uint32{4})
	assert.Equal(t, options.Uint64Arr, []uint64{5})
}

func TestParseFor_defaultValueIsUintArrayAndSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=[1,2]"`
		Uint8Arr  []uint8  `optcfg:"=[2,3]"`
		Uint16Arr []uint16 `optcfg:"=[3,4]"`
		Uint32Arr []uint32 `optcfg:"=[4,5]"`
		Uint64Arr []uint64 `optcfg:"=[5,6]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{1, 2})
	assert.Equal(t, options.Uint8Arr, []uint8{2, 3})
	assert.Equal(t, options.Uint16Arr, []uint16{3, 4})
	assert.Equal(t, options.Uint32Arr, []uint32{4, 5})
	assert.Equal(t, options.Uint64Arr, []uint64{5, 6})
}

func TestParseFor_overwriteUintArrayWithDefaultValueIfSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=[1,2]"`
		Uint8Arr  []uint8  `optcfg:"=[2,3]"`
		Uint16Arr []uint16 `optcfg:"=[3,4]"`
		Uint32Arr []uint32 `optcfg:"=[4,5]"`
		Uint64Arr []uint64 `optcfg:"=[5,6]"`
	}
	options := MyOptions{
		UintArr:   []uint{11},
		Uint8Arr:  []uint8{22},
		Uint16Arr: []uint16{33},
		Uint32Arr: []uint32{44},
		Uint64Arr: []uint64{55},
	}

	os.Args = []string{"./app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{1, 2})
	assert.Equal(t, options.Uint8Arr, []uint8{2, 3})
	assert.Equal(t, options.Uint16Arr, []uint16{3, 4})
	assert.Equal(t, options.Uint32Arr, []uint32{4, 5})
	assert.Equal(t, options.Uint64Arr, []uint64{5, 6})
}

func TestParseFor_defaultValueIsUintArraySeparatedByColons(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr   []uint   `optcfg:"=:[1:2]"`
		Uint8Arr  []uint8  `optcfg:"=/[2/3]"`
		Uint16Arr []uint16 `optcfg:"=![3!4]"`
		Uint32Arr []uint32 `optcfg:"=|[4|5]"`
		Uint64Arr []uint64 `optcfg:"='[5'6]"`
	}
	options := MyOptions{}

	os.Args = []string{"app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.Equal(t, options.UintArr, []uint{1, 2})
	assert.Equal(t, options.Uint8Arr, []uint8{2, 3})
	assert.Equal(t, options.Uint16Arr, []uint16{3, 4})
	assert.Equal(t, options.Uint32Arr, []uint32{4, 5})
	assert.Equal(t, options.Uint64Arr, []uint64{5, 6})
}

func TestParseFor_defaultValueIsFloatArrayAndSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[]"`
		Float64Arr []float64 `optcfg:"=[]"`
	}
	options := MyOptions{}

	os.Args = []string{"./app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{})
	assert.Equal(t, options.Float64Arr, []float64{})
}

func TestParseFor_overwriteFloatArrayWithDefaultValueIfSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[]"`
		Float64Arr []float64 `optcfg:"=[]"`
	}
	options := MyOptions{
		Float32Arr: []float32{0.999},
		Float64Arr: []float64{0.888},
	}

	os.Args = []string{"./app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{})
	assert.Equal(t, options.Float64Arr, []float64{})
}

func TestParseFor_defaultValueIsFloatArrayAndSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[0.1]"`
		Float64Arr []float64 `optcfg:"=[0.2]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{0.1})
	assert.Equal(t, options.Float64Arr, []float64{0.2})
}

func TestParseFor_overwriteFloatArrayWithDefaultValueIfSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[0.1]"`
		Float64Arr []float64 `optcfg:"=[0.2]"`
	}
	options := MyOptions{
		Float32Arr: []float32{0.99},
		Float64Arr: []float64{0.88},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{0.1})
	assert.Equal(t, options.Float64Arr, []float64{0.2})
}

func TestParseFor_defaultValueIsFloatArrayAndSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[0.1,0.2]"`
		Float64Arr []float64 `optcfg:"=[0.3,0.4]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{0.1, 0.2})
	assert.Equal(t, options.Float64Arr, []float64{0.3, 0.4})
}

func TestParseFor_overwriteFloatArrayWithDefaultValueIfSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[0.1,0.2]"`
		Float64Arr []float64 `optcfg:"=[0.3,0.4]"`
	}
	options := MyOptions{
		Float32Arr: []float32{0.99},
		Float64Arr: []float64{0.88},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{0.1, 0.2})
	assert.Equal(t, options.Float64Arr, []float64{0.3, 0.4})
}

func TestParseFor_defaultValueIsNegativeFloatArray(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=[-0.1,-0.2]"`
		Float64Arr []float64 `optcfg:"=[-0.3,-0.4]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{-0.1, -0.2})
	assert.Equal(t, options.Float64Arr, []float64{-0.3, -0.4})
}

func TestParseFor_defaultValueIsFloatArraySeparatedByColons(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float32Arr []float32 `optcfg:"=|[-0.1|-0.2]"`
		Float64Arr []float64 `optcfg:"='[-0.3'-0.4]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 2)
	assert.Equal(t, options.Float32Arr, []float32{-0.1, -0.2})
	assert.Equal(t, options.Float64Arr, []float64{-0.3, -0.4})
}

func TestParseFor_defaultValueIsStringAndSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=[]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{})
}

func TestParseFor_overwriteStringArrayWithDefaultValueIfSize0(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=[]"`
	}
	options := MyOptions{
		StringArr: []string{"ZZZ"},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{})
}

func TestParseFor_defaultValueIsStringArrayAndSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=[ABC]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{"ABC"})
}

func TestParseFor_overwriteStringArrayWithDefaultValueIfSize1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=[ABC]"`
	}
	options := MyOptions{
		StringArr: []string{"ZZZ"},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{"ABC"})
}

func TestParseFor_defaultValueIsStringArrayAndSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=[ABC,DEF]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{"ABC", "DEF"})
}

func TestParseFor_overwriteStringArrayWithDefaultValueIfSize2(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=[ABC,DEF]"`
	}
	options := MyOptions{
		StringArr: []string{"ZZZ"},
	}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{"ABC", "DEF"})
}

func TestParseFor_defaultValueIsStringArraySeparatedByColons(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"=|[ABC|DEF]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{"ABC", "DEF"})
}

func TestParseFor_ignoreEmptyDefaultValueIfOptionIsBool(t *testing.T) {
	defer reset()

	type MyOptions struct {
		BoolVar bool `optcfg:"="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.False(t, options.BoolVar)
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsInt(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntVar int `optcfg:"int-var="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{"+
		"StoreKey:IntVar,Option:int-var,OptArg:,TypeKind:int,"+
		"Cause:strconv.ParseInt: parsing \"\": invalid syntax}",
	)
	assert.NotNil(t, errors.Unwrap(err))
	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "IntVar")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "int-var")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Int)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsUint(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintVar uint `optcfg:"uint-var="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{"+
		"StoreKey:UintVar,Option:uint-var,OptArg:,TypeKind:uint,"+
		"Cause:strconv.ParseUint: parsing \"\": invalid syntax}",
	)
	assert.NotNil(t, errors.Unwrap(err))
	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "UintVar")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "uint-var")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Uint)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsFloat(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float64Var float64 `optcfg:"float-var="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{"+
		"StoreKey:Float64Var,Option:float-var,OptArg:,TypeKind:float64,"+
		"Cause:strconv.ParseFloat: parsing \"\": invalid syntax}",
	)
	assert.NotNil(t, errors.Unwrap(err))
	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "Float64Var")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "float-var")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Float64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsString(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringVar string `optcfg:"str-var="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringVar, "")
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsIntArray(t *testing.T) {
	defer reset()

	type MyOptions struct {
		IntArr []int `optcfg:"int-arr="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{"+
		"StoreKey:IntArr,Option:int-arr,OptArg:,TypeKind:int,"+
		"Cause:strconv.ParseInt: parsing \"\": invalid syntax}",
	)
	assert.NotNil(t, errors.Unwrap(err))
	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "IntArr")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "int-arr")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Int)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsUintArray(t *testing.T) {
	defer reset()

	type MyOptions struct {
		UintArr []uint `optcfg:"uint-arr="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{"+
		"StoreKey:UintArr,Option:uint-arr,OptArg:,TypeKind:uint,"+
		"Cause:strconv.ParseUint: parsing \"\": invalid syntax}",
	)
	assert.NotNil(t, errors.Unwrap(err))
	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "UintArr")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "uint-arr")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Uint)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsFloatArray(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Float64Arr []float64 `optcfg:"float-arr="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{"+
		"StoreKey:Float64Arr,Option:float-arr,OptArg:,TypeKind:float64,"+
		"Cause:strconv.ParseFloat: parsing \"\": invalid syntax}",
	)
	assert.NotNil(t, errors.Unwrap(err))
	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "Float64Arr")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "float-arr")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Float64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_optionIsStringArrayAndSetOneEmptyStringByDefaultArray(t *testing.T) {
	defer reset()

	type MyOptions struct {
		StringArr []string `optcfg:"str-arr="`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.Equal(t, options.StringArr, []string{""})
}

func TestParseFor_defaultValueIsIgnoreWhenTypeIsBool(t *testing.T) {
	defer reset()

	type MyOptions struct {
		BoolVar bool `optcfg:"bool-var=true"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 1)
	assert.False(t, options.BoolVar)
}

func TestParseFor_errorIfDefaultValueIsInvalidType(t *testing.T) {
	defer reset()

	type MyOptions struct {
		BoolArr []bool
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 0) // because of the error
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "BadFieldType{Option:BoolArr,Field:BoolArr,Type:[]bool}")
	switch err.(type) {
	case errs.BadFieldType:
		assert.Equal(t, err.(errs.BadFieldType).Option, "BoolArr")
		assert.Equal(t, err.(errs.BadFieldType).Field, "BoolArr")
		assert.Equal(t, err.(errs.BadFieldType).Type, reflect.TypeOf(options.BoolArr))
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_multipleOptsAndMultipleArgs(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f"`
		Baz    int      `optcfg:"baz,b=99"`
		Qux    string   `optcfg:"=XXX"`
		Quux   []string `optcfg:"quux=/[A/B/C]"`
		Corge  []int
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"--foo-bar", "c1", "-b", "12", "--Qux", "ABC", "c2",
		"--Corge", "20", "--Corge=21",
	}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"c1", "c2"})
	assert.Equal(t, len(cmd.OptCfgs), 5)
	assert.True(t, options.FooBar)
	assert.Equal(t, options.Baz, 12)
	assert.Equal(t, options.Qux, "ABC")
	assert.Equal(t, options.Quux, []string{"A", "B", "C"})
	assert.Equal(t, options.Corge, []int{20, 21})
}

func TestMakeOptCfgsFor_multipleOptsAndMultipleArgs(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f"`
		Baz    int      `optcfg:"baz,b=99"`
		Qux    string   `optcfg:"=XXX"`
		Quux   []string `optcfg:"quux=/[A/B/C]"`
		Corge  []int
	}
	options := MyOptions{}

	os.Args = []string{
		"path/to/app",
		"--foo-bar", "c1", "-b", "12", "--Qux", "ABC", "c2",
		"--Corge", "20", "--Corge=21",
	}

	optCfgs, err0 := cliargs.MakeOptCfgsFor(&options)
	assert.Nil(t, err0)
	assert.Equal(t, optCfgs[0].StoreKey, "FooBar")
	assert.Equal(t, optCfgs[0].Names, []string{"foo-bar", "f"})
	assert.False(t, optCfgs[0].HasArg)
	assert.False(t, optCfgs[0].IsArray)
	assert.Nil(t, optCfgs[0].Defaults)
	assert.Equal(t, optCfgs[1].StoreKey, "Baz")
	assert.Equal(t, optCfgs[1].Names, []string{"baz", "b"})
	assert.True(t, optCfgs[1].HasArg)
	assert.False(t, optCfgs[1].IsArray)
	assert.Equal(t, optCfgs[1].Defaults, []string{"99"})
	assert.Equal(t, optCfgs[2].StoreKey, "Qux")
	assert.Equal(t, optCfgs[2].Names, []string{"Qux"})
	assert.True(t, optCfgs[2].HasArg)
	assert.False(t, optCfgs[2].IsArray)
	assert.Equal(t, optCfgs[2].Defaults, []string{"XXX"})
	assert.Equal(t, optCfgs[3].StoreKey, "Quux")
	assert.Equal(t, optCfgs[3].Names, []string{"quux"})
	assert.True(t, optCfgs[3].HasArg)
	assert.True(t, optCfgs[3].IsArray)
	assert.Equal(t, optCfgs[3].Defaults, []string{"A", "B", "C"})
	assert.Equal(t, optCfgs[4].StoreKey, "Corge")
	assert.Equal(t, optCfgs[4].Names, []string{"Corge"})
	assert.True(t, optCfgs[4].HasArg)
	assert.True(t, optCfgs[4].IsArray)
	assert.Equal(t, optCfgs[4].Defaults, []string(nil))

	cmd := cliargs.NewCmd()
	err1 := cmd.ParseWith(optCfgs)
	assert.Nil(t, err1)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{"c1", "c2"})
	assert.True(t, cmd.HasOpt("FooBar"))
	assert.True(t, cmd.HasOpt("Baz"))
	assert.True(t, cmd.HasOpt("Qux"))
	assert.True(t, cmd.HasOpt("Quux"))
	assert.True(t, cmd.HasOpt("Corge"))
	assert.Equal(t, cmd.OptArg("FooBar"), "")
	assert.Equal(t, cmd.OptArg("Baz"), "12")
	assert.Equal(t, cmd.OptArg("Qux"), "ABC")
	assert.Equal(t, cmd.OptArg("Quux"), "A")
	assert.Equal(t, cmd.OptArg("Corge"), "20")
	assert.Equal(t, cmd.OptArgs("FooBar"), []string(nil))
	assert.Equal(t, cmd.OptArgs("Baz"), []string{"12"})
	assert.Equal(t, cmd.OptArgs("Qux"), []string{"ABC"})
	assert.Equal(t, cmd.OptArgs("Quux"), []string{"A", "B", "C"})
	assert.Equal(t, cmd.OptArgs("Corge"), []string{"20", "21"})
	assert.True(t, options.FooBar)
	assert.Equal(t, options.Baz, 12)
	assert.Equal(t, options.Qux, "ABC")
	assert.Equal(t, options.Quux, []string{"A", "B", "C"})
	assert.Equal(t, options.Corge, []int{20, 21})
}

func TestParseFor_emptyArrayOfDefaultValueWithNotCommaSeparator(t *testing.T) {
	defer reset()

	type MyOptions struct {
		Foo []int     `optcfg:"foo=/[]"`
		Bar []uint    `optcfg:"bar=|[]"`
		Baz []float64 `optcfg:"baz=@[]"`
		Qux []string  `optcfg:"qux=![]"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app"}

	cmd := cliargs.NewCmd()
	err := cmd.ParseFor(&options)
	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.Equal(t, len(cmd.OptCfgs), 4)
	assert.Equal(t, options.Foo, []int{})
	assert.Equal(t, options.Bar, []uint{})
	assert.Equal(t, options.Baz, []float64{})
	assert.Equal(t, options.Qux, []string{})
}

func TestMakeOptCfgsFor_optionDescriptions(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optdesc:"FooBar description"`
		Baz    int      `optcfg:"baz,b=99" optdesc:"Baz description"`
		Qux    string   `optcfg:"=XXX" optdesc:"Qux description"`
		Quux   []string `optcfg:"quux=/[A/B/C]" optdesc:"Quux description"`
		Corge  []int
	}
	options := MyOptions{}

	optCfgs, err0 := cliargs.MakeOptCfgsFor(&options)
	assert.Nil(t, err0)
	assert.Equal(t, optCfgs[0].Desc, "FooBar description")
	assert.Equal(t, optCfgs[1].Desc, "Baz description")
	assert.Equal(t, optCfgs[2].Desc, "Qux description")
	assert.Equal(t, optCfgs[3].Desc, "Quux description")
	assert.Equal(t, optCfgs[4].Desc, "")
}

func TestMakeOptCfgsFor_optionParam(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optarg:"aaa"`
		Baz    int      `optcfg:"baz,b=99" optarg:"bbb"`
		Qux    string   `optcfg:"=XXX" optarg:"ccc"`
		Quux   []string `optcfg:"quux=/[A/B/C]" optarg:"ddd (multiple)"`
		Corge  []int
	}
	options := MyOptions{}

	optCfgs, err0 := cliargs.MakeOptCfgsFor(&options)
	assert.Nil(t, err0)
	assert.Equal(t, optCfgs[0].ArgInHelp, "")
	assert.Equal(t, optCfgs[1].ArgInHelp, "bbb")
	assert.Equal(t, optCfgs[2].ArgInHelp, "ccc")
	assert.Equal(t, optCfgs[3].ArgInHelp, "ddd (multiple)")
	assert.Equal(t, optCfgs[4].ArgInHelp, "")
}

func TestParseFor_optCfgHasUnsupportedType(t *testing.T) {
	defer reset()

	type A struct{ Name string }
	type MyOptions struct {
		FooBar A `optcfg:"foo-bar,f" optdesc:"FooBar description"`
	}

	options := MyOptions{}

	_, err := cliargs.MakeOptCfgsFor(&options)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "BadFieldType{"+
		"Option:foo-bar,Field:FooBar,Type:cliargs_test.A}",
	)
	switch err.(type) {
	case errs.BadFieldType:
		assert.Equal(t, err.(errs.BadFieldType).Option, "foo-bar")
		assert.Equal(t, err.(errs.BadFieldType).Field, "FooBar")
		assert.Equal(t, err.(errs.BadFieldType).Type.(reflect.Type).Name(), "A")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_argIsNotPointer(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optdesc:"FooBar description"`
		Baz    int      `optcfg:"baz,b=99" optdesc:"Baz description"`
		Qux    string   `optcfg:"=XXX" optdesc:"Qux description"`
		Quux   []string `optcfg:"quux=/[A/B/C]" optdesc:"Quux description"`
		Corge  []int
	}
	options := MyOptions{}

	_, err := cliargs.MakeOptCfgsFor(options)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "OptionStoreIsNotChangeable{}")
	switch err.(type) {
	case errs.OptionStoreIsNotChangeable:
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseUntilSubCmdFor_hasSubCmd(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool `optcfg:"foo-bar,f"`
		Corge  bool `optcfg:"corge"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app", "--foo-bar", "baz", "qux", "--corge"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("FooBar"))
	assert.False(t, cmd.HasOpt("Corge"))

	assert.True(t, options.FooBar)
	assert.False(t, options.Corge)

	assert.Equal(t, subCmd.Name, "baz")
	assert.Equal(t, subCmd.Args, []string{})
	assert.False(t, subCmd.HasOpt("corge"))

	err = subCmd.Parse()
	assert.Nil(t, err)

	assert.Equal(t, subCmd.Name, "baz")
	assert.Equal(t, subCmd.Args, []string{"qux"})
	assert.True(t, subCmd.HasOpt("corge"))
}

func TestParseUntilSubCmdFor_noSubCmd(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar bool `optcfg:"foo-bar,f"`
		Corge  bool `optcfg:"corge"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app", "--foo-bar", "--corge"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdFor(&options)

	assert.Nil(t, err)
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("FooBar"))
	assert.True(t, cmd.HasOpt("Corge"))

	assert.True(t, options.FooBar)
	assert.True(t, options.Corge)

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))

	err = subCmd.Parse()
	assert.Nil(t, err)

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))
}

func TestParseUntilSubCmdFor_error0(t *testing.T) {
	defer reset()

	type A struct{}
	type MyOptions struct {
		FooBar A `optcfg:"foo-bar,f"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app", "--foo-bar", "baz", "qux", "--corge"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdFor(&options)

	switch err.(type) {
	case errs.BadFieldType:
		assert.Equal(t, err.(errs.BadFieldType).Option, "foo-bar")
		assert.Equal(t, err.(errs.BadFieldType).Field, "FooBar")
		assert.Equal(t, err.(errs.BadFieldType).Type, reflect.TypeOf(options.FooBar))
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.False(t, cmd.HasOpt("FooBar"))

	assert.Equal(t, options.FooBar, A{})

	assert.Equal(t, subCmd.Name, "")
	assert.Equal(t, subCmd.Args, []string(nil))
}

func TestParseUntilSubCmdFor_error1(t *testing.T) {
	defer reset()

	type MyOptions struct {
		FooBar int `optcfg:"foo-bar,f"`
	}
	options := MyOptions{}

	os.Args = []string{"/path/to/app", "--foo-bar", "baz", "qux", "--corge"}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdFor(&options)

	switch err.(type) {
	case errs.OptionArgIsInvalid:
		assert.Equal(t, err.(errs.OptionArgIsInvalid).StoreKey, "FooBar")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Option, "foo-bar")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).OptArg, "baz")
		assert.Equal(t, err.(errs.OptionArgIsInvalid).TypeKind, reflect.Int)
		assert.Equal(t, err.(errs.OptionArgIsInvalid).Cause.Error(), "strconv.ParseInt: parsing \"baz\": invalid syntax")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args, []string{})
	assert.True(t, cmd.HasOpt("FooBar"))

	assert.Equal(t, options.FooBar, 0)

	assert.Equal(t, subCmd.Name, "qux")
	assert.Equal(t, subCmd.Args, []string{})
}
