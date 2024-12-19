# NotForOil - A pipeline runner in go

I hate java. Every decent CI/CD system seems to require java, or something proprietary.

That needs to change.

## License

- Selected AGPLv3 license for this project. Only files included in the repository
  or mentioned as part of the project must abide by the terms of the license

## Contributors Agreement

- Anyone can contribute on a technical level. It's up to maintainers to approve
  any Merge Requests that become part of the project
- Any contribution made as part of the project becomes owned by the project.
- Any contribution made as part of the project has copyright assigned to the
  project
- The current form of governance is BDFL, owned by Matthew Stobbs,
  with project ownership belonging to Sprouting Communications in a commercial
  form

## Goals

- Be feature complete, able to run jobs on any system that can run go
- Dynamic agent installation, via ssh connection to install the binary,
  configure the agent, and make it available to the main server
- REST API that is used as the first class citizen, meaning all actions
  must be able to be performed using the REST API. Nothing is hidden

## Features

- [ ] HTTP server for easy reverse-proxy creation
  - [ ] Configurable with HTTPS using given certificate and private key
- [ ] Versioned REST API with access to every possible code path
- [ ] Job creation should be easy to manage, allowing different configuration
  file types. Using [viper](https://github.com/spf13/viper), can support
  yaml, json, toml, etc.
  - [ ] Jobs must be validated before they are accepted

## Design

Initially will be done as a single binary, running the REST API, managing jobs,
assigning jobs to agents/runners, managing logs of running jobs.

Will eventually be able to be split into the API server, managing user requests
and the jobs server, handling the job scheduling and such. This would mostly
be for stability reasons, and federation, with a single API server communicating
with multip Jobs servers.

Overcomplicating from the beginning is asking for trouble. Right now, jobs are
submitted, run, logs collected and reported via the same binary.

### API server

- HTTP Rest API
- Validates jobs

### Jobs server

- HTTP Rest API
- Manages jobs and job queue
- Schedules jobs
- Access secrets from external providers:
  - Should support Vault early on
  - Initially available using symmetric encryption
    done on the jobs server

### Messege server

- Reports job status via numerous notification methods
  - Obviously should support SMTP notifications
  - Other notification services to consider:
    - Prometheus
    - Grafana
    - ntfy.sh

### Job agent

- At first, jobs will be run from the central server on the same
  machine, or on remote machines using SSH.
- Jobs will, at first, be simple CLI instructions, similar to
  shell scripts, but in a structured format

#### Definition of a job

- A job is a collection of tasks
- Tasks can either depend on the output of a previous task (pipeline) or;
- Tasks can be an ordered set of instructions

#### What makes up a task

- Tasks are steps in a job. Their definition should be simple
  enough that any shell script can be interpreted into a
  task
- The structure of a task makes getting output from each step is possible
- That output can be recorded (logs), used in the following task (pipelined)
  or stored as an artifact
