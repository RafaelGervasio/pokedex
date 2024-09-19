package main

import (
	"fmt"
	"pokedex/internal/commands"
	"pokedex/internal/models"
	"bufio"
	"strings"
	"os"
	"pokedex/internal/pokecache"
	"time"
)


func main() {
	cliCommands := commands.CreateCommands()
	myConfig := models.NewConfig()
	myCache := pokecache.NewCache(time.Hour)
	pokedex := make(map[string]models.Pokemon)

	for {
		arrayUserInput := getInput()
		
		command, ok := cliCommands[arrayUserInput[0]]
		
		if arrayUserInput[0] == "exit" {
			fmt.Println("Exiting program.")
			break
		} else if ok {
			command.ExecFunc(myConfig, arrayUserInput, myCache, pokedex)
		} else {
			fmt.Printf("Invalid command!\n")
		}
	}
}



func getInput() (arrayUserInput []string) {
	fmt.Printf("Pokedex > ")
	
	var userInput string
	myScanner := bufio.NewScanner(os.Stdin)
	
	if myScanner.Scan() {
		userInput = myScanner.Text()
		arrayUserInput = strings.Fields(userInput)
	}

	return arrayUserInput
}


