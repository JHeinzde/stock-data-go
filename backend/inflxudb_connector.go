package main

import (
	"fmt"
	av "github.com/cmckee-dev/go-alpha-vantage"
)

func insertTimePoint(input <-chan av.TimeSeriesValue) {
	for tp := range input {
		fmt.Println(tp)
	}
}
