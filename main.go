package main

import (
	"phptogo/events"
	"phptogo/levels"
)

func main() {
	events.IntroEvent()
	levels.SetDifficulty()
	levels.Levels()
}
