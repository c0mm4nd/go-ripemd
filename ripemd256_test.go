package ripemd

import (
	"fmt"
	"io"
	"testing"
)

func Test256Vectors(t *testing.T) {
	var vectors = [...]mdTest{
		{"77093b1266befed58d512e67b3a8a15398c3ce5c1333d66a190becc9baa329e9", "123456"},
	}
	for i := 0; i < len(vectors); i++ {
		tv := vectors[i]
		md := New256()
		//for j := 0; j < 3; j++ {
		//	if j < 2 {
				io.WriteString(md, tv.in)
			//} else {
			//	io.WriteString(md, tv.in[0:len(tv.in)/2])
				//md.Sum(nil)
				//io.WriteString(md, tv.in[len(tv.in)/2:])
			//}
			s := fmt.Sprintf("%x", md.Sum(nil))
			if s != tv.out {
				t.Fatalf("RIPEMD-320(%s) = %s, expected %s", tv.in, s, tv.out)
			}
			md.Reset()
		//}
	}
}