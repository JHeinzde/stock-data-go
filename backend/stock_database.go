package main

import (
	"bufio"
	"log"
	"os"
)

func LoadFile() []string {
	file, err := os.Open("./stockdb")
	defer file.Close()

	if err != nil {
		log.Fatalln("Could not open database for reading!")
	}

	resultList := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resultList = append(resultList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Could not read database!")
	}

	return resultList
}

func AddLineToFile(stockSymbol string) {
	file, err := os.OpenFile("stockdb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	if err != nil {
		log.Fatalln("Could not open database for writing")
	}

	_, err = file.WriteString(stockSymbol + "\n")
	if err != nil {
		log.Fatalln("Could not write to file.", err)
	}
	log.Println("Wrote Symbol to stock database")
	defer file.Close()
}
