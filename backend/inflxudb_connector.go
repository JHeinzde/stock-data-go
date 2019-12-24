package main

import "fmt"

func insertTimePoint(input <-chan TimePoint) {
	for tp := range input {
		fmt.Println(tp)
	}
}
