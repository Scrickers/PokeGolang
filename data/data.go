package data

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Bag struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Heal        int    `json:"Heal"`
	DropRate    int    `json:"DropRate"`
	Id          int    `json:"Id"`
}
type Pokemon struct {
	Name   string   `json:"Name"`
	Type   []string `json:"Type"`
	HP     int      `json:"HP"`
	MaxHP  int      `json:"MaxHP"`
	Dmg    int      `json:"Dmg"`
	Active bool     `json:"Active"`
}

var turn = 1

func WriteJSONFile(fileName string, data []Pokemon) error {
	jsonFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonParse := json.NewEncoder(jsonFile)
	err = jsonParse.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func WriteJSONFileBag(fileName string, data []Bag) error {
	jsonFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonParse := json.NewEncoder(jsonFile)
	err = jsonParse.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func ReadJSONFilePoke(fileName string) ([]Pokemon, error) {
	var poke []Pokemon
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonParse := json.NewDecoder(jsonFile)
	err = jsonParse.Decode(&poke)
	if err != nil {
		fmt.Println(err)
	}
	return poke, err
}

func ReadJSONFileBag(fileName string) ([]Bag, error) {
	var bag []Bag
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonParse := json.NewDecoder(jsonFile)
	err = jsonParse.Decode(&bag)
	if err != nil {
		fmt.Println(err)
	}
	return bag, err
}

func GetTurn() string {
	return strconv.Itoa(turn)
}
func NextTurn() {
	switch turn {
	case 1:
		fmt.Println("It's turn of player 2 !")
		turn = 2
	case 2:
		fmt.Println("It's turn of player 1 !")
		turn = 1
	}
}
func ClearChat() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Wait() {
	var input string
	fmt.Scanln(&input)
}
