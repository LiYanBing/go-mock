package node

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/liyanbing/go-mock/rules"
	"github.com/liyanbing/go-mock/variables"
)

type TypeOfNode int

const (
	TypeOfNodeMap    TypeOfNode = iota // map 或者结构体
	TypeOfNodeArray                    // 数组类型
	TypeOfNodeSingle                   // 基础字段类型
)

type Node struct {
	Name  string      `json:"n"`
	Type  TypeOfNode  `json:"t"`
	Rule  string      `json:"r"`
	Value interface{} `json:"v"`
	Items []*Node     `json:"its"`
}

func (s *Node) MockJson(ctx context.Context) string {
	ret, _ := json.Marshal(s.Mock(ctx))
	return string(ret)
}

func (s *Node) Mock(ctx context.Context) interface{} {
	switch s.Type {
	case TypeOfNodeMap:
		return s.mockMap(ctx)
	case TypeOfNodeArray:
		return s.mockArray(ctx)
	default:
		return s.mockSingle(ctx)
	}
}

func (s *Node) selectItems(n int) []*Node {
	if n >= len(s.Items) {
		return s.Items
	}

	items := s.Items
	var newItems []*Node
	for i := 0; i < n; i++ {
		j := rand.Intn(len(items))
		newItems = append(newItems, items[j])
		items = append(items[0:j], items[j+1:]...)
	}
	return newItems
}

func (s *Node) mockSingle(ctx context.Context) interface{} {
	ret, err := variables.Factory.Get(ctx, s.Rule, s.Value)
	if err != nil {
		return nil
	}
	return ret
}

func (s *Node) mockMap(ctx context.Context) interface{} {
	if len(s.Items) <= 0 {
		return nil
	}

	items := s.Items
	if s.Rule != "" {
		parser := rules.GetDefaultRuleParser()
		rule, err := parser.Parse(s.Rule)
		if err != nil {
			return nil
		}
		n := rule.Random()
		itemsNum := int64(len(s.Items))
		if n >= itemsNum {
			items = s.Items
		} else {
			items = s.selectItems(int(n))
		}
	}

	ret := make(map[string]interface{})
	for _, item := range items {
		ret[item.Name] = item.Mock(ctx)
	}
	return ret
}

func (s *Node) mockArray(ctx context.Context) interface{} {
	items := s.Items
	repeat := false
	repeatedNum := 0

	if s.Rule != "" {
		parser := rules.GetDefaultRuleParser()
		rule, err := parser.Parse(s.Rule)
		if err != nil {
			return nil
		}

		repeatedNum = int(rule.Random())
		if repeatedNum == 1 && rule.RuleType() == rules.RuleTypeSingle {
			items = s.selectItems(1)
		} else {
			repeat = true
		}
	}

	temp := make([]interface{}, 0, len(items))
	for _, item := range items {
		temp = append(temp, item.Mock(ctx))
	}

	if repeat {
		ret := make([]interface{}, 0, len(temp)*repeatedNum)
		for i := 0; i < repeatedNum; i++ {
			ret = append(ret, temp...)
		}
		return ret
	}

	if len(temp) == 1 {
		return temp[0]
	}
	return temp
}
