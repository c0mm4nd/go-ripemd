package ripemd

const (
	// The size of the checksum in bytes.
	Ripemd256Size = 32

	// The block size of the hash algorithm in bytes.
	Ripemd256BlockSize = 64
)

type Ripemd256 struct {
	s  [8]uint32                // running context
	x  [Ripemd256BlockSize]byte // temporary buffer
	nx int                      // index into buffer
	tc uint64                   // total count of bytes processed
}

func New256() *Ripemd256 {
	r := new(Ripemd256)
	r.Reset()
	return r
}

func (r *Ripemd256) Reset() {
	r.s[0], r.s[1], r.s[2], r.s[3], r.s[4], r.s[5], r.s[6], r.s[7] = _s0, _s1, _s2, _s3, _s4, _s5, _s6, _s7
	r.nx = 0
	r.tc = 0
}

func (r *Ripemd256) Size() int {
	return Ripemd256Size
}

func (r *Ripemd256) BlockSize() int {
	return Ripemd256BlockSize
}
