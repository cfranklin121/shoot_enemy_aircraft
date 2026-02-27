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

	ac := []aircraft{}
	num_ac := 0
	num_shots := 0
	shoot := false
	var result []string
	var sky []string
	var rng int

	for i := 0; i < n-1; i++ {
		sky = append(sky, input[i])
	}

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

	/*
		sam := aircraft{}
		for x, space := range input[n-1] {
			if space == '^' {
				sam = aircraft{
					dir: 0,
					x:   x,
					y:   n,
				}
			}
		}
	*/

	sam := aircraft{
		dir: 0,
		x:   9,
		y:   n,
	}

	num_turns := 0
	for num_shots = 0; num_shots < num_ac; {
		for j := 0; j < num_ac; j++ {
			rng = sam.y - ac[j].y
			if ac[j].x == sam.x-(rng*ac[j].dir) {
				shoot = true
			}
			ac[j].x = ac[j].x + (1 * ac[j].dir)
		}

		if shoot {
			fmt.Println("SHOOT")
			result = append(result, "SHOOT")
			num_shots++
			shoot = false
		} else {
			fmt.Println("WAIT")
			result = append(result, "WAIT")
		}
		num_turns++
		if num_turns >= 20 {
			break
		}
	}
	return result
}
