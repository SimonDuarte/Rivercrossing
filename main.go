package main

import (
	"fmt"
	"strings"
)


var west = "(Kylling, Korn, Person, Rev ---W___"
var east = "___E--- )"
var boat = "\\__/"

func main() {
	ViewState()
	PutInBoatWest("Rev")
	ViewState()


}

/* Viser status på verden */
func ViewState() {
	fmt.Println(west + boat + east)
}

// Tar et objekt i båten og flytter det til vestsiden
func PutInBoatWest(item string) {
	switch item {
	case "Kylling":
		west = strings.ReplaceAll(west, "Kylling", "")
		boat = boat[:len(boat)-3] + " Kylling " + boat[len(boat)-3:]
		break
	case "Person":
		west = strings.ReplaceAll(west, "Person", "")
		boat = boat[:len(boat)-3] + " Person " + boat[len(boat)-3:]
	case "Rev":
		west = strings.ReplaceAll(west, "Rev", "")
		boat = boat[:len(boat)-3] + " Rev " + boat[len(boat)-3:]
	case "Korn":
		west = strings.ReplaceAll(west, "Korn", "")
		boat = boat[:len(boat)-3] + " Korn " + boat[len(boat)-3:]
	}
}

// Tar et objekt i båten og flytter det til østsiden
func CrossWaterEast(item string) {
	switch item {
	case "Korn":
		boat = strings.ReplaceAll(boat, "Korn", "")
		east = east[:len(east)-1] + " Korn " + east[len(east)-1:]
		break
	case "Rev":
		boat = strings.ReplaceAll(boat, "Rev", "")
		east = east[:len(east)-1] + " Rev " + east[len(east)-1:]
		break
	case "Person":
		boat = strings.ReplaceAll(boat, "Person", "")
		east = east[:len(east)-1] + " Person " + east[len(east)-1:]
		break
	case "Kylling":
		boat = strings.ReplaceAll(boat, "Kylling", "")
		east = east[:len(east)-1] + " Kylling " + east[len(east)-1:]
		break
	}
}

// Henter ut west verdien for å bruke i test på rc.test.go
func GetWest() string {
	return west
}

func GetBoat() string {
	return boat
}

/**
// Ved evt implementering av putInBoatEast, vil også GetEast være en måte å teste varriabler.
func GetEast() string {
	return east
}
*/
