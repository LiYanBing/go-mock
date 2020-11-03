package tools

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/liyanbing/go-mock/constant"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func BenchmarkBoolean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ret := Boolean()
		assert.Equal(&testing.T{}, ret == true || ret == false, true)
	}
}

func TestNatural(t *testing.T) {
	cases := []struct {
		Min int64
		Max int64
	}{
		{
			Min: 0,
			Max: 0,
		},
		{
			Min: 0,
			Max: 100,
		},
		{
			Min: -100,
			Max: 100,
		},
		{
			Min: -100,
			Max: -100,
		},
	}

	for _, v := range cases {
		ret := Natural(v.Min, v.Max)
		if v.Max != v.Min {
			assert.Equal(t, ret >= v.Min && ret < v.Max, true, ret, v.Min, v.Max)
		}
	}
}

func TestCharacter(t *testing.T) {
	cases := []struct {
		Type     CharType
		Expected func(in string) bool
	}{
		{
			Type: CharTypeNumber,
			Expected: func(in string) bool {
				return strings.Contains(constant.Number, in)
			},
		},
		{
			Type: CharTypeUpper,
			Expected: func(in string) bool {
				return strings.Contains(constant.Upper, in)
			},
		},
		{
			Type: CharTypeLower,
			Expected: func(in string) bool {
				return strings.Contains(constant.Lower, in)
			},
		},
		{
			Type: CharTypeSymbol,
			Expected: func(in string) bool {
				return strings.Contains(constant.Symbol, in)
			},
		},
		{
			Type: CharTypeRand,
			Expected: func(in string) bool {
				return strings.Contains(constant.Pool, in)
			},
		},
	}

	for _, v := range cases {
		ret := Character(v.Type)
		assert.Equal(t, v.Expected(ret), true, ret)
	}
}

func TestFloat(t *testing.T) {
	cases := []struct {
		Min  int64
		Max  int64
		Dmin int64
		Dmax int64
	}{
		{
			Min:  100,
			Max:  200,
			Dmin: 1,
			Dmax: 5,
		},
		{
			Min:  100,
			Max:  200,
			Dmin: 5,
			Dmax: 5,
		},
		{
			Min:  100,
			Max:  200,
			Dmin: 0,
			Dmax: 0,
		},
		{
			Min:  100,
			Max:  200,
			Dmin: 1,
			Dmax: 5,
		},
	}

	for _, v := range cases {
		ret := Float(v.Min, v.Max, v.Dmin, v.Dmax)
		fmt.Println("Param:", v, "ret:", ret)
	}
}

func TestBoolean(t *testing.T) {
	aa := []interface{}{1, 2, 3}
	t.Log(PickMany(2, aa...))
}

func BenchmarkStringWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(StringWithPool(constant.CommonChineseWords+"MaryPatriciaLindaBarbaraElizabeth", 5))
	}
}

func TestRGB(t *testing.T) {
	for i := 0; i < 100000; i++ {
		t.Log(fmt.Sprintf("%.2f", rand.Float64()))
	}
}
