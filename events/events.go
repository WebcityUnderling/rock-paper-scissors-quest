package events

import (
	"fmt"
	"os"
)

func IntroEvent() {
	b, err := os.ReadFile("intro.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	fmt.Println("Welcome to Rock Paper Scissors Quest. It's simple, defeat all the monsters using either Rock, Paper or Scissors.")
	fmt.Println("Type your move when prompted. Let's begin!\n\n")
}

func TieEvent() {
	b, err := os.ReadFile("tie.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func TriumphEvent() {
	b, err := os.ReadFile("triumph.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func ExitEvent() {
	b, err := os.ReadFile("end.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	fmt.Println("Congratulations! You conquered the dungeon! Now you can brag to all your friends you are a Rock Paper Scissor Wizard")
}

func DeathEvent(level int) {
	b, err := os.ReadFile("death.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	fmt.Println("You Died. You got to level", level)
}
