# todo
A simple CLI task manager written in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/todo)](https://goreportcard.com/report/github.com/DataDrake/todo) [![license](https://img.shields.io/github/license/DataDrake/todo.svg)]()

## Motivation

There are a few well-known CLI task managers, but none of them quite fit my workflow, so I decided to write one.

## Goals

 * Easy to use
 * Stays out of the way
 * No external dependencies
 * A+ Rating on [Report Card](https://goreportcard.com/report/github.com/DataDrake/todo)
 
## Requirements

#### Compile-Time
* Go 1.15 (tested)
* Make

## Installation

1. Clone repo and enter its directory
2. `make`
3. `sudo make install`

## Usage

### Typical Workflow

* New tasks are created and added to the Backlog
* Tasks can be claimed and moved to the TODO list
* Tasks may be returned to the Backlog
* When a task is finished, it is marked as done and moves to the Complete list

### Projects and Labels

Each task can optionally be assigned to a Project and/or marked with a Label. Examples:
```
:Label
:"Longer Label"
@Project
@"Longer Project"
@"Why Not" :both?
```

```
Under Construction
```

## License
 
Copyright 2021 Bryan T. Meyers <root@datadrake.com>
 
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 
http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
