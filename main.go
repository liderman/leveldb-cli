// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

// ATTENTION! This version of the software is an experimental!
// There is no guarantee that the application will work correctly
// This code will be refactored, so do not rely on its structure

package main

import (
	"bitbucket.org/creachadair/shell"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/liderman/leveldb-cli/commands"
	"os"
	"path"
	"runtime"
	"strconv"
)

// Software version number
const VERSION = "0.2.0"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("show",
		readline.PcItem("prefix"),
		readline.PcItem("range"),
	),
	readline.PcItem("exit"),
	readline.PcItem("quit"),
	readline.PcItem("help"),
	readline.PcItem("get"),
	readline.PcItem("put"),
	readline.PcItem("set"),
	readline.PcItem("delete"),
	readline.PcItem("version"),
)

// Main function
func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:       "\033[31m»\033[0m ",
		HistoryFile:  "/tmp/leveldb-cli.tmp",
		AutoComplete: completer,
	})
	if err != nil {
		panic(err)
	}

	defer l.Close()

	fmt.Println("*********************************************************")
	fmt.Println("**                     LevelDB CLI                     **")
	fmt.Println("*********************************************************")
	fmt.Println("")
	fmt.Println("Run command 'help' for help.")

	for {
		line, err := l.Readline()
		if err != nil {
			break
		}

		args, _ := shell.Split(line)
		switch {
		// Command: version
		case line == "version":
			fmt.Printf("Version %s. %s %s %s\n", VERSION, runtime.Compiler, runtime.GOARCH, runtime.GOOS)
			break
		// Command: help
		case line == "help":
			fmt.Println("Enter one of the commands to get help: show, set, get, open, close, put, delete, version")
			break
		// Command: quit and exit
		case line == "quit":
		case line == "exit":
			goto exit
		// Command: show
		case args[0] == "show":
			if len(args) == 1 {
				fmt.Println("Bad format. Please use 'show prefix|range'")
				break
			}

			switch args[1] {
			// Sub-command: range
			case "range":
				if len(args) < 4 || len(args) > 5 {
					fmt.Println("Bad format. Please use 'show range START LIMIT [FORMAT]'")
					break
				}

				format := ""
				if len(args) == 5 {
					format = args[4]
				}

				fmt.Println(commands.ShowByRange(args[2], args[3], format))
				break
			// Sub-command: prefix
			case "prefix":
				if len(args) < 3 || len(args) > 4 {
					fmt.Println("Bad format. Please use 'show prefix PREFIX [FORMAT]'")
					break
				}

				format := ""
				if len(args) == 4 {
					format = args[3]
				}

				fmt.Println(commands.ShowByPrefix(args[2], format))
				break
			}

			break
		// Command: put and set
		case args[0] == "put":
		case args[0] == "set":
			if len(args) != 3 {
				fmt.Printf("Bad format. Please use '%s KEY VALUE'\n", args[0])
				break
			}

			fmt.Println(commands.Set(args[1], args[2]))
			break
		// Command: get
		case args[0] == "get":
			if len(args) < 2 || len(args) > 3 {
				fmt.Println("Bad format. Please use 'get KEY FORMAT'")
				break
			}

			format := ""
			if len(args) == 3 {
				format = args[2]
			}

			fmt.Println(commands.Get(args[1], format))
			break
		// Command: delete
		case args[0] == "delete":
			if len(args) != 2 {
				fmt.Print("Bad format. Please use 'delete KEY'")
				break
			}

			fmt.Println(commands.Delete(args[1]))
			break
		// Command: close
		case args[0] == "close":
			if len(args) != 1 {
				fmt.Print("Bad format. Please use 'close'")
				break
			}

			l.SetPrompt("\033[31m»\033[0m ")
			fmt.Println(commands.Close())
			break
		// Command: open
		case args[0] == "open":
			_, err := os.Stat(args[1])
			if err != nil {
				fmt.Println("Database not exist! Create new database.")
			}

			if len(args) != 2 {
				fmt.Println("Bad format. Please use 'open DATABASE_NAME'")
				break
			}

			l.SetPrompt(fmt.Sprintf("\033[31m%s»\033[0m ", path.Base(args[1])))
			fmt.Println(commands.Open(args[1]))
			break
		default:
			fmt.Println("Unknown command: ", strconv.Quote(line))
		}
	}

exit:
}
