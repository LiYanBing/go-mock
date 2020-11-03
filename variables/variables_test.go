package variables

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @date
func TestDateVariable(t *testing.T) {
	cases := []struct {
		Params []interface{}
	}{
		{
			Params: nil,
		},
		{
			Params: []interface{}{"yyyy-MM-dd A HH:mm:ss"},
		},
		{
			Params: []interface{}{"yy-MM-dd a HH:mm:ss"},
		},
		{
			Params: []interface{}{"y-MM-dd HH:mm:ss"},
		},
		{
			Params: []interface{}{"y-M-d H:m:s"},
		},
		{
			Params: []interface{}{"y-M-d H:m:s"},
		},
		{
			Params: []interface{}{"yyyy yy y MM M dd d HH H hh h mm m ss s SS S A a T"},
		},
	}

	for _, value := range cases {
		ret, err := DateFunc(context.Background(), &VariableInput{
			Params: value.Params,
		})
		assert.Nil(t, err)
		fmt.Println(ret)
	}
}

func TestBoolVariable(t *testing.T) {
	format := fmt.Sprintf("%x", 255)
	t.Log(format)

	i, err := strconv.ParseInt(format, 16, 64)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(i)
}
