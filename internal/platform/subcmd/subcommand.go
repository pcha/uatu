package subcmd

type Runnable func([]string) error

type SubCommand struct {
	name             string
	flagsDefinitions []FlagDefinition
	Logic            Runnable
	Flags            map[string]interface{}
}

func NewSubCommand(name string, flags []FlagDefinition, action Runnable) *SubCommand {
	return &SubCommand{
		name:             name,
		flagsDefinitions: flags,
		Logic:            action,
	}
}

func (c SubCommand) Init() error {

}
