package commands


import (
	"pokedex/internal/models"
	"pokedex/internal/pokecache"
)

type cliCommand struct {
	Name string
	Description string
	ExecFunc func(*models.Config, []string, *pokecache.Cache, map[string]models.Pokemon) error
}


func CreateCommands() map[string]cliCommand {
	return map[string]cliCommand {
	    "help": {
	        Name: "help",
	        Description: "Displays a help message",
	        ExecFunc: commandHelp,
	    },
	    "exit": {
	        Name: "exit",
	        Description: "Exit the Pokedex",
	        ExecFunc: commandExit,
	    },
	    "clear": {
	    	Name: "clear",
	    	Description: "Clears all text on screen",
	    	ExecFunc: commandClear,
	    },
	    "map": {
	    	Name: "map",
	    	Description: "Displays the next 20 locations in the Pokemon world",
	    	ExecFunc: commandMap,
	    },
	    "mapb": {
	    	Name: "mapb",
	    	Description: "Displays the previous 20 locations in the Pokemon world",
	    	ExecFunc: commandMapb,
	    },
	    "explore": {
	    	Name: "explore",
	    	Description: "Explore the pokemons in a given area.",
	    	ExecFunc: commandExplore,
	    },
	    "catch": {
	    	Name: "catch",
	    	Description: "Attempts to catch a pokemon adding it to the Pokedex",
	    	ExecFunc: commandCatch,
	    },
	    "inspect": {
	    	Name: "inspect",
	    	Description: "Displays name, height, weight, stats and type(s) of a Pokemon in your Pokedex.",
	    	ExecFunc: commandInspect,
	    },
	    "pokedex": {
	    	Name: "pokedex",
	    	Description: "Displays all pokemons in your Pokedex.",
	    	ExecFunc: commandPokedex,
	    },
	}
}
