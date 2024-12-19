package jobs

import (
	"io"
)

// The jobs package defines the expected fields for a job
// as well as interfaces needed for jobs to run

type Job struct {
	tasks []Task
}

type Runner interface{}
type Logger interface{}

type Task struct {
	name    string
	chdir   string
	command string
	argv    []string
	out     io.Writer
	log     io.Reader
	vars    map[string]interface{}
}

type TaskOpts map[string]string

// NewTask create a new task with the given options
func NewTask(name string, opts TaskOpts, command string, args ...string) *Task {
	t := &Task{
		name:    name,
		command: command,
		vars:    make(map[string]interface{}),
	}
	copy(t.argv, args)
	for k, v := range opts {
		t.vars[k] = v
	}
	return t
}

// SetOutput sets an io.Writer to write output to
func (t *Task) SetOutput(w io.Writer) {
	t.out = w
}
