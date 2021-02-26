package notforoil

import (
	"bufio"
	"context"
	"io"
	"log"
	"os/exec"
	"sync"
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
	Out     chan string
	Err     chan string
}

var commandpool = sync.Pool{
	New: func() interface{} {
		return &Command{}
	},
}

// NewCommand sets up a new command object
func NewCommand(wg *sync.WaitGroup, cmd string, args ...string) *Command {
	c := commandpool.Get().(*Command)
	c.Reset()
	c.Cmd = cmd
	c.Args = args

	c.cmd = exec.Command(cmd, args...)

	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}

	go scanOutput(wg, stdout, c.Out)

	stderr, err := c.cmd.StderrPipe()
	if err != nil {
		log.Println(err)
	}
	go scanOutput(wg, stderr, c.Err)

	return c
}

func scanOutput(wg *sync.WaitGroup, rc io.ReadCloser, ch chan string) {
	wg.Add(1)
	log.Println("Starting channel reader")
	scanner := bufio.NewScanner(rc)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	log.Println("Closing channel")
	close(ch)
	wg.Done()
}

// Do a given command
func (c *Command) Do(ctx context.Context) error {
	var errchan = make(chan error)

	c.Start = time.Now()
	defer c.markEnd()

	go func() {
		errchan <- c.cmd.Start()
		errchan <- c.cmd.Wait()
		close(errchan)
	}()
	log.Print("Waiting for command to end")
	for {
		select {
		case <-ctx.Done():
			log.Printf("Command ended, Exit code: %b", c.cmd.ProcessState.ExitCode())
			return ctx.Err()
		case err, ok := <-errchan:
			if err != nil {
				log.Print(err)
			}
			if !ok {
				log.Printf("Command ended, Exit code: %b", c.cmd.ProcessState.ExitCode())
				return err
			}
			continue
		}
	}
}

// Reset the command to defaults, including new channels and UUID
func (c *Command) Reset() {
	c.cmd = nil
	c.ID = uuid.NewV4()
	c.Out = make(chan string)
	c.Err = make(chan string)
	c.Args = []string{}
	c.Cmd = ""
	c.Created = time.Now()
	c.Start = time.Time{}
	c.End = time.Time{}
}

func (c *Command) markEnd() {
	c.End = time.Now()
}
