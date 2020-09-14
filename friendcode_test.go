package csfriendcode

import (
	"testing"
)

func TestFriendCode(t *testing.T) {
	got := FriendCode(76561197960287930)
	if got != "SUCVS-FADA" {
		t.Errorf("GetFriendCode(76561197960287930) = %s; want 'SUCVS-FADA'", got)
	}
	got = FriendCode(76561197960265729)
	if got != "AJJJS-ABAA" {
		t.Errorf("GetFriendCode(76561197960265729) = %s; want 'AJJJS-ABAA'", got)
	}
	got = FriendCode(76561198259812645)
	if got != "SN7N4-D5HG" {
		t.Errorf("GetFriendCode(76561198259812645) = %s; want 'SN7N4-D5HG'", got)
	}
}

func TestRb32(t *testing.T) {
	x := rb32("AAAA-SUCVS-FADA")
	if x != 711231307777 {
		t.Errorf("rb32(AAAA-SUCVS-FADA) = %d; want 711231307777", x)
	}
}

func TestSteamID(t *testing.T) {
	got := SteamID("SUCVS-FADA")
	if got != 76561197960287930 {
		t.Errorf("SteamID(SUCVS-FADA) = %d; want 76561197960287930", got)
	}
	got = SteamID("AJJJS-ABAA")
	if got != 76561197960265729 {
		t.Errorf("SteamID(AJJJS-ABAA) = %d; want 76561197960265729", got)
	}
	got = SteamID("SN7N4-D5HG")
	if got != 76561198259812645 {
		t.Errorf("SteamID(SN7N4-D5HG) = %d; want 76561198259812645", got)
	}

}
