package main

import (
	"fmt"
	"log"

	goxymemmory "github.com/Xustyx/goxymemory"
	"github.com/colbyleiske/pokemonpartytracker/model"
	"github.com/colbyleiske/pokemonpartytracker/util"
)

func main() {
	// trainerName := "COLB"
	trainerName := "AAAA" // BBBBBBBB 00FFBBBB
	// trainerName := "BCC6C9BD"

	//Init
	dm := goxymemmory.DataManager("mGBA.exe") //Get the DataManager with the process passed.
	if !dm.IsOpen {                           //Check if process was opened.
		fmt.Printf("Failed opening process.\n")
		return
	}

	uTrainerName, err := util.ConvertStringToHex(trainerName)
	if err != nil {
		log.Println(err)
	}
	hTrainerName := fmt.Sprintf("%X", uTrainerName)
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

	var mask uint = 0x00004500
	prefix := finalAddr ^ mask

	var partyStart uint = 0x000044ec
	startingSlot := prefix + partyStart

	partySlots := make([]model.PokemonBytes, 6)

	for p := 0; p < 6; p++ {
		byteData := make([]byte, 100)
		for i := 0; i < 100; i++ {
			slot := startingSlot + uint(i*0x01) + uint(100 * p) 

			data, err := dm.Read(slot, goxymemmory.BYTE)
			if err != nil { //Check if not failed.
				fmt.Printf("Failed reading memory. %s", err)
			}

			byteData[i] = data.Value.(byte)
		}
		partySlots[p] = model.CreatePokemonBytes(byteData)
	}

	fmt.Println(partySlots[2].GetNickname())

}
