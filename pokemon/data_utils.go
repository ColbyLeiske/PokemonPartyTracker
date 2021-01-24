package pokemon

import (
	"encoding/binary"
	"fmt"

	"github.com/colbyleiske/pokemonpartytracker/lib"
)

/*


 */

// type DecryptedData []Substructure {
// 	GrowthData []byte
// 	AttackData []byte
// 	EVData     []byte
// 	MiscData   []byte
// }

type Substructure []byte
type DecryptedData map[string]Substructure

type SubstructureOrder []string

const (
	GAME        = "GAME"
	ATTACKS     = "ATTACKS"
	EVCONDITION = "EVCONDITIONS"
	MISC        = "MISC"
)

// https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_data_substructures_in_Generation_III
var Orders = []SubstructureOrder{
	{GAME, ATTACKS, EVCONDITION, MISC}, // GAEM
	{GAME, ATTACKS, MISC, EVCONDITION}, // GAME
	{GAME, EVCONDITION, ATTACKS, MISC}, // GEAM
	{GAME, EVCONDITION, MISC, ATTACKS}, // GEMA
	{GAME, MISC, ATTACKS, EVCONDITION}, // GMAE
	{GAME, MISC, EVCONDITION, ATTACKS}, // GMEA
	{ATTACKS, GAME, EVCONDITION, MISC}, // AGEM
	{ATTACKS, GAME, MISC, EVCONDITION}, // AGME
	{ATTACKS, EVCONDITION, GAME, MISC}, // AEGM
	{ATTACKS, EVCONDITION, MISC, GAME}, // AEMG
	{ATTACKS, MISC, GAME, EVCONDITION}, // AMGE
	{ATTACKS, MISC, EVCONDITION, GAME}, // AMEG
	{EVCONDITION, GAME, ATTACKS, MISC}, // EGAM
	{EVCONDITION, GAME, MISC, ATTACKS}, // EGMA
	{EVCONDITION, ATTACKS, GAME, MISC}, // EAGM
	{EVCONDITION, ATTACKS, MISC, GAME}, // EAMG
	{EVCONDITION, MISC, GAME, ATTACKS}, // EMGA
	{EVCONDITION, MISC, ATTACKS, GAME}, // EMAG
	{MISC, GAME, ATTACKS, EVCONDITION}, // MGAE
	{MISC, GAME, EVCONDITION, ATTACKS}, // MGEA
	{MISC, ATTACKS, GAME, EVCONDITION}, // MAGE
	{MISC, ATTACKS, EVCONDITION, GAME}, // MAEG
	{MISC, EVCONDITION, GAME, ATTACKS}, // MEGA
	{MISC, EVCONDITION, ATTACKS, GAME}, // MEAG
}

func (p PokemonBytes) getDecryptionKey() uint32 {
	return uint32(p.OriginalTrainerID) ^ p.Personality
}

func getNickname(nickIntOne uint64, nickIntTwo uint16) string {
	return BytesToString(nickIntOne) + BytesToString(uint64(nickIntTwo))
}

//https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_data_substructures_in_Generation_III
func decryptData(encryptedData []byte, decryptionKey uint32, personalityValue uint32, checksum uint16) (decryptedData DecryptedData, err error) {
	//determine order of substructures
	order := Orders[personalityValue%24]
	fmt.Printf("suspected order is %v\n", order)
	var runningChecksumTotal uint16
	for i := 0; i < 48; i += 4 {
		// encrytpedSubStructure := encryptedData[i*12 : (i+1)*12]
		encrytpedSubStructure := binary.LittleEndian.Uint32(encryptedData[i:(i + 4)])
		decryptedSubStructure := encrytpedSubStructure ^ decryptionKey

		runningChecksumTotal += uint16(decryptedSubStructure)
	}
	fmt.Printf("ORIGINAL   CHECKSUM: %X\n", checksum)
	fmt.Printf("CALCULATED CHECKSUM: %X\n", runningChecksumTotal)
	// decryptedData[GAME] = Substructure{}
	// decryptedData[ATTACKS] = Substructure{}
	// decryptedData[EVCONDITION] = Substructure{}
	// decryptedData[MISC] = Substructure{}
	return DecryptedData{}, nil
}

func (p PokemonBytes) IsChecksumValid() (bool, error) {
	return false, nil
}

func BytesToString(bytes uint64) (out string) {
	for i := 0; ; i++ {
		shiftedBytes := bytes >> (8 * i)
		letter := shiftedBytes & 0xFF
		if letter == 0 {
			return
		}
		y := (letter & 0xF0) >> 4
		x := (letter & 0x0F)
		letterASCII := lib.CharacterTable[y][x]
		out = fmt.Sprintf("%s%s", out, letterASCII)
	}
	return
}

// func (p PokemonBytes) getSubstructureOrder()
