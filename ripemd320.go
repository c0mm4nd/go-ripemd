package ripemd

const (
	// The size of the checksum in bytes.
	Ripemd320Size = 40

	// The block size of the hash algorithm in bytes.
	Ripemd320BlockSize = 128
)

type Ripemd320 struct {
	s  [10]uint32               // running context
	x  [Ripemd320BlockSize]byte // temporary buffer
	nx int                      // index into buffer
	tc uint64                   // total count of bytes processed
}

func New320() *Ripemd320 {
	r := new(Ripemd320)
	r.Reset()
	return r
}

func (r *Ripemd320) Reset() {
	r.s[0], r.s[1], r.s[2], r.s[3], r.s[4], r.s[5], r.s[6], r.s[7], r.s[8], r.s[9] = _s0, _s1, _s2, _s3, _s4, _s5, _s6, _s7, _s8, _s9
	r.nx = 0
	r.tc = 0
}

func (r *Ripemd320) Size() int {
	return Ripemd320Size
}

func (r *Ripemd320) BlockSize() int {
	return Ripemd320BlockSize
}
