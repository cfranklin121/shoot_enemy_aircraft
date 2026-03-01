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
		{ //case 1
			input_n: 6,
			input: []string{
				"....................",
				".>..................",
				"...................<",
				"....................",
				"....................",
				"_________^__________",
			},
			expected: []string{
				"WAIT",
				"WAIT",
				"WAIT",
				"SHOOT",
				"WAIT",
				"WAIT",
				"SHOOT",
			},
		},
		{ //case 2
			input_n: 6,
			input: []string{
				">...................",
				">...................",
				">...................",
				"....................",
				"....................",
				"_________^__________",
			},
			expected: []string{
				"WAIT",
				"WAIT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
			},
		},
		{ //case 3
			input_n: 2,
			input: []string{
				"...<<<<<<<<<<<<<<<<<<<<<<<..<<<",
				"^______________________________",
			},
			expected: []string{
				"WAIT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
			},
		},
		{ //case 4
			input_n: 10,
			input: []string{
				"...........................<...........",
				".......................................",
				".......................................",
				".......................................",
				"............<..<..<..<..<..<..<..<..<..",
				".......................................",
				".......................................",
				".......................................",
				"..>......<..<..<..<..<..<..<..<..<..<..",
				"______^________________________________",
			},
			expected: []string{
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"SHOOT",
				"SHOOT",
				"WAIT",
				"SHOOT",
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
		fmt.Println("=====================")
	}
}
