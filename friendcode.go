package csfriendcode

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"math/bits"
)

const (
	alnum          = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	defaultSteamID = 0x110000100000000
)

var (
	ralnum = map[byte]uint64{
		'A': 0,
		'B': 1,
		'C': 2,
		'D': 3,
		'E': 4,
		'F': 5,
		'G': 6,
		'H': 7,
		'J': 8,
		'K': 9,
		'L': 10,
		'M': 11,
		'N': 12,
		'P': 13,
		'Q': 14,
		'R': 15,
		'S': 16,
		'T': 17,
		'U': 18,
		'V': 19,
		'W': 20,
		'X': 21,
		'Y': 22,
		'Z': 23,
		'2': 24,
		'3': 25,
		'4': 26,
		'5': 27,
		'6': 28,
		'7': 29,
		'8': 30,
		'9': 31,
	}
)

// b32 is some kind of base32 that is used
// to encode the value that is created from the initial
// function. For some reason that I cant figure out
// it creates a different result to using go's base32 function
// even though they should basically do exactly the same thing

func b32(input uint64) string {
	result := []byte{}

	// big endian the number
	input = bits.ReverseBytes64(input)

	for i := 0; i < 13; i++ {
		if i == 4 || i == 9 {
			result = append(result, '-')
		}
		result = append(result, alnum[input&0x1F])
		input >>= 5
	}

	return string(result)
}

func rb32(input string) uint64 {
	result := uint64(0)

	for i := 0; i < 13; i++ {
		if i == 4 || i == 9 {
			input = input[1:]
		}
		result |= ralnum[input[0]] << (5 * i)

		input = input[1:]
	}

	return bits.ReverseBytes64(result)
}

func hashSteamID(id uint64) (uint32, error) {
	accountID := uint32(id)
	strangeSteamID := uint64(accountID) | 0x4353474F00000000

	steamIDBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(steamIDBytes, strangeSteamID)

	h := md5.New()
	n, err := h.Write(steamIDBytes)

	if err != nil {
		return 0, err
	}

	if n != 8 {
		return 0, errors.New("Couldnt hash steamid")
	}

	hashSteamID := h.Sum(nil)
	return binary.LittleEndian.Uint32(hashSteamID), nil
}

// makeU64 takes a high and low uint32
// and makes a uint64 out of them
func makeU64(hi uint32, lo uint32) uint64 {
	return uint64((uint64(hi) << 32) | uint64(lo))
}

func friendCode(id uint64) string {
	h, err := hashSteamID(id)
	if err != nil {
		return ""
	}

	r := uint64(0)
	for i := 0; i < 8; i++ {
		idNibble := byte(id & 0xF)
		id >>= 4

		hashNibble := (h >> i) & 1

		a := uint32(r<<4) | uint32(idNibble)

		r = makeU64(uint32(r>>28), a)
		r = makeU64(uint32(r>>31), a<<1|hashNibble)
	}

	return b32(r)
}

// Encode gets a friend code based on a provided steamid64
func Encode(id uint64) string {
	fc := friendCode(id)
	if fc[:5] == "AAAA-" {
		fc = fc[5:]
	}
	return fc
}

func steamID(fc string) uint64 {
	val := rb32(fc)
	id := uint32(0)

	for i := 0; i < 8; i++ {
		// hashBit := val & 0x1
		val >>= 1
		idNibble := uint32(val & 0xF)
		val >>= 4

		id <<= 4
		id |= idNibble
	}

	return uint64(id) | defaultSteamID
}

// Decode gets a steamid from a friendcode
func Decode(friendCode string) uint64 {
	if len(friendCode) != 10 {
		return 0
	}
	friendCode = "AAAA-" + friendCode

	return steamID(friendCode)
}
