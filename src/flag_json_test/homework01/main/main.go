package main

import (
	"homewoek01/util"
)

func main() {
	monster := util.Monster{
		Name:  "Goblin",
		Age:   200,
		Skill: "Stealth",
	}
	monster.Store()
	monster.Restore()

}
