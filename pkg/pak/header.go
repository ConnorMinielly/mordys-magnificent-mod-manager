package pak

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ReadHeader(reader *bytes.Reader) (header Header, err error) {
	err = binary.Read(reader, binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("Failed to read header of pak file:", err)
		return Header{}, err
	}
	return header, nil
}

// Larian Studios Pak Header
type Header struct {
	Version        uint32
	FileListOffset uint64
	FileListSize   uint32
	Flags          byte
	Priority       byte
	Md5            [16]byte
	NumParts       uint16
}
