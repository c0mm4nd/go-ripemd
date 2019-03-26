package ripemd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/bamiaux/iobit"
	"math/bits"
)

const (
	Ripemd160BufferSize = 64 * 1024 // int32 64kb
	// The size of the checksum in bytes.
	Ripemd160Size = 20

	// The block size of the hash algorithm in bytes.
	Ripemd160BlockSize = 64 // a_block_size
	Ripemd160HashSize = 5 // a_hash_size

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
	state           [5]uint32                // running context
	//buffer          [Ripemd160BlockSize]byte // tempary buffer
	//nx              int                      // index into buffer
	buffer   *bytes.Buffer

	processed_bytes uint64                   // total count of bytes processed
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

func (r *Ripemd160) TransformBlock(in []byte) []byte {
	var a, b, c, d, e uint32
	var data []uint32
	a = r.state[0]
	b = r.state[1]
	c = r.state[2]
	d = r.state[3]
	e = r.state[4]
	aa := a
	bb := b
	cc := c
	dd := d
	ee := e

	//length := int(len(in) / 4)
	//if len(in) % 4 != 0{
	//	length ++
	//}
	//var data [16]uint32
	//var indata [4*16]byte
	//copy(indata, in)

	//for i:=0; i<16; i=i+1{
	//	if len(in) > 4 && 4*i+4 < 64 {
	//		//fmt.Println(in[4*i: 4*i+3])
	//		key := binary.LittleEndian.Uint32(in[:4])
	//		data[i] = key
	//	}else {
	//		data[i] = binary.LittleEndian.Uint32(in[:])
	//	}
	//}

	//dataWritter := iobit.NewWriter(data)

	inReader := iobit.NewReader(in)
	data = make([]uint32, 0)
	for i:=0; inReader.LeftBits() != 0; i++ {
		 data = append(data, inReader.Le32())
	}
	for len(data) < Ripemd160Size {
		data = append(data, 0x0)
	}

	a = a + (data[0] + (b ^ c ^ d))
	a = bits.RotateLeft32(a, 11) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[1] + (a ^ b ^ c))
	e = bits.RotateLeft32(e, 14) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[2] + (e ^ a ^ b))
	d = bits.RotateLeft32(d, 15) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[3] + (d ^ e ^ a))
	c = bits.RotateLeft32(c, 12) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[4] + (c ^ d ^ e))
	b = bits.RotateLeft32(b, 5) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[5] + (b ^ c ^ d))
	a = bits.RotateLeft32(a, 8) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[6] + (a ^ b ^ c))
	e = bits.RotateLeft32(e, 7) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[7] + (e ^ a ^ b))
	d = bits.RotateLeft32(d, 9) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[8] + (d ^ e ^ a))
	c = bits.RotateLeft32(c, 11) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[9] + (c ^ d ^ e))
	b = bits.RotateLeft32(b, 13) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[10] + (b ^ c ^ d))
	a = bits.RotateLeft32(a, 14) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[11] + (a ^ b ^ c))
	e = bits.RotateLeft32(e, 15) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[12] + (e ^ a ^ b))
	d = bits.RotateLeft32(d, 6) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[13] + (d ^ e ^ a))
	c = bits.RotateLeft32(c, 7) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[14] + (c ^ d ^ e))
	b = bits.RotateLeft32(b, 9) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[15] + (b ^ c ^ d))
	a = bits.RotateLeft32(a, 8) + e
	c = bits.RotateLeft32(c, 10)



	aa = aa + (data[5] + C1 + (bb ^ (cc | ^dd)))
	aa = bits.RotateLeft32(aa, 8) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[14] + C1 + (aa ^ (bb | ^cc)))
	ee = bits.RotateLeft32(ee, 9) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[7] + C1 + (ee ^ (aa | ^bb)))
	dd = bits.RotateLeft32(dd, 9) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[0] + C1 + (dd ^ (ee | ^aa)))
	cc = bits.RotateLeft32(cc, 11) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[9] + C1 + (cc ^ (dd | ^ee)))
	bb = bits.RotateLeft32(bb, 13) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[2] + C1 + (bb ^ (cc | ^dd)))
	aa = bits.RotateLeft32(aa, 15) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[11] + C1 + (aa ^ (bb | ^cc)))
	ee = bits.RotateLeft32(ee, 15) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[4] + C1 + (ee ^ (aa | ^bb)))
	dd = bits.RotateLeft32(dd, 5) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[13] + C1 + (dd ^ (ee | ^aa)))
	cc = bits.RotateLeft32(cc, 7) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[6] + C1 + (cc ^ (dd | ^ee)))
	bb = bits.RotateLeft32(bb, 7) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[15] + C1 + (bb ^ (cc | ^dd)))
	aa = bits.RotateLeft32(aa, 8) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[8] + C1 + (aa ^ (bb | ^cc)))
	ee = bits.RotateLeft32(ee, 11) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[1] + C1 + (ee ^ (aa | ^bb)))
	dd = bits.RotateLeft32(dd, 14) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[10] + C1 + (dd ^ (ee | ^aa)))
	cc = bits.RotateLeft32(cc, 14) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[3] + C1 + (cc ^ (dd | ^ee)))
	bb = bits.RotateLeft32(bb, 12) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[12] + C1 + (bb ^ (cc | ^dd)))
	aa = bits.RotateLeft32(aa, 6) + ee
	cc = bits.RotateLeft32(cc, 10)

	e = e + (data[7] + C2 + ((a & b) | (^a & c)))
	e = bits.RotateLeft32(e, 7) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[4] + C2 + ((e & a) | (^e & b)))
	d = bits.RotateLeft32(d, 6) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[13] + C2 + ((d & e) | (^d & a)))
	c = bits.RotateLeft32(c, 8) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[1] + C2 + ((c & d) | (^c & e)))
	b = bits.RotateLeft32(b, 13) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[10] + C2 + ((b & c) | (^b & d)))
	a = bits.RotateLeft32(a, 11) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[6] + C2 + ((a & b) | (^a & c)))
	e = bits.RotateLeft32(e, 9) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[15] + C2 + ((e & a) | (^e & b)))
	d = bits.RotateLeft32(d, 7) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[3] + C2 + ((d & e) | (^d & a)))
	c = bits.RotateLeft32(c, 15) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[12] + C2 + ((c & d) | (^c & e)))
	b = bits.RotateLeft32(b, 7) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[0] + C2 + ((b & c) | (^b & d)))
	a = bits.RotateLeft32(a, 12) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[9] + C2 + ((a & b) | (^a & c)))
	e = bits.RotateLeft32(e, 15) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[5] + C2 + ((e & a) | (^e & b)))
	d = bits.RotateLeft32(d, 9) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[2] + C2 + ((d & e) | (^d & a)))
	c = bits.RotateLeft32(c, 11) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[14] + C2 + ((c & d) | (^c & e)))
	b = bits.RotateLeft32(b, 7) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[11] + C2 + ((b & c) | (^b & d)))
	a = bits.RotateLeft32(a, 13) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[8] + C2 + ((a & b) | (^a & c)))
	e = bits.RotateLeft32(e, 12) + d
	b = bits.RotateLeft32(b, 10)

	ee = ee + (data[6] + C3 + ((aa & cc) | (bb & ^cc)))
	ee = bits.RotateLeft32(ee, 9) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[11] + C3 + ((ee & bb) | (aa & ^bb)))
	dd = bits.RotateLeft32(dd, 13) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[3] + C3 + ((dd & aa) | (ee & ^aa)))
	cc = bits.RotateLeft32(cc, 15) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[7] + C3 + ((cc & ee) | (dd & ^ee)))
	bb = bits.RotateLeft32(bb, 7) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[0] + C3 + ((bb & dd) | (cc & ^dd)))
	aa = bits.RotateLeft32(aa, 12) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[13] + C3 + ((aa & cc) | (bb & ^cc)))
	ee = bits.RotateLeft32(ee, 8) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[5] + C3 + ((ee & bb) | (aa & ^bb)))
	dd = bits.RotateLeft32(dd, 9) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[10] + C3 + ((dd & aa) | (ee & ^aa)))
	cc = bits.RotateLeft32(cc, 11) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[14] + C3 + ((cc & ee) | (dd & ^ee)))
	bb = bits.RotateLeft32(bb, 7) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[15] + C3 + ((bb & dd) | (cc & ^dd)))
	aa = bits.RotateLeft32(aa, 7) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[8] + C3 + ((aa & cc) | (bb & ^cc)))
	ee = bits.RotateLeft32(ee, 12) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[12] + C3 + ((ee & bb) | (aa & ^bb)))
	dd = bits.RotateLeft32(dd, 7) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[4] + C3 + ((dd & aa) | (ee & ^aa)))
	cc = bits.RotateLeft32(cc, 6) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[9] + C3 + ((cc & ee) | (dd & ^ee)))
	bb = bits.RotateLeft32(bb, 15) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[1] + C3 + ((bb & dd) | (cc & ^dd)))
	aa = bits.RotateLeft32(aa, 13) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[2] + C3 + ((aa & cc) | (bb & ^cc)))
	ee = bits.RotateLeft32(ee, 11) + dd
	bb = bits.RotateLeft32(bb, 10)

	d = d + (data[3] + C4 + ((e | ^a) ^ b))
	d = bits.RotateLeft32(d, 11) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[10] + C4 + ((d | ^e) ^ a))
	c = bits.RotateLeft32(c, 13) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[14] + C4 + ((c | ^d) ^ e))
	b = bits.RotateLeft32(b, 6) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[4] + C4 + ((b | ^c) ^ d))
	a = bits.RotateLeft32(a, 7) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[9] + C4 + ((a | ^b) ^ c))
	e = bits.RotateLeft32(e, 14) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[15] + C4 + ((e | ^a) ^ b))
	d = bits.RotateLeft32(d, 9) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[8] + C4 + ((d | ^e) ^ a))
	c = bits.RotateLeft32(c, 13) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[1] + C4 + ((c | ^d) ^ e))
	b = bits.RotateLeft32(b, 15) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[2] + C4 + ((b | ^c) ^ d))
	a = bits.RotateLeft32(a, 14) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[7] + C4 + ((a | ^b) ^ c))
	e = bits.RotateLeft32(e, 8) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[0] + C4 + ((e | ^a) ^ b))
	d = bits.RotateLeft32(d, 13) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[6] + C4 + ((d | ^e) ^ a))
	c = bits.RotateLeft32(c, 6) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[13] + C4 + ((c | ^d) ^ e))
	b = bits.RotateLeft32(b, 5) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[11] + C4 + ((b | ^c) ^ d))
	a = bits.RotateLeft32(a, 12) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[5] + C4 + ((a | ^b) ^ c))
	e = bits.RotateLeft32(e, 7) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[12] + C4 + ((e | ^a) ^ b))
	d = bits.RotateLeft32(d, 5) + c
	a = bits.RotateLeft32(a, 10)

	dd = dd + (data[15] + C5 + ((ee | ^aa) ^ bb))
	dd = bits.RotateLeft32(dd, 9) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[5] + C5 + ((dd | ^ee) ^ aa))
	cc = bits.RotateLeft32(cc, 7) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[1] + C5 + ((cc | ^dd) ^ ee))
	bb = bits.RotateLeft32(bb, 15) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[3] + C5 + ((bb | ^cc) ^ dd))
	aa = bits.RotateLeft32(aa, 11) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[7] + C5 + ((aa | ^bb) ^ cc))
	ee = bits.RotateLeft32(ee, 8) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[14] + C5 + ((ee | ^aa) ^ bb))
	dd = bits.RotateLeft32(dd, 6) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[6] + C5 + ((dd | ^ee) ^ aa))
	cc = bits.RotateLeft32(cc, 6) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[9] + C5 + ((cc | ^dd) ^ ee))
	bb = bits.RotateLeft32(bb, 14) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[11] + C5 + ((bb | ^cc) ^ dd))
	aa = bits.RotateLeft32(aa, 12) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[8] + C5 + ((aa | ^bb) ^ cc))
	ee = bits.RotateLeft32(ee, 13) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[12] + C5 + ((ee | ^aa) ^ bb))
	dd = bits.RotateLeft32(dd, 5) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[2] + C5 + ((dd | ^ee) ^ aa))
	cc = bits.RotateLeft32(cc, 14) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[10] + C5 + ((cc | ^dd) ^ ee))
	bb = bits.RotateLeft32(bb, 13) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[0] + C5 + ((bb | ^cc) ^ dd))
	aa = bits.RotateLeft32(aa, 13) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[4] + C5 + ((aa | ^bb) ^ cc))
	ee = bits.RotateLeft32(ee, 7) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[13] + C5 + ((ee | ^aa) ^ bb))
	dd = bits.RotateLeft32(dd, 5) + cc
	aa = bits.RotateLeft32(aa, 10)

	c = c + (data[1] + C6 + ((d & a) | (e & ^a)))
	c = bits.RotateLeft32(c, 11) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[9] + C6 + ((c & e) | (d & ^e)))
	b = bits.RotateLeft32(b, 12) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[11] + C6 + ((b & d) | (c & ^d)))
	a = bits.RotateLeft32(a, 14) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[10] + C6 + ((a & c) | (b & ^c)))
	e = bits.RotateLeft32(e, 15) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[0] + C6 + ((e & b) | (a & ^b)))
	d = bits.RotateLeft32(d, 14) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[8] + C6 + ((d & a) | (e & ^a)))
	c = bits.RotateLeft32(c, 15) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[12] + C6 + ((c & e) | (d & ^e)))
	b = bits.RotateLeft32(b, 9) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[4] + C6 + ((b & d) | (c & ^d)))
	a = bits.RotateLeft32(a, 8) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[13] + C6 + ((a & c) | (b & ^c)))
	e = bits.RotateLeft32(e, 9) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[3] + C6 + ((e & b) | (a & ^b)))
	d = bits.RotateLeft32(d, 14) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[7] + C6 + ((d & a) | (e & ^a)))
	c = bits.RotateLeft32(c, 5) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[15] + C6 + ((c & e) | (d & ^e)))
	b = bits.RotateLeft32(b, 6) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[14] + C6 + ((b & d) | (c & ^d)))
	a = bits.RotateLeft32(a, 8) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[5] + C6 + ((a & c) | (b & ^c)))
	e = bits.RotateLeft32(e, 6) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[6] + C6 + ((e & b) | (a & ^b)))
	d = bits.RotateLeft32(d, 5) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[2] + C6 + ((d & a) | (e & ^a)))
	c = bits.RotateLeft32(c, 12) + b
	e = bits.RotateLeft32(e, 10)

	cc = cc + (data[8] + C7 + ((dd & ee) | (^dd & aa)))
	cc = bits.RotateLeft32(cc, 15) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[6] + C7 + ((cc & dd) | (^cc & ee)))
	bb = bits.RotateLeft32(bb, 5) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[4] + C7 + ((bb & cc) | (^bb & dd)))
	aa = bits.RotateLeft32(aa, 8) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[1] + C7 + ((aa & bb) | (^aa & cc)))
	ee = bits.RotateLeft32(ee, 11) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[3] + C7 + ((ee & aa) | (^ee & bb)))
	dd = bits.RotateLeft32(dd, 14) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[11] + C7 + ((dd & ee) | (^dd & aa)))
	cc = bits.RotateLeft32(cc, 14) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[15] + C7 + ((cc & dd) | (^cc & ee)))
	bb = bits.RotateLeft32(bb, 6) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[0] + C7 + ((bb & cc) | (^bb & dd)))
	aa = bits.RotateLeft32(aa, 14) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[5] + C7 + ((aa & bb) | (^aa & cc)))
	ee = bits.RotateLeft32(ee, 6) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[12] + C7 + ((ee & aa) | (^ee & bb)))
	dd = bits.RotateLeft32(dd, 9) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[2] + C7 + ((dd & ee) | (^dd & aa)))
	cc = bits.RotateLeft32(cc, 12) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[13] + C7 + ((cc & dd) | (^cc & ee)))
	bb = bits.RotateLeft32(bb, 9) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[9] + C7 + ((bb & cc) | (^bb & dd)))
	aa = bits.RotateLeft32(aa, 12) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[7] + C7 + ((aa & bb) | (^aa & cc)))
	ee = bits.RotateLeft32(ee, 5) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[10] + C7 + ((ee & aa) | (^ee & bb)))
	dd = bits.RotateLeft32(dd, 15) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[14] + C7 + ((dd & ee) | (^dd & aa)))
	cc = bits.RotateLeft32(cc, 8) + bb
	ee = bits.RotateLeft32(ee, 10)

	b = b + (data[4] + C8 + (c ^ (d | ^e)))
	b = bits.RotateLeft32(b, 9) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[0] + C8 + (b ^ (c | ^d)))
	a = bits.RotateLeft32(a, 15) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[5] + C8 + (a ^ (b | ^c)))
	e = bits.RotateLeft32(e, 5) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[9] + C8 + (e ^ (a | ^b)))
	d = bits.RotateLeft32(d, 11) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[7] + C8 + (d ^ (e | ^a)))
	c = bits.RotateLeft32(c, 6) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[12] + C8 + (c ^ (d | ^e)))
	b = bits.RotateLeft32(b, 8) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[2] + C8 + (b ^ (c | ^d)))
	a = bits.RotateLeft32(a, 13) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[10] + C8 + (a ^ (b | ^c)))
	e = bits.RotateLeft32(e, 12) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[14] + C8 + (e ^ (a | ^b)))
	d = bits.RotateLeft32(d, 5) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[1] + C8 + (d ^ (e | ^a)))
	c = bits.RotateLeft32(c, 12) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[3] + C8 + (c ^ (d | ^e)))
	b = bits.RotateLeft32(b, 13) + a
	d = bits.RotateLeft32(d, 10)
	a = a + (data[8] + C8 + (b ^ (c | ^d)))
	a = bits.RotateLeft32(a, 14) + e
	c = bits.RotateLeft32(c, 10)
	e = e + (data[11] + C8 + (a ^ (b | ^c)))
	e = bits.RotateLeft32(e, 11) + d
	b = bits.RotateLeft32(b, 10)
	d = d + (data[6] + C8 + (e ^ (a | ^b)))
	d = bits.RotateLeft32(d, 8) + c
	a = bits.RotateLeft32(a, 10)
	c = c + (data[15] + C8 + (d ^ (e | ^a)))
	c = bits.RotateLeft32(c, 5) + b
	e = bits.RotateLeft32(e, 10)
	b = b + (data[13] + C8 + (c ^ (d | ^e)))
	b = bits.RotateLeft32(b, 6) + a
	d = bits.RotateLeft32(d, 10)

	bb = bb + (data[12] + (cc ^ dd ^ ee))
	bb = bits.RotateLeft32(bb, 8) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[15] + (bb ^ cc ^ dd))
	aa = bits.RotateLeft32(aa, 5) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[10] + (aa ^ bb ^ cc))
	ee = bits.RotateLeft32(ee, 12) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[4] + (ee ^ aa ^ bb))
	dd = bits.RotateLeft32(dd, 9) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[1] + (dd ^ ee ^ aa))
	cc = bits.RotateLeft32(cc, 12) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[5] + (cc ^ dd ^ ee))
	bb = bits.RotateLeft32(bb, 5) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[8] + (bb ^ cc ^ dd))
	aa = bits.RotateLeft32(aa, 14) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[7] + (aa ^ bb ^ cc))
	ee = bits.RotateLeft32(ee, 6) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[6] + (ee ^ aa ^ bb))
	dd = bits.RotateLeft32(dd, 8) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[2] + (dd ^ ee ^ aa))
	cc = bits.RotateLeft32(cc, 13) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[13] + (cc ^ dd ^ ee))
	bb = bits.RotateLeft32(bb, 6) + aa
	dd = bits.RotateLeft32(dd, 10)
	aa = aa + (data[14] + (bb ^ cc ^ dd))
	aa = bits.RotateLeft32(aa, 5) + ee
	cc = bits.RotateLeft32(cc, 10)
	ee = ee + (data[0] + (aa ^ bb ^ cc))
	ee = bits.RotateLeft32(ee, 15) + dd
	bb = bits.RotateLeft32(bb, 10)
	dd = dd + (data[3] + (ee ^ aa ^ bb))
	dd = bits.RotateLeft32(dd, 13) + cc
	aa = bits.RotateLeft32(aa, 10)
	cc = cc + (data[9] + (dd ^ ee ^ aa))
	cc = bits.RotateLeft32(cc, 11) + bb
	ee = bits.RotateLeft32(ee, 10)
	bb = bb + (data[11] + (cc ^ dd ^ ee))
	bb = bits.RotateLeft32(bb, 11) + aa
	dd = bits.RotateLeft32(dd, 10)

	dd = dd + c + r.state[1]
	r.state[1] = r.state[2] + d + ee
	r.state[2] = r.state[3] + e + aa
	r.state[3] = r.state[4] + a + bb
	r.state[4] = r.state[0] + b + cc
	r.state[0] = dd

	var dataBytes []byte
	dataBytesWritter := iobit.NewWriter(dataBytes)
	for i:=0; i<len(r.state); i++ {
		fmt.Println(r.state[i])
		dataBytesWritter.PutLe32(r.state[i])
	}

	//fmt.Println(dataBytesWritter.Bytes())
	dataBytesWritter.Flush()
	fmt.Println(dataBytesWritter.Bytes())
	return dataBytes
}

//func (r *Ripemd160) Final() []byte {
//	var result = make([]byte, 0)
//	var addition = make([]byte, 4)
//	finalReader := iobit.NewReader(r.state)
//
//	return result
//}

func (r *Ripemd160) Finish() (thebits uint64, padindex int32, pad []byte){
	thebits = r.processed_bytes * 8

	//if r.buffer.Pos < 7 * 8{
	//	padindex = 56 - r.buffer.Pos
	//} else{
	//	padindex = 120 - r.buffer.Pos
	//}
//	System.SetLength(pad, padindex + 8)
//
	pad[0] = 0x80
//
//	bits = TConverters.le2me_64(bits)
	bitsReader := []byte{}
	binary.LittleEndian.PutUint64(bitsReader, thebits)
	thebits, _ = binary.Uvarint(bitsReader)
//	TConverters.ReadUInt64AsBytesLE(bits, pad, padindex)
	padReader := iobit.NewReader(pad)
	thebits = padReader.Le64()

	//padindex = padindex + 8
	padReader.Skip(8)

	//TransformBytes(pad, 0, padindex)

	return
}

func (r *Ripemd160)TransformBytes(a_data []byte) []byte {
	ptr_a_data := a_data[:]

	// a_index => At
	// a_length => LeftBits
	adataReader := iobit.NewReader(a_data)
	//if (not Fm_buffer.IsEmpty) then
	//begin
	//	if (Fm_buffer.Feed(ptr_a_data, System.Length(a_data), a_index, a_length,
	//		Fm_processed_bytes)) then
	//	begin
	//	TransformBuffer();
	//	end;
	//end;

	//if r.buffer.Len()!=0 {
	//	l, err := r.buffer.(ptr_a_data)
	//	if (){
	//		r.TransformBlock(r.buffer.Bytes())
	//	}
	//}


	//while (a_length >= (Fm_buffer.Length)) do
	//begin
	//	Fm_processed_bytes := Fm_processed_bytes + UInt64(Fm_buffer.Length);
	//	TransformBlock(ptr_a_data, Fm_buffer.Length, a_index);
	//	a_index := a_index + (Fm_buffer.Length);
	//	a_length := a_length - (Fm_buffer.Length);
	//end;

	for adataReader.LeftBits() > Ripemd160BlockSize {
		r.processed_bytes = r.processed_bytes + uint64(Ripemd160BlockSize)
		r.TransformBlock(ptr_a_data)
		adataReader.Skip(Ripemd160BlockSize)
	}

	//if (a_length > 0) then
	//begin
	//	Fm_buffer.Feed(ptr_a_data, System.Length(a_data), a_index, a_length,
	//		Fm_processed_bytes);
	//end;

	return ptr_a_data

//func Create(a_state_length int32, a_hash_size int32){
//	hash_size = 64
//	System.SetLength(Fm_state, a_state_length);
}

//func GetResult() []byte {
//	System.SetLength(result, System.Length(Fm_state) * System.SizeOf(UInt32));
//
//	TConverters.le32_copy(PCardinal(Fm_state), 0, PByte(result), 0, System.Length(result));
//	return result
//}

//func Clone() (HashInstance *Ripemd160) {
//	HashInstance = Create()
//	HashInstance.Fm_state := System.Copy(Fm_state);
//	HashInstance.Fm_buffer := Fm_buffer.Clone();
//	HashInstance.Fm_processed_bytes := Fm_processed_bytes;
//	result := HashInstance as IHash;
//	result.BufferSize := BufferSize;
//}

//func (r *Ripemd160) TransformFinal() {
//	r.Finish()
//	//tempresult := GetResult();
//	//result :=  r.GetResult()
//	r.Initialize()
//	return result
//}

func (r *Ripemd160) ComputeBytes(in []byte) []byte {
	r.Initialize() // Initialize();
	r.TransformBytes(in)
	// result := THashResult.Create(tempresult);
	//r.TransformBlock(in)

	var result =  []byte{}
	addition:= make([]byte, 8)
	//var addition []byte

	for i:=0; i<len(r.state); i++ {
		fmt.Println(r.state[i])
		binary.LittleEndian.PutUint32(addition, r.state[i])
		result = bytes.Join([][]byte{result, addition}, nil)
	}

	return result
}