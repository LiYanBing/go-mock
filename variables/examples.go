package variables

import "encoding/json"

type VariableExample struct {
	Name        string          `json:"name"`        // 变量名
	Description string          `json:"description"` // 描述
	Examples    json.RawMessage `json:"examples"`    // 示例
}

type VariableExamples []*VariableExample

var (
	Examples = VariableExamples{
		{
			Name:        VariableString,
			Description: "生成指定长度字符串(默认:3~10)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"string":{
			"type": "string",
            "mock": "Mock.mock(@string)"
		},
		"string()":{
			"type": "string",
            "mock": "Mock.mock(@string())"
		},
		"string(5)":{
			"type": "string",
            "mock": "Mock.mock(@string(5))"
		},
		"string(lower,5)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"lower\",5))"
		},
		"string(upper,5)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"upper\",5))"
		},
		"string(number,5)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"number\",5))"
		},
		"string(symbol,5)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"symbol\",5))"
		},
		"string(aeiou,5)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"aeiou\",5))"
		},
		"string(7, 10)":{
			"type": "string",
            "mock": "Mock.mock(@string(7, 10))"
		},
		"string(lower,5)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"lower\",5))"
		},
		"string(upper,1,3)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"upper\",1,3))"
		},
		"string(number,1,3)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"number\",1,3))"
		},
		"string(symbol,1,3)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"symbol\",1,3))"
		},
		"string(aeiou,1,3)":{
			"type": "string",
            "mock": "Mock.mock(@string(\"aeiou\",1,3))"
		}
	}
}`),
		},
		{
			Name:        VariableInteger,
			Description: "生成指定范围数字(默认int64范围内的数字)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"integer":{
			"type": "string",
            "mock": "Mock.mock(@integer)"
		},
		"integer()":{
			"type": "string",
            "mock": "Mock.mock(@integer())"
		},
		"integer_10000":{
			"type": "string",
            "mock": "Mock.mock(@integer(10000))"
		},
		"integer_60_100":{
			"type": "string",
            "mock": "Mock.mock(@integer(60,100))"
		}
	}
}`),
		},
		{
			Name:        VariableFloat,
			Description: "生成指定范围的浮点数(默认情况下整数部分:0~MaxInt32;小数部分:0~17)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"float":{
			"type": "string",
            "mock": "Mock.mock(@float)"
		},
		"float()":{
			"type": "string",
            "mock": "Mock.mock(@float())"
		},
		"float_0":{
			"type": "string",
            "mock": "Mock.mock(@float(0))"
		},
		"float_60_100":{
			"type": "string",
            "mock": "Mock.mock(@float(60,100))"
		},
		"float_60_100_3":{
			"type": "string",
            "mock": "Mock.mock(@float(60,100,3))"
		},
		"float_60_100_3_5":{
			"type": "string",
            "mock": "Mock.mock(@float(60,100,3,5))"
		}
	}
}`),
		},
		{
			Name:        VariableCharacter,
			Description: "生成一个字符(默认所有字符中随机一个)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"character":{
			"type": "string",
            "mock": "Mock.mock(@character)"
		},
		"character()":{
			"type": "string",
            "mock": "Mock.mock(@character())"
		},
		"character(upper)":{
			"type": "string",
            "mock": "Mock.mock(@character(\"upper\"))"
		},
		"character(lower)":{
			"type": "string",
            "mock": "Mock.mock(@character(\"lower\"))"
		},
		"character(number)":{
			"type": "string",
            "mock": "Mock.mock(@character(\"number\"))"
		},
		"character(symbol)":{
			"type": "string",
            "mock": "Mock.mock(@character(\"symbol\"))"
		},
		"character(aeiou)":{
			"type": "string",
            "mock": "Mock.mock(@character(\"aeiou\"))"
		}
	}
}`),
		},
		{
			Name:        VariableBool,
			Description: "随机一个bool值",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"boolean":{
			"type": "string",
            "mock": "Mock.mock(@boolean)"
		},
		"boolean()":{
			"type": "string",
            "mock": "Mock.mock(@boolean())"
		},
		"boolean_1_2_true":{
			"type": "string",
            "mock": "Mock.mock(@boolean(1,9,true))"
		}
	}
}`),
		},
		{
			Name:        VariableBoolean,
			Description: "随机一个bool值",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"boolean":{
			"type": "string",
            "mock": "Mock.mock(@boolean)"
		},
		"boolean()":{
			"type": "string",
            "mock": "Mock.mock(@boolean())"
		},
		"boolean_1_2_true":{
			"type": "string",
            "mock": "Mock.mock(@boolean(1,9,true))"
		}
	}
}`),
		},
		{
			Name:        VariableNatural,
			Description: "生成指定范围内的自然数(默认:1~MaxInt64)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"natural":{
			"type": "string",
            "mock": "Mock.mock(@natural)"
		},
		"natural()":{
			"type": "string",
            "mock": "Mock.mock(@natural())"
		},
		"natural_10000":{
			"type": "string",
            "mock": "Mock.mock(@natural(10000))"
		},
		"natural_60_100":{
			"type": "string",
            "mock": "Mock.mock(@natural(60,100))"
		}
	}
}`),
		},
		{
			Name:        VariableDate,
			Description: "随机一个日期字符串(默认layout:yyyy-MM-dd)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"data":{
			"type": "string",
            "mock": "Mock.mock(@date)"
		},
		"data()":{
			"type": "string",
            "mock": "Mock.mock(@date())"
		},
		"date(yyyy-MM-dd)":{
			"type": "string",
            "mock": "Mock.mock(@date(\"yyyy-MM-dd\"))"
		},
		"date(yy-MM-dd)":{
			"type": "string",
            "mock": "Mock.mock(@date(\"yy-MM-dd\"))"
		},
		"date(y-MM-dd)":{
			"type": "string",
            "mock": "Mock.mock(@date(\"y-MM-dd\"))"
		},
		"@date(y-M-d)":{
			"type": "string",
            "mock": "Mock.mock(@date(\"y-M-d\"))"
		},
		"date(yyyy yy y MM M dd d)":{
			"type": "string",
            "mock": "Mock.mock(@date(\"yyyy yy y MM M dd d\"))"
		}
	}
}`),
		},
		{
			Name:        VariableTime,
			Description: "随机一个时间(默认格式:15:04:05)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"time":{
			"type": "string",
            "mock": "Mock.mock(@time)"
		},
		"time()":{
			"type": "string",
            "mock": "Mock.mock(@time())"
		},
		"time(A HH:mm:ss)":{
			"type": "string",
            "mock": "Mock.mock(@time(\"A HH:mm:ss\"))"
		},
		"time(a HH:mm:ss)":{
			"type": "string",
            "mock": "Mock.mock(@time(\"a HH:mm:ss\"))"
		},
		"time(HH:mm:ss)":{
			"type": "string",
            "mock": "Mock.mock(@time(\"HH:mm:ss\"))"
		},
		"@time(H:m:s)":{
			"type": "string",
            "mock": "Mock.mock(@time(\"H:m:s\"))"
		}
	}
}`),
		},
		{
			Name:        VariableDateTime,
			Description: "随机一个详细时间(默认格式:2006-01-02 15:04:05)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"datetime":{
			"type": "string",
            "mock": "Mock.mock(@datetime)"
		},
		"datetime()":{
			"type": "string",
            "mock": "Mock.mock(@datetime())"
		},
		"datetime(yyyy-MM-dd A HH:mm:ss)":{
			"type": "string",
            "mock": "Mock.mock(@datetime(\"yyyy-MM-dd A HH:mm:ss\"))"
		},
		"datetime(yy-MM-dd a HH:mm:ss)":{
			"type": "string",
            "mock": "Mock.mock(@datetime(\"yy-MM-dd a HH:mm:ss\"))"
		},
		"datetime(y-MM-dd HH:mm:ss)":{
			"type": "string",
            "mock": "Mock.mock(@datetime(\"y-MM-dd HH:mm:ss\"))"
		},
		"@datetime(y-M-d H:m:s)":{
			"type": "string",
            "mock": "Mock.mock(@datetime(\"y-M-d H:m:s\"))"
		},
		"@datetime(yyyy yy y MM M dd d HH H hh h mm m ss s SS S A a T)":{
			"type": "string",
            "mock": "Mock.mock(@datetime(\"yyyy yy y MM M dd d HH H hh h mm m ss s SS S A a T\"))"
		}
	}
}`),
		},
		{
			Name:        VariableNow,
			Description: "返回当前时间指定数据(默认:2006-01-02 15:04:05)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"now":{
			"type": "string",
            "mock": "Mock.mock(@now)"
		},
		"now()":{
			"type": "string",
            "mock": "Mock.mock(@now())"
		},
		"now(year)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"year\"))"
		},
		"now(month)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"month\"))"
		},
		"now(week)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"week\"))"
		},
		"@now(day)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"day\"))"
		},
		"now(hour)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"hour\"))"
		},
		"now(minute)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"minute\"))"
		},
		"now(second)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"second\"))"
		},
		"@now(yyyy-MM-dd HH:mm:ss SS)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"yyyy-MM-dd HH:mm:ss SS\"))"
		},
		"@now(day,yyyy-MM-dd HH:mm:ss SS)":{
			"type": "string",
            "mock": "Mock.mock(@now(\"day\",\"yyyy-MM-dd HH:mm:ss SS\"))"
		}
	}
}`),
		},
		{
			Name:        VariableParagraph,
			Description: "生成指定范围的英文段落(默认:1~4)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"paragraph":{
			"type": "string",
            "mock": "Mock.mock(@paragraph)"
		},
		"paragraph()":{
			"type": "string",
            "mock": "Mock.mock(@paragraph())"
		},
		"paragraph(2)":{
			"type": "string",
            "mock": "Mock.mock(@paragraph(2))"
		},
		"paragraph(1,3)":{
			"type": "string",
            "mock": "Mock.mock(@paragraph(1,3))"
		}
	}
}`),
		},
		{
			Name:        VariableSentence,
			Description: "生成指定范围个数的句子(默认:5~20)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"sentence":{
			"type": "string",
            "mock": "Mock.mock(@sentence)"
		},
		"sentence()":{
			"type": "string",
            "mock": "Mock.mock(@sentence())"
		},
		"sentence(5)":{
			"type": "string",
            "mock": "Mock.mock(@sentence(5))"
		},
		"sentence(3,5)":{
			"type": "string",
            "mock": "Mock.mock(@sentence(3,5))"
		}
	}
}`),
		},
		{
			Name:        VariableWord,
			Description: "生成指定范围个数字母的单词(默认:3~10)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"word":{
			"type": "string",
            "mock": "Mock.mock(@word)"
		},
		"word()":{
			"type": "string",
            "mock": "Mock.mock(@word())"
		},
		"word(5)":{
			"type": "string",
            "mock": "Mock.mock(@word(5))"
		},
		"word(3,5)":{
			"type": "string",
            "mock": "Mock.mock(@word(3,5))"
		}
	}
}`),
		},
		{
			Name:        VariableTitle,
			Description: "生成指定范围个单词的标题(首字母大小)(默认:3~10)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"title":{
			"type": "string",
            "mock": "Mock.mock(@title)"
		},
		"title()":{
			"type": "string",
            "mock": "Mock.mock(@title())"
		},
		"title(5)":{
			"type": "string",
            "mock": "Mock.mock(@title(5))"
		},
		"title(3,5)":{
			"type": "string",
            "mock": "Mock.mock(@title(3,5))"
		}
	}
}`),
		},
		{
			Name:        VariableCParagraph,
			Description: "生成指定范围个数的中文段落(默认:2~5)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"cparagraph":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph)"
		},
		"cparagraph()":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph())"
		},
		"cparagraph(2)":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph(2))"
		},
		"cparagraph(1,3)":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph(1,3))"
		}
	}
}`),
		},
		{
			Name:        VariableCSentence,
			Description: "生成指定范围汉字的句子(默认:5~20)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"csentence":{
			"type": "string",
            "mock": "Mock.mock(@csentence)"
		},
		"csentence()":{
			"type": "string",
            "mock": "Mock.mock(@csentence())"
		},
		"csentence(5)":{
			"type": "string",
            "mock": "Mock.mock(@csentence(5))"
		},
		"csentence(3,5)":{
			"type": "string",
            "mock": "Mock.mock(@csentence(3,5))"
		}
	}
}`),
		},
		{
			Name:        VariableCWord,
			Description: "生成指定范围的中文汉字(默认:1)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"cword":{
			"type": "string",
            "mock": "Mock.mock(@cword)"
		},
		"cword()":{
			"type": "string",
            "mock": "Mock.mock(@cword())"
		},
		"cword(3)":{
			"type": "string",
            "mock": "Mock.mock(@cword(3))"
		},
		"cword(3,5)":{
			"type": "string",
            "mock": "Mock.mock(@cword(3,5))"
		},
		"cword(5,7)":{
			"type": "string",
            "mock": "Mock.mock(@cword(\"零一二三四五六七八九十\",5,7))"
		}
	}
}`),
		},
		{
			Name:        VariableCTitle,
			Description: "生成指定范围的中文标题(默认:3~6)",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"ctitle":{
			"type": "string",
            "mock": "Mock.mock(@ctitle)"
		},
		"ctitle()":{
			"type": "string",
            "mock": "Mock.mock(@ctitle())"
		},
		"ctitle(5)":{
			"type": "string",
            "mock": "Mock.mock(@ctitle(5))"
		},
		"ctitle(3,5)":{
			"type": "string",
            "mock": "Mock.mock(@ctitle(3,5))"
		}
	}
}`),
		},
		{
			Name:        VariableFirst,
			Description: "随机一个英文名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"first":{
			"type": "string",
            "mock": "Mock.mock(@first)"
		},
		"first()":{
			"type": "string",
            "mock": "Mock.mock(@first())"
		}
	}
}`),
		},
		{
			Name:        VariableLast,
			Description: "随机一个英文姓",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"last":{
			"type": "string",
            "mock": "Mock.mock(@last)"
		},
		"last()":{
			"type": "string",
            "mock": "Mock.mock(@last())"
		}
	}
}`),
		},
		{
			Name:        VariableName,
			Description: "生成一个英文全名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"name":{
			"type": "string",
            "mock": "Mock.mock(@name)"
		},
		"name()":{
			"type": "string",
            "mock": "Mock.mock(@name())"
		},
		"name(true)":{
			"type": "string",
            "mock": "Mock.mock(@name(true))"
		}
	}
}`),
		},
		{
			Name:        VariableCFirst,
			Description: "生成一个中文姓",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"cfirst":{
			"type": "string",
            "mock": "Mock.mock(@cfirst)"
		},
		"cfirst()":{
			"type": "string",
            "mock": "Mock.mock(@cfirst())"
		}
	}
}`),
		},
		{
			Name:        VariableCLast,
			Description: "生成一个中文名字",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"clast":{
			"type": "string",
            "mock": "Mock.mock(@clast)"
		},
		"clast()":{
			"type": "string",
            "mock": "Mock.mock(@clast())"
		}
	}
}`),
		},
		{
			Name:        VariableCname,
			Description: "生成一个中文全名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"cname":{
			"type": "string",
            "mock": "Mock.mock(@cname)"
		},
		"cname()":{
			"type": "string",
            "mock": "Mock.mock(@cname())"
		}
	}
}`),
		},
		{
			Name:        VariableUrl,
			Description: "生成一个地址",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"url":{
			"type": "string",
            "mock": "Mock.mock(@url)"
		},
		"url()":{
			"type": "string",
            "mock": "Mock.mock(@url())"
		}
	}
}`),
		},
		{
			Name:        VariableDomain,
			Description: "生成一个域名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"domain":{
			"type": "string",
            "mock": "Mock.mock(@domain)"
		},
		"domain()":{
			"type": "string",
            "mock": "Mock.mock(@domain())"
		}
	}
}`),
		},
		{
			Name:        VariableProtocol,
			Description: "生成一个协议",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"protocol":{
			"type": "string",
            "mock": "Mock.mock(@protocol)"
		},
		"protocol()":{
			"type": "string",
            "mock": "Mock.mock(@protocol())"
		}
	}
}`),
		},
		{
			Name:        VariableTld,
			Description: "生成一个域名后缀",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"tld":{
			"type": "string",
            "mock": "Mock.mock(@tld)"
		},
		"tld()":{
			"type": "string",
            "mock": "Mock.mock(@tld())"
		}
	}
}`),
		},
		{
			Name:        VariableEmail,
			Description: "生成一个邮箱",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"email":{
			"type": "string",
            "mock": "Mock.mock(@email)"
		},
		"email()":{
			"type": "string",
            "mock": "Mock.mock(@email())"
		}
	}
}`),
		},
		{
			Name:        VariableIp,
			Description: "生成一个IP",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"ip":{
			"type": "string",
            "mock": "Mock.mock(@ip)"
		},
		"ip()":{
			"type": "string",
            "mock": "Mock.mock(@ip())"
		}
	}
}`),
		},
		{
			Name:        VariableRegion,
			Description: "生成一个区域",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"region":{
			"type": "string",
            "mock": "Mock.mock(@region)"
		},
		"region()":{
			"type": "string",
            "mock": "Mock.mock(@region())"
		}
	}
}`),
		},
		{
			Name:        VariableProvince,
			Description: "生成一个省份名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"province":{
			"type": "string",
            "mock": "Mock.mock(@province)"
		},
		"province()":{
			"type": "string",
            "mock": "Mock.mock(@province())"
		}
	}
}`),
		},
		{
			Name:        VariableCity,
			Description: "生成一个市级名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"city":{
			"type": "string",
            "mock": "Mock.mock(@city)"
		},
		"city()":{
			"type": "string",
            "mock": "Mock.mock(@city())"
		},
		"city(true)":{
			"type": "string",
            "mock": "Mock.mock(@city(true))"
		}
	}
}`),
		},
		{
			Name:        VariableCounty,
			Description: "生成一个县级名",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"county":{
			"type": "string",
            "mock": "Mock.mock(@county)"
		},
		"county()":{
			"type": "string",
            "mock": "Mock.mock(@county())"
		},
		"county(true)":{
			"type": "string",
            "mock": "Mock.mock(@county(true))"
		}
	}
}`),
		},
		{
			Name:        VariableZip,
			Description: "生成一个6位邮编",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"zip":{
			"type": "string",
            "mock": "Mock.mock(@zip)"
		},
		"zip()":{
			"type": "string",
            "mock": "Mock.mock(@zip())"
		}
	}
}`),
		},
		{
			Name:        VariableCapitalize,
			Description: "对输入值首字母大写",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"capitalize":{
			"type": "string",
            "mock": "Mock.mock(@capitalize(\"hello\"))"
		},
		"capitalize()":{
			"type": "string",
            "mock": "Mock.mock(@capitalize(\"hello,world\"))"
		}
	}
}`),
		},
		{
			Name:        VariableUpper,
			Description: "对输入值转大写",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"upper":{
			"type": "string",
            "mock": "Mock.mock(@upper(\"hello\"))"
		}
	}
}`),
		},
		{
			Name:        VariableLower,
			Description: "对输入值转小写",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"lower":{
			"type": "string",
            "mock": "Mock.mock(@lower(\"HELLO,World\"))"
		}
	}
}`),
		},
		{
			Name:        VariablePick,
			Description: "从输入的值中选择其中一项",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"pick([a,e,i,o,u])":{
			"type": "string",
            "mock": "Mock.mock(@pick([\"a\", \"e\", \"i\", \"o\", \"u\"]))"
		}
	}
}`),
		},
		{
			Name:        VariableShuffle,
			Description: "对输入的数组重新洗牌",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"shuffle([a,e,i,o,u])":{
			"type": "string",
            "mock": "Mock.mock(@shuffle([\"a\", \"e\", \"i\", \"o\", \"u\"]))"
		}
	}
}`),
		},
		{
			Name:        VariableGuid,
			Description: "生成一个GUID",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"guid":{
			"type": "string",
            "mock": "Mock.mock(@guid)"
		},
		"guid()":{
			"type": "string",
            "mock": "Mock.mock(@guid())"
		}
	}
}`),
		},
		{
			Name:        VariableId,
			Description: "生成全是数字的字符串的ID",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"id":{
			"type": "string",
            "mock": "Mock.mock(@id)"
		},
		"id()":{
			"type": "string",
            "mock": "Mock.mock(@id())"
		}
	}
}`),
		},
		{
			Name:        VariableColor,
			Description: "生成一个16进制表示的颜色",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"color":{
			"type": "string",
            "mock": "Mock.mock(@color)"
		},
		"color()":{
			"type": "string",
            "mock": "Mock.mock(@color())"
		}
	}
}`),
		},
		{
			Name:        VariableHex,
			Description: "生成一个16进制表示的数字",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"hex":{
			"type": "string",
            "mock": "Mock.mock(@hex)"
		},
		"hex()":{
			"type": "string",
            "mock": "Mock.mock(@hex())"
		}
	}
}`),
		},
		{
			Name:        VariableRGB,
			Description: "生成一个RGB颜色值",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"rgb":{
			"type": "string",
            "mock": "Mock.mock(@rgb)"
		},
		"rgb()":{
			"type": "string",
            "mock": "Mock.mock(@rgb())"
		}
	}
}`),
		},
		{
			Name:        VariableRGBA,
			Description: "生成一个RGBA的颜色值",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"rgba":{
			"type": "string",
            "mock": "Mock.mock(@rgba)"
		},
		"rgba()":{
			"type": "string",
            "mock": "Mock.mock(@rgba())"
		}
	}
}`),
		},
		{
			Name:        VariableHSL,
			Description: "生成HSL的颜色值",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"hsl":{
			"type": "string",
            "mock": "Mock.mock(@hsl)"
		},
		"hsl()":{
			"type": "string",
            "mock": "Mock.mock(@hsl())"
		}
	}
}`),
		},
		{
			Name:        VariableImage,
			Description: "生成一个图片",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"image(200x100)":{
			"type": "string",
            "mock": "Mock.mock(@image(\"200x100\"))"
		},
		"image(200*100,#FF6600)":{
			"type": "string",
            "mock": "Mock.mock(@image(\"200x100\",\"#FF6600\"))"
		},
		"image(200x100,#FF6600,Hello)":{
			"type": "string",
            "mock": "Mock.mock(@image(\"200x100\",\"#FF6600\",\"Hello\"))"
		},
		"image(200x100,#FF6600,#FFFFFF,Hello)":{
			"type": "string",
            "mock": "Mock.mock(@image(\"200x100\",\"#FF6600\",\"#FFFFFF\",\"Hello\"))"
		},
		"image(200x100,#FF6600,#FFFFFF,png,Hello)":{
			"type": "string",
            "mock": "Mock.mock(@image(\"200x100\",\"#FF6600\",\"#FFFFFF\",\"png\",\"Hello\"))"
		}
	}
}`),
		},
		{
			Name:        VariableDataImage,
			Description: "生成一个图片",
			Examples: json.RawMessage(`{
    "type": "object",
    "properties": {
		"dataImage(200x100)":{
			"type": "string",
            "mock": "Mock.mock(@dataImage(\"200x100\"))"
		},
		"dataImage(200*100,#FF6600)":{
			"type": "string",
            "mock": "Mock.mock(@dataImage(\"200x100\",\"hello Mock.js!\"))"
		}
	}
}`),
		},
	}
)
