package util

import (
	"github.com/colbyleiske/pokemonpartytracker/tracker/lib"
)

func ConvertStringToBinary(val string) (conv uint64) {
	for _, letter := range val {
		//find letter in the characerTable
		for x, v := range lib.CharacterTable {
			for y, v2 := range v {
				if v2 == string(letter) {
					conv = (conv << 8) + uint64((x<<4)+y)
					break
				}
			}
		}
	}
	return
}
