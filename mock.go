package mock

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/liyanbing/go-mock/node"
)

var mockFuncRegex = regexp.MustCompile(`^Mock.mock\((.*)\)`)

var (
	ErrInvalidJsonSchema     = errors.New("invalid json_schema")
	ErrInvalidMockFuncType   = errors.New("invalid mock func")
	ErrInvalidMockSchemaData = errors.New("invalid mock schema data")
)

type Mock struct {
	Root *node.Node `json:"r"`
}

func (s *Mock) Mock(ctx context.Context) interface{} {
	if s.Root == nil {
		return nil
	}
	return s.Root.Mock(ctx)
}

func (s *Mock) MockString(ctx context.Context) string {
	if s.Root == nil {
		return ""
	}
	return s.Root.MockJson(ctx)
}

// 入参必须是json格式的mock对象
/** 例如：
{
  "object|2": {
    "310000": "上海市",
    "320000": "江苏省",
    "330000": "浙江省",
    "340000": "安徽省"
  }
}
*/
func NewMockWithMockSchemaData(data interface{}) (*Mock, error) {
	_, ok := data.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidMockSchemaData
	}

	root, err := parseMockSchemaData("", data)
	if err != nil {
		return nil, err
	}

	return &Mock{
		Root: root,
	}, nil
}

// 入参必须是json_schema的对象
/** 例如
  {
    "type": "object",
    "properties": {
        "mock_test": {
            "type": "string",
          	"mock": "Mock.mock({\"name|1-9\": \"张\"})"
        },
        "mock_object_test1": {
            "type": "string",
            "mock": "Mock.mock({\"object|2\":{\"310000\":\"上海市\",\"320000\":\"江苏省\",\"330000\":\"浙江省\",\"340000\":\"安徽省\"}})"
        }
     }
  }
*/
func NewMockWithJsonSchemaData(data interface{}) (*Mock, error) {
	jsonData, ok := data.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidMockSchemaData
	}

	root, err := parseJsonSchemaData("", jsonData)
	if err != nil {
		return nil, err
	}

	return &Mock{
		Root: root,
	}, nil
}

// 入参可以是json_schema 的字符串或者是以Mock.mock开头的mock字符串
func NewMock(schema string) (*Mock, error) {
	if strings.HasPrefix(schema, "Mock.mock") {
		root, err := parseMockSchema("", schema)
		if err != nil {
			return nil, err
		}

		return &Mock{
			Root: root,
		}, nil
	}

	var schemaData map[string]interface{}
	err := json.Unmarshal([]byte(schema), &schemaData)
	if err != nil {
		return nil, ErrInvalidJsonSchema
	}

	root, err := parseJsonSchemaData("", schemaData)
	if err != nil {
		return nil, err
	}

	return &Mock{
		Root: root,
	}, nil
}

func parseJsonArray(name string, arrayJson map[string]interface{}) (*node.Node, error) {
	arrayItems, ok := arrayJson["items"]
	if !ok {
		return nil, ErrInvalidJsonSchema
	}

	items, ok := arrayItems.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidJsonSchema
	}

	item, err := parseJsonSchemaData("", items)
	if err != nil {
		return nil, err
	}

	na, rule := parseFieldName(name)
	return &node.Node{
		Name:  na,
		Rule:  rule,
		Type:  node.TypeOfNodeArray,
		Items: []*node.Node{item},
	}, nil
}

func parseJsonObject(name string, objectJson map[string]interface{}) (*node.Node, error) {
	var (
		properties interface{}
		ok         bool
	)

	properties, ok = objectJson["properties"]
	additional := false
	if !ok {
		properties, ok = objectJson["additionalProperties"]
		if !ok {
			return nil, ErrInvalidJsonSchema
		}
		additional = true
	}

	propertiesMap, ok := properties.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidJsonSchema
	}

	if additional {
		return parseJsonObject(name, propertiesMap)
	}

	items := make([]*node.Node, 0, len(propertiesMap))
	for key, value := range propertiesMap {
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, ErrInvalidJsonSchema
		}

		item, err := parseJsonSchemaData(key, valueMap)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	na, rule := parseFieldName(name)
	return &node.Node{
		Name:  na,
		Rule:  rule,
		Type:  node.TypeOfNodeMap,
		Items: items,
	}, nil
}

func parseJsonString(name string, objectJson map[string]interface{}) (*node.Node, error) {
	enum, ok := objectJson["enum"]
	enumArr, ok1 := enum.([]interface{})

	if ok && ok1 && len(enumArr) > 0 {
		items := make([]*node.Node, 0, len(enumArr))
		for _, enum := range enumArr {
			enumStr, ok := enum.(string)
			if !ok {
				continue
			}

			items = append(items, &node.Node{
				Name:  enumStr,
				Type:  node.TypeOfNodeSingle,
				Value: enumStr,
			})
		}

		// 这里不需要解析name;这里是一个枚举rule是一个固定的选择枚举中的一个即可
		return &node.Node{
			Name:  name,
			Type:  node.TypeOfNodeArray,
			Rule:  "1",
			Value: "",
			Items: items,
		}, nil
	}
	return parseJsonSingleFileWithDefault(name, "Mock.mock(@string(1,10))", objectJson)
}

func parseJsonSingleFileWithDefault(name, mock string, obj map[string]interface{}) (*node.Node, error) {
	mockField, ok := obj["mock"]
	if ok {
		mockFieldStr, ok := mockField.(string)
		if !ok {
			return nil, ErrInvalidMockFuncType
		}
		mock = mockFieldStr
	}
	return parseMockSchema(name, mock)
}

func parseJsonSchemaData(name string, in map[string]interface{}) (*node.Node, error) {
	tt, ok := in["type"]
	if !ok {
		return nil, ErrInvalidJsonSchema
	}

	typeStr, ok := tt.(string)
	if !ok {
		return nil, ErrInvalidJsonSchema
	}
	mockValue := ""

	switch strings.ToLower(typeStr) {
	case "object":
		return parseJsonObject(name, in)
	case "array":
		return parseJsonArray(name, in)
	case "string":
		return parseJsonString(name, in)
	case "number":
		mockValue = "Mock.mock(@integer(1,10))"
	case "integer":
		mockValue = "Mock.mock(@integer(1,10))"
	case "bool", "boolean":
		mockValue = "Mock.mock(@bool(1,2))"
	case "any":
		mockValue = "Mock.mock(@string(1,10))"
	case "null":
		mockValue = "Mock.mock(@null)"
	}
	return parseJsonSingleFileWithDefault(name, mockValue, in)
}

func parseMockSchema(name string, mock string) (*node.Node, error) {
	subs := mockFuncRegex.FindStringSubmatch(mock)
	mockSchema := ""
	if len(subs) == 2 {
		mockSchema = subs[1]
	}

	mockSchema = strings.TrimSuffix(strings.TrimSuffix(mockSchema, `"`), "'")
	mockSchema = strings.TrimPrefix(strings.TrimPrefix(mockSchema, `"`), "'")

	if strings.HasPrefix(mockSchema, "@") {
		return parseMockLambda(name, mockSchema)
	}

	if mockSchema == "" {
		mockSchema = "Mock.mock({\"mock|1\": \"sobe\"})"
	}

	var schemaData map[string]interface{}
	err := json.Unmarshal([]byte(mockSchema), &schemaData)
	if err != nil {
		return nil, ErrInvalidJsonSchema
	}
	return parseMockSchemaData(name, schemaData)
}

func parseMockSchemaData(name string, data interface{}) (*node.Node, error) {
	items := make([]*node.Node, 0)

	tt := node.TypeOfNodeSingle
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map, reflect.Struct:
		tt = node.TypeOfNodeMap
		for key, value := range data.(map[string]interface{}) {
			item, err := parseMockSchemaData(key, value)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}

	case reflect.Slice, reflect.Array:
		tt = node.TypeOfNodeArray
		for _, value := range data.([]interface{}) {
			item, err := parseMockSchemaData("", value)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
	default:
		return parseMockLambda(name, data)
	}

	name, rule := parseFieldName(name)
	return &node.Node{
		Name:  name,
		Rule:  rule,
		Value: data,
		Type:  tt,
		Items: items,
	}, nil
}

func parseMockLambda(name string, value interface{}) (*node.Node, error) {
	name, rule := parseFieldName(name)
	return &node.Node{
		Name:  name,
		Type:  node.TypeOfNodeSingle,
		Rule:  rule,
		Value: value,
	}, nil
}

func parseFieldName(in string) (name, rule string) {
	splits := strings.Split(in, "|")
	switch len(splits) {
	case 2:
		name = splits[0]
		rule = splits[1]
	case 1:
		name = splits[0]
	}
	return
}
