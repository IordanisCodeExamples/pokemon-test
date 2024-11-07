// Package pokeapiadapter provides an adapter for interacting with the PokeAPI.
// It allows searching for Pokémon by name, returning structured Pokémon data,
// and handles errors for cases where data is not found or the API request fails.
package pokeapiadapter

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/pkg/errors"

	internalError "github.com/IordanisCodeExamples/pokemon-test/internal/errors"
	"github.com/IordanisCodeExamples/pokemon-test/internal/model"
)

const (
	pokemonEndpoint = "pokemon"
)


//TODO Maybe create a wrpper as this is annoyinly not good  for testing

// PokeAPIAdapter provides access to Pokémon data from the PokeAPI.
type PokeAPIAdapter struct{}

// NewPokeAPIAdapter initializes and returns a new PokeAPIAdapter instance.
func NewPokeAPIAdapter() *PokeAPIAdapter {
	return &PokeAPIAdapter{}
}

// SearchPokemon retrieves Pokémon data by name. 
// Returns a Pokémon model or an error if not found or if the request fails.
func (p *PokeAPIAdapter) SearchPokemon(name string) (*model.Pokemon, error) {
	result, err := pokeapi.Search(pokemonEndpoint, name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search for Pokémon")
	}

	if len(result.Results) == 0 {
		return nil, internalError.ErrNotFoundRequest
	}

	return &model.Pokemon{
		Name: result.Results[0].Name,
		URL:  result.Results[0].URL,
	}, nil
}
