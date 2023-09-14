package main

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/state"
)

var commands = []api.CreateCommandData{
	{
		Name:        "ping",
		Description: "pong",
	},
	/*
		{
			Name:        "echo",
			Description: "echo back the argument",
			Options: []discord.CommandOption{
				&discord.StringOption{
					OptionName:  "argument",
				},
			},
		},
	*/
	{
		Name:        "thonk",
		Description: "biiiig thonk",
	},
	{
		Name:        "fun",
		Description: "show a random event happening soon",
	},
}

func overwriteCommands(s *state.State) error {
	return cmdroute.OverwriteCommands(s, commands)
}
