package tron

import "testing"

func TestAddress(t *testing.T) {
	s := NewTronService()
	//ethAddr := "0x418aa06d692d0f98f8d3c24cf9901371b021ef555b"
	ethAddr := "0x8aa06d692d0f98f8d3c24cf9901371b021ef555b"
	tronAddr, err := s.ToTronAddress(ethAddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tronAddr = %s", tronAddr)

	ethAddr, err = s.ToEthAddress(tronAddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ethAddr = %s", ethAddr)
}
