package main

import (
	"gobot/commands"
	"gobot/config"
)

func registerActions(i *Invoker, conf *config.Config) {

	i.Register(commands.RestCommand{Url: conf.BotUrl()})
	i.Register(commands.MenuCommand{Url: conf.BotUrl()})
	i.Register(commands.SaveCommand{Url: conf.BotUrl()})
	//i.Register(callbacks.EndRestCallback{Url: conf.BotUrl()})

	i.RegisterUndefined(commands.UndefinedCommand{Url: conf.BotUrl()})
}
