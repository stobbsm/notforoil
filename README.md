# NotForOil - A pipeline runner in go

I hate java. Every decent CI/CD system seems to require java, or something proprietary.

That needs to change.

## Goals

Be feature complete, able to run jobs on any system that can run go

## Features
_Unchecked features are not considered complete. Dash items are in progress_

- [ ] Have server that can speak to agent over simple https. Require http2.
- [ ] GraphQL is the ONLY API in use, and can be extended via plugins
- [ ] Easy to define pipelines:
    - [ ] Templated jobs
    - [ ] Templated Pipelines
    - [ ] Stored in source control
- [ ] Manageable configuration:
    - [ ] Stored in source control
    - [ ] Automatic rollback on failure
    - [ ] Easy to edit and manage via gui or editing
    - [ ] Viewable diffs
    - [ ] Everything is configurable
    - [ ] Etcd?
- [ ] Source management:
    - [ ] Git is first class
    - [ ] Common interfaces to build other SCM modules
    - [ ] Service that polls for changes
- [ ] Secrets management:
    - [ ] Common interfaces to build secrets manager
    - [ ] Default should be good to use, with sane defaults
    - [ ] Separate service
- [ ] Containers first:
    - [ ] When on Linux, implement containers by default
- [ ] Build system integrations:
    - [ ] Default templates for different build systems
    - [ ] Enables containers by default
    - [ ] Use a standard container tool (docker, podman, lxc, systemd-nspawn) by default
- [ ] ENV Management:
    - [ ] Protected by default
    - [ ] Opt-in to r/w
    - [ ] Ability to modify from build steps, if r/w
- [ ] GraphQL API:
    - [ ] Query anything
    - [ ] Separate micro-service
- [ ] Authentication:
    - [ ] Basic management using http auth
    - [ ] Token based API access
- [ ] WebUI:
    - [ ] Frontend to build pipelines
- [ ] Pipeline:
    - [ ] Steps; Many per pipeline
    - [ ] Jobs; Many per step, concurrent possible when defined
- [ ] Plugins:
    - [ ] Micro-service architecture that can be extended
    - [ ] Common interfaces to facilitate communication
    - [ ] Must implement via graphql
