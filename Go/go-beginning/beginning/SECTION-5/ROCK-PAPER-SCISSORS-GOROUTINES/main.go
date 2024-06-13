package main

import "myapp/game"

func main() {

	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}
	go game.Rounds()
	game.CleanScreens()
	game.PrintIntro()

	for {
		game.RoundChan <- 1
		<-game.RoundChan // wait until RoundChan process
		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() { // if this round invalid, again
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}
	game.PrintSummary()
}
