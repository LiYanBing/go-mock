package rules

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const (
	DefaultParser = "defaultParser"
)

// rule parser
type RuleParser interface {
	Parse(rule string) (Rule, error)
}

var (
	ruleParser = map[string]RuleParser{
		DefaultParser: &defaultRuleParser{},
	}
)

func RegisterRuleParser(name string, parser RuleParser) error {
	if _, ok := ruleParser[name]; ok {
		return fmt.Errorf("%v rule parser already exists", name)
	}
	ruleParser[name] = parser
	return nil
}

func GetDefaultRuleParser() RuleParser {
	parser, _ := GetRuleParser(DefaultParser)
	return parser
}

func GetRuleParser(name string) (RuleParser, error) {
	if parser, ok := ruleParser[name]; ok {
		return parser, nil
	}
	return nil, fmt.Errorf("%v rule parser not found", name)
}

type defaultRuleParser struct{}

func (s *defaultRuleParser) Parse(rule string) (Rule, error) {
	splits := strings.Split(rule, ".")

	start := rule
	end := ""
	if len(splits) > 0 {
		start = splits[0]
	}

	if len(splits) > 1 {
		end = splits[1]
	}

	ret := &defaultRule{
		rule: rule,
		t:    RuleTypeSingle,
	}

	var err error
	if len(start) > 0 {
		ret.min, ret.max, ret.count, err = s.parse(start)
		if err != nil {
			return nil, err
		}

		if ret.max > 0 {
			ret.t = RuleTypeTwo
		}
	}

	if len(end) > 0 {
		ret.dMin, ret.dMax, _, err = s.parse(end)
		if err != nil {
			return nil, err
		}

		if ret.dMax > 0 {
			ret.t = RuleTypeThree
		}
	}
	return ret, nil
}

func (s *defaultRuleParser) parse(in string) (min, max, count int64, err error) {
	splits := strings.Split(in, "-")
	if len(splits) > 0 {
		min, err = strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return
		}
		count = min
	}

	if len(splits) >= 2 {
		max, err = strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			return
		}
	}
	return
}

/*===================rule===================*/
type RuleType int

const (
	RuleTypeSingle RuleType = iota
	RuleTypeTwo
	RuleTypeThree
)

type Rule interface {
	Rule() string
	Random() int64
	DRandom() int64
	RuleType() RuleType
	Generate(string) string
}

type defaultRule struct {
	rule  string
	count int64
	min   int64
	max   int64
	dMin  int64
	dMax  int64
	t     RuleType
}

func (s *defaultRule) Rule() string {
	return s.rule
}

func (s *defaultRule) Random() int64 {
	if s.max <= s.min {
		return s.count
	}

	n := s.max - s.min
	return int64(rand.Intn(int(n))) + s.min
}

func (s *defaultRule) DRandom() int64 {
	if s.dMax <= s.dMin {
		return 0
	}

	n := s.dMax - s.dMin
	return int64(rand.Intn(int(n))) + s.dMin
}

func (s *defaultRule) Generate(in string) string {
	if s.rule == "" {
		return in
	}

	n := s.Random()
	ret := ""
	for i := 0; i < int(n); i++ {
		ret += in
	}
	return ret
}

func (s *defaultRule) RuleType() RuleType {
	return s.t
}
