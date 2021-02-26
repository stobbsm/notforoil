package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/stobbsm/notforoil"
)

func main() {
	wg := &sync.WaitGroup{}
	c := notforoil.NewCommand(wg, `ls`, `/home/stobbsm/`, `-la`)

	ctx := context.Background()
	var output, errors []string

	go func() {
		for {
			select {
			case out, ok := <-c.Out:
				if out != "" {
					output = append(output, out)
				}
				if !ok {
					break
				}
			case err := <-c.Err:
				if err != "" {
					errors = append(errors, err)
					break
				}
			}
		}
	}()

	c.Do(ctx)

	wg.Wait()
	if len(errors) > 0 {
		log.Println("Errors:")
		printStrings("\t%s\n", errors)
	} else {
		log.Println("No errors")
	}

	log.Println("Output from command:")
	printStrings("\t%s\n", output)
}

func printStrings(format string, sl []string) {
	for _, l := range sl {
		fmt.Printf(format, l)
	}
}
