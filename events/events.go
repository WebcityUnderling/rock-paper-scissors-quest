package events

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func IntroEvent() {
	asciiByLine("ascii/intro.txt")
	time.Sleep(2 * time.Second)
	fmt.Println("Welcome to Rock Paper Scissors Quest. It's simple, defeat all the monsters using either Rock, Paper or Scissors.")
	time.Sleep(2 * time.Second)
	fmt.Println("Type your move when prompted. Let's begin!")
	time.Sleep(2 * time.Second)
}

func TieEvent() {
	asciiByLine("ascii/tie.txt")
	time.Sleep(2 * time.Second)
}

func TriumphEvent() {
	asciiByLine("ascii/triumph.txt")
}

func ExitEvent() {
	asciiByLine("ascii/end.txt")
	fmt.Println("Congratulations! You conquered the dungeon! Now you can brag to all your friends you are a Rock Paper Scissor Wizard")
}

func DeathEvent(level int) {
	asciiByLine("ascii/death.txt")

}

func asciiByLine(path string) {
	f, e := os.Open(path)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		println(s.Text())
		time.Sleep(75 * time.Millisecond)
	}
}
