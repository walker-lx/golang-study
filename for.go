package main

import "fmt"

func main() {

	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for k := range(3) {
		fmt.Println(k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)
	}
}
