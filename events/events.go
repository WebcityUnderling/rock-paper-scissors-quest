package events

import (
	_ "embed"
	"fmt"
	"rpsq/beastiary"
	"rpsq/rooms"
	"rpsq/setup"
	"rpsq/utils"
	"sync"
	"time"
)

type Timeout int

var sleepDuration = 2000 * time.Millisecond

func SetEventTimeout() {
	if setup.Environment == "development" {
		sleepDuration = 0
	}
}

//go:embed ascii/intro.txt
var introTxt string

func PrintIntroEvent(wg *sync.WaitGroup) string {
	defer wg.Done()
	wg.Add(1)
	utils.PrintAsciiByLine(introTxt)
	time.Sleep(sleepDuration)
	fmt.Println("Welcome to Rock Paper Scissors Quest. It's simple, defeat all the monsters using either Rock, Paper or Scissors.")
	time.Sleep(sleepDuration)
	fmt.Println("Type your move when prompted. Let's begin!")
	time.Sleep(sleepDuration)
	return "event:intro"
}

func PrintEnterRoomEvent(level int, opponent *beastiary.Beast, room *rooms.Room) {
	fmt.Println("\n\nLevel", level, room.Name)
	time.Sleep(sleepDuration)
	fmt.Println(room.Enter)
	time.Sleep(sleepDuration)
	fmt.Println(opponent.EntryMessage)
	time.Sleep(sleepDuration)
}

//go:embed ascii/tie.txt
var TieTxt string

func PrintTieEvent() {
	utils.PrintAsciiByLine(TieTxt)
	time.Sleep(sleepDuration)
	fmt.Println("Go Again!")
}

//go:embed ascii/triumph.txt
var TriumphTxt string

func PrintTriumphEvent(opponent *beastiary.Beast, room *rooms.Room) {
	utils.PrintAsciiByLine(TriumphTxt)
	fmt.Println("You defeated", opponent.Name)
	time.Sleep(sleepDuration)
	fmt.Println(opponent.DefeatMessage)
	time.Sleep(sleepDuration)
	fmt.Println(room.Leave)
	time.Sleep(sleepDuration)
}

//go:embed ascii/death.txt
var DeathTxt string

func PrintDeathEvent(level int, opponent *beastiary.Beast, room *rooms.Room) {
	utils.PrintAsciiByLine(DeathTxt)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You Died. You got to level", level)
	time.Sleep(sleepDuration)
	fmt.Println(opponent.WinMessage)
	time.Sleep(sleepDuration)
	fmt.Println(room.Defeat)
}

//go:embed ascii/end.txt
var EndTxt string

func PrintExitEvent() {
	utils.PrintAsciiByLine(EndTxt)
	fmt.Println("Congratulations! You conquered the dungeon! Now you can brag to all your friends you are a Rock Paper Scissor Wizard")
}
