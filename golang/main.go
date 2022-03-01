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

	payload []byte
}

//Splitter splits arbitrary byte payload into chunks. mtu w/o jumbograms atm
type Splitter interface {
	Split(payload []byte, id uint32, mtu uint16) ([]Fragment, error)
}

//NewFragment new fragment object
func NewFragment(offset uint16, moreFollow bool, id uint32, payload []byte) *Fragment {
	if offset > 0b1111_1111_1111_1 {
		panic(fmt.Sprintf("fragment offset limited to 13 bits, but %d", offset))
	}

	return &Fragment{nextHeader: HEADER_FRAGMENT,
		offsetAndFlag: getOffsetAndFlag(offset, moreFollow),
		id:            id,
		payload:       payload}
}

func getOffsetAndFlag(offset uint16, moreFollow bool) uint16 {
	offsetAndFlag := offset << 3
	if moreFollow {
		offsetAndFlag += 1
	}
	return offsetAndFlag
}

func main() {

}
