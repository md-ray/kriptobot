package main

import (
	"time"

	_ "github.com/saviourcat/kriptobot/messaging/service"
)

func main() {
	// logger := log.NewLogfmtLogger(os.Stderr)

	go forever()
	select {}
}

func forever() {
	for {
		//fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
