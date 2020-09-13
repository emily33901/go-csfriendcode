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
