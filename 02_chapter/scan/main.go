package main

import "fmt"

func main() {
	fmt.Println("Input numbers. To exit type 0.")

	var min, max, curr int
	fmt.Scan(&curr)
	min, max = curr, curr
	for {
		fmt.Scan(&curr)
		if curr == 0 {
			break
		}
		if curr < min {
			min = curr
		}
		if curr > max {
			max = curr
		}
	}

	fmt.Printf("Min nr: %v, Max nr: %v", min, max)
}
