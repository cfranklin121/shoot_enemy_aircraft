package main

import "fmt"

func main() {
	var sum int
	x := 3
	y := 7

	ptr := &x

	fmt.Println("ptr:", *ptr)

	sum = x + y
	fmt.Printf("Sum of %v and %v is %v\n", x, y, sum)
}

func shootEnemyAircraft(n int, input []string) []string {
	type aircraft struct {
		direction string
		x         int
		y         int
	}

	var sky []string
	for i := 0; i < n-1; i++ {
		sky = append(sky, input[i])
	}
	ground := input[n-1]

	fmt.Println("n:", n)
	fmt.Println("sky:", sky)
	fmt.Println("ground:", ground)

	num_ac := 0

	ac := []aircraft{}
	for y, line := range sky {
		for x, space := range line {
			switch space {
			case '>':
				num_ac++
				ac = append(ac, aircraft{
					direction: "r",
					x:         x,
					y:         y,
				})
			case '<':
				num_ac++
				ac = append(ac, aircraft{
					direction: "l",
					x:         x,
					y:         y,
				})
			}
		}
	}

	return []string{}
}
