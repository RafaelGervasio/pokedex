package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"pokedex/internal/models"
	"pokedex/internal/pokecache"
	"errors"

)


func commandHelp(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("\n")

	cliCommands := CreateCommands()
	for _, command := range cliCommands {
		fmt.Printf(command.Name + ": ")
		fmt.Printf(command.Description)
		fmt.Printf("\n")
	}

	return nil
}

func commandExit(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	return nil
}

func commandClear(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		for i := 0; i < 100; i++ {
			fmt.Println()
		}
	}

	return nil
}	

func commandInspect(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	if len(arrayUserInput) != 2 {
		fmt.Printf("Inspect command must have exactly one pokemon area as an argument!\n")
		return errors.New("Invalid number of args")
	}

	pokemonName := arrayUserInput[1]

	pokemon, ok := pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not yet caught %s!\n", pokemonName)
		return errors.New("Pokemon not caught.")
	}

	pokemon.Print()
	return nil
}


func commandPokedex(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	if len(pokedex) == 0 {
		fmt.Println("Your pokedex is empty! Go catch some pokemon first.")
		return errors.New("Empty pokedex")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}




