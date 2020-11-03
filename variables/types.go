package variables

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/liyanbing/go-mock/rules"
)

var (
	funcRegex = regexp.MustCompile(`^(@[a-zA-Z0-9]+)(\(.*\))?`)
)

// ---------------factory--------------
type VariableInput struct {
	Name   string
	Rule   rules.Rule
	Params []interface{}
	Value  interface{}
}

type Operator func(context.Context, *VariableInput) (interface{}, error)

type Creator func(ctx context.Context, rule, name string, params []interface{}, value interface{}) (interface{}, error)

func DefaultRuleParserVariable(op Operator) Creator {
	return NewVariableCreator(rules.GetDefaultRuleParser(), op)
}

func NewVariableCreator(parser rules.RuleParser, op Operator) Creator {
	return func(ctx context.Context, rule, name string, params []interface{}, value interface{}) (interface{}, error) {
		ru, err := parser.Parse(rule)
		if err != nil {
			return nil, err
		}

		return op(ctx, &VariableInput{
			Name:   name,
			Rule:   ru,
			Params: params,
			Value:  value,
		})
	}
}

type factory struct {
	creators map[string]Creator
}

func (s *factory) Get(ctx context.Context, rule string, value interface{}) (interface{}, error) {
	if value == nil {
		return nil, nil
	}

	name, params, err := parseValue(value)
	if err != nil {
		return nil, err
	}

	creator, ok := s.creators[name]
	// 如果是没有预定义的名字则直接返回 空字符串
	if !ok {
		return "", fmt.Errorf("invalid %v ", name)
	}

	return creator(ctx, rule, name, params, value)
}

func (s *factory) Register(name string, creator Creator) error {
	if _, ok := s.creators[name]; ok {
		return fmt.Errorf("%v variable already exists", name)
	}

	s.creators[name] = creator
	return nil
}

func parseValue(in interface{}) (name string, parameters []interface{}, err error) {
	valueType := reflect.TypeOf(in)
	switch valueType.Kind() {
	case reflect.Bool:
		name = VariableTypeBool
	case reflect.String:
		return parseStringFuncSign(in.(string))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		name = VariableTypeNumber
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		name = VariableTypeNumber
	case reflect.Float32, reflect.Float64:
		name = VariableTypeNumber
	default:
		err = errors.New("invalid Value")
	}
	return
}

func parseStringFuncSign(in string) (name string, parameters []interface{}, err error) {
	if !strings.HasPrefix(in, "@") {
		name = VariableTypeString
		return
	}

	subs := funcRegex.FindStringSubmatch(in)
	if len(subs) >= 2 {
		name = subs[1]
	}

	if len(subs) >= 3 {
		parStr := subs[2]
		parStr = strings.TrimSuffix(strings.TrimPrefix(parStr, "("), ")")
		parStr = "[" + parStr + "]"
		err = json.Unmarshal([]byte(parStr), &parameters)
	}
	return
}
