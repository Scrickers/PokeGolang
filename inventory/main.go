package inventory

import (
	"fmt"
	"pokemon/data"
	"pokemon/maths"
)

func SetPokemon() {
	nbPoke := 5
	pokeData, _ := data.ReadJSONFilePoke("../data/src/PokeData.json")
	pokePLayer := make([]data.Pokemon, nbPoke)
	for i := 0; i < 2; i++ {
		for i := 0; i < nbPoke; i++ {
			o := maths.GetRand(0, len(pokeData))
			for {
				if checkExist(pokePLayer, pokeData[o].Name) {
					o = maths.GetRand(0, len(pokeData))
				} else {
					break
				}
			}
			if i == 0 {
				pokeData[o].Active = true
			} else {
				pokeData[o].Active = false
			}
			pokePLayer[i] = pokeData[o]
		}
		fileName := "../data/player" + string(rune('0'+i+1)) + "/Pokemon.json"
		data.WriteJSONFile(fileName, pokePLayer)
	}
}
func SetBag() {
	BagData, _ := data.ReadJSONFileBag("../data/src/BagData.json")
	bagPlayer := make([]data.Bag, 7)
	for o := 0; o < 2; o++ {
		for i := 0; i < 7; i++ {
			RandNbr := maths.GetRand(0, 100)
			if RandNbr < 60 {
				bagPlayer[i] = BagData[0]
			} else if RandNbr < 80 {
				bagPlayer[i] = BagData[1]
			} else if RandNbr < 90 {
				bagPlayer[i] = BagData[2]
			} else if RandNbr < 95 {
				bagPlayer[i] = BagData[3]
			} else {
				bagPlayer[i] = BagData[4]
			}
		}
		fileName := "../data/player" + string(rune('0'+o+1)) + "/Bag.json"
		data.WriteJSONFileBag(fileName, bagPlayer)
	}
}
func checkExist(b []data.Pokemon, strCheck string) bool {
	for i := 0; i < len(b); i++ {
		if b[i].Name == strCheck {
			return true
		}
	}
	return false
}
func ChangePokemon(s int) {
	fileName := "../data/player" + data.GetTurn() + "/Pokemon.json"

	fmt.Println("Changing Pokemon...")
	s -= 1
	pokePLayer, _ := data.ReadJSONFilePoke(fileName)
	if s >= 0 && s <= len(pokePLayer)-1 {
		for i := 0; i < len(pokePLayer); i++ {
			if i == s {
				pokePLayer[i].Active = true
			} else {
				pokePLayer[i].Active = false
			}
		}
		fmt.Println("Change Pokemon")
		data.WriteJSONFile(fileName, pokePLayer)
		return
	}
	fmt.Println("Vous devez choisir un Pokemon entre 1 et ", len(pokePLayer))
	return
}
func Revive() {
	ListPokemonDeath()
	fmt.Println("Which Pokemon do you want to revive ?")
	var s int
	fmt.Scanln(&s)
	s -= 1
	fileName := "../data/player" + data.GetTurn() + "/PokeDeath.json"
	pokePLayer, _ := data.ReadJSONFilePoke(fileName)
	fileNames := "../data/player" + data.GetTurn() + "/PokeDeath.json"
	pokePLayers, _ := data.ReadJSONFilePoke(fileNames)
	if s >= 0 && s <= len(pokePLayer)-1 {
		pokePLayer[s].HP = pokePLayer[s].MaxHP
		pokePLayer[s].Active = false
		pokePLayers = append(pokePLayers, pokePLayer[s])
		data.WriteJSONFile(fileName, pokePLayers)
		data.WriteJSONFile(fileName, append(pokePLayer[:s], pokePLayer[s+1:]...))
		fmt.Println("Pokemon revived")
		return
	}
}
func HealthPotion(i int) {
	fileName := "../data/player" + data.GetTurn() + "/Pokemon.json"
	pokePLayer, _ := data.ReadJSONFilePoke(fileName)
	for o := 0; o < len(pokePLayer); o++ {
		if pokePLayer[o].Active {
			pokePLayer[o].HP += i
			if pokePLayer[o].HP > pokePLayer[o].MaxHP {
				pokePLayer[o].HP = pokePLayer[o].MaxHP
			}
		}
	}
	fmt.Println("Ypu used a Health Potion on %s, it now has %d HP", pokePLayer[i].Name, pokePLayer[i].HP)

	data.WriteJSONFile(fileName, pokePLayer)
	return
}
func ListPokemonDeath() {
	fmt.Println("Listing all Pokemon Death...")
	UserData, _ := data.ReadJSONFilePoke("../data/player" + data.GetTurn() + "/PokeDeath.json")
	for i := 0; i < len(UserData); i++ {
		n := ""
		if UserData[i].Active {
			n = "*"
		}
		Pb := maths.ProgressBar(float64(UserData[i].HP * 10 / UserData[i].MaxHP))
		str := fmt.Sprintf("%d%s. %s (%s) [%d/%d]", i+1, n, UserData[i].Name, Pb, UserData[i].HP, UserData[i].MaxHP)
		fmt.Println(str)
	}
}
