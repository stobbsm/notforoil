package jobs

// The jobs package defines the expected fields for a job
// as well as interfaces needed for jobs to run

type Job struct {
	name    string
	chdir   string
	command string
	argv    []string
}

type Runner interface{}
type Logger interface{}
