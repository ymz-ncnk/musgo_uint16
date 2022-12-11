package musgo_uint16

import "testing"

func TestMusgoUint16(t *testing.T) {
	var n uint16 = 5
	buf := make([]byte, Uint16(n).SizeMUS())
	Uint16(n).MarshalMUS(buf)

	var an uint16
	(*Uint16)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
