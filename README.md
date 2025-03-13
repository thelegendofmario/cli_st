# CLIst
## (CLI list)

## What is this?

## Install
PS: this assumes you have linux :)

to install, you can either download the binary from the releases, or you can clone the repository, and ( after you install go), you can run `go install` in the cloned folder, then add the go/ folder in your home directory to the path. 

## How Do I use It?

if you run it with no arguments, it won't do anything. (we know, that will be fixed!) the arguments you can choose are: 
| command | function | aliases |
|-----|-----|-------|
| `help` | display help | `--help`, `-h` |
| `create` | create `.plan.txt` file in home directory | `new` |
| `add ITEM` | adds `ITEM` to the plan file |  |
| `check-off ITEM (number)` | checks off an item in the plan file | `done`, `chkoff` |
| `delete ITEM (number)` | deletes `ITEM` (which is a number) from the todo list | `del` |
