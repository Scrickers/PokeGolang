package pokemon

func PokemonData() map[string]map[string]interface{} {
	pokemons := map[string]map[string]interface{}{
		"Bulbasaur": {
			"Type":  []string{"Grass", "Poison"},
			"HP":    22,
			"MaxHP": 22,
			"Dmg":   4,
		},
		"Charmander": {
			"Type":  []string{"Fire"},
			"HP":    24,
			"MaxHP": 24,
			"Dmg":   6,
		},
	}
	return pokemons
}
