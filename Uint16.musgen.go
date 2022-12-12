// Code generated by musgen. DO NOT EDIT.

package musgo_uint16

import "github.com/ymz-ncnk/musgo/errs"

// MarshalMUS fills buf with the MUS encoding of v.
func (v Uint16) MarshalMUS(buf []byte) int {
	i := 0
	{
		for v >= 0x80 {
			buf[i] = byte(v) | 0x80
			v >>= 7
			i++
		}
		buf[i] = byte(v)
		i++
	}
	return i
}

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *Uint16) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		}
		shift := 0
		done := false
		for l, b := range buf[i:] {
			if l == 2 && b > 3 {
				return i, errs.ErrOverflow
			}
			if b < 0x80 {
				(*v) = (*v) | Uint16(b)<<shift
				done = true
				i += l + 1
				break
			}
			(*v) = (*v) | Uint16(b&0x7F)<<shift
			shift += 7
		}
		if !done {
			return i, errs.ErrSmallBuf
		}
	}
	return i, err
}

// SizeMUS returns the size of the MUS-encoded v.
func (v Uint16) SizeMUS() int {
	size := 0
	{
		for v >= 0x80 {
			v >>= 7
			size++
		}
		size++
	}
	return size
}