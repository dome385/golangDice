package main

import (
	"fmt"
	"math/rand"
)

func roll(sides int) int{
	return rand.Intn(sides) + 1
}

func main() {
	//rand.Seed(time.Now().UnixNano())

	dice, sides := 2, 6
	rolls := 1

	for r:=1; r<=rolls; r++{
		sum:=0

		for d:=1; d<=dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Wurf #", r, ", Würfel #", d, ":", rolled)
		}
		fmt.Println("Ingesamte Augenzahl:", sum)

	}
}
