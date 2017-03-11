package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`
You must provide exactly two arguments.
The first is the folder to watch.
The second is the command to run when a change is seen.

    E.g.: ./bin/blink /some/folder some_script.sh
    
`)
		os.Exit(1)
	}
	fmt.Printf("Will run '%s' when anything changes\n", os.Args[2])
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer watcher.Close()

	err = watcher.Add(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Watching:", os.Args[1])

	for {
		select {
		case <-watcher.Events:
			fmt.Print("*")
			c := exec.Command(os.Args[2])
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
