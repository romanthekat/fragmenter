package main

import "fmt"

const (
	COMMON_MTU      = 1480
	HEADER_FRAGMENT = 44
)

//Fragment defines IPv6 fragment
type Fragment struct {
	nextHeader byte
	reserved   byte

	offsetAndFlag uint16

	id uint32
}

//NewFragment new fragment object
func NewFragment(offset uint16, moreFollow bool, id uint32) *Fragment {
	if offset > 0b1111_1111_1111_1 {
		panic(fmt.Sprintf("fragment offset limited to 13 bits, but %d", offset))
	}

	return &Fragment{nextHeader: HEADER_FRAGMENT, offsetAndFlag: getOffsetAndFlag(offset, moreFollow), id: id}
}

func getOffsetAndFlag(offset uint16, moreFollow bool) uint16 {
	offsetAndFlag := offset << 3
	if moreFollow {
		offsetAndFlag += 1
	}
	return offsetAndFlag
}

//Split splits arbitrary byte payload into chunks. mtu w/o jumbograms atm
func Split(payload []byte, id uint32, mtu uint16) ([]Fragment, error) {
	err := validateMtu(mtu)
	if err != nil {
		return nil, err
	}

	return nil, nil
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

func main() {

}
