package ripemd

const (
	// The size of the checksum in bytes.
	Ripemd128Size = 16

	// The block size of the hash algorithm in bytes.
	Ripemd128BlockSize = 32
)

type Ripemd128 struct {
	s  [4]uint32                // running context
	x  [Ripemd128BlockSize]byte // temporary buffer
	nx int                      // index into buffer
	tc uint64                   // total count of bytes processed
}

func New128() *Ripemd128 {
	r := new(Ripemd128)
	r.Reset()
	return r
}

func (r *Ripemd128) Reset() {
	r.s[0], r.s[1], r.s[2], r.s[3] = _s0, _s1, _s2, _s3
	r.nx = 0
	r.tc = 0
}

func (r *Ripemd128) Size() int {
	return Ripemd128Size
}

func (r *Ripemd128) BlockSize() int {
	return Ripemd128BlockSize
}

// func (r *Ripemd128) Sum(in []byte) []byte {
// 	// Make a copy of r so that caller can keep writing and summing.
// 	d := *r

// 	// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.
// 	processed_bytes := d.processed_bytes
// 	var tmp [32]byte
// 	tmp[0] = 0x80
// 	if processed_bytes%32 < 32 {
// 		d.Write(tmp[0 : 56-processed_bytes%64])
// 	} else {
// 		d.Write(tmp[0 : 64+56-processed_bytes%64])
// 	}

// 	// Length in bits.
// 	processed_bytes <<= 3
// 	for i := uint(0); i < 8; i++ {
// 		tmp[i] = byte(processed_bytes >> (8 * i))
// 	}
// 	d.Write(tmp[0:8])

// 	if d.nx != 0 {
// 		panic("d.nx != 0")
// 	}

// 	var digest [Size]byte
// 	for i, state := range d.state {
// 		digest[i*4] = byte(state)
// 		digest[i*4+1] = byte(state >> 8)
// 		digest[i*4+2] = byte(state >> 16)
// 		digest[i*4+3] = byte(state >> 24)
// 	}

// 	return append(in, digest[:]...)
// }
