package main

import (
	"fmt"
	"time"
)

func main() {
	var index uint64
	fmt.Print("Введите индекс: ")
	fmt.Scanln(&index)
	save := make(map[uint64]uint64)
	start := time.Now()
	var numFibMemo uint64 = fibonacciMemo(index, save)
	duration := time.Since(start)
	fmt.Println("Число фибоначи: ", numFibMemo)
	fmt.Printf("Время выполнения задачи с оптимизацией:  %v\n", duration.Nanoseconds())
	start = time.Now()
	fibon(index)
	duration = time.Since(start)
	fmt.Printf("Время выполнения задачи без оптимизации:  %v\n", duration.Nanoseconds())
}

func fibonacciMemo(x uint64, save map[uint64]uint64) uint64 {
	if x < 2 {
		return x
	}
	if _, ok := save[x-1]; !ok {
		save[x-1] = fibonacciMemo(x-1, save)
	}
	if _, ok := save[x-1]; !ok {
		save[x-2] = fibonacciMemo(x-2, save)
	}
	return save[x-1] + save[x-2]
}

func fibon(x uint64) uint64 {
	if x < 2 {
		return x
	} else {
		return fibon(x-1) + fibon(x-2)
	}
}
