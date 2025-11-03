package codec

import (
	"bytes"
	img "image"
	"image/png"
	"math"
)


// Embed arbitrary bytes into an image.
// The resulting image pixels buffer contains a 4 bytes header
// for the length of the payload, followed by the payload itself,
// then optionally extra padding to be disgarded.
func Encode(payload []byte) (m *img.NRGBA, err error) {
	header := BuildHeader(payload)
	bytesLen := len(header) + len(payload)

	// Generate a square image.
	// We can store 4 bytes in each pixel.
	stride := int(math.Ceil(math.Sqrt(float64(bytesLen / 4))))

	// This allocates the image buffer; we write to it directly.
	m = img.NewNRGBA(img.Rect(0, 0, stride, stride))
	buf := SliceWriter[byte]{Buf: m.Pix}
	_, err = buf.Write(header)
	if err != nil { return }
	_, err = buf.Write(payload)
	if err != nil { return }

	// Debug: mark all the unused extra pixels red.
	/*
	println(len(header))
	println(len(payload))
	n := len(header) + len(payload)
	for i := (n &^ 0b11) + 4; i < len(m.Pix); i += 4 {
		m.Pix[i] = 0xff
		m.Pix[i+1] = 0
		m.Pix[i+2] = 0
		m.Pix[i+3] = 0xff
	}
	*/

	return
}

func EncodePNG(payload []byte) ([]byte, error) {
	m, err := Encode(payload)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	encoder := png.Encoder{CompressionLevel: png.NoCompression}
	err = encoder.Encode(&buf, m)
	return buf.Bytes(), err
}


// Extracts the payload bytes embedded in the image pixels buffer.
// The returned slice will be backed by the same image buffer.
func Decode(m *img.NRGBA) (payload []byte, err error) {
	return ExtractPayload(m.Pix)
}


func DecodePNG(pngData []byte) ([]byte, error) {
	m, err := png.Decode(bytes.NewBuffer(pngData))
	if err != nil {
		return nil, err
	}
	return Decode(m.(*img.NRGBA))
}
