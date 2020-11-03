package match

import "regexp"

type Match struct {
	Rule   string `json:"rule"`
	On     bool   `json:"on"`
	Target string `json:"target"`
}

type Matcher interface {
	Match(name string) string
}

func NewMatcher(m *Match) (Matcher, error) {
	return nil, nil
}

type match struct {
	Rule      string         // 原始规则
	Reg       *regexp.Regexp // 正则模式
	fullMatch bool           // 是否完全匹配
	target    string         // 目标
}

func (s *match) Match(name string) string {
	panic("implement me")
}
