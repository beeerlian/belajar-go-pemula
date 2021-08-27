package main

import "fmt"

type Gamer struct {
	Nama  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)

}

func main() {
	gamer := Gamer{Nama: "Naufal"}
	gamer.AddGame("Mario")
	gamer.AddGame("Pes 2021")
	gamer.AddGame("PUBG Mobile")
	fmt.Println(gamer.Nama)
	fmt.Println("Games : ")
	for index, game := range gamer.Games {
		fmt.Println(index+1, " : ", game)
	}
}
