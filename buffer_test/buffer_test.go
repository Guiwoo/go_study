package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

type ABC struct {
	A uint32
	B uint16
	D uint8
}

func SetBinaryData() []byte {
	data := ABC{
		A: 1293,
		B: 233,
		D: 8,
	}
	bin := new(bytes.Buffer)
	// 데이터를 BigEndian으로 bin에 쓴다.
	binary.Write(bin, binary.LittleEndian, data.A)
	binary.Write(bin, binary.LittleEndian, data.B)
	binary.Write(bin, binary.LittleEndian, data.D)

	return bin.Bytes()
}

// 버퍼 테스트 바이너리 데이터를 읽고 파싱
func BenchmarkBufferTest(b *testing.B) {
	data := SetBinaryData()

	b.ResetTimer()

	b.Run("readbinary", func(b *testing.B) {
		reader := bytes.NewReader(data)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var c ABC
				binary.Read(reader, binary.LittleEndian, &c.A)
				binary.Read(reader, binary.LittleEndian, &c.B)
				binary.Read(reader, binary.LittleEndian, &c.D)
			}
		})
	})

	b.Run("readbinary2", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var parse ABC
				parse.A = binary.LittleEndian.Uint32(data[0:4])
				parse.B = binary.LittleEndian.Uint16(data[4:6])
				parse.D = data[6]
				if parse.A != 1293 || parse.B != 233 || parse.D != 8 {
					b.Error(fmt.Errorf("fail to parse %+v", parse))
				}
			}
		})
	})
}
