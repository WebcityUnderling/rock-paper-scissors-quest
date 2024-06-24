package events

import (
	"fmt"
	"phptogo/beastiary"
	"phptogo/rooms"
	"phptogo/utils"
	"time"
)

func PrintIntroEvent() {
	utils.PrintAsciiByLine("ascii/intro.txt")
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

func PrintTieEvent() {
	utils.PrintAsciiByLine("ascii/tie.txt")
	time.Sleep(2 * time.Second)
	fmt.Println("Go Again!")
}

func PrintTriumphEvent(opponent *beastiary.Beast, room *rooms.Room) {
	utils.PrintAsciiByLine("ascii/triumph.txt")
	fmt.Println("You defeated", opponent.Name)
	time.Sleep(2 * time.Second)
	fmt.Println(opponent.DefeatMessage)
	time.Sleep(2 * time.Second)
	fmt.Println(room.Leave)
	time.Sleep(2 * time.Second)
}

func PrintDeathEvent(level int, opponent *beastiary.Beast, room *rooms.Room) {
	utils.PrintAsciiByLine("ascii/death.txt")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You Died. You got to level", level)
	time.Sleep(2 * time.Second)
	fmt.Println(opponent.WinMessage)
	time.Sleep(2 * time.Second)
	fmt.Println(room.Defeat)
}

func PrintExitEvent() {
	utils.PrintAsciiByLine("ascii/end.txt")
	fmt.Println("Congratulations! You conquered the dungeon! Now you can brag to all your friends you are a Rock Paper Scissor Wizard")
}
