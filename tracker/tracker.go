package main

import (
	"fmt"
	"log"
	"os"
	"time"

	goxymemmory "github.com/Xustyx/goxymemory"
	"github.com/colbyleiske/pokemonpartytracker/tracker/pokemon"
	"github.com/colbyleiske/pokemonpartytracker/tracker/util"
)

type Tracker interface {
	RefreshParty() error
	TrackParty() error
}

// func NewTracker() *Tracker {
// 	return Tracker{}
// }

func main() {
	trackParty("AAAA")
}

func trackParty(trainerName string) {

	//Init
	dm := goxymemmory.DataManager("mGBA.exe") //Get the DataManager with the process passed.
	if !dm.IsOpen {                           //Check if process was opened.
		fmt.Printf("Failed opening process.\n")
		return
	}

	uTrainerName := util.ConvertStringToBinary(trainerName)

	hTrainerName := fmt.Sprintf("%X", uTrainerName)

	c1 := make(chan uint, 1)

	go func() {
		start := 0x00004500
		var finalAddr uint
		for i := 0x0; ; i += 0x00010000 {
			currentAddress := start + i
			data, err := dm.Read(uint(currentAddress), goxymemmory.UINT) //Reads the string.
			if err != nil {                                              //Check if not failed.
				fmt.Printf("Failed reading memory. %s", err)
			}
			h := fmt.Sprintf("%X", data.Value)
			if h == hTrainerName {
				fmt.Printf("ADDRESS IS %X \n", currentAddress)
				finalAddr = uint(currentAddress)
				break
			}
		}
		c1 <- finalAddr
	}()

	var addr uint
	select {
	case addr = <-c1:
		break
	case <-time.After(5 * time.Second):
		log.Println("Timed out searching for pokemon party. Ensure the game is loaded and you have atleast ONE pokemon, then press Try Again.")
		os.Exit(3)
	}

	var mask uint = 0x00004500
	prefix := addr ^ mask

	var partyStart uint = 0x000044ec
	startingSlot := prefix + partyStart
	const sizeOfPokemon int = 100
	partySlots := make([]pokemon.PokemonBytes, 6)

	for p := 0; p < 6; p++ {
		byteData := make([]byte, 100)
		for i := 0; i < 100; i++ {
			slot := startingSlot + uint(i*0x01) + uint(sizeOfPokemon*p)

			data, err := dm.Read(slot, goxymemmory.BYTE)
			if err != nil { //Check if not failed.
				fmt.Printf("Failed reading memory. %s", err)
			}

			byteData[i] = data.Value.(byte)
		}
		pB, err := pokemon.CreatePokemonBytes(byteData)
		if err != nil {
			// log.Println(err)
			continue
		}
		partySlots[p] = pB
	}

	for _, v := range partySlots {
		if v.Nickname == "" {
			continue
		}
		// fmt.Println(v.Nickname)
		// fmt.Println(v.OriginalTrainerName)
		// p, _ := v.GetPersonalityValue()
		// fmt.Println(p % 24)
	}
}
