package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input in channels
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 0    // back to 0
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			// *** changed here to send info back when done
			g.DisplayChan <- ""
		}
	}
}

// ClearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	// *** changed here
	g.DisplayChan <- fmt.Sprintf(`
Rock, Paper & Scissors
----------------------
Game is played for three rounds, and best of three wins the game. Good luck!

`)
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1
	// *** changed here
	g.DisplayChan <- fmt.Sprintf(`

Round %d
-------
`, g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	// *** changed here
	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan
	switch computerValue {
	case ROCK:
		// *** changed here
		g.DisplayChan <- "Computer chose ROCK"
		<-g.DisplayChan
		break
	case PAPER:
		// *** changed here
		g.DisplayChan <- "Computer chose PAPER"
		<-g.DisplayChan
		break
	case SCISSORS:
		// *** changed here
		g.DisplayChan <- "Computer chose SCISSORS"
		<-g.DisplayChan
		break
	default:
	}

	// *** changed here
	g.DisplayChan <- ""
	<-g.DisplayChan

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		// *** changed here
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid choice!"
			// *** changed here
			<-g.DisplayChan
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	// *** changed here
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	// *** changed here
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	// *** changed here
	g.DisplayChan <- fmt.Sprintf(`
Final Score
-----------
Player: %d/3, Computer %d/3
`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChan

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins game!"
		// *** changed here
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Computer wins game!"
		// *** changed here
		<-g.DisplayChan
	}

	// *** changed here
	g.DisplayChan <- ""
	<-g.DisplayChan
}
