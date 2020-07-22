package main

import (
	"time"

	"paste-corral/pastebin"
)

func main() {

	time.Sleep(2 * time.Hour)

	// Run web crawler in concurrent goroutine:
	go pastebin.Crawl()

	time.Sleep(2 * time.Hour)
}
