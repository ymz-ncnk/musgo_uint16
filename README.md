# MusGo_uint16
Provides serialization of the Golang's `uint16` type.

# How to use
Simply cast `uint16` to `musgo_uint16.Uint16`:
```go
	var n uint16 = 5
	// Marshal
	buf := make([]byte, musgo_uint16.Uint16(n).SizeMUS())
	musgo_uint16.Uint16(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_uint16.Uint16)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

