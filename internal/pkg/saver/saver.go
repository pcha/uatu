package saver

type Saver interface {
	Save(fact *Fact, bucket string) error
}
