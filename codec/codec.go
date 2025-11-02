package codec

import (
	"io";
	. "image";
	"math";
)


const HeaderSize = 4

func BuildHeader(src []byte) []byte {
	length := uint(len(src))
	header := make([]byte, 4)
	header[0] = byte(0xff & length)
	header[1] = byte(0xff & (length >> 8))
	header[2] = byte(0xff & (length >> 16))
	header[3] = byte(0xff & (length >> 24))
	return header
}


func ExtractHeader(data []byte) uint {
	return uint(data[0] | data[1] << 8 | data[2] << 16 | data[3] << 24)
}


func WritePacket(dst io.Writer, payload []byte) (n int, err error) {
	if n, err = dst.Write(BuildHeader(payload)); err != nil {
		return
	}
	m, err := dst.Write(payload)
	return n + m, err
}


func Encode(payload []byte) (img Image, err error) {
	bytesLen := uint(len(payload)) + HeaderSize
	// We can store 4 bytes in each pixel.
	stride := int(math.Ceil(math.Sqrt(float64(bytesLen / 4))))

	// This allocates the image buffer; we write to it directly.
	nrgba := NewNRGBA(Rect(0, 0, stride, stride))
	n, err := WritePacket(SliceWriter[byte]{Buf: nrgba.Pix}, payload)

	// Debug: mark all the unused extra pixels red.
	_ = n
	/*
	for i := (n &^ 0b11) + 4; i < len(nrgba.Pix); i += 4 {
		nrgba.Pix[i] = 0xff
		nrgba.Pix[i+1] = 0
		nrgba.Pix[i+2] = 0
		nrgba.Pix[i+3] = 0xff
	}
	*/

	return nrgba, err
}
