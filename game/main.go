package main

import (
	"fmt"
	"pokemon/data"
	"pokemon/hit"
	"pokemon/inventory"
	"pokemon/maths"
)

func main() {
	// set Inventory
	inventory.SetPokemon()
	inventory.SetBag()
	// set Inventory
	//ListBag()
	data.ClearChat()
	StartGame()
}

func printLogo() {
	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[33m")
	fmt.Println("██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗     ██████╗██╗     ██╗")
	fmt.Println("██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║    ██╔════╝██║     ██║")
	fmt.Println("██████╔╝██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║    ██║     ██║     ██║")
	fmt.Println("██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║    ██║     ██║     ██║")
	fmt.Println("██║     ╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║    ╚██████╗███████╗██║")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝╚══════╝╚═╝")
	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[0m")
	fmt.Print("\n")
}
func PrintList() {
	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[33m")
	fmt.Println("██╗     ██╗███████╗████████╗    ██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗")
	fmt.Println("██║     ██║██╔════╝╚══██╔══╝    ██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║")
	fmt.Println("██║     ██║███████╗   ██║       ██████╔╝██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║")
	fmt.Println("██║     ██║╚════██║   ██║       ██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║")
	fmt.Println("███████╗██║███████║   ██║       ██║     ╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║")
	fmt.Println("╚══════╝╚═╝╚══════╝   ╚═╝       ╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝")
	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[0m")
	fmt.Print("\n")
}
func PrintBag() {
	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[33m")
	fmt.Println("██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗    ██████╗  █████╗  ██████╗ ")
	fmt.Println("██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║    ██╔══██╗██╔══██╗██╔════╝ ")
	fmt.Println("██████╔╝██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║    ██████╔╝███████║██║  ███╗")
	fmt.Println("██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║    ██╔══██╗██╔══██║██║   ██║")
	fmt.Println("██║     ╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║    ██████╔╝██║  ██║╚██████╔╝")
	fmt.Println("╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ")
	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[0m")
	fmt.Print("\n")

}

func StartGame() {
	printLogo()
	fmt.Println("Please select an option:")
	fmt.Println("1. List all Pokemon")
	fmt.Println("2. fight")
	fmt.Println("3. Inventory")
	fmt.Println("4. Change Pokemon")
	fmt.Println("5. Exit")

	var option int
	fmt.Scanln(&option)
	data.ClearChat()
	switch option {
	case 1:
		ListPokemon()
		data.Wait()
		data.ClearChat()
	case 2:
		hit.PokemonAttack()
		data.NextTurn()
		data.Wait()
		data.ClearChat()
	case 3:
		ListBag()
		data.Wait()
		data.ClearChat()
	case 4:
		ChangePokemon()
		var option int
		fmt.Scanln(&option)
		inventory.ChangePokemon(option)
		data.Wait()
		data.ClearChat()
	case 5:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid option. Please try again.")
	}
	StartGame()

}

func ListPokemon() {
	PrintList()
	fmt.Print("\n")
	fmt.Println("Your Pokemon:")
	fileName := "../data/player" + data.GetTurn() + "/Pokemon.json"
	UserData, _ := data.ReadJSONFilePoke(fileName)
	for i := 0; i < len(UserData); i++ {
		n := ""
		if UserData[i].Active {
			n = "*"
		}
		Pb := maths.ProgressBar(float64(UserData[i].HP * 10 / UserData[i].MaxHP))
		str := fmt.Sprintf("%d%s. %s (%s) [%d/%d]", i+1, n, UserData[i].Name, Pb, UserData[i].HP, UserData[i].MaxHP)
		fmt.Println(str)
	}

	fmt.Println("\033[34m")
	fmt.Println(".........................................................................................")
	fmt.Println("\033[0m")
	fmt.Println("Enemy Pokemon:")
	if data.GetTurn() == "1" {
		fileName = "../data/player2/Pokemon.json"
	} else {
		fileName = "../data/player1/Pokemon.json"
	}
	UserData, _ = data.ReadJSONFilePoke(fileName)
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

func ChangePokemon() {
	PrintList()
	fmt.Print("\n")
	fmt.Println("Choisissez un pokemon qui sera deffinit comme actif :")
	fileName := "../data/player" + data.GetTurn() + "/Pokemon.json"
	UserData, _ := data.ReadJSONFilePoke(fileName)
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
func ListBag() {
	PrintBag()

	id := []int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	UserData, _ := data.ReadJSONFileBag("../data/player" + data.GetTurn() + "/Bag.json")
	for i := 0; i < len(UserData); i++ {
		id[UserData[i].Id]++
	}
	if id[1] > 0 {
		fmt.Println("1. Potion (x", id[1], ")")
	}
	if id[2] > 0 {
		fmt.Println("2. Super Potion (x", id[2], ")")
	}
	if id[3] > 0 {
		fmt.Println("3. Hyper Potion (x", id[3], ")")
	}
	if id[4] > 0 {
		fmt.Println("4. Max Potion (x", id[4], ")")
	}
	if id[5] > 0 {
		fmt.Println("5. Revive (x", id[5], ")")
	}
	fmt.Println("6. Exit")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		if id[1] == 0 {
			fmt.Println("You do not have any potions.")
			return
		}
		fmt.Println("Using Potion...")
		ItemRemove(1)
	case 2:
		if id[2] == 0 {
			fmt.Println("You do not have any Super potions.")
			return
		}
		fmt.Println("Using Super Potion...")
		ItemRemove(2)
	case 3:
		if id[3] == 0 {
			fmt.Println("You do not have any Hyper potions.")
			return
		}
		fmt.Println("Using Max Potion...")
		ItemRemove(3)
	case 4:
		if id[4] == 0 {
			fmt.Println("You do not have any Hyper potions.")
			return
		}
		fmt.Println("Using Max Potion...")
		ItemRemove(4)
	case 5:
		if id[5] == 0 {
			fmt.Println("You do not have any reveive.")
			return
		}
		fmt.Println("Using Revive...")
		ItemRemove(5)
	case 6:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invallid option. Please try again.")
	}
}
func ItemRemove(s int) {
	UserData, _ := data.ReadJSONFileBag("../data/player" + data.GetTurn() + "/Bag.json")
	for i := 0; i < len(UserData); i++ {
		switch s {
		case 1:
			if UserData[i].Id == 1 {
				UserData = append(UserData[:i], UserData[i+1:]...)
				data.WriteJSONFileBag("../data/player"+data.GetTurn()+"/Bag.json", UserData)
				inventory.HealthPotion(6)
				return
			}
		case 2:
			if UserData[i].Id == 2 {
				UserData = append(UserData[:i], UserData[i+1:]...)
				data.WriteJSONFileBag("../data/player"+data.GetTurn()+"/Bag.json", UserData)
				inventory.HealthPotion(10)
				return
			}
		case 3:
			if UserData[i].Id == 3 {
				UserData = append(UserData[:i], UserData[i+1:]...)
				data.WriteJSONFileBag("../data/player"+data.GetTurn()+"/Bag.json", UserData)
				inventory.HealthPotion(20)
				return
			}
		case 4:
			if UserData[i].Id == 4 {
				UserData = append(UserData[:i], UserData[i+1:]...)
				data.WriteJSONFileBag("../data/player"+data.GetTurn()+"/Bag.json", UserData)
				inventory.HealthPotion(-1)
				return
			}
		case 5:
			if UserData[i].Id == 5 {
				UserData = append(UserData[:i], UserData[i+1:]...)
				data.WriteJSONFileBag("../data/player"+data.GetTurn()+"/Bag.json", UserData)
				inventory.Revive()
				return
			}

		}
	}
}
