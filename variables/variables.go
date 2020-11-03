package variables

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/liyanbing/go-mock/constant"
	"github.com/liyanbing/go-mock/images"
	"github.com/liyanbing/go-mock/tools"

	uuid "github.com/satori/go.uuid"
)

var (
	ErrInvalidMockMethod = errors.New("invalid mock method or Params")
)

/*=============variable================*/
// type.number
func TypeNumberFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		min, dMin = int64(0), int64(0)
		max, dMax = int64(10), int64(10)
		err       error
	)
	paramNum := len(in.Params)

	if paramNum <= 1 {
		return in.Value, nil
	}

	if paramNum >= 1 {
		min, err = tools.ToInt(in.Params[0])
		if err != nil {
			return nil, err
		}
	}

	if paramNum >= 2 {
		max, err = tools.ToInt(in.Params[1])
		if err != nil {
			return nil, err
		}
	}

	if paramNum >= 3 {
		dMin, err = tools.ToInt(in.Params[2])
		if err != nil {
			return nil, err
		}
	}

	if paramNum >= 4 {
		dMax, err = tools.ToInt(in.Params[2])
		if err != nil {
			return nil, err
		}
	}

	if paramNum == 2 {
		return tools.Natural(min, max), nil
	}

	return tools.Float(min, max, dMin, dMax), nil
}

// type.string
func TypeStringFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	return in.Rule.Generate(in.Value.(string)), nil
}

// @string
func StringFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		tt  = tools.CharTypeRand
		min = int64(3)
		max = int64(10)
		err error
	)
	paramsNum := len(in.Params)

	setMin := false
	if paramsNum >= 1 {
		par1, ok := in.Params[0].(string)
		if ok {
			tt = tools.CharType(par1)
		} else {
			setMin = true
			min, err = tools.ToInt(in.Params[0])
			if err != nil {
				return nil, ErrInvalidMockMethod
			}
			max = min
		}
	}

	if paramsNum >= 2 {
		var cur int64
		cur, err = tools.ToInt(in.Params[1])
		if err != nil {
			return nil, ErrInvalidMockMethod
		}

		if setMin {
			max = cur
		} else {
			min = cur
			max = min
		}
	}

	if paramsNum >= 3 {
		max, err = tools.ToInt(in.Params[2])
		if err != nil {
			return nil, ErrInvalidMockMethod
		}
	}

	switch tt {
	case tools.CharTypeLower, tools.CharTypeUpper, tools.CharTypeNumber, tools.CharTypeSymbol, tools.CharTypeRand:
		return tools.String(tt, min, max), nil
	default:
		var maxs []int64
		if max > 0 {
			maxs = append(maxs, max)
		}
		return tools.StringWithPool(string(tt), min, maxs...), nil
	}
}

// @integer
func IntegerFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(math.MinInt64, math.MaxInt64, in)
	if err != nil {
		return nil, err
	}
	return tools.Natural(min, max), nil
}

// @float
func FloatFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		min, dMin = int64(0), int64(0)
		max, dMax = int64(math.MaxInt32), int64(17)
		err       error
	)
	paramNum := len(in.Params)

	if paramNum >= 1 {
		min, err = tools.ToInt(in.Params[0])
		if err != nil {
			return nil, err
		}
	}

	if paramNum >= 2 {
		max, err = tools.ToInt(in.Params[1])
		if err != nil {
			return nil, err
		}
	}

	if paramNum >= 3 {
		dMin, err = tools.ToInt(in.Params[2])
		if err != nil {
			return nil, err
		}
	}

	if paramNum >= 4 {
		dMax, err = tools.ToInt(in.Params[3])
		if err != nil {
			return nil, err
		}
	}
	return tools.Float(min, max, dMin, dMax), nil
}

// @character
func CharacterFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	tt := tools.CharTypeRand
	paramNum := len(in.Params)

	switch paramNum {
	case 1:
		ct := tools.ToString(in.Params[0])
		withPool := false
		switch tools.CharType(strings.ToLower(ct)) {
		case tools.CharTypeLower:
			tt = tools.CharTypeLower
		case tools.CharTypeUpper:
			tt = tools.CharTypeUpper
		case tools.CharTypeNumber:
			tt = tools.CharTypeNumber
		case tools.CharTypeSymbol:
			tt = tools.CharTypeSymbol
		default:
			withPool = true
		}

		if withPool && len(ct) > 0 {
			return tools.CharacterWithPool(ct), nil
		}
		return tools.Character(tt), nil
	default:
		return tools.Character(tt), nil
	}
}

// @bool(随机生成一个布尔值，值为 value 的概率是 min / (min + max)，值为 !value 的概率是 max / (min + max))
func BoolFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	if len(in.Params) >= 3 {
		var (
			min     = int64(0)
			max     = int64(0)
			current = false
			err     error
		)
		paramNum := len(in.Params)

		if paramNum >= 1 {
			min, err = tools.ToInt(in.Params[0])
			if err != nil {
				return "", err
			}
		}

		if paramNum >= 2 {
			max, err = tools.ToInt(in.Params[1])
			if err != nil {
				return "", err
			}
		}

		if paramNum >= 3 {
			current = tools.ToBool(in.Params[2])
		}

		n := tools.Natural(0, min+max)
		if n < min {
			return current, nil
		}
		return !current, nil
	}
	return tools.Boolean(), nil
}

// @natural
func NaturalFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(1, math.MaxInt64, in)
	if err != nil {
		return nil, err
	}
	return tools.Natural(min, max), nil
}

// @date
func DateFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	format := ""
	if len(in.Params) >= 1 {
		format = tools.ToString(in.Params[0])
	}
	return tools.DatetimeFormat(tools.LayoutDate, tools.RandDatetime(), format), nil
}

// @time
func TimeFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	format := ""
	if len(in.Params) >= 1 {
		format = tools.ToString(in.Params[0])
	}
	return tools.DatetimeFormat(tools.LayoutTime, tools.RandDatetime(), format), nil
}

// @datetime
func DatetimeFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	format := ""
	if len(in.Params) >= 1 {
		format = tools.ToString(in.Params[0])
	}
	return tools.DatetimeFormat(tools.LayoutDateTime, tools.RandDatetime(), format), nil
}

// @now
func NowFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		format = ""
		dur    = ""
	)
	paramNum := len(in.Params)

	if paramNum >= 1 {
		dur = tools.ToString(in.Params[0])
	}

	if paramNum >= 2 {
		format = tools.ToString(in.Params[1])
	}

	now := time.Now()
	switch dur {
	case "year":
		now = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	case "month":
		now = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	case "week":
		now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, -int(now.Weekday()))
	case "day":
		now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	case "hour":
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	case "minute":
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC)
	case "second":
		now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	default:
		format = dur
	}

	return tools.DatetimeFormat(tools.LayoutDateTime, now, format), nil
}

// @region
func RegionFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.Region...), nil
}

// @province
func ProvinceFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return constant.RandProvince().Name, nil
}

// @city
func CityFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	prefix := false
	if len(in.Params) > 0 {
		prefix = true
	}

	province := constant.RandProvince()
	if len(province.Children) <= 0 {
		return province.Name, nil
	}

	city, ok := province.PickCity()
	if !ok {
		return province.Name, nil
	}

	if prefix {
		return province.Name + " " + city.Name, nil
	}
	return city.Name, nil
}

// @county
func CountryFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	prefix := false
	if len(in.Params) > 0 {
		prefix = true
	}

	province := constant.RandProvince()
	if len(province.Children) <= 0 {
		return province.Name, nil
	}

	city, ok := province.PickCity()
	if !ok {
		return province.Name, nil
	}

	distinct, ok := city.PickDistinct()
	if !ok {
		return province.Name + " " + city.Name, nil
	}

	if prefix {
		return province.Name + " " + city.Name + " " + distinct.Name, nil
	}
	return distinct.Name, nil
}

// @word
func WordFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(3, 10, in)
	if err != nil {
		return nil, err
	}

	return tools.String(tools.CharTypeLower, min, max), nil
}

// @title
func TitleFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(3, 10, in)
	if err != nil {
		return nil, err
	}

	n := randFromMinAndMax(adjustMinAndMax(min, max))
	ret := ""
	for i := 0; i < int(n); i++ {
		ret += " " + tools.FirstUpper(tools.String(tools.CharTypeLower, 3, 10))
	}
	return ret, nil
}

// @sentence
func SentenceFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(5, 20, in)
	if err != nil {
		return nil, err
	}

	n := randFromMinAndMax(adjustMinAndMax(min, max))
	ret := ""
	for i := 0; i < int(n); i++ {
		item := tools.String(tools.CharTypeLower, 2, 10)
		if i == 0 {
			item = tools.FirstUpper(item)
		}
		ret += " " + item
	}
	return ret, nil
}

// @paragraph
func ParagraphFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(1, 4, in)
	if err != nil {
		return nil, err
	}

	n := randFromMinAndMax(adjustMinAndMax(min, max))
	ret := ""
	for i := 0; i < int(n); i++ {
		paragraph := ""
		wordN := tools.Natural(10, 40)
		for i := 0; i < int(wordN); i++ {
			item := tools.String(tools.CharTypeLower, 3, 10)
			if i == 0 {
				paragraph = tools.FirstUpper(item)
			}
			paragraph += " " + item
		}
		paragraph += "."

		if i > 0 {
			ret += "\n\n"
		}
		ret += paragraph
	}
	return ret, nil
}

// @cword
func CWordFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		pool = constant.CommonChineseWords
		min  = int64(1)
		max  = int64(0)
		err  error
	)
	paramNum := len(in.Params)
	firstPool := false

	if paramNum >= 1 {
		min, err = tools.ToInt(in.Params[0])
		if err != nil {
			min = 1
			firstPool = true
			pool = tools.ToString(in.Params[0])
		}
	}

	if paramNum >= 2 {
		ret, err := tools.ToInt(in.Params[1])
		if err != nil {
			return "", err
		}

		if firstPool {
			min = ret
		} else {
			max = ret
		}
	}

	if paramNum >= 3 {
		max, err = tools.ToInt(in.Params[2])
		if err != nil {
			return "", err
		}
	}

	var maxs []int64
	if max > 0 {
		maxs = append(maxs, max)
	}
	return tools.StringWithPool(pool, min, maxs...), nil
}

// @ctitle
func CTitleFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(3, 6, in)
	if err != nil {
		return nil, err
	}

	var maxs []int64
	if max > 0 {
		maxs = append(maxs, max)
	}
	return tools.StringWithPool(constant.CommonChineseWords, min, maxs...), nil
}

// @csentence
func CSentenceFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(5, 20, in)
	if err != nil {
		return nil, err
	}

	var maxs []int64
	if max > 0 {
		maxs = append(maxs, max)
	}
	return tools.StringWithPool(constant.CommonChineseWords, min, maxs...) + "。", nil
}

// @cparagraph
func CParagraphFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	min, max, err := parseMinAndMax(2, 5, in)
	if err != nil {
		return nil, err
	}

	n := randFromMinAndMax(min, max)
	ret := ""
	for i := 0; i < int(n); i++ {
		parag := ""
		titleN := tools.Natural(5, 10)
		for j := 0; j < int(titleN); j++ {
			parag += tools.StringWithPool(constant.CommonChineseWords, 5, 20) + "。"
		}
		ret += parag + "\n\n"
	}
	return ret, nil
}

// @first
func FirstFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.First...), nil
}

// @last
func LastFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.Last...), nil
}

// @name
func NameFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		re = false
	)

	if len(in.Params) >= 1 {
		re = true
	}

	name := fmt.Sprintf("%v %v", tools.Pick(constant.First...), tools.Pick(constant.Last...))
	if re {
		name = fmt.Sprintf("%v %v", tools.Pick(constant.First...), name)
	}
	return name, nil
}

// @cfirst
func CFirstFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.CFirst...), nil
}

// @last
func CLastFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.CLast...), nil
}

// @name
func CNameFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	name := fmt.Sprintf("%v%v", tools.Pick(constant.CFirst...), tools.Pick(constant.CLast...))
	return name, nil
}

// @protocol
func ProtocolFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.Protocol...), nil
}

// @domain
func DomainFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return fmt.Sprintf("%v.%v", tools.String(tools.CharTypeLower, 3, 10), tools.Pick(constant.TLD...)), nil
}

// @tld
func TLDFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.Pick(constant.TLD...), nil
}

// @url
func URLFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	domain := fmt.Sprintf("%v.%v",
		tools.String(tools.CharTypeLower, 3, 10),
		tools.Pick(constant.TLD...))

	return fmt.Sprintf("%v://%v/%v",
		tools.Pick(constant.Protocol...),
		domain,
		tools.String(tools.CharTypeLower, 3, 10)), nil
}

// @email
func EmailFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return fmt.Sprintf("%v.%v@%v.%v",
		tools.Character(tools.CharTypeLower),
		tools.String(tools.CharTypeLower, 3, 10),
		tools.String(tools.CharTypeLower, 3, 10),
		tools.Pick(constant.TLD...)), nil
}

// @ip
func IPFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return fmt.Sprintf("%v.%v.%v.%v",
		tools.Natural(1, 256),
		tools.Natural(1, 256),
		tools.Natural(1, 256),
		tools.Natural(1, 256)), nil
}

// @zip
func ZIPFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return tools.String(tools.CharTypeNumber, 6, 6), nil
}

// @capitalize
func CapitalizeFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	if len(in.Params) <= 0 {
		return "", nil
	}
	return tools.FirstUpper(tools.ToString(in.Params[0])), nil
}

// @upper
func UpperFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	if len(in.Params) <= 0 {
		return "", nil
	}

	return strings.ToUpper(tools.ToString(in.Params[0])), nil
}

// @lower
func LowerFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	if len(in.Params) <= 0 {
		return "", nil
	}
	return strings.ToLower(tools.ToString(in.Params[0])), nil
}

// @pick
func PickFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	if len(in.Params) <= 0 {
		return "", nil
	}

	param, ok := in.Params[0].([]interface{})
	if !ok {
		return "", nil
	}
	return tools.Pick(param...), nil
}

// @shuffle
func ShuffleFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	if len(in.Params) <= 0 {
		return "", nil
	}

	param, ok := in.Params[0].([]interface{})
	if !ok {
		return "", nil
	}
	return tools.Shuffle(param...), nil
}

// @guid
func GUIDFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return uuid.NewV4().String(), nil
}

// @id
func IDFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return fmt.Sprintf("%v", tools.Natural(0, 0)), nil
}

// @color
func ColorFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return images.ColorHex(), nil
}

// @rgb
func RGBFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return images.ColorRGB(), nil
}

// @rgba
func RGBAFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return images.ColorRGBA(), nil
}

// @hsl
func HSLFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return fmt.Sprintf("hsl(%v, %v, %v)",
		rand.Intn(361),
		rand.Intn(101),
		rand.Intn(101)), nil
}

// @image
func ImageFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		size       = "200x100"
		background = "#FF6600"
		foreground = "#4A7BF7"
		format     = "png"
		text       = size
	)
	paramsNum := len(in.Params)

	if paramsNum >= 1 {
		size = tools.ToString(in.Params[0])
	}

	if paramsNum >= 2 {
		background = tools.ToString(in.Params[1])
	}

	if paramsNum >= 3 {
		text = tools.ToString(in.Params[2])
	}

	if paramsNum >= 4 {
		format = text
		text = tools.ToString(in.Params[3])
	}

	imageID := images.GetImageId(size, background, foreground, format, text)
	if imageAdaptor != nil {
		return imageAdaptor.GenImageURL(imageID), nil
	}
	return imageID, nil
}

// @dataImage
func DataImageFunc(_ context.Context, in *VariableInput) (interface{}, error) {
	var (
		size       = "200x100"
		background = "#FF6600"
		foreground = "#FFF"
		format     = "png"
		text       = size
	)

	paramsNum := len(in.Params)

	if paramsNum >= 1 {
		size = tools.ToString(in.Params[0])
	}

	if paramsNum >= 2 {
		text = tools.ToString(in.Params[1])
	}

	imageID := images.GetImageId(size, background, foreground, format, text)
	if imageAdaptor != nil {
		return imageAdaptor.GenImageURL(imageID), nil
	}
	return imageID, nil
}

// @null
func NullFunc(_ context.Context, _ *VariableInput) (interface{}, error) {
	return nil, nil
}

func parseMinAndMax(defaultMin, defaultMax int64, in *VariableInput) (int64, int64, error) {
	var (
		err error
	)
	paramNum := len(in.Params)

	if paramNum >= 1 {
		defaultMin, err = tools.ToInt(in.Params[0])
		if err != nil {
			return 0, 0, err
		}
		defaultMax = defaultMin
	}

	if paramNum >= 2 {
		defaultMax, err = tools.ToInt(in.Params[1])
		if err != nil {
			return 0, 0, err
		}
	}

	return defaultMin, defaultMax, nil
}

func randFromMinAndMax(min, max int64) int64 {
	ret := min
	if min != max {
		ret = tools.Natural(min, max)
	}

	if ret > max {
		ret = max
	}
	return ret
}

// 调整在1-30之间
func adjustMinAndMax(min, max int64) (mi, ma int64) {
	if max > 30 {
		max = 30
	}

	if max < 1 {
		max = 10
	}

	if min < 1 {
		min = 1
	}

	if min > 30 {
		min = 10
	}

	if min > max {
		min, max = max, min
	}
	return min, max
}
