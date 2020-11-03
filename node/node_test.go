package node

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestNode_Mock(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	root := &Node{
		Name: "",
		Type: TypeOfNodeMap,
		//Rule:  "1-5",
		Value: nil,
		Items: []*Node{
			{
				Name:  "name",
				Type:  TypeOfNodeSingle,
				Rule:  "4",
				Value: "golang",
			},
			{
				Name: "city",
				Type: TypeOfNodeArray,
				Rule: "1",
				Items: []*Node{
					{
						Name: "city",
						Type: TypeOfNodeMap,
						Rule: "1",
						Items: []*Node{
							{
								Name:  "shanghai",
								Type:  TypeOfNodeSingle,
								Rule:  "1",
								Value: "shanghai",
							},
							{
								Name:  "beiging",
								Type:  TypeOfNodeSingle,
								Rule:  "1",
								Value: "beiging",
							},
							{
								Name:  "guangzhou",
								Type:  TypeOfNodeSingle,
								Rule:  "1",
								Value: "guangzhou",
							},
						},
					},
				},
			},
			{
				Name: "province",
				Type: TypeOfNodeArray,
				Rule: "1",
				Items: []*Node{
					{
						Name:  "shanghai",
						Type:  TypeOfNodeSingle,
						Rule:  "1",
						Value: "shanghai",
					},
					{
						Name:  "beiging",
						Type:  TypeOfNodeSingle,
						Rule:  "1",
						Value: "beiging",
					},
					{
						Name:  "guangzhou",
						Type:  TypeOfNodeSingle,
						Rule:  "1",
						Value: "guangzhou",
					},
				},
			},
		},
	}

	t.Log(root.MockJson(context.Background()))
}
