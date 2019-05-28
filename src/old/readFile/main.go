package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile("/Users/V729168/Desktop/h-scope-gpsfleet.log")
}

func readFile(file string) {
	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {

		fmt.Println(scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("File Contents: %s", content)
}
