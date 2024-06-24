package events

import (
	"fmt"
	"phptogo/utils"
	"time"
)

func IntroEvent() {
	utils.PrintAsciiByLine("ascii/intro.txt")
	time.Sleep(2 * time.Second)
	fmt.Println("Welcome to Rock Paper Scissors Quest. It's simple, defeat all the monsters using either Rock, Paper or Scissors.")
	time.Sleep(2 * time.Second)
	fmt.Println("Type your move when prompted. Let's begin!")
	time.Sleep(2 * time.Second)
}

func TieEvent() {
	utils.PrintAsciiByLine("ascii/tie.txt")
	time.Sleep(2 * time.Second)
}

func TriumphEvent() {
	utils.PrintAsciiByLine("ascii/triumph.txt")
}

func ExitEvent() {
	utils.PrintAsciiByLine("ascii/end.txt")
	fmt.Println("Congratulations! You conquered the dungeon! Now you can brag to all your friends you are a Rock Paper Scissor Wizard")
}

func DeathEvent(level int) {

	utils.PrintAsciiByLine("ascii/death.txt")
}
