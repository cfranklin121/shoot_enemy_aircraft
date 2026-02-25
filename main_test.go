package main

import (
	"fmt"
	"testing"
)

func TestShootEnemyAircraft(t *testing.T) {
	cases := []struct {
		input_n  int
		input    []string
		expected []string
	}{
		{
			input_n: 6,
			input: []string{
				"....................",
				".>..................",
				"...................<",
				"....................",
				"....................",
				"_________^__________",
			},
			expected: []string{"WAIT",
				"WAIT",
				"WAIT",
				"SHOOT",
				"WAIT",
				"WAIT",
				"SHOOT",
			},
		},
	}
	for _, c := range cases {
		fmt.Println("Run test")
		pass := true
		//fmt.Println("Input:")
		//for _, input := range c.input {
		//fmt.Println(input)
		//}
		result := shootEnemyAircraft(c.input_n, c.input)

		fmt.Println("Expected:")
		for _, expected := range c.expected {
			fmt.Println("   ", expected)
		}

		fmt.Println("Actual:")
		for _, actual := range result {
			fmt.Println("   ", actual)
		}

		if len(result) != len(c.expected) {
			pass = false
		} else {
			for i, exp := range c.expected {
				if result[i] != exp {
					pass = false
					break
				}
			}
		}
		if pass == true {
			fmt.Println("PASS")
		} else {
			t.Errorf("FAIL")
		}
	}
}
