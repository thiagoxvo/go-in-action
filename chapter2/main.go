package main

import (
	"log"
	"os"

	_ "github.com/thiagoxvo/goinaction/chapter2/matchers"
	"github.com/thiagoxvo/goinaction/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("EUA")
}
