package ripemd

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRipemd160(t *testing.T) {
	// return DoRipemd160("") == "9c1185a5c5e9fc54612808977ee8f548b2258d31"
	// RIPEMD-160("The quick brown fox jumps over the lazy dog") = 37f332f68db77bd9d7edd4969571ad671cf9dd3b
	//r := New160()
	//result := r.ComputeBytes([]byte(""))
	//t.Log(hex.EncodeToString(result))
	message := []byte("")
	buf := MDinit()
	//MDfinish(buf)
	length := len(message)

	var X [16]uint64
	/* process message in 16-word chunks */
	for nbytes := length; nbytes > 63; nbytes -= 64 {
		for i := 0; i < 16; i++ {
			X[i] = bytes2uint64(message)
			//message += 4
		}
		Rounds(buf, X)
	} /* length mod 64 bytes left */

	/* finish: */
	msgBuf := bytes.NewBuffer(message)
	MDfinish(buf, msgBuf, uint64(length), 0)

	//var hashcode [160 / 8]byte

	var hash []byte
	for i := 0; i < 160/8; i += 4 {
		fmt.Println("buf", buf[i>>2])
		var subhash []byte
		binary.LittleEndian.PutUint64(subhash, buf[i>>2]) /* implicit cast to byte  */
		hash = bytes.Join([][]byte{hash, subhash}, nil)
		binary.LittleEndian.PutUint64(hash, (buf[i>>2] >> 8)) /*  extracts the 8 least  */
		hash = bytes.Join([][]byte{hash, subhash}, nil)
		binary.LittleEndian.PutUint64(hash, (buf[i>>2] >> 16)) /*  significant bits.     */
		hash = bytes.Join([][]byte{hash, subhash}, nil)
		binary.LittleEndian.PutUint64(hash, (buf[i>>2] >> 24))
		hash = bytes.Join([][]byte{hash, subhash}, nil)
	}

	fmt.Println(hex.EncodeToString(hash))

}
