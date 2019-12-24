package database

import "time"

type TimePoint struct {
	Date   time.Time
	Open   float32
	High   float32
	Low    float32
	Close  float32
	Volume int64
}

type InfluxTimePoint struct {
	Measurement string
	Time        time.Time
	Tags        map[string]string
	Stocks      IntervalResult
}

type IntervalResult struct {
	Open   float32
	High   float32
	Low    float32
	Close  float32
	Volume int64
}
