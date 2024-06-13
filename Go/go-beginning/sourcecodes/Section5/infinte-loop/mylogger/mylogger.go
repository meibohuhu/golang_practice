package mylogger

import "log"

// ListenForLog listens to a channel and logs whatever
// that channel receives
func ListenForLog(ch chan string) {
	// this loop runs forever, and since we called this function
	// using the "go" keyword, it runs in the background, forever
	for {
		msg := <-ch
		log.Println(msg)
	}
}
