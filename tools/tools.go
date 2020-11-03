package tools

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/liyanbing/go-mock/constant"
)

const (
	LayoutDate     = "2006-01-02"
	LayoutTime     = "15:04:05"
	LayoutDateTime = "2006-01-02 15:04:05"
)

type CharType string

const (
	CharTypeUpper  CharType = "upper"
	CharTypeLower  CharType = "Lower"
	CharTypeNumber CharType = "number"
	CharTypeSymbol CharType = "symbol"
	CharTypeRand   CharType = "rand"
)

func IsString(in interface{}) bool {
	_, ok := in.(string)
	return ok
}

func IsNumber(in interface{}) bool {
	switch in.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}

func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(value.(int64)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(value.(float64), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return fmt.Sprintf("%+v", value)
	}
}

func ToFloat(input interface{}) (float64, error) {
	switch input.(type) {
	case int:
		return float64(input.(int)), nil
	case uint:
		return float64(input.(uint)), nil
	case int8:
		return float64(input.(int8)), nil
	case uint8:
		return float64(input.(uint8)), nil
	case int16:
		return float64(input.(int16)), nil
	case uint16:
		return float64(input.(uint16)), nil
	case int32:
		return float64(input.(int32)), nil
	case uint32:
		return float64(input.(uint32)), nil
	case int64:
		return float64(input.(int64)), nil
	case uint64:
		return float64(input.(uint64)), nil
	case float32:
		return float64(input.(float32)), nil
	case float64:
		return input.(float64), nil
	case string:
		return strconv.ParseFloat(input.(string), 64)
	}
	return 0, fmt.Errorf("invalid value<%v> ToFloat", input)
}

func ToInt(input interface{}) (int64, error) {
	switch input.(type) {
	case int:
		return int64(input.(int)), nil
	case uint:
		return int64(input.(uint)), nil
	case int8:
		return int64(input.(int8)), nil
	case uint8:
		return int64(input.(uint8)), nil
	case int16:
		return int64(input.(int16)), nil
	case uint16:
		return int64(input.(uint16)), nil
	case int32:
		return int64(input.(int32)), nil
	case uint32:
		return int64(input.(uint32)), nil
	case int64:
		return input.(int64), nil
	case uint64:
		return int64(input.(uint64)), nil
	case float32:
		return int64(input.(float32)), nil
	case float64:
		return int64(input.(float64)), nil
	case string:
		return strconv.ParseInt(input.(string), 10, 64)
	}
	return 0, fmt.Errorf("invalid value<%v> ToInt", input)
}

func ToBool(in interface{}) bool {
	switch in.(type) {
	case bool:
		return in.(bool)
	case string:
		ret, err := strconv.ParseBool(in.(string))
		if err != nil {
			return false
		}
		return ret
	case int:
		return in.(int) > 0
	case uint:
		return in.(uint) > 0
	case int8:
		return in.(int8) > 0
	case uint8:
		return in.(uint8) > 0
	case int16:
		return in.(int16) > 0
	case uint16:
		return in.(uint16) > 0
	case int32:
		return in.(int32) > 0
	case uint32:
		return in.(uint32) > 0
	case int64:
		return in.(int64) > 0
	case uint64:
		return in.(uint64) > 0
	case float32:
		return in.(float32) > 0
	case float64:
		return in.(float64) > 0
	}
	return false
}

func FirstLower(input string) string {
	if len(input) <= 0 {
		return input
	}

	return strings.ToLower(input[:1]) + input[1:]
}

func FirstUpper(input string) string {
	if len(input) <= 0 {
		return input
	}

	return strings.ToUpper(input[:1]) + input[1:]
}

func Boolean() bool {
	return rand.Intn(2) > 0
}

// 如果 min <= max 则会返回一个max -> min 之间的数字
// 如果 min = max 且min >0, 会生成 min- ∞ 之间的数字
func Natural(min, max int64) int64 {
	if min == max {
		if min == math.MaxInt64 {
			return math.MaxInt64
		}
		max = math.MaxInt64
	}

	if min > max {
		min, max = max, min
	}

	if min >= 0 {
		return rand.Int63n(max-min) + min
	}

	// 这里蒋不能随机到 -9223372036854775808
	if min <= -9223372036854775807 {
		min = -9223372036854775807
	}
	positiveMin := -min
	absMin := positiveMin
	if max < absMin {
		absMin, max = max, absMin
	}

	ret := rand.Int63n(max)
	if ret <= absMin {
		n := rand.Intn(2)
		if n < 1 {
			return -ret
		}
	}
	return ret
}

// 小数点后小数部分不能超过17位
func Float(min, max, dmin, dmax int64) float64 {
	num1 := Natural(min, max)

	if dmin >= dmax {
		return float64(num1)
	}

	if dmax-dmin > 17 {
		dmax = dmin + 17
	}

	num2 := Natural(dmin, dmax)
	if num2 <= 0 {
		return float64(num1)
	}

	if num2 > 17 {
		num2 = 17
	}

	ret := fmt.Sprintf("%v.", num1)
	for i := 0; i < int(num2); i++ {
		n := Character(CharTypeNumber)
		if i == int(num2)-1 && n == "0" {
			continue
		}
		ret += n
	}
	nn, _ := strconv.ParseFloat(ret, 64)
	return nn
}

func CharacterWithPool(pool string) string {
	n := Natural(0, int64(len(pool)))
	return pool[n : n+1]
}

func Character(t CharType) string {
	switch t {
	case CharTypeUpper:
		n := Natural(0, int64(len(constant.Upper)))
		return constant.Upper[n : n+1]
	case CharTypeLower:
		n := Natural(0, int64(len(constant.Lower)))
		return constant.Lower[n : n+1]
	case CharTypeNumber:
		n := Natural(0, int64(len(constant.Number)))
		return constant.Number[n : n+1]
	case CharTypeSymbol:
		n := Natural(0, int64(len(constant.Symbol)))
		return constant.Symbol[n : n+1]
	default:
		n := Natural(0, int64(len(constant.Pool)))
		return constant.Pool[n : n+1]
	}
}

func String(t CharType, min, max int64) string {
	var length int64
	if min == max && min > 0 {
		length = min
	} else {
		if min == max && min <= 0 {
			min = 0
			max = 30
		}
		length = Natural(min, max)
	}

	if length > max {
		length = max
	}

	ret := ""
	for i := 0; i < int(length); i++ {
		ret += Character(t)
	}
	return ret
}

// 如果没有max则返回min个汉字
func CWordWithNumber(min int64, maxs ...int64) string {
	return StringWithPool(constant.CommonChineseWords, min, maxs...)
}

// 从pool中生成min个汉字，如果有传递maxs则生成min-maxs[0]中间随机数个汉字
func StringWithPool(pool string, min int64, maxs ...int64) string {
	if pool == "" {
		return ""
	}

	n := min
	if len(maxs) > 0 {
		max := maxs[0]
		if max < min {
			min, max = max, min
		}

		n = Natural(min, max)
		if n > max {
			n = max
		}
	}

	poolRune := []rune(pool)
	poolLen := utf8.RuneCountInString(pool)
	ret := make([]rune, 0, int(n))
	for i := 0; i < int(n); i++ {
		index := Natural(0, int64(poolLen))
		ret = append(ret, poolRune[index:index+1]...)
	}
	return string(ret)
}

func DatetimeFormat(selfFormat string, in time.Time, format string) string {
	if format == "" {
		return in.Format(selfFormat)
	}

	format = strings.Replace(format, "yyyy", "2006", -1)
	format = strings.Replace(format, "yy", "06", -1)
	format = strings.Replace(format, "y", "6", -1)
	format = strings.Replace(format, "MM", "01", -1)
	format = strings.Replace(format, "M", "1", -1)
	format = strings.Replace(format, "dd", "02", -1)
	format = strings.Replace(format, "d", "2", -1)
	format = strings.Replace(format, "HH", "15", -1)
	format = strings.Replace(format, "H", "3", -1)
	format = strings.Replace(format, "hh", "03", -1)
	format = strings.Replace(format, "h", "3", -1)
	format = strings.Replace(format, "mm", "04", -1)
	format = strings.Replace(format, "m", "4", -1)
	format = strings.Replace(format, "ss", "05", -1)
	format = strings.Replace(format, "s", "5", -1)
	format = in.Format(format)
	format = strings.Replace(format, "T", fmt.Sprintf("%v", in.Unix()), -1)
	format = strings.Replace(format, "SS", fmt.Sprintf("%v", in.Nanosecond()/1e6), -1)
	format = strings.Replace(format, "S", "0", -1)

	if in.Hour() < 12 {
		format = strings.Replace(format, "A", "AM", -1)
		format = strings.Replace(format, "a", "am", -1)
	} else {
		format = strings.Replace(format, "A", "PM", -1)
		format = strings.Replace(format, "a", "pm", -1)
	}
	return format
}

func RandDatetime() time.Time {
	return time.Unix(rand.Int63n(time.Now().Unix()), 0)
}

func Pick(in ...interface{}) interface{} {
	if len(in) <= 0 {
		return nil
	}

	if len(in) == 1 {
		return in[0]
	}
	return in[rand.Intn(len(in))]
}

func PickMany(n int, in ...interface{}) []interface{} {
	if n < 1 {
		return nil
	}

	if n == 1 {
		return []interface{}{Pick(n, in)}
	}

	result := make([]interface{}, 0, len(in))
	for i := 0; i < n; i++ {
		n := rand.Intn(len(in))
		result = append(result, in[n])
		in = append(in[0:n], in[n+1:]...)
	}
	return result
}

func Shuffle(in ...interface{}) []interface{} {
	if len(in) <= 1 {
		return in
	}

	result := make([]interface{}, 0, len(in))
	for len(in) > 0 {
		n := rand.Intn(len(in))
		result = append(result, in[n])
		in = append(in[0:n], in[n+1:]...)
	}
	return result
}
