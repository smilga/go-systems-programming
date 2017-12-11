package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s [3]string
	s[0] = "1 b 3 a a d b 3 2"
	s[1] = "11 8 a h n n d w 2"
	s[2] = "3 8 l d s a 2 d m"

	counts := make(map[string]int)

	args := os.Args
	min, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("No treshold! Pass number!")
	}

	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		for _, word := range data {
			_, ok := counts[word]
			if ok {
				counts[word] = counts[word] + 1
			} else {
				counts[word] = 1
			}
		}
	}

	for k, v := range counts {
		if v >= min {
			fmt.Printf("%s -> %d \n", k, v)
		}
	}

}
