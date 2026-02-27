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
		dir int
		x   int
		y   int
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
					dir: 1,
					x:   x,
					y:   y,
				})
			case '<':
				num_ac++
				ac = append(ac, aircraft{
					dir: -1,
					x:   x,
					y:   y,
				})
			}
		}
	}

	sam := aircraft{}
	y := n
	for x, space := range ground {
		if space == '^' {
			sam = aircraft{
				dir: 1,
				x:   x,
				y:   y,
			}
		}
	}

	fmt.Println(sam)
	var rng int
	num_shots := num_ac
	for i := 0; i < num_shots; {
		for j := 0; j < num_ac; j++ {
			rng = sam.y - ac[j].y
			if ac[j].x == sam.x-(rng*ac[j].dir) {
				fmt.Println("SHOOT")
				i++
			} else {
				fmt.Println("WAIT")

				ac[j].x = ac[j].x + (1 * ac[j].dir)

			}
		}
	}

	fmt.Println(rng)
	return []string{}
}
