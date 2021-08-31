package main

import (
	"black-hat-go/chapter3/shodan/shodan"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	// .env file is located inside chapter 3 outside of basic and shadon
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error in loading .env")
	}
	apikey := os.Getenv("SHODAN_API_KEY")

	s := shodan.New(apikey)

	info, err := s.APIInfo()

	fmt.Println(info)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])

	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
