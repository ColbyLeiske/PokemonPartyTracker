package pokemon

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type PokemonBytes struct {
	Personality         uint32
	OriginalTrainerID   uint32
	Nickname            string
	Language            []byte
	OriginalTrainerName string
	Markings            []byte
	Checksum            uint16
	Unknowns            []byte // ???? - Padding?
	Data                DecryptedData
	StatusCondition     []byte
	Level               []byte
	Pokerus             []byte
	CurrentHP           []byte
	TotalHP             []byte
	Attack              []byte
	Defense             []byte
	Speed               []byte
	SpAttack            []byte
	SpDefense           []byte
}

func CreatePokemonBytes(p []byte) (pokemon PokemonBytes, err error) {
	byteReader := binary.LittleEndian

	// nickname := util.
	pokemon = PokemonBytes{
		Personality:         byteReader.Uint32(p[0:4]),
		OriginalTrainerID:   byteReader.Uint32(p[4:8]),
		Nickname:            getNickname(byteReader.Uint64(p[8:16]), byteReader.Uint16(p[16:18])), //Little Endian for easier string parsing
		Language:            p[18:20],
		OriginalTrainerName: BytesToString((byteReader.Uint64(p[20:28])), 7),
		Markings:            p[27:28],
		Checksum:            byteReader.Uint16(p[28:30]),
		Unknowns:            p[30:32],
		// Data:                p[32:80],
		StatusCondition: p[80:84],
		Level:           p[84:85],
		Pokerus:         p[85:86],
		CurrentHP:       p[86:88],
		TotalHP:         p[88:90],
		Attack:          p[90:92],
		Defense:         p[92:94],
		Speed:           p[94:96],
		SpAttack:        p[96:98],
		SpDefense:       p[98:100],
	}

	//temp
	if pokemon.Nickname == "" {
		return PokemonBytes{}, errors.New("pokemon not found")
	}

	pokemon.Data, err = decryptData(p[32:80], pokemon.getDecryptionKey(), pokemon.Personality, pokemon.Checksum)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
