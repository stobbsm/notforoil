package notforoil

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Command is the command type
type Command struct {
	cmd     *exec.Cmd // Raw command object
	ID      uuid.UUID
	Cmd     string
	Args    []string
	Created time.Time
	Start   time.Time
	End     time.Time
	Out     []string
	Err     []string
}

var commands = make(map[uuid.UUID]Command)

// NewCommand sets up a new command object
func NewCommand(cmd string, args ...string) *Command {
	c := &Command{
		ID:   uuid.NewV4(),
		Cmd:  cmd,
		Args: args,
		Out:  []string{},
		Err:  []string{},
	}
	c.cmd = exec.Command(cmd, args...)

	var o, e = make(chan string), make(chan string)

	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}

	go scanOutput(stdout, o)
	go linkChanToSlice(o, c.Out)

	stderr, err := c.cmd.StderrPipe()
	if err != nil {
		log.Println(err)
	}
	go scanOutput(stderr, e)
	go linkChanToSlice(e, c.Err)

	return c
}

func scanOutput(rc io.ReadCloser, ch chan<- string) {
	log.Println("Starting channel reader")
	scanner := bufio.NewScanner(rc)
	for scanner.Scan() {
		log.Print(scanner.Text())
		ch <- scanner.Text()
	}
}

func linkChanToSlice(ch <-chan string, sl []string) {
	log.Println("Linking slice to reader")
	select {
	case <-ch:
		sl = append(sl, <-ch)
	}
}

// Do a given command
func (c *Command) Do() <-chan error {
	var errchan = make(chan error)

	c.Start = time.Now()

	errchan <- c.cmd.Start()
	errchan <- c.cmd.Wait()
	c.End = time.Now()
	return errchan
}
