package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Input numbers. To exit type 0.")

	args := os.Args
	var min, max int

	for i, n := range args[1:] {
		curr, err := strconv.Atoi(n)
		if i == 0 {
			min, max = curr, curr
		}
		if err != nil {
			log.Fatal(err)
		}
		if curr < min {
			min = curr
		}
		if curr > max {
			max = curr
		}
	}

	fmt.Printf("min: %v, max: %v", min, max)

}
