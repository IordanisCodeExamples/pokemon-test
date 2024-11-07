package main

import pokeapiadapter "github.com/IordanisCodeExamples/pokemon-test/internal/adapters/pokeapi"


func main() {
	client := pokeapiadapter.NewPokeAPIAdapter()
	x,_ := client.SearchPokemon("pikachus")
	if x == nil {
		println("nil")
	}else{
		println(x.Name)
	}
}
		
