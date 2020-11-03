package variables

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncRegex(t *testing.T) {
	cases := []struct {
		Func         string
		ExpectedName string
		ParamLen     int
	}{
		{
			Func:         "golang",
			ExpectedName: VariableTypeString,
		},
		{
			Func:         "@string",
			ExpectedName: "@string",
		},
		{
			Func:         "@string()",
			ExpectedName: "@string",
		},
		{
			Func:         "@string([0,1,2,3,4])",
			ExpectedName: "@string",
			ParamLen:     1,
		},
		{
			Func:         "@string([[1,2],[2,3],[4,5],[6,7]],4)",
			ExpectedName: "@string",
			ParamLen:     2,
		},
		{
			Func:         "@string([[1,2],[2,3],[4,5],[6,7]],4)",
			ExpectedName: "@string",
			ParamLen:     2,
		},
		{
			Func:         `@string([{"Name":"张三"},[2,3],[4,5],[6,7]],4)`,
			ExpectedName: "@string",
			ParamLen:     2,
		},
	}

	for _, value := range cases {
		name, params, err := parseStringFuncSign(value.Func)
		assert.Nil(t, err)
		assert.Equal(t, value.ExpectedName, name)
		assert.Equal(t, value.ParamLen, len(params))
		fmt.Println("Name:", name, "Params:", params)
	}
}
