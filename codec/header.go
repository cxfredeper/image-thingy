package codec

import (
	"hash/crc32"
	"encoding/binary"
	"slices"
	"errors"
)


// Build a header from payload.
// The header consists of a 4-byte CRC32 checksum,
// followed by uvarint encoded length field of the length of the payload.
func BuildHeader(payload []byte) []byte {
	// Encode the variable length field.
	lenField := binary.AppendUvarint(nil, uint64(len(payload)))

	crc := crc32.NewIEEE()
	crc.Write(lenField)
	crc.Write(payload)

	hashField := make([]byte, 4)
	binary.BigEndian.PutUint32(hashField, crc.Sum32())

	return slices.Concat(hashField, lenField)
}


type header struct {
	checksum      uint32
	offset, length int  // Payload offset and length.
}


func ExtractPayload(data []byte) (payload []byte, err error) {
	h, err := parsePacket(data)
	if err != nil {
		return
	}

	end := h.offset + h.length
	if h.checksum != crc32.ChecksumIEEE(data[4:end]) {
		err = errors.New("Checksum failed")
		return
	}

	return data[h.offset:end], err
}


func parsePacket(data []byte) (h header, err error) {
	h = header{}
	h.checksum = binary.BigEndian.Uint32(data[:4])
	length, n := binary.Uvarint(data[4:])
	h.length = int(length)

	if n <= 0 {
		err = errors.New("Malformed header")
		return
	}

	h.offset = 4 + n

	if end := h.offset + h.length; end > len(data) {
		err = errors.New("Truncated packet")
	}

	return
}
