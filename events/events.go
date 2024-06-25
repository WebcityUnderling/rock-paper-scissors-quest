package events

import (
	_ "embed"
	"fmt"
	"phptogo/beastiary"
	"phptogo/rooms"
	"phptogo/utils"
	"time"
)

//go:embed ascii/intro.txt
var introTxt string

func PrintIntroEvent() {
	utils.PrintAsciiByLine(introTxt)
	time.Sleep(2 * time.Second)
	fmt.Println("Welcome to Rock Paper Scissors Quest. It's simple, defeat all the monsters using either Rock, Paper or Scissors.")
	time.Sleep(2 * time.Second)
	fmt.Println("Type your move when prompted. Let's begin!")
	time.Sleep(2 * time.Second)
}

func PrintEnterRoomEvent(level int, opponent *beastiary.Beast, room *rooms.Room) {
	fmt.Println("\n\nLevel", level, room.Name)
	time.Sleep(2 * time.Second)
	fmt.Println(room.Enter)
	time.Sleep(2 * time.Second)
	fmt.Println(opponent.EntryMessage)
	time.Sleep(2 * time.Second)
}

//go:embed ascii/tie.txt
var TieTxt string

func PrintTieEvent() {
	utils.PrintAsciiByLine(TieTxt)
	time.Sleep(2 * time.Second)
	fmt.Println("Go Again!")
}

//go:embed ascii/triumph.txt
var TriumphTxt string

func PrintTriumphEvent(opponent *beastiary.Beast, room *rooms.Room) {
	utils.PrintAsciiByLine(TriumphTxt)
	fmt.Println("You defeated", opponent.Name)
	time.Sleep(2 * time.Second)
	fmt.Println(opponent.DefeatMessage)
	time.Sleep(2 * time.Second)
	fmt.Println(room.Leave)
	time.Sleep(2 * time.Second)
}

//go:embed ascii/death.txt
var DeathTxt string

func PrintDeathEvent(level int, opponent *beastiary.Beast, room *rooms.Room) {
	utils.PrintAsciiByLine(DeathTxt)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You Died. You got to level", level)
	time.Sleep(2 * time.Second)
	fmt.Println(opponent.WinMessage)
	time.Sleep(2 * time.Second)
	fmt.Println(room.Defeat)
}

//go:embed ascii/end.txt
var EndTxt string

func PrintExitEvent() {
	utils.PrintAsciiByLine(EndTxt)
	fmt.Println("Congratulations! You conquered the dungeon! Now you can brag to all your friends you are a Rock Paper Scissor Wizard")
}
