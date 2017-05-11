package main

import (
	_ "learn-go/goinaction/code/chapter2/sample/matchers"
	"learn-go/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
