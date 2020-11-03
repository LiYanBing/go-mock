package constant

import "math/rand"

type City struct {
	Name     string
	Children []*Distinct
}

func (s *City) PickDistinct() (*Distinct, bool) {
	if len(s.Children) <= 0 {
		return nil, false
	}
	return s.Children[rand.Intn(len(s.Children))], true
}

type Distinct struct {
	Name string
}

type Province struct {
	Name     string
	Children []*City
}

func (s *Province) PickCity() (*City, bool) {
	if len(s.Children) <= 0 {
		return nil, false
	}
	return s.Children[rand.Intn(len(s.Children))], true
}

func RandProvince() *Province {
	return Provinces[rand.Intn(len(Provinces))]
}
