package main

import "fmt"

func main() {

	for i := 2; i < 100; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
	}

}
