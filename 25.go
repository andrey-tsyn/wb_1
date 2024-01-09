package main

import "time"

func main() {
	println("Sleeping...")
	sleep(5000)
	println("Done!")
}

func sleep(ms int64) {
	timer := time.NewTimer(time.Duration(ms) * time.Millisecond)
	<-timer.C
}
