package saver

type Fact map[string]interface{}

func NewFact() *Fact {
	return &Fact{}
}
