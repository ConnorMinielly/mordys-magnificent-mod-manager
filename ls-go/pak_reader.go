package lsgo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

// PAK READER ORDER OF OPERATIONS: Header Data
// 1. read .pak file into go as binary.
// 2. grab first 4 bytes and check that they match the signature string "LSPK" or [0x4C,0x53,0x50,0x4B]
// 3. if match, seek past the 4 byte signature block.
// 4. read ext block of binary into the Header struct (see code fo data shape)
// 5. confirm version code is = 18, other engine versions are unsupported at this time
// 6. If version matches 18, congrats you've pulled the meta data successfully from the pak file!

func ReadSignature(reader *bytes.Reader) ([]byte, error) {
	buf := make([]byte, 4)
	_, err := io.ReadAtLeast(reader, buf, 4)
	if err != nil {
		log.Fatal("Failed to read signature of pak file:", err)
		return nil, err
	}
	return buf, nil
}

func ReadHeader(reader *bytes.Reader) (*LSPKHeader, error) {
	var header LSPKHeader
	reader.Seek(4, 0)
	err := binary.Read(reader, binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("Failed to read header of pak file:", err)
		return nil, err
	}
	return &header, nil
}

func ReadPak(file_path string) *LSPK {
	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)

	var pak_results = new(LSPK)

	sig, err := ReadSignature(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", sig)

	if !bytes.Equal(sig, []byte("LSPK")) {
		fmt.Println("WARNING: pak file signature doesn't match \"LSPK\", this is not a valid Larian Studios pak file.")
	} else {
		header, err := ReadHeader(reader)
		if err != nil {
			panic(err)
		}
		pak_results.Header = *header
	}
	return pak_results
}
