# `to-do`

Manage a to-do list using a simple command line tool

## Installation

```
go install
```

Verify your installation by running `to-do`
```
% to-do
A tool that enables users to manage a simple to-do list in a command line workflow

Usage:
  to-do [command]

Available Commands:
  add         Adds a task to the to-do list
  complete    Removes a task from the to-do list
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  show        Prints the current to-do list

Flags:
  -h, --help   help for to-do

Use "to-do [command] --help" for more information about a command.
```

## Commands

dd a task to your list
```
to-do add
```

Complete a task
```
to-do complete
```

Show your to-do list
```
to-do show
```

`to-do` stores your current list in your home directory in a file called `to-do.json`

## Example workflow

```
% to-do add 
What do you need to do?: Schedule a call with teammate
% to-do add
What do you need to do?: Fill out important survey by Friday
% to-do show

To-do: 
  ● Schedule a call with teammate
  ● Fill out important survey by Friday

% to-do complete
✔ Schedule a call with teammate
kylemartin@MB-FMQYX21HCJ to-do % to-do show    

To-do: 
  ● Fill out important survey by Friday

% to-do complete
Use the arrow keys to navigate: ↓ ↑ → ← 
? What have you completed?: 
  ▸ Fill out important survey by Friday
```
