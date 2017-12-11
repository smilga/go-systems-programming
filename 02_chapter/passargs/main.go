package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for i, a := range args {
		switch a {
		case "-i":
			fmt.Printf("Argument -i passed: %v", args[i+1])
			fallthrough
		case "-k":
			fmt.Printf("Argument -k passed: %v", args[i+1])
		}
	}

}
