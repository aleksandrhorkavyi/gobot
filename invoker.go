package main

import (
	"gobot/commands"
	"gobot/models"
)

type Invoker struct {
	api *BotApi
	update models.Update
	commands []commands.Command
	undefinedCommand commands.Command
}

func (i *Invoker) Register(c commands.Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) RegisterUndefined(c commands.Command) {
	i.undefinedCommand = c
}

func (i *Invoker) ExecuteActions() {
	var handled bool
	for _, cmd := range i.commands {
		sent, err := cmd.Execute(i.update)
		if err != nil {
			panic(err)
		}

		if sent {
			handled = true
			break
		}
	}

	if !handled {
		_, err := i.undefinedCommand.Execute(i.update)
		if err != nil {
			panic(err)
		}
	}
	i.api.offset = i.update.UpdateId + 1
}
