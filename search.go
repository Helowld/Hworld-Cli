package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func search() {
	commando.
		Register("search").
		SetShortDescription("this command searches for language argument").
		AddArgument("language name", "name of the language", "").
		SetAction(func(action map[string]commando.ArgValue, flag map[string]commando.FlagValue) {
			fmt.Println("yet to implement")
		})
}
