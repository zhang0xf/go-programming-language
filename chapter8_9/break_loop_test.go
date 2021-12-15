package chapter8_9

import (
	"fmt"
	"testing"
)

func TestBreakLoop(t *testing.T) {

loop:
	for {
		for {
			fmt.Printf("in loop\n")
			break loop
		}
	}
	fmt.Printf("after break loop\n")

	fmt.Printf("------------")
}
