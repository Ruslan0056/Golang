package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 20
	height = 20
)

type Universe [][]bool
type Neighbors [][]int

func NewUniverse() Universe {
	Universe := make([][]bool, width, width)
	for i := range Universe {
		Universe[i] = make([]bool, height, height)

	}
	return Universe
}

func (u Universe) Show() Universe {

	for i := range u {
		fmt.Print("|")
		for j := range u[i] {
			if u[j][i] {
				fmt.Print("∗|")
			} else {
				fmt.Print("-|")
			}
		}
		fmt.Println()

	}
	return u
}

func (u Universe) Seed() Universe {
	quantityOfTrue, a := 0, 0
	amount := height * width / 5
	for i := range u {
		for j := range u[i] {
			a = rand.Intn(100) % 4
			if a == 1 && quantityOfTrue < amount {
				u[j][i] = true
				quantityOfTrue++
			}
		}
	}
	return u
}

func (u Universe) Neighbors() Neighbors {
	Neighbors := make([][]int, width, width)
	for i := range Neighbors {
		Neighbors[i] = make([]int, height, height)
	}

	for i := range u {
		for j := range u[(i-1+width)%width] {
			if u[(j-1+height)%height][(i-1+width)%width] {
				Neighbors[j][i]++
				//	fmt.Print(Neighbors[i])
			}
			if u[(j+height)%height][(i-1+width)%width] {
				Neighbors[j][i]++
			}
			if u[(j+1+height)%height][(i-1+width)%width] {
				Neighbors[j][i]++
			}
		}
		for j := range u[i] {
			if u[(j-1+height)%height][(i+width)%width] {
				Neighbors[j][i]++
			}
			if u[(j+1+height)%height][(i+width)%width] {
				Neighbors[j][i]++
			}
		}
		for j := range u[(i+1+width)%width] {
			if u[(j-1+height)%height][(i+1+width)%width] {
				Neighbors[j][i]++
			}
			if u[(j+height)%height][(i+1+width)%width] {
				Neighbors[j][i]++
			}
			if u[(j+1+height)%height][(i+1+width)%width] {
				Neighbors[j][i]++
			}
		}

	}
	return Neighbors
}

func Next(u Universe, n Neighbors) Universe {
	nextUniverse := make([][]bool, width, width)
	for i := range nextUniverse {
		nextUniverse[i] = make([]bool, height, height)

	}
	for i := range n {
		for j := range n[i] {
			if u[j][i] {
				if n[j][i] == 2 || n[j][i] == 3 {
					nextUniverse[j][i] = true
				}
			} else if !u[j][i] && n[j][i] == 3 {
				nextUniverse[j][i] = true
			} else {
				nextUniverse[j][i] = false
			}
		}
	}
	return nextUniverse
}

func clear() {
	time.Sleep(time.Duration(300 * time.Millisecond))
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Game(a Universe) {
	var value int = 16
	for value > 0 {
		value = 0
		for i := range a {
			for j := range a[i] {
				if a[i][j] {
					value++
				}
			}
		}
		clear()
		b := Next(a, a.Neighbors())
		b.Show()
		a, b = b, a
	}
}

func main() {
	s := NewUniverse()
	a := s.Seed().Show()
	Game(a)
}
