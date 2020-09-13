package csfriendcode

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"math/bits"
	"strings"
)

const (
	alnum = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
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

func friendCode(id uint64) (string, error) {
	h, err := hashSteamID(id)
	if err != nil {
		return "", err
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

	return b32(r), nil
}

// FriendCode gets a friend code based on a provided steamid64
func FriendCode(id uint64) string {
	fc, err := friendCode(id)
	if err != nil {
		return ""
	}
	if strings.Contains(fc, "AAAA-") {
		fc = fc[5:]
	}
	return fc
}
