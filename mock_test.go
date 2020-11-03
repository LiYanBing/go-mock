package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liyanbing/go-mock/images"
	"github.com/liyanbing/go-mock/variables"
)

const (
	schema = `
{
    "type": "object",
    "properties": {
		"pick":{
			"type": "string",
            "mock": "Mock.mock(@pick([1,2,3,4,5,6,7]))"
		},
		"shuffle":{
			"type": "string",
            "mock": "Mock.mock(@shuffle([1,2,3,4,5,6,7]))"
		},
		"lower":{
			"type": "string",
            "mock": "Mock.mock(@lower(\"Hello\"))"
		},
		"lower1":{
			"type": "string",
            "mock": "Mock.mock(@lower(\"\"))"
		},
		"upper":{
			"type": "string",
            "mock": "Mock.mock(@upper(\"hello\"))"
		},
		"upper1":{
			"type": "string",
            "mock": "Mock.mock(@upper(\"\"))"
		},
		"capitalize":{
			"type": "string",
            "mock": "Mock.mock(@capitalize(\"hello\"))"
		},
		"capitalize1":{
			"type": "string",
            "mock": "Mock.mock(@capitalize(\"\"))"
		},
		"ip":{
			"type": "string",
            "mock": "Mock.mock(@ip)"
		},
		"ip1":{
			"type": "string",
            "mock": "Mock.mock(@ip())"
		},
		"email":{
			"type": "string",
            "mock": "Mock.mock(@email)"
		},
		"email1":{
			"type": "string",
            "mock": "Mock.mock(@email())"
		},
		"tld":{
			"type": "string",
            "mock": "Mock.mock(@tld)"
		},
		"tld1":{
			"type": "string",
            "mock": "Mock.mock(@tld())"
		},
		"protocol":{
			"type": "string",
            "mock": "Mock.mock(@protocol)"
		},
		"protocol1":{
			"type": "string",
            "mock": "Mock.mock(@protocol())"
		},
		"domain":{
			"type": "string",
            "mock": "Mock.mock(@domain)"
		},
		"domain1":{
			"type": "string",
            "mock": "Mock.mock(@domain())"
		},
		"url":{
			"type": "string",
            "mock": "Mock.mock(@url)"
		},
		"url1":{
			"type": "string",
            "mock": "Mock.mock(@url())"
		},
		"cfirst":{
			"type": "string",
            "mock": "Mock.mock(@cfirst)"
		},
		"cfirst1":{
			"type": "string",
            "mock": "Mock.mock(@cfirst())"
		},
		"clast":{
			"type": "string",
            "mock": "Mock.mock(@clast)"
		},
		"clast1":{
			"type": "string",
            "mock": "Mock.mock(@clast())"
		},
		"cname":{
			"type": "string",
            "mock": "Mock.mock(@cname)"
		},
		"cname1":{
			"type": "string",
            "mock": "Mock.mock(@cname())"
		},
		"first":{
			"type": "string",
            "mock": "Mock.mock(@first)"
		},
		"first1":{
			"type": "string",
            "mock": "Mock.mock(@first())"
		},
		"last":{
			"type": "string",
            "mock": "Mock.mock(@last)"
		},
		"last1":{
			"type": "string",
            "mock": "Mock.mock(@last())"
		},
		"name":{
			"type": "string",
            "mock": "Mock.mock(@name)"
		},
		"name1":{
			"type": "string",
            "mock": "Mock.mock(@name())"
		},
		"name2":{
			"type": "string",
            "mock": "Mock.mock(@name(true))"
		},
		"cparagraph":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph)"
		},
		"cparagraph1":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph())"
		},
		"cparagraph_2":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph(2))"
		},
		"cparagraph_1_3":{
			"type": "string",
            "mock": "Mock.mock(@cparagraph(1,3))"
		},
		"csentence": {
			"type": "string",
            "mock": "Mock.mock(@csentence)"
		},
		"csentence1": {
			"type": "string",
            "mock": "Mock.mock(@csentence())"
		},
		"csentence_5": {
			"type": "string",
            "mock": "Mock.mock(@csentence(5))"
		},
		"csentence_3_5": {
			"type": "string",
            "mock": "Mock.mock(@csentence(3,5))"
		},
		"ctitle": {
			"type": "string",
            "mock": "Mock.mock(@ctitle)"
		},
		"ctitle1": {
			"type": "string",
            "mock": "Mock.mock(@ctitle())"
		},
		"ctitle_3_5": {
			"type": "string",
            "mock": "Mock.mock(@ctitle(3,5))"
		},
		"cword5": {
			"type": "string",
            "mock": "Mock.mock(@cword(5))"
		},
		"cword1": {
			"type": "string",
            "mock": "Mock.mock(@cword())"
		},
		"cword_pool": {
			"type": "string",
            "mock": "Mock.mock(@cword(\"零一二三四五六七八九十\"))"
		},
		"cword__3": {
			"type": "string",
            "mock": "Mock.mock(@cword(3))"
		},
		"cword_pool_3": {
			"type": "string",
            "mock": "Mock.mock(@cword(\"零一二三四五六七八九十\",3))"
		},
		"cword_3_5": {
			"type": "string",
            "mock": "Mock.mock(@cword(3,5))"
		},
		"cword_pool_5_7": {
			"type": "string",
            "mock": "Mock.mock(@cword(\"零一二三四五六七八九十\",5,7))"
		},
		"paragraph": {
			"type": "string",
            "mock": "Mock.mock(@paragraph)"
		},
		"paragraph1": {
			"type": "string",
            "mock": "Mock.mock(@paragraph())"
		},
		"paragraph_2": {
			"type": "string",
            "mock": "Mock.mock(@paragraph(2))"
		},
		"paragraph_1_3": {
			"type": "string",
            "mock": "Mock.mock(@paragraph(1,3))"
		},
        "sentence": {
			"type": "string",
            "mock": "Mock.mock(@sentence)"
		},
        "sentence1": {
			"type": "string",
            "mock": "Mock.mock(@sentence())"
		},
        "sentence_5": {
			"type": "string",
            "mock": "Mock.mock(@sentence(5))"
		},
        "sentence_3_5": {
			"type": "string",
            "mock": "Mock.mock(@sentence(3,5))"
		},
        "title": {
			"type": "string",
            "mock": "Mock.mock(@title)"
		},
        "title1": {
			"type": "string",
            "mock": "Mock.mock(@title())"
		},
        "title_5": {
			"type": "string",
            "mock": "Mock.mock(@title(5))"
		},
        "title_3_5": {
			"type": "string",
            "mock": "Mock.mock(@title(3,5))"
		},
        "word": {
			"type": "string",
            "mock": "Mock.mock(@word)"
		},
        "word1": {
			"type": "string",
            "mock": "Mock.mock(@word())"
		},
        "word_5": {
			"type": "string",
            "mock": "Mock.mock(@word(5))"
		},
        "word_3_5": {
			"type": "string",
            "mock": "Mock.mock(@word(3,5))"
		},
        "now": {
			"type": "string",
            "mock": "Mock.mock(@now)"
        },
        "now1": {
			"type": "string",
            "mock": "Mock.mock(@now())"
        },
        "now_year": {
			"type": "string",
            "mock": "Mock.mock(@now(\"year\"))"
        },
        "now_month": {
			"type": "string",
            "mock": "Mock.mock(@now(\"month\"))"
        },
        "now_week": {
			"type": "string",
            "mock": "Mock.mock(@now(\"week\"))"
        },
        "now_day": {
			"type": "string",
            "mock": "Mock.mock(@now(\"day\"))"
        },
        "now_hour": {
			"type": "string",
            "mock": "Mock.mock(@now(\"hour\"))"
        },
        "now_minute": {
			"type": "string",
            "mock": "Mock.mock(@now(\"minute\"))"
        },
        "now_second": {
			"type": "string",
            "mock": "Mock.mock(@now(\"second\"))"
        },
        "now_month_format": {
			"type": "string",
            "mock": "Mock.mock(@now(\"month\",\"yyyy-MM-dd HH:mm:ss SS\"))"
        },
        "now_format": {
			"type": "string",
            "mock": "Mock.mock(@now(\"yyyy-MM-dd HH:mm:ss SS\"))"
        },
        "time": {
			"type": "string",
            "mock": "Mock.mock(@time)"
        },
        "time1": {
			"type": "string",
            "mock": "Mock.mock(@time(\"A HH:mm:ss\"))"
        },
        "time2": {
			"type": "string",
            "mock": "Mock.mock(@time(\"a HH:mm:ss\"))"
        },
        "time3": {
			"type": "string",
            "mock": "Mock.mock(@time(\"H:m:s\"))"
        },
        "datetime": {
			"type": "string",
            "mock": "Mock.mock(@datetime)"
        },
        "datetime1": {
			"type": "string",
            "mock": "Mock.mock(@datetime())"
        },
        "datetime2": {
			"type": "string",
            "mock": "Mock.mock(@datetime(\"yyyy-MM-dd A HH:mm:ss\"))"
        },
        "datetime3": {
			"type": "string",
            "mock": "Mock.mock(@datetime(\"yy-MM-dd a HH:mm:ss\"))"
        },
        "datetime3": {
			"type": "string",
            "mock": "Mock.mock(@datetime(\"y-MM-dd HH:mm:ss\"))"
        },
        "datetime4": {
			"type": "string",
            "mock": "Mock.mock(@datetime(\"y-M-d H:m:s\"))"
        },
        "region": {
            "type": "string",
            "mock": "Mock.mock(@region())"
        },
        "province": {
            "type": "string",
            "mock": "Mock.mock(@province())"
        },
        "city": {
			"type": "string",
            "mock": "Mock.mock(@city)"
        },
        "city1": {
			"type": "string",
            "mock": "Mock.mock(@city(true))"
        },
        "country": {
			"type": "string",
            "mock": "Mock.mock(@county)"
        },
        "country1": {
			"type": "string",
            "mock": "Mock.mock(@county(true))"
        },
        "mock_test": {
            "type": "string",
          	"mock": "Mock.mock({\"name|1-9\": \"张\"})"
        },
        "mock_object_test1": {
            "type": "string",
            "mock": "Mock.mock({\"object|2\":{\"310000\":\"上海市\",\"320000\":\"江苏省\",\"330000\":\"浙江省\",\"340000\":\"安徽省\"}})"
        },
        "mock_object_test2": {
            "type": "string",
            "mock": "Mock.mock({\"object|2-4\":{\"110000\":\"北京市\",\"120000\":\"天津市\",\"130000\":\"河北省\",\"140000\":\"山西省\"}})"
        },
        "mock_object_test3": {
            "type": "string",
            "mock": "Mock.mock({\"array|1\":[\"AMD\",\"CMD\",\"UMD\"]})"
        },
        "mock_object_test4": {
            "type": "string",
            "mock": "Mock.mock({\"array|1-10\":[{\"name|+1\":[\"Hello\",\"Mock.js\",\"!\"]}]})"
        },
        "mock_object_test5": {
            "type": "string",
            "mock": "Mock.mock({\"array|1-10\":[\"Mock.js\"]})"
        },
        "mock_object_test5": {
            "type": "string",
            "mock": "Mock.mock({\"array|1-10\":[\"Hello\",\"Mock.js\",\"!\"]})"
        },
        "mock_object_test6": {
            "type": "string",
            "mock": "Mock.mock({\"array|3\":[\"Mock.js\"]})"
        },
        "mock_object_test7": {
             "type": "string",
             "mock": "Mock.mock({\"array|3\":[\"Hello\",\"Mock.js\",\"!\"]})"
        },
        "string_upper": {
            "type": "string",
            "mock": "Mock.mock(@string(\"upper\"))"
        },
        "string_lower": {
            "type": "string",
            "mock": "Mock.mock(@string(\"\"))"
        },
        "string_symbol": {
            "type": "string",
            "mock": "Mock.mock(@string(\"symbol\"))"
        },
        "string_number": {
            "type": "string",
            "mock": "Mock.mock(@string(\"number\"))"
        },
        "integer": {
            "type": "integer",
            "mock": "Mock.mock(@integer(10,60))"
        },
		"date": {
            "type": "string",
            "mock": "Mock.mock(@date(\"yyyy-MM-dd A HH:mm:ss\"))"
        },
        "interface_method_id": {
            "type": "string"
        },
        "method": {
            "type": "string",
            "enum": [
                "Get",
                "Put",
                "Post",
                "Delete",
                "Options",
                "Head",
                "Patch",
                "Trace"
            ],
            "default": "Get"
        },
        "parameters": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "reference": {
                        "type": "string"
                    },
                    "name": {
                        "type": "string"
                    },
                    "in": {
                        "type": "string",
                        "enum": [
                            "Header",
                            "Query",
                            "Path",
                            "Cookie"
                        ],
                        "default": "Header"
                    },
                    "description": {
                        "type": "string"
                    },
                    "required": {
                        "type": "boolean",
                        "format": "boolean"
                    },
                    "deprecated": {
                        "type": "boolean",
                        "format": "boolean"
                    },
                    "allow_empty_value": {
                        "type": "boolean",
                        "format": "boolean"
                    },
                    "style": {
                        "type": "string",
                        "enum": [
                            "Matrix",
                            "Label",
                            "Form",
                            "Simple",
                            "SpaceDelimited",
                            "PipeDelimited",
                            "DeepObject"
                        ],
                        "default": "Matrix"
                    },
                    "explode": {
                        "type": "boolean",
                        "format": "boolean"
                    },
                    "allow_reserved": {
                        "type": "boolean",
                        "format": "boolean"
                    },
                    "schema": {
                        "type": "string"
                    },
                    "examples": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "object",
                            "properties": {
                                "reference": {
                                    "type": "string"
                                },
                                "summary": {
                                    "type": "string"
                                },
                                "description": {
                                    "type": "string"
                                },
                                "value": {
                                    "type": "string"
                                },
                                "external_value": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "content": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "object",
                            "properties": {
                                "schema": {
                                    "type": "string"
                                },
                                "examples": {
                                    "type": "object",
                                    "additionalProperties": {
                                        "type": "object",
                                        "properties": {
                                            "reference": {
                                                "type": "string"
                                            },
                                            "summary": {
                                                "type": "string"
                                            },
                                            "description": {
                                                "type": "string"
                                            },
                                            "value": {
                                                "type": "string"
                                            },
                                            "external_value": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                "title": "Header 跟Parameter同结构体，只是Header不包含 name跟in 字段"
            }
        },
        "request_body": {
            "type": "object",
            "properties": {
                "reference": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "required": {
                    "type": "boolean",
                    "format": "boolean"
                },
                "content": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "object",
                        "properties": {
                            "schema": {
                                "type": "string"
                            },
                            "examples": {
                                "type": "object",
                                "additionalProperties": {
                                    "type": "object",
                                    "properties": {
                                        "reference": {
                                            "type": "string"
                                        },
                                        "summary": {
                                            "type": "string"
                                        },
                                        "description": {
                                            "type": "string"
                                        },
                                        "value": {
                                            "type": "string"
                                        },
                                        "external_value": {
                                            "type": "string"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "responses": {
            "type": "object",
            "additionalProperties": {
                "type": "object",
                "properties": {
                    "reference": {
                        "type": "string"
                    },
                    "description": {
                        "type": "string"
                    },
                    "headers": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "object",
                            "properties": {
                                "reference": {
                                    "type": "string"
                                },
                                "name": {
                                    "type": "string"
                                },
                                "in": {
                                    "type": "string",
                                    "enum": [
                                        "Header",
                                        "Query",
                                        "Path",
                                        "Cookie"
                                    ],
                                    "default": "Header"
                                },
                                "description": {
                                    "type": "string"
                                },
                                "required": {
                                    "type": "boolean",
                                    "format": "boolean"
                                },
                                "deprecated": {
                                    "type": "boolean",
                                    "format": "boolean"
                                },
                                "allow_empty_value": {
                                    "type": "boolean",
                                    "format": "boolean"
                                },
                                "style": {
                                    "type": "string",
                                    "enum": [
                                        "Matrix",
                                        "Label",
                                        "Form",
                                        "Simple",
                                        "SpaceDelimited",
                                        "PipeDelimited",
                                        "DeepObject"
                                    ],
                                    "default": "Matrix"
                                },
                                "explode": {
                                    "type": "boolean",
                                    "format": "boolean"
                                },
                                "allow_reserved": {
                                    "type": "boolean",
                                    "format": "boolean"
                                },
                                "schema": {
                                    "type": "string"
                                },
                                "examples": {
                                    "type": "object",
                                    "additionalProperties": {
                                        "type": "object",
                                        "properties": {
                                            "reference": {
                                                "type": "string"
                                            },
                                            "summary": {
                                                "type": "string"
                                            },
                                            "description": {
                                                "type": "string"
                                            },
                                            "value": {
                                                "type": "string"
                                            },
                                            "external_value": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                },
                                "content": {
                                    "type": "object",
                                    "additionalProperties": {
                                        "type": "object",
                                        "properties": {
                                            "schema": {
                                                "type": "string"
                                            },
                                            "examples": {
                                                "type": "object",
                                                "additionalProperties": {
                                                    "type": "object",
                                                    "properties": {
                                                        "reference": {
                                                            "type": "string"
                                                        },
                                                        "summary": {
                                                            "type": "string"
                                                        },
                                                        "description": {
                                                            "type": "string"
                                                        },
                                                        "value": {
                                                            "type": "string"
                                                        },
                                                        "external_value": {
                                                            "type": "string"
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            },
                            "title": "Header 跟Parameter同结构体，只是Header不包含 name跟in 字段"
                        }
                    },
                    "content": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "object",
                            "properties": {
                                "schema": {
                                    "type": "string"
                                },
                                "examples": {
                                    "type": "object",
                                    "additionalProperties": {
                                        "type": "object",
                                        "properties": {
                                            "reference": {
                                                "type": "string"
                                            },
                                            "summary": {
                                                "type": "string"
                                            },
                                            "description": {
                                                "type": "string"
                                            },
                                            "value": {
                                                "type": "string"
                                            },
                                            "external_value": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "description": {
            "type": "string"
        },
        "security": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "security_id": {
                        "type": "string"
                    },
                    "scopes": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "title": "about interface method"
}
`
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

// 当前test用于debug
func TestMock01(t *testing.T) {
	m, err := NewMock(schema)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(m.MockString(context.Background()))
}

// 每添加一个variable需要再次添加对应的单元测试查看结果
func TestMockAll(t *testing.T) {
	variables.RegisterImageURLAdaptor(variables.ImageURLFunc(func(s string) string {
		return fmt.Sprintf("%v/images/%v", "http://localhost:7070", s)
	}))

	cases := []struct {
		Name   string
		Schema string
	}{
		// boolean
		{
			Name: "boolean",
			Schema: `{
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
}`,
		},
		// natural
		{
			Name: "natural",
			Schema: `{
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
}`,
		},
		// integer
		{
			Name: "integer",
			Schema: `{
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
}`,
		},
		// float
		{
			Name: "float",
			Schema: `{
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
}`,
		},
		// character
		{
			Name: "character",
			Schema: `{
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
}`,
		},
		// string
		{
			Name: "string",
			Schema: `{
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
}`,
		},
		// date
		{
			Name: "data",
			Schema: `{
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
}`,
		},
		// time
		{
			Name: "time",
			Schema: `{
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
}`,
		},
		// datetime
		{
			Name: "datetime",
			Schema: `{
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
}`,
		},
		// now
		{
			Name: "now",
			Schema: `{
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
}`,
		},
		// paragraph
		{
			Name: "paragraph",
			Schema: `{
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
}`,
		},
		// sentence
		{
			Name: "sentence",
			Schema: `{
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
}`,
		},
		// word
		{
			Name: "word",
			Schema: `{
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
}`,
		},
		// title
		{
			Name: "title",
			Schema: `{
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
}`,
		},
		// cparagraph
		{
			Name: "cparagraph",
			Schema: `{
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
}`,
		},
		// csentence
		{
			Name: "csentence",
			Schema: `{
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
}`,
		},
		// cword
		{
			Name: "cword",
			Schema: `{
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
}`,
		},
		// ctitle
		{
			Name: "ctitle",
			Schema: `{
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
}`,
		},
		// first
		{
			Name: "first",
			Schema: `{
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
}`,
		},
		// last
		{
			Name: "last",
			Schema: `{
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
}`,
		},
		// name
		{
			Name: "name",
			Schema: `{
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
}`,
		},
		// cfirst
		{
			Name: "cfirst",
			Schema: `{
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
}`,
		},
		// clast
		{
			Name: "clast",
			Schema: `{
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
}`,
		},
		// cname
		{
			Name: "cname",
			Schema: `{
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
}`,
		},
		// url
		{
			Name: "url",
			Schema: `{
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
}`,
		},
		// domain
		{
			Name: "domain",
			Schema: `{
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
}`,
		},
		// protocol
		{
			Name: "protocol",
			Schema: `{
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
}`,
		},
		// tld
		{
			Name: "tld",
			Schema: `{
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
}`,
		},
		// email
		{
			Name: "email",
			Schema: `{
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
}`,
		},
		// ip
		{
			Name: "ip",
			Schema: `{
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
}`,
		},
		// region
		{
			Name: "region",
			Schema: `{
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
}`,
		},
		// province
		{
			Name: "province",
			Schema: `{
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
}`,
		},
		// city
		{
			Name: "city",
			Schema: `{
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
}`,
		},
		// county
		{
			Name: "county",
			Schema: `{
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
}`,
		},
		// zip
		{
			Name: "zip",
			Schema: `{
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
}`,
		},
		// capitalize
		{
			Name: "capitalize",
			Schema: `{
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
}`,
		},
		// upper
		{
			Name: "upper",
			Schema: `{
    "type": "object",
    "properties": {
		"upper":{
			"type": "string",
            "mock": "Mock.mock(@upper(\"hello\"))"
		}
	}
}`,
		},
		// lower
		{
			Name: "lower",
			Schema: `{
    "type": "object",
    "properties": {
		"lower":{
			"type": "string",
            "mock": "Mock.mock(@lower(\"HELLO,World\"))"
		}
	}
}`,
		},
		// pick
		{
			Name: "pick",
			Schema: `{
    "type": "object",
    "properties": {
		"pick([a,e,i,o,u])":{
			"type": "string",
            "mock": "Mock.mock(@pick([\"a\", \"e\", \"i\", \"o\", \"u\"]))"
		}
	}
}`,
		},
		// shuffle
		{
			Name: "shuffle",
			Schema: `{
    "type": "object",
    "properties": {
		"shuffle([a,e,i,o,u])":{
			"type": "string",
            "mock": "Mock.mock(@shuffle([\"a\", \"e\", \"i\", \"o\", \"u\"]))"
		}
	}
}`,
		},
		// guid
		{
			Name: "guid",
			Schema: `{
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
}`,
		},
		// id
		{
			Name: "id",
			Schema: `{
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
}`,
		},
		// color
		{
			Name: "color",
			Schema: `{
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
}`,
		},
		// hex
		{
			Name: "hex",
			Schema: `{
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
}`,
		},
		// rgb
		{
			Name: "rgb",
			Schema: `{
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
}`,
		},
		// rgba
		{
			Name: "rgba",
			Schema: `{
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
}`,
		},
		// hsl
		{
			Name: "hsl",
			Schema: `{
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
}`,
		},
		// @image
		{
			Name: "image",
			Schema: `{
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
}`,
		},
		// @dataImage
		{
			Name: "dataImage",
			Schema: `{
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
}`,
		},
	}

	for _, value := range cases {
		m, err := NewMock(value.Schema)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("name: ", value.Name, "=>", m.MockString(context.Background()))
	}

}

func BenchmarkMock_MockString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, err := NewMock(schema)
		if err != nil {
			b.Error(err)
			return
		}
		m.MockString(context.Background())
	}
}

func TestMock(t *testing.T) {
	cases := []struct {
		Schema string
	}{
		{
			Schema: schema,
		},
		{
			Schema: "Mock.mock({\"name|1-9\": \"张\"})",
		},
		{
			Schema: "Mock.mock({\"object|2\":{\"310000\":\"上海市\",\"320000\":\"江苏省\",\"330000\":\"浙江省\",\"340000\":\"安徽省\"}})",
		},
	}

	for _, value := range cases {
		m, err := NewMock(value.Schema)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(m.MockString(context.Background()))
	}
}

func TestNewMockWithJsonSchemaData(t *testing.T) {
	var schemaData map[string]interface{}
	err := json.Unmarshal([]byte(schema), &schemaData)
	if err != nil {
		t.Error(err)
		return
	}

	m, err := NewMockWithJsonSchemaData(schemaData)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(m.MockString(context.Background()))
}

func TestNewMockWithMockSchemaData(t *testing.T) {
	cases := []struct {
		Data map[string]interface{}
	}{
		{
			Data: map[string]interface{}{
				"object|2": map[string]interface{}{
					"310000": "上海市",
					"320000": "江苏省",
					"330000": "浙江省",
					"340000": "安徽省",
				},
			},
		},
	}

	for _, value := range cases {
		m, err := NewMockWithMockSchemaData(value.Data)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(m.MockString(context.Background()))
	}
}

func TestImages(t *testing.T) {
	t.Skipf("images")
	address := "127.0.0.1:7070"
	r := gin.Default()
	variables.RegisterImageURLAdaptor(variables.ImageURLFunc(func(s string) string {
		return fmt.Sprintf("%v/images/%v", address, s)
	}))

	r.GET("/images", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"image_id": images.GetImageId("200x100", "#50B347", "#FFFFFF", "png", "golang")})
		return
	})

	r.GET("/images/:image_id", func(c *gin.Context) {
		imageID := c.Param("image_id")
		_, err := images.GenImageWithID(imageID, c.Writer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.Header("Content-Type", "image/png")
		return
	})
	err := r.Run(address)
	if err != nil {
		t.Error(err)
	}
}
