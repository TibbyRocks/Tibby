package commands

type Command struct {
	Name string
	Help string
}

var BotCommands map[string]*Command = make(map[string]*Command)
