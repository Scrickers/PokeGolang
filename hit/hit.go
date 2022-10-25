package hit

import (
	"fmt"
	"math"
	"pokemon/data"
)

func PokemonAttack() {
	PokeData1, _ := data.ReadJSONFilePoke("../data/player1/Pokemon.json")
	PokeData2, _ := data.ReadJSONFilePoke("../data/player2/Pokemon.json")
	j, k := getPokemonActive()
	var PokeAttack data.Pokemon
	var PokeDefend data.Pokemon
	if data.GetTurn() == "1" {
		PokeAttack = PokeData1[j]
		PokeDefend = PokeData2[k]
	} else {
		PokeAttack = PokeData2[k]
		PokeDefend = PokeData1[j]
	}
	dmg := PokeAttack.Dmg
	for i := 0; i < len(PokeAttack.Type); i++ {
		for o := 0; o < len(PokeDefend.Type); o++ {
			dmg = int(math.Ceil(dmgTypePokemon(PokeAttack.Type[i], PokeDefend.Type[o], float64(dmg))))
		}
	}
	PokeDefend.HP -= dmg
	if data.GetTurn() == "1" {
		PokeData2[k] = PokeDefend
		if PokeDefend.HP <= 0 {
			PokeData2 = remove(PokeData2, k)
			if len(PokeData2) == 0 {
				fmt.Println("Player 1 win")
				return
			}
			PokeData2[0].Active = true
		}
		data.WriteJSONFile("../data/player2/Pokemon.json", PokeData2)
	} else {
		PokeData1[j] = PokeDefend
		if PokeDefend.HP <= 0 {
			PokeData1 = remove(PokeData1, k)
			if len(PokeData1) == 0 {
				fmt.Println("Player 2 win")
				return
			}
			PokeData1[0].Active = true
		}
		data.WriteJSONFile("../data/player1/Pokemon.json", PokeData1)
	}
}
func remove(slice []data.Pokemon, s int) []data.Pokemon {
	PokeData, _ := data.ReadJSONFilePoke("../data/player" + data.GetTurn() + "/PokeDeath.json")
	t := append(PokeData, slice[s])
	data.WriteJSONFile("../data/player"+data.GetTurn()+"/PokeDeath.json", t)
	return append(slice[:s], slice[s+1:]...)
}
func getPokemonActive() (int, int) {
	var active1 int
	var active2 int

	UserData1, _ := data.ReadJSONFilePoke("../data/player1/Pokemon.json")
	for i := 0; i < len(UserData1); i++ {
		if UserData1[i].Active {
			active1 = i
		}
	}
	UserData2, _ := data.ReadJSONFilePoke("../data/player2/Pokemon.json")
	for i := 0; i < len(UserData2); i++ {
		if UserData2[i].Active {
			active2 = i
		}
	}
	return active1, active2

}

func dmgTypePokemon(attack, defend string, dmg float64) float64 {
	switch attack {
	case "Normal":
		switch defend {
		case "Rock":
			return math.Ceil(dmg * 0.5)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		case "Ghost":
			return 0
		}
	case "Fire":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 0.5)
		case "Water":
			return math.Ceil(dmg * 0.5)
		case "Grass":
			return math.Ceil(dmg * 2)
		case "Ice":
			return math.Ceil(dmg * 2)
		case "Bug":
			return math.Ceil(dmg * 2)
		case "Rock":
			return math.Ceil(dmg * 0.5)
		case "Dragon":
			return math.Ceil(dmg * 0.5)
		case "Steel":
			return math.Ceil(dmg * 2)
		}
	case "Water":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 2)
		case "Water":
			return math.Ceil(dmg * 0.5)
		case "Grass":
			return math.Ceil(dmg * 0.5)
		case "Ground":
			return math.Ceil(dmg * 2)
		case "Rock":
			return math.Ceil(dmg * 2)
		case "Dragon":
			return math.Ceil(dmg * 0.5)
		}
	case "Electric":
		switch defend {
		case "Water":
			return math.Ceil(dmg * 2)
		case "Electric":
			return math.Ceil(dmg * 0.5)
		case "Grass":
			return math.Ceil(dmg * 0.5)
		case "Ground":
			return 0
		case "Flying":
			return math.Ceil(dmg * 2)
		case "Dragon":
			return math.Ceil(dmg * 0.5)
		}
	case "Grass":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 0.5)
		case "Water":
			return math.Ceil(dmg * 2)
		case "Electric":
			return math.Ceil(dmg * 2)
		case "Grass":
			return math.Ceil(dmg * 0.5)
		case "Ice":
			return math.Ceil(dmg * 2)
		case "Poison":
			return math.Ceil(dmg * 0.5)
		case "Ground":
			return math.Ceil(dmg * 2)
		case "Flying":
			return math.Ceil(dmg * 0.5)
		case "Bug":
			return math.Ceil(dmg * 0.5)
		case "Rock":
			return math.Ceil(dmg * 2)
		case "Dragon":
			return math.Ceil(dmg * 0.5)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	case "Ice":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 0.5)
		case "Water":
			return math.Ceil(dmg * 0.5)
		case "Grass":
			return math.Ceil(dmg * 2)
		case "Ice":
			return math.Ceil(dmg * 0.5)
		case "Ground":
			return math.Ceil(dmg * 2)
		case "Flying":
			return math.Ceil(dmg * 2)
		case "Dragon":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	case "Fighting":
		switch defend {
		case "Normal":
			return math.Ceil(dmg * 2)
		case "Ice":
			return math.Ceil(dmg * 2)
		case "Poison":
			return math.Ceil(dmg * 0.5)
		case "Flying":
			return math.Ceil(dmg * 0.5)
		case "Psychic":
			return math.Ceil(dmg * 0.5)
		case "Bug":
			return math.Ceil(dmg * 0.5)
		case "Rock":
			return math.Ceil(dmg * 2)
		case "Ghost":
			return 0
		case "Dark":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 2)
		case "Fairy":
			return math.Ceil(dmg * 0.5)
		}
	case "Poison":
		switch defend {
		case "Grass":
			return math.Ceil(dmg * 2)
		case "Poison":
			return math.Ceil(dmg * 0.5)
		case "Ground":
			return math.Ceil(dmg * 0.5)
		case "Rock":
			return math.Ceil(dmg * 0.5)
		case "Ghost":
			return math.Ceil(dmg * 0.5)
		case "Steel":
			return 0
		case "Fairy":
			return math.Ceil(dmg * 2)
		}
	case "Ground":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 2)
		case "Electric":
			return math.Ceil(dmg * 2)
		case "Grass":
			return math.Ceil(dmg * 0.5)
		case "Poison":
			return math.Ceil(dmg * 2)
		case "Flying":
			return 0
		case "Bug":
			return math.Ceil(dmg * 0.5)
		case "Rock":
			return math.Ceil(dmg * 0.5)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	case "Flying":
		switch defend {
		case "Electric":
			return math.Ceil(dmg * 0.5)
		case "Grass":
			return math.Ceil(dmg * 2)
		case "Fighting":
			return math.Ceil(dmg * 2)
		case "Bug":
			return math.Ceil(dmg * 2)
		case "Rock":
			return math.Ceil(dmg * 0.5)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	case "Psychic":
		switch defend {
		case "Fighting":
			return math.Ceil(dmg * 2)
		case "Poison":
			return math.Ceil(dmg * 2)
		case "Psychic":
			return math.Ceil(dmg * 0.5)
		case "Dark":
			return 0
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	case "Bug":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 0.5)
		case "Grass":
			return math.Ceil(dmg * 2)
		case "Fighting":
			return math.Ceil(dmg * 0.5)
		case "Poison":
			return math.Ceil(dmg * 0.5)
		case "Flying":
			return math.Ceil(dmg * 0.5)
		case "Psychic":
			return math.Ceil(dmg * 2)
		case "Ghost":
			return math.Ceil(dmg * 0.5)
		case "Dark":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		case "Fairy":
			return math.Ceil(dmg * 0.5)
		}
	case "Rock":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 2)
		case "Ice":
			return math.Ceil(dmg * 2)
		case "Fighting":
			return math.Ceil(dmg * 0.5)
		case "Ground":
			return math.Ceil(dmg * 0.5)
		case "Flying":
			return math.Ceil(dmg * 2)
		case "Bug":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	case "Ghost":
		switch defend {
		case "Normal":
			return 0
		case "Psychic":
			return math.Ceil(dmg * 2)
		case "Ghost":
			return math.Ceil(dmg * 2)
		case "Dark":
			return math.Ceil(dmg * 0.5)
		}
	case "Dragon":
		switch defend {
		case "Dragon":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		case "Fairy":
			return 0
		}
	case "Dark":
		switch defend {
		case "Fighting":
			return math.Ceil(dmg * 0.5)
		case "Psychic":
			return math.Ceil(dmg * 2)
		case "Ghost":
			return math.Ceil(dmg * 2)
		case "Dark":
			return math.Ceil(dmg * 0.5)
		case "Fairy":
			return math.Ceil(dmg * 0.5)
		}
	case "Steel":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 0.5)
		case "Water":
			return math.Ceil(dmg * 0.5)
		case "Electric":
			return math.Ceil(dmg * 0.5)
		case "Ice":
			return math.Ceil(dmg * 2)
		case "Rock":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		case "Fairy":
			return math.Ceil(dmg * 2)
		}
	case "Fairy":
		switch defend {
		case "Fire":
			return math.Ceil(dmg * 0.5)
		case "Fighting":
			return math.Ceil(dmg * 2)
		case "Poison":
			return math.Ceil(dmg * 0.5)
		case "Dragon":
			return math.Ceil(dmg * 2)
		case "Dark":
			return math.Ceil(dmg * 2)
		case "Steel":
			return math.Ceil(dmg * 0.5)
		}
	}
	return dmg
}
