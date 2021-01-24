package util

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/colbyleiske/pokemonpartytracker/lib"
)

//Assumes addr val (in hexadecimal) is even number
func ConvertHexToString(hexVal []byte) (string, error) {
	hexStr := fmt.Sprintf("%X", hexVal)
	if len(strings.ReplaceAll(hexStr, "0", "")) == 0 {
		return "", errors.New("missing name data")
	}
	strLen := len(hexStr) / 2
	var convertedString string
	for i := 0; i < strLen; i++ {
		//mask byte by byte to get true value
		character := hexStr[i*2 : (i*2)+2]
		x, y := ConvertByteToIndex(character)
		convertedString = fmt.Sprintf("%v%s", convertedString, lib.CharacterTable[x][y])
	}

	return convertedString, nil
}

// func ConvertBytesToString(bytes uint64) (string, error){

// }

// uses little endian
func ConvertStringToHex(val string) (uint64, error) {
	var convertedHex string
	for _, v := range val {
		letter := fmt.Sprintf("%c", v)

		var finalX, finalY int
		//find letter in the characerTable
		for k, v := range lib.CharacterTable {
			for k2, v2 := range v {
				if v2 == letter {
					finalX = k
					finalY = k2
				}
			}
		}
		hexVal := fmt.Sprintf("%X%X", finalX, finalY)
		convertedHex = fmt.Sprintf("%v%v", hexVal, convertedHex)
	}
	fmt.Println(convertedHex)
	return strconv.ParseUint(convertedHex, 16, 64)
}

func ConvertByteToIndex(byteVal string) (int, int) {
	xStr := byteVal[:1]
	yStr := byteVal[1:]
	x, err := strconv.ParseInt(xStr, 16, 64)
	if err != nil {
		log.Println(err)
	}
	y, err := strconv.ParseInt(yStr, 16, 64)
	if err != nil {
		log.Println(err)
	}
	return int(x), int(y)
}
