package main

import (
	"log"

	"github.com/stobbsm/notforoil"
)

func main() {
	c := notforoil.NewCommand(`echo`, `hello`, `world`)

	log.Print(c)
}
