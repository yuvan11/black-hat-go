package main

import (
	"fmt"
	"log"
	"os"

	"black-hat-go/chapter3/metasploit/rpc"

	"github.com/joho/godotenv"
)

func main() {

	// .env file is located inside chapter 3 outside of basic and metasploit
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Println("error in loading .env file")
	}
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")

	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required environment variable MSFHOST or MSFPASS")
	}

	msf, err := rpc.New(host, user, pass)

	fmt.Println("MSF", msf)

	if err != nil {

		log.Fatalln("Cannot generate new user access")
	}
	defer msf.Logout()

	sessions, err := msf.SessionList()

	fmt.Println("Sessions", sessions)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}

}
