package main

import (
	"time"

	"paste-corral/pastebin"
)

func main() {

	// Run web crawler in concurrent goroutine:
	go pastebin.Crawl()

	time.Sleep(2 * time.Minute)
}
