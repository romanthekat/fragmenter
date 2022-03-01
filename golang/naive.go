package main

import (
	"fmt"
)

type NaiveSplitter struct{}

func (n NaiveSplitter) Split(payload []byte, id uint32, mtu uint16) ([]Fragment, error) {
	err := validateMtu(mtu)
	if err != nil {
		return nil, err
	}

	fragmentsCount := (uint16(len(payload)) / mtu) + 1

	var fragments []Fragment

	var num uint16
	for num = 0; num < fragmentsCount; num++ {
		fragment := getFragment(num, payload, id, mtu, num < fragmentsCount-1)
		fragments = append(fragments, fragment)
	}

	return fragments, nil
}

func getFragment(num uint16, payload []byte, id uint32, mtu uint16, more bool) Fragment {
	fragment = NewFragment(num * mtu, more, id, payload)
}

func validateMtu(mtu uint16) error {
	if mtu < 1280 {
		return fmt.Errorf("mtu must be 1280 at least")
	}

	if mtu%8 != 0 {
		return fmt.Errorf("mtu must be multiple of 8")
	}

	return nil
}
