package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/Nupgod/CLI-Reader/internal/reader"
	"github.com/Nupgod/CLI-Reader/internal/logger"
)

func main() {
	filePath := flag.String("file", " ", "Read data from file: file или stdin")
	url := flag.String("url", " ", "URL")
	flag.Parse()

	// Logger initialization
	logger, err := logger.New()
	if err != nil {
		fmt.Printf("Logger initialization error: %v\n", err)
		return
	}
	defer logger.Close()

	// Read data from stdin by default. If flag -file is set, read data from JSON file if exists.
	// To successfully read from a file. The JSON file must have a "numbers" parameter that stores a list of integers. 
	// {
	//  "name": "John",
	//	"born": 1990,
	// 	"numbers": [1, 2, 3, 4]
	// }
	var numbers []int
	if *filePath != " " {
		numbers, err = reader.ReadJSONFile(*filePath)
		if err != nil {
			logger.Log(fmt.Sprintf("Reаd file error: %v", err))
			fmt.Printf("Reаd file error: %v", err)
			return
		}
		logger.Log("Read data from JSON file")
	} else {
		numbers, err = reader.ReadStdin()
		if err != nil {
			logger.Log(fmt.Sprintf("Error reading from stdin: %v", err))
			fmt.Printf("Error reading from stdin: %v", err)
			return
		}
		logger.Log("Success read data from stdin")
		fmt.Printf("Success read data from stdin\n")
	}
	// Calculation of total sum
	sum := Summarizer(numbers)
	logger.Log(fmt.Sprintf("Total sum: %d", sum))
	fmt.Printf("Total sum: %d\n", sum)
	
	// Make HTTP GET request ifflag -url is set.
	if *url != " " {
		resp, err := http.Get(*url)
		if err != nil {
			logger.Log(fmt.Sprintf("Error URL: %v", err))
			fmt.Printf("Error URL: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logger.Log(fmt.Sprintf("Bad HTTP request on URL: %s Status code: %d", *url, resp.StatusCode))
			fmt.Printf("Bad HTTP request on URL: %s Status code: %d\n", *url, resp.StatusCode)
			return
		}
		logger.Log(fmt.Sprintf("Success HTTP request on URL: %s", *url))
		fmt.Printf("Success HTTP request on URL: %s\n", *url)
	} 
}

func Summarizer(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}