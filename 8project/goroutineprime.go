package main

import (
	"fmt"
	"math"
	"sync"
)

func checkRange(start, end, number int, result chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if number%i == 0 {
			result <- false
			return
		}
	}
	result <- true
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}

	limit := int(math.Sqrt(float64(number)))

	result := make(chan bool)

	var wg sync.WaitGroup

	numGoroutines := 4
	rangeSize := (limit + 1) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 2
		end := (i + 1) * rangeSize
		if i == numGoroutines-1 {
			end = limit
		}
		wg.Add(1)
		go checkRange(start, end, number, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		if !res {
			return false 
		}
	}
	return true 
}

func main() {
	var number int
	fmt.Print("Enter a number to check if it's prime: ")
	fmt.Scanln(&number)

	if isPrime(number) {
		fmt.Printf("%d is a prime number.\n", number)
	} else {
		fmt.Printf("%d is not a prime number.\n", number)
	}
}
