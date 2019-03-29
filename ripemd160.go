package ripemd

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	Ripemd160BufferSize = 64 * 1024 // int32 64kb
	// The size of the checksum in bytes.
	Ripemd160Size = 20

	// The block size of the hash algorithm in bytes.
	Ripemd160BlockSize = 64 // a_block_size
	Ripemd160HashSize  = 5  // a_hash_size

	C1 = uint32(0x50A28BE6)
	C2 = uint32(0x5A827999)
	C3 = uint32(0x5C4DD124)
	C4 = uint32(0x6ED9EBA1)
	C5 = uint32(0x6D703EF3)
	C6 = uint32(0x8F1BBCDC)
	C7 = uint32(0x7A6D76E9)
	C8 = uint32(0xA953FD4E)
)

type Ripemd160 struct {
	state [5]uint32 // running context
	//buffer          [Ripemd160BlockSize]byte // tempary buffer
	//nx              int                      // index into buffer
	buffer *bytes.Buffer

	processed_bytes uint64 // total count of bytes processed
}

func New160() *Ripemd160 {
	// Create
	r := new(Ripemd160)
	r.Initialize()
	return r
}

func (r *Ripemd160) Initialize() {
	r.state[0], r.state[1], r.state[2], r.state[3], r.state[4] = _s0, _s1, _s2, _s3, _s4
	//r.nx = 0
	r.processed_bytes = 0
}

func (r *Ripemd160) Size() int {
	return Ripemd160Size
}

func (r *Ripemd160) BlockSize() int {
	return Ripemd160BlockSize
}

// BYTES_TO_DWORD
func bytes2uint64(b []byte) uint64 {
	result, _ := binary.Uvarint(b)
	return result
}

func byte2uint64(b byte) uint64 {
	result, _ := binary.Uvarint([]byte{b})
	return result
}

func MDinit() []uint64 {
	var MDbuf []uint64
	MDbuf = append(MDbuf, 0x67452301)
	MDbuf = append(MDbuf, 0xefcdab89)
	MDbuf = append(MDbuf, 0x98badcfe)
	MDbuf = append(MDbuf, 0x10325476)
	MDbuf = append(MDbuf, 0xc3d2e1f0)
	return MDbuf
}

func FF(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += ((b) ^ (c) ^ (d)) + (x)
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func GG(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += (((b) & (c)) | (^b & (d))) + (x) + 0x5a827999
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func HH(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	(a) += (((b) | (^c)) ^ (d)) + (x) + 0x6ed9eba1
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func II(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += (((b) & (d)) | ((c) & (^d))) + (x) + 0x8f1bbcdc
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func JJ(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += ((b) ^ ((c) | (^d))) + (x) + 0xa953fd4e
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func FFF(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	return FF(a, b, c, d, e, x, s)
}

func GGG(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += (((b) & (c)) | (^b & (d))) + (x) + 0x7a6d76e9
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func HHH(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	(a) += (((b) | (^c)) ^ (d)) + (x) + 0x6d703ef3
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func III(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += (((b) & (d)) | ((c) & (^d))) + (x) + 0x5c4dd124
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func JJJ(a, b, c, d, e, x, s uint64) (aa, bb, cc, dd, ee, xx, ss uint64) {
	a += ((b) ^ ((c) | (^d))) + (x) + 0x50a28be6
	a = (((a) << (s)) | ((a) >> (32 - (s)))) + (e)
	c = ((c) << (10)) | ((c) >> (32 - (10)))
	return a, b, c, d, e, x, s
}

func Rounds(MDbuf []uint64, X [16]uint64) []uint64 {
	aa := MDbuf[0]
	bb := MDbuf[1]
	cc := MDbuf[2]
	dd := MDbuf[3]
	ee := MDbuf[4]

	aaa := MDbuf[0]
	bbb := MDbuf[1]
	ccc := MDbuf[2]
	ddd := MDbuf[3]
	eee := MDbuf[4]

	/* round 1 */
	aa, bb, cc, dd, ee, X[0], _ = FF(aa, bb, cc, dd, ee, X[0], 11)
	ee, aa, bb, cc, dd, X[1], _ = FF(ee, aa, bb, cc, dd, X[1], 14)
	dd, ee, aa, bb, cc, X[2], _ = FF(dd, ee, aa, bb, cc, X[2], 15)
	cc, dd, ee, aa, bb, X[3], _ = FF(cc, dd, ee, aa, bb, X[3], 12)
	bb, cc, dd, ee, aa, X[4], _ = FF(bb, cc, dd, ee, aa, X[4], 5)
	aa, bb, cc, dd, ee, X[5], _ = FF(aa, bb, cc, dd, ee, X[5], 8)
	ee, aa, bb, cc, dd, X[6], _ = FF(ee, aa, bb, cc, dd, X[6], 7)
	dd, ee, aa, bb, cc, X[7], _ = FF(dd, ee, aa, bb, cc, X[7], 9)
	cc, dd, ee, aa, bb, X[8], _ = FF(cc, dd, ee, aa, bb, X[8], 11)
	bb, cc, dd, ee, aa, X[9], _ = FF(bb, cc, dd, ee, aa, X[9], 13)
	aa, bb, cc, dd, ee, X[10], _ = FF(aa, bb, cc, dd, ee, X[10], 14)
	ee, aa, bb, cc, dd, X[11], _ = FF(ee, aa, bb, cc, dd, X[11], 15)
	dd, ee, aa, bb, cc, X[12], _ = FF(dd, ee, aa, bb, cc, X[12], 6)
	cc, dd, ee, aa, bb, X[13], _ = FF(cc, dd, ee, aa, bb, X[13], 7)
	bb, cc, dd, ee, aa, X[14], _ = FF(bb, cc, dd, ee, aa, X[14], 9)
	aa, bb, cc, dd, ee, X[15], _ = FF(aa, bb, cc, dd, ee, X[15], 8)

	/* round 2 */
	ee, aa, bb, cc, dd, X[7], _ = GG(ee, aa, bb, cc, dd, X[7], 7)
	dd, ee, aa, bb, cc, X[4], _ = GG(dd, ee, aa, bb, cc, X[4], 6)
	cc, dd, ee, aa, bb, X[13], _ = GG(cc, dd, ee, aa, bb, X[13], 8)
	bb, cc, dd, ee, aa, X[1], _ = GG(bb, cc, dd, ee, aa, X[1], 13)
	aa, bb, cc, dd, ee, X[10], _ = GG(aa, bb, cc, dd, ee, X[10], 11)
	ee, aa, bb, cc, dd, X[6], _ = GG(ee, aa, bb, cc, dd, X[6], 9)
	dd, ee, aa, bb, cc, X[15], _ = GG(dd, ee, aa, bb, cc, X[15], 7)
	cc, dd, ee, aa, bb, X[3], _ = GG(cc, dd, ee, aa, bb, X[3], 15)
	bb, cc, dd, ee, aa, X[12], _ = GG(bb, cc, dd, ee, aa, X[12], 7)
	aa, bb, cc, dd, ee, X[0], _ = GG(aa, bb, cc, dd, ee, X[0], 12)
	ee, aa, bb, cc, dd, X[9], _ = GG(ee, aa, bb, cc, dd, X[9], 15)
	dd, ee, aa, bb, cc, X[5], _ = GG(dd, ee, aa, bb, cc, X[5], 9)
	cc, dd, ee, aa, bb, X[2], _ = GG(cc, dd, ee, aa, bb, X[2], 11)
	bb, cc, dd, ee, aa, X[14], _ = GG(bb, cc, dd, ee, aa, X[14], 7)
	aa, bb, cc, dd, ee, X[11], _ = GG(aa, bb, cc, dd, ee, X[11], 13)
	ee, aa, bb, cc, dd, X[8], _ = GG(ee, aa, bb, cc, dd, X[8], 12)

	/* round 3 */
	dd, ee, aa, bb, cc, X[3], _ = HH(dd, ee, aa, bb, cc, X[3], 11)
	cc, dd, ee, aa, bb, X[10], _ = HH(cc, dd, ee, aa, bb, X[10], 13)
	bb, cc, dd, ee, aa, X[14], _ = HH(bb, cc, dd, ee, aa, X[14], 6)
	aa, bb, cc, dd, ee, X[4], _ = HH(aa, bb, cc, dd, ee, X[4], 7)
	ee, aa, bb, cc, dd, X[9], _ = HH(ee, aa, bb, cc, dd, X[9], 14)
	dd, ee, aa, bb, cc, X[15], _ = HH(dd, ee, aa, bb, cc, X[15], 9)
	cc, dd, ee, aa, bb, X[8], _ = HH(cc, dd, ee, aa, bb, X[8], 13)
	bb, cc, dd, ee, aa, X[1], _ = HH(bb, cc, dd, ee, aa, X[1], 15)
	aa, bb, cc, dd, ee, X[2], _ = HH(aa, bb, cc, dd, ee, X[2], 14)
	ee, aa, bb, cc, dd, X[7], _ = HH(ee, aa, bb, cc, dd, X[7], 8)
	dd, ee, aa, bb, cc, X[0], _ = HH(dd, ee, aa, bb, cc, X[0], 13)
	cc, dd, ee, aa, bb, X[6], _ = HH(cc, dd, ee, aa, bb, X[6], 6)
	bb, cc, dd, ee, aa, X[13], _ = HH(bb, cc, dd, ee, aa, X[13], 5)
	aa, bb, cc, dd, ee, X[11], _ = HH(aa, bb, cc, dd, ee, X[11], 12)
	ee, aa, bb, cc, dd, X[5], _ = HH(ee, aa, bb, cc, dd, X[5], 7)
	dd, ee, aa, bb, cc, X[12], _ = HH(dd, ee, aa, bb, cc, X[12], 5)

	/* round 4 */
	cc, dd, ee, aa, bb, X[1], _ = II(cc, dd, ee, aa, bb, X[1], 11)
	bb, cc, dd, ee, aa, X[9], _ = II(bb, cc, dd, ee, aa, X[9], 12)
	aa, bb, cc, dd, ee, X[11], _ = II(aa, bb, cc, dd, ee, X[11], 14)
	ee, aa, bb, cc, dd, X[10], _ = II(ee, aa, bb, cc, dd, X[10], 15)
	dd, ee, aa, bb, cc, X[0], _ = II(dd, ee, aa, bb, cc, X[0], 14)
	cc, dd, ee, aa, bb, X[8], _ = II(cc, dd, ee, aa, bb, X[8], 15)
	bb, cc, dd, ee, aa, X[12], _ = II(bb, cc, dd, ee, aa, X[12], 9)
	aa, bb, cc, dd, ee, X[4], _ = II(aa, bb, cc, dd, ee, X[4], 8)
	ee, aa, bb, cc, dd, X[13], _ = II(ee, aa, bb, cc, dd, X[13], 9)
	dd, ee, aa, bb, cc, X[3], _ = II(dd, ee, aa, bb, cc, X[3], 14)
	cc, dd, ee, aa, bb, X[7], _ = II(cc, dd, ee, aa, bb, X[7], 5)
	bb, cc, dd, ee, aa, X[15], _ = II(bb, cc, dd, ee, aa, X[15], 6)
	aa, bb, cc, dd, ee, X[14], _ = II(aa, bb, cc, dd, ee, X[14], 8)
	ee, aa, bb, cc, dd, X[5], _ = II(ee, aa, bb, cc, dd, X[5], 6)
	dd, ee, aa, bb, cc, X[6], _ = II(dd, ee, aa, bb, cc, X[6], 5)
	cc, dd, ee, aa, bb, X[2], _ = II(cc, dd, ee, aa, bb, X[2], 12)

	/* round 5 */
	bb, cc, dd, ee, aa, X[4], _ = JJ(bb, cc, dd, ee, aa, X[4], 9)
	aa, bb, cc, dd, ee, X[0], _ = JJ(aa, bb, cc, dd, ee, X[0], 15)
	ee, aa, bb, cc, dd, X[5], _ = JJ(ee, aa, bb, cc, dd, X[5], 5)
	dd, ee, aa, bb, cc, X[9], _ = JJ(dd, ee, aa, bb, cc, X[9], 11)
	cc, dd, ee, aa, bb, X[7], _ = JJ(cc, dd, ee, aa, bb, X[7], 6)
	bb, cc, dd, ee, aa, X[12], _ = JJ(bb, cc, dd, ee, aa, X[12], 8)
	aa, bb, cc, dd, ee, X[2], _ = JJ(aa, bb, cc, dd, ee, X[2], 13)
	ee, aa, bb, cc, dd, X[10], _ = JJ(ee, aa, bb, cc, dd, X[10], 12)
	dd, ee, aa, bb, cc, X[14], _ = JJ(dd, ee, aa, bb, cc, X[14], 5)
	cc, dd, ee, aa, bb, X[1], _ = JJ(cc, dd, ee, aa, bb, X[1], 12)
	bb, cc, dd, ee, aa, X[3], _ = JJ(bb, cc, dd, ee, aa, X[3], 13)
	aa, bb, cc, dd, ee, X[8], _ = JJ(aa, bb, cc, dd, ee, X[8], 14)
	ee, aa, bb, cc, dd, X[11], _ = JJ(ee, aa, bb, cc, dd, X[11], 11)
	dd, ee, aa, bb, cc, X[6], _ = JJ(dd, ee, aa, bb, cc, X[6], 8)
	cc, dd, ee, aa, bb, X[15], _ = JJ(cc, dd, ee, aa, bb, X[15], 5)
	bb, cc, dd, ee, aa, X[13], _ = JJ(bb, cc, dd, ee, aa, X[13], 6)

	/* parallel round 1 */
	aaa, bbb, ccc, ddd, eee, X[5], _ = JJJ(aaa, bbb, ccc, ddd, eee, X[5], 8)
	eee, aaa, bbb, ccc, ddd, X[14], _ = JJJ(eee, aaa, bbb, ccc, ddd, X[14], 9)
	ddd, eee, aaa, bbb, ccc, X[7], _ = JJJ(ddd, eee, aaa, bbb, ccc, X[7], 9)
	ccc, ddd, eee, aaa, bbb, X[0], _ = JJJ(ccc, ddd, eee, aaa, bbb, X[0], 11)
	bbb, ccc, ddd, eee, aaa, X[9], _ = JJJ(bbb, ccc, ddd, eee, aaa, X[9], 13)
	aaa, bbb, ccc, ddd, eee, X[2], _ = JJJ(aaa, bbb, ccc, ddd, eee, X[2], 15)
	eee, aaa, bbb, ccc, ddd, X[11], _ = JJJ(eee, aaa, bbb, ccc, ddd, X[11], 15)
	ddd, eee, aaa, bbb, ccc, X[4], _ = JJJ(ddd, eee, aaa, bbb, ccc, X[4], 5)
	ccc, ddd, eee, aaa, bbb, X[13], _ = JJJ(ccc, ddd, eee, aaa, bbb, X[13], 7)
	bbb, ccc, ddd, eee, aaa, X[6], _ = JJJ(bbb, ccc, ddd, eee, aaa, X[6], 7)
	aaa, bbb, ccc, ddd, eee, X[15], _ = JJJ(aaa, bbb, ccc, ddd, eee, X[15], 8)
	eee, aaa, bbb, ccc, ddd, X[8], _ = JJJ(eee, aaa, bbb, ccc, ddd, X[8], 11)
	ddd, eee, aaa, bbb, ccc, X[1], _ = JJJ(ddd, eee, aaa, bbb, ccc, X[1], 14)
	ccc, ddd, eee, aaa, bbb, X[10], _ = JJJ(ccc, ddd, eee, aaa, bbb, X[10], 14)
	bbb, ccc, ddd, eee, aaa, X[3], _ = JJJ(bbb, ccc, ddd, eee, aaa, X[3], 12)
	aaa, bbb, ccc, ddd, eee, X[12], _ = JJJ(aaa, bbb, ccc, ddd, eee, X[12], 6)

	/* parallel round 2 */
	eee, aaa, bbb, ccc, ddd, X[6], _ = III(eee, aaa, bbb, ccc, ddd, X[6], 9)
	ddd, eee, aaa, bbb, ccc, X[11], _ = III(ddd, eee, aaa, bbb, ccc, X[11], 13)
	ccc, ddd, eee, aaa, bbb, X[3], _ = III(ccc, ddd, eee, aaa, bbb, X[3], 15)
	bbb, ccc, ddd, eee, aaa, X[7], _ = III(bbb, ccc, ddd, eee, aaa, X[7], 7)
	aaa, bbb, ccc, ddd, eee, X[0], _ = III(aaa, bbb, ccc, ddd, eee, X[0], 12)
	eee, aaa, bbb, ccc, ddd, X[13], _ = III(eee, aaa, bbb, ccc, ddd, X[13], 8)
	ddd, eee, aaa, bbb, ccc, X[5], _ = III(ddd, eee, aaa, bbb, ccc, X[5], 9)
	ccc, ddd, eee, aaa, bbb, X[10], _ = III(ccc, ddd, eee, aaa, bbb, X[10], 11)
	bbb, ccc, ddd, eee, aaa, X[14], _ = III(bbb, ccc, ddd, eee, aaa, X[14], 7)
	aaa, bbb, ccc, ddd, eee, X[15], _ = III(aaa, bbb, ccc, ddd, eee, X[15], 7)
	eee, aaa, bbb, ccc, ddd, X[8], _ = III(eee, aaa, bbb, ccc, ddd, X[8], 12)
	ddd, eee, aaa, bbb, ccc, X[12], _ = III(ddd, eee, aaa, bbb, ccc, X[12], 7)
	ccc, ddd, eee, aaa, bbb, X[4], _ = III(ccc, ddd, eee, aaa, bbb, X[4], 6)
	bbb, ccc, ddd, eee, aaa, X[9], _ = III(bbb, ccc, ddd, eee, aaa, X[9], 15)
	aaa, bbb, ccc, ddd, eee, X[1], _ = III(aaa, bbb, ccc, ddd, eee, X[1], 13)
	eee, aaa, bbb, ccc, ddd, X[2], _ = III(eee, aaa, bbb, ccc, ddd, X[2], 11)

	/* parallel round 3 */
	ddd, eee, aaa, bbb, ccc, X[15], _ = HHH(ddd, eee, aaa, bbb, ccc, X[15], 9)
	ccc, ddd, eee, aaa, bbb, X[5], _ = HHH(ccc, ddd, eee, aaa, bbb, X[5], 7)
	bbb, ccc, ddd, eee, aaa, X[1], _ = HHH(bbb, ccc, ddd, eee, aaa, X[1], 15)
	aaa, bbb, ccc, ddd, eee, X[3], _ = HHH(aaa, bbb, ccc, ddd, eee, X[3], 11)
	eee, aaa, bbb, ccc, ddd, X[7], _ = HHH(eee, aaa, bbb, ccc, ddd, X[7], 8)
	ddd, eee, aaa, bbb, ccc, X[14], _ = HHH(ddd, eee, aaa, bbb, ccc, X[14], 6)
	ccc, ddd, eee, aaa, bbb, X[6], _ = HHH(ccc, ddd, eee, aaa, bbb, X[6], 6)
	bbb, ccc, ddd, eee, aaa, X[9], _ = HHH(bbb, ccc, ddd, eee, aaa, X[9], 14)
	aaa, bbb, ccc, ddd, eee, X[11], _ = HHH(aaa, bbb, ccc, ddd, eee, X[11], 12)
	eee, aaa, bbb, ccc, ddd, X[8], _ = HHH(eee, aaa, bbb, ccc, ddd, X[8], 13)
	ddd, eee, aaa, bbb, ccc, X[12], _ = HHH(ddd, eee, aaa, bbb, ccc, X[12], 5)
	ccc, ddd, eee, aaa, bbb, X[2], _ = HHH(ccc, ddd, eee, aaa, bbb, X[2], 14)
	bbb, ccc, ddd, eee, aaa, X[10], _ = HHH(bbb, ccc, ddd, eee, aaa, X[10], 13)
	aaa, bbb, ccc, ddd, eee, X[0], _ = HHH(aaa, bbb, ccc, ddd, eee, X[0], 13)
	eee, aaa, bbb, ccc, ddd, X[4], _ = HHH(eee, aaa, bbb, ccc, ddd, X[4], 7)
	ddd, eee, aaa, bbb, ccc, X[13], _ = HHH(ddd, eee, aaa, bbb, ccc, X[13], 5)

	/* parallel round 4 */
	ccc, ddd, eee, aaa, bbb, X[8], _ = GGG(ccc, ddd, eee, aaa, bbb, X[8], 15)
	bbb, ccc, ddd, eee, aaa, X[6], _ = GGG(bbb, ccc, ddd, eee, aaa, X[6], 5)
	aaa, bbb, ccc, ddd, eee, X[4], _ = GGG(aaa, bbb, ccc, ddd, eee, X[4], 8)
	eee, aaa, bbb, ccc, ddd, X[1], _ = GGG(eee, aaa, bbb, ccc, ddd, X[1], 11)
	ddd, eee, aaa, bbb, ccc, X[3], _ = GGG(ddd, eee, aaa, bbb, ccc, X[3], 14)
	ccc, ddd, eee, aaa, bbb, X[11], _ = GGG(ccc, ddd, eee, aaa, bbb, X[11], 14)
	bbb, ccc, ddd, eee, aaa, X[15], _ = GGG(bbb, ccc, ddd, eee, aaa, X[15], 6)
	aaa, bbb, ccc, ddd, eee, X[0], _ = GGG(aaa, bbb, ccc, ddd, eee, X[0], 14)
	eee, aaa, bbb, ccc, ddd, X[5], _ = GGG(eee, aaa, bbb, ccc, ddd, X[5], 6)
	ddd, eee, aaa, bbb, ccc, X[12], _ = GGG(ddd, eee, aaa, bbb, ccc, X[12], 9)
	ccc, ddd, eee, aaa, bbb, X[2], _ = GGG(ccc, ddd, eee, aaa, bbb, X[2], 12)
	bbb, ccc, ddd, eee, aaa, X[13], _ = GGG(bbb, ccc, ddd, eee, aaa, X[13], 9)
	aaa, bbb, ccc, ddd, eee, X[9], _ = GGG(aaa, bbb, ccc, ddd, eee, X[9], 12)
	eee, aaa, bbb, ccc, ddd, X[7], _ = GGG(eee, aaa, bbb, ccc, ddd, X[7], 5)
	ddd, eee, aaa, bbb, ccc, X[10], _ = GGG(ddd, eee, aaa, bbb, ccc, X[10], 15)
	ccc, ddd, eee, aaa, bbb, X[14], _ = GGG(ccc, ddd, eee, aaa, bbb, X[14], 8)

	/* parallel round 5 */
	bbb, ccc, ddd, eee, aaa, X[12], _ = FFF(bbb, ccc, ddd, eee, aaa, X[12], 8)
	aaa, bbb, ccc, ddd, eee, X[15], _ = FFF(aaa, bbb, ccc, ddd, eee, X[15], 5)
	eee, aaa, bbb, ccc, ddd, X[10], _ = FFF(eee, aaa, bbb, ccc, ddd, X[10], 12)
	ddd, eee, aaa, bbb, ccc, X[4], _ = FFF(ddd, eee, aaa, bbb, ccc, X[4], 9)
	ccc, ddd, eee, aaa, bbb, X[1], _ = FFF(ccc, ddd, eee, aaa, bbb, X[1], 12)
	bbb, ccc, ddd, eee, aaa, X[5], _ = FFF(bbb, ccc, ddd, eee, aaa, X[5], 5)
	aaa, bbb, ccc, ddd, eee, X[8], _ = FFF(aaa, bbb, ccc, ddd, eee, X[8], 14)
	eee, aaa, bbb, ccc, ddd, X[7], _ = FFF(eee, aaa, bbb, ccc, ddd, X[7], 6)
	ddd, eee, aaa, bbb, ccc, X[6], _ = FFF(ddd, eee, aaa, bbb, ccc, X[6], 8)
	ccc, ddd, eee, aaa, bbb, X[2], _ = FFF(ccc, ddd, eee, aaa, bbb, X[2], 13)
	bbb, ccc, ddd, eee, aaa, X[13], _ = FFF(bbb, ccc, ddd, eee, aaa, X[13], 6)
	aaa, bbb, ccc, ddd, eee, X[14], _ = FFF(aaa, bbb, ccc, ddd, eee, X[14], 5)
	eee, aaa, bbb, ccc, ddd, X[0], _ = FFF(eee, aaa, bbb, ccc, ddd, X[0], 15)
	ddd, eee, aaa, bbb, ccc, X[3], _ = FFF(ddd, eee, aaa, bbb, ccc, X[3], 13)
	ccc, ddd, eee, aaa, bbb, X[9], _ = FFF(ccc, ddd, eee, aaa, bbb, X[9], 11)
	bbb, ccc, ddd, eee, aaa, X[11], _ = FFF(bbb, ccc, ddd, eee, aaa, X[11], 11)

	/* combine results */
	ddd += cc + MDbuf[1] /* final result for MDbuf[0] */
	MDbuf[1] = MDbuf[2] + dd + eee
	MDbuf[2] = MDbuf[3] + ee + aaa
	MDbuf[3] = MDbuf[4] + aa + bbb
	MDbuf[4] = MDbuf[0] + bb + ccc
	MDbuf[0] = ddd

	return MDbuf
}

func MDfinish(MDbuf []uint64, strptr *bytes.Buffer, lswlen, mswlen uint64) []byte {
	//var i uint                                 /* counter       */
	var X [16]uint64 /* message words */
	for i := uint64(0); i < (lswlen & 63); i++ {
		/* byte i goes into word X[i div 4] at pos.  8*(i mod 4)  */
		b, err := strptr.ReadByte()
		if err != nil {
			fmt.Println(err)
			break
		}

		X[i>>2] ^= byte2uint64(b) << (8 * (i & 3))
	}

	/* append the bit m_n == 1 */
	X[(lswlen>>2)&15] ^= 1 << (8*(lswlen&3) + 7)

	if (lswlen & 63) > 55 {
		/* length goes to next block */
		Rounds(MDbuf, X)
	}

	/* append length in bits*/
	X[14] = lswlen << 3
	X[15] = (lswlen >> 29) | (mswlen << 3)
	r := Rounds(MDbuf, X)

	data := make([]byte, 0)

	for _, i := range r {
		add64 := make([]byte, 8)
		fmt.Println(i)
		binary.LittleEndian.PutUint64(add64, i)
		data = bytes.Join([][]byte{data, add64}, nil)
	}
	return data
}

//func (r *Ripemd160) TransformBlock(in []byte) []byte {

//}

//func (r *Ripemd160) ComputeBytes(in []byte) []byte {
//	r.Initialize() // Initialize();
//	r.TransformBytes(in)
//	// result := THashResult.Create(tempresult);
//	//r.TransformBlock(in)
//
//	var result =  []byte{}
//	addition:= make([]byte, 8)
//	//var addition []byte
//
//	for i:=0; i<len(r.state); i++ {
//		fmt.Println(r.state[i])
//		binary.LittleEndian.PutUint32(addition, r.state[i])
//		result = bytes.Join([][]byte{result, addition}, nil)
//	}
//
//	return result
//}
