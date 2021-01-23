package model

import (
	"fmt"
	"log"

	"github.com/colbyleiske/pokemonpartytracker/util"
)

type PokemonBytes struct {
	Personality         []byte
	OriginalTrainerID   []byte
	Nickname            []byte
	Language            []byte
	OriginalTrainerName []byte
	Markings            []byte
	Checksum            []byte
	Unknowns            []byte // ???? - Padding?
	Data                []byte
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

func CreatePokemonBytes(p []byte) PokemonBytes {
	return PokemonBytes{
		Personality:         p[0:4],
		OriginalTrainerID:   p[4:8],
		Nickname:            p[8:18],
		Language:            p[18:20],
		OriginalTrainerName: p[20:27],
		Markings:            p[27:28],
		Checksum:            p[28:30],
		Unknowns:            p[30:32],
		Data:                p[32:80],
		StatusCondition:     p[80:84],
		Level:               p[84:85],
		Pokerus:             p[85:86],
		CurrentHP:           p[86:88],
		TotalHP:             p[88:90],
		Attack:              p[90:92],
		Defense:             p[92:94],
		Speed:               p[94:96],
		SpAttack:            p[96:98],
		SpDefense:           p[98:100],
	}
}

func (p *PokemonBytes) GetNickname() string {
	nameData := fmt.Sprintf("%X", p.Nickname)
	name, err := util.ConvertHexToString(nameData)
	if err != nil {
		log.Printf("Failed to create nickname from bytes because %v\n", err)
		return "MISSING"
	}
	return name
}
