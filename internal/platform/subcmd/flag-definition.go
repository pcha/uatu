package subcmd

type FlagType int

const (
	String FlagType = iota
	Int
)

type FlagDefinition struct {
	Type         FlagType
	Name         string
	DefaultValue interface{}
	Description  string
}

func NewFlagDefinition(flagType FlagType, name string, defaultVal interface{}, description string) *FlagDefinition {
	return &FlagDefinition{
		flagType,
		name,
		defaultVal,
		description,
	}
}

func (d *FlagDefinition) Get() {
	switch d.Type {

	}
}
