package commands


import(
	"io"
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
	"pokedex/internal/models"
	"pokedex/internal/pokecache"
	"time"
	"math/rand"	


)


func commandMap(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	var res *http.Response
	var err error
	var fullURL string

	if c.Next == "" {
		fullURL = "https://pokeapi.co/api/v2/location-area"
	} else {
		fullURL = c.Next
	}

	data, ok := cache.Get(fullURL)
	if ok {
		// Cache hit
		fmt.Println("Cache hit")
		var pokeLocation models.PokemonLocation
		err := json.Unmarshal(data, &pokeLocation)
		if err != nil {
			return err
		}

		allLocations := pokeLocation.Results

		for _, location := range allLocations {
			fmt.Println(location.Name)
		}

		c.Next = pokeLocation.Next
		c.Previous = pokeLocation.Previous

		return nil
	}

	fmt.Println("Cache miss")
	res, err = http.Get(fullURL)

	
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(fullURL, body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Reponse returned with status code: %i", res.StatusCode)
	}

	
	// fmt.Printf("%s", body)

	var pokeLocation models.PokemonLocation
	err = json.Unmarshal(body, &pokeLocation)
	if err != nil {
		return err
	}

	allLocations := pokeLocation.Results

	for _, location := range allLocations {
		fmt.Println(location.Name)
	}

	c.Next = pokeLocation.Next
	c.Previous = pokeLocation.Previous

	return nil
}


func commandMapb(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	if c.Previous == nil {
		fmt.Printf("Can't go more back!")
		return fmt.Errorf("Can't go more back!")
	}
	c.Next = c.Previous.(string)
	return commandMap(c, arrayUserInput, cache, pokedex)
}


func commandExplore(c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	if len(arrayUserInput) != 2 {
		fmt.Printf("Explore command must have exactly one location area as an argument!\n")
		return errors.New("Invalid number of args")
	}

	locationName := arrayUserInput[1]

	fullURL := "https://pokeapi.co/api/v2/location-area/" + locationName
	
	
	data, ok := cache.Get(fullURL)
	if ok {
		// Cache hit
		fmt.Println("Cache hit")
		var myExploringLocation models.ExploringLocation
		err := json.Unmarshal(data, &myExploringLocation)

		if err != nil {
			return err
		}

		for _, pokemon := range myExploringLocation.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}

		return nil
	}


	fmt.Println("Cache miss")
	res, err := http.Get(fullURL)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(fullURL, body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Reponse returned with status code: %i", res.StatusCode)
	}

	var myExploringLocation models.ExploringLocation
	err = json.Unmarshal(body, &myExploringLocation)

	if err != nil {
		return err
	}

	for _, pokemon := range myExploringLocation.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}



func commandCatch (c *models.Config, arrayUserInput []string, cache *pokecache.Cache, pokedex map[string]models.Pokemon) error {
	if len(arrayUserInput) != 2 {
		fmt.Printf("Catch command must have exactly one pokemon as an argument!\n")
		return errors.New("Invalid number of args")
	}

	pokemonName := arrayUserInput[1]

	_, ok := pokedex[pokemonName]
	if ok {
		fmt.Printf("%s has already been caught!\n", pokemonName)
		return errors.New("Already caught")
	}

	fullURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	
	
	data, ok := cache.Get(fullURL)
	if ok {
		// Cache hit
		fmt.Printf("Cache hit\n")
		var pokemon models.Pokemon
		err := json.Unmarshal(data, &pokemon)

		if err != nil {
			return err
		}

		fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)


		rand.Seed(time.Now().UnixNano())

		minExp := 36
		maxExp := 636

		difficulty := float64(pokemon.BaseExperience-minExp) / float64(maxExp-minExp)
		catchChance := rand.Float64()

		if catchChance > difficulty {
			pokedex[pokemonName] = pokemon
			fmt.Printf("%s was caught!\n", pokemonName)
			fmt.Printf("You may now inspect it with the inspect command.\n")
		} else {
			fmt.Printf("%s escaped!\n", pokemonName)
		}

		return nil

	}


	fmt.Println("Cache miss")
	res, err := http.Get(fullURL)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(fullURL, body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Reponse returned with status code: %i", res.StatusCode)
	}

	var pokemon models.Pokemon
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	rand.Seed(time.Now().UnixNano())

	minExp := 36
	maxExp := 636

	difficulty := float64(pokemon.BaseExperience-minExp) / float64(maxExp-minExp)
	catchChance := rand.Float64()

	if catchChance > difficulty {
		pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}


