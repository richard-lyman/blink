package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`
You must provide at least two arguments.
The first is the command to run when a change is seen.
The second, *and all remaining arguments*, are the items that will be watched - whose changes trigger the given command.

    E.g.: ./bin/blink some_script.sh /some/folder some_file

`)
		os.Exit(1)
	}
	script := os.Args[1]
	fmt.Printf("Will run '%s' when anything changes\n", script)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer watcher.Close()

	for _, v := range os.Args[2:] {
		err = watcher.Add(v)
		if err != nil {
			fmt.Println("Failed to watch given argument:", v, err)
			os.Exit(1)
		}
		fmt.Println("Watching:", v)
	}

	for {
		select {
		case <-watcher.Events:
			fmt.Print("*")
			c := exec.Command(script)
			err := c.Run()
			if err != nil {
				fmt.Printf("\n\nERROR:\n%s\n\n", err)
				os.Exit(1)
			}
		case err := <-watcher.Errors:
			fmt.Println("\n\nERROR:", err)
			os.Exit(1)
		}
	}
}
