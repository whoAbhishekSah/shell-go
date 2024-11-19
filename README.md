This repo is a shell program written in Golang. It builds a POSIX compliant shell that's capable of
interpreting shell commands, running external programs and builtin commands like
cd, pwd, echo and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

## Run the shell

```
go run cmd/myshell/main.go
```

#### Example commands to try

```
$ cd /tmp/blueberry/pineapple/raspberry
$ pwd
$ cd ~
$ pwd
$ cd /tmp/grape
$ pwd
$ cd ./banana/blueberry
$ pwd
$ cd ../../../
$ pwd
$ cd /tmp/orange/orange/pineapple
$ pwd
$ cd /non-existing-directory
$ type pwd
$ pwd
$ my_exe David
$ type cat
$ type cp
$ type mkdir
$ type my_exe
$ type nonexistent
$ type echo
$ type exit
$ type type
$ type nonexistent
$ type nonexistentcommand
$ echo blueberry pineapple
$ echo pineapple pear apple
$ echo orange pineapple
$ invalid_command_1
$ exit 0
$ invalid_command_1
$ invalid_command_2
$ invalid_command_3
$ invalid_command_4
$ nonexistent
```
