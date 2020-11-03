package match

// integrate
type Integration interface {
	Match(name string) string
}

type integrate struct {
	matchers []Matcher
}

func (i *integrate) Match(name string) string {
	panic("implement me")
}

func NewMatch(matchers []*Match) (Integration, error) {
	return nil, nil
}
