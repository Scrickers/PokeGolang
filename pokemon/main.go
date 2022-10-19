package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var fileName1 = "Pokemon.json"
var fileName2 = "Bag.json"
var nb = 5
var GI = 0
var DevMode = true

func main() {
	getData("player1/" + fileName1)
	GameName()
	//fmt.Println(getMapWithRandomKey(PokemonData()))
	//action()
}
func getMapWithRandomKey(m map[string]map[string]interface{}) map[string]interface{} {
	r := getRand(0, len(m))
	i := 0
	for _, v := range m {
		if i == r {
			return v
		}
		i++
	}
	return nil
}

func getRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func PokemonData() map[string]map[string]interface{} {
	pokemons := map[string]map[string]interface{}{
		"Bulbasaur": {
			"Name":  "Bulbasaur",
			"Type":  []string{"Grass", "Poison"},
			"HP":    22,
			"MaxHP": 22,
			"Dmg":   4,
		},
		"Charmander": {
			"Name":  "Charmander",
			"Type":  []string{"Fire"},
			"HP":    24,
			"MaxHP": 24,
			"Dmg":   6,
		}, "Squirtle": {
			"Name":  "Squirtle",
			"Type":  []string{"Water"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   5,
		},
		"Pikachu": {
			"Name":  "Pikachu",
			"Type":  []string{"Electric"},
			"HP":    25,
			"MaxHP": 25,
			"Dmg":   5,
		},
		"Weedle": {
			"Name":  "Weedle",
			"Type":  []string{"Bug", "Poison"},
			"HP":    16,
			"MaxHP": 16,
			"Dmg":   2,
		},
		"Ratata": {
			"Name":  "Ratata",
			"Type":  []string{"Normal"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   4,
		},
	}
	return pokemons
}
func GameName() {
	fmt.Println(".........................................................................................")
	fmt.Println("██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗     ██████╗██╗     ██╗")
	fmt.Println("██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║    ██╔════╝██║     ██║")
	fmt.Println("██████╔╝██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║    ██║     ██║     ██║")
	fmt.Println("██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║    ██║     ██║     ██║")
	fmt.Println("██║     ╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║    ╚██████╗███████╗██║")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝╚══════╝╚═╝")
	fmt.Println(".........................................................................................")
	/*

				██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗     ██████╗██╗     ██╗
				██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║    ██╔════╝██║     ██║
				██████╔╝██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║    ██║     ██║     ██║
				██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║    ██║     ██║     ██║
				██║     ╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║    ╚██████╗███████╗██║
				╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝╚══════╝╚═╝
		https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=par%20scricker%20ss
	*/

	createFile()
	setPokemon()
}

func createFile() {
	fileName := []string{fileName1, fileName2}
	for _, v := range fileName {
		_, err := os.Stat(v)

		if !os.IsNotExist(err) {
			os.Remove("player1/" + v)
			log("File player1/" + v + " supprimé")
			os.Remove("player2/" + v)
			log("File player2/" + v + " supprimé")
		}
		file, err := os.Create("player1/" + v)
		if err != nil {
			log(err.Error())
		}
		log("File player1/" + v + " créé")
		file.Close()
		file, err = os.Create("player2/" + v)
		if err != nil {
			log(err.Error())
			fmt.Println(err)
		}
		log("File player2/" + v + " créé")
		file.Close()
	}
}

func waitCommande() string {
	var input string
	fmt.Scanln(&input)
	return input
}
func checkCommande(input string) {
	switch input {
	case "attack":
		fmt.Println("attack")
	default:
		fmt.Println("default")
		checkCommande(waitCommande())
	}

}
func writeDataInFile(fileName, data string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log(err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		log(err.Error())
	}

}
func setPokemon() {
	for i := 1; i < 3; i++ {
		c := strconv.Itoa(i)
		log("Ajout des pokemon dans player" + c)
		writeDataInFile("player"+c+"/"+fileName1, "[\n")
		for j := 0; j < nb; j++ {
			writeDataInFile("player"+c+"/"+fileName1, convertMapToString(getMapWithRandomKey(PokemonData())))
		}
		writeDataInFile("player"+c+"/"+fileName1, "]")
	}
}
func getData(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log(err.Error())
	}
	log(string(content))
}
func convertMapToString(m map[string]interface{}) string {
	var s = "{"
	i := 0
	for k, v := range m {
		i++
		if i == len(m) {
			s += `"` + k + `"` + ":" + `"` + fmt.Sprintf("%v", v) + `"` + "\n"
		} else {
			s += `"` + k + `"` + ":" + `"` + fmt.Sprintf("%v", v) + `"` + ",\n"
		}
	}
	GI++
	if GI == nb {
		s += "}"
	} else {
		s += "},\n"
	}
	return s
}
func log(s string) {
	if DevMode {

		now := time.Now()
		hour, min, sec := now.Clock()
		fmt.Println("\033[31m[DevLog] \033[36m(", hour, ":", min, ":", sec, ") : \033[34m", s, "\033[0m")
	}
}
