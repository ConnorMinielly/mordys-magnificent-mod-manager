package lsgo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/pierrec/lz4"
)

// PAK READER ORDER OF OPERATIONS: Header Data
// 1. read .pak file into go as binary.
// 2. grab first 4 bytes and check that they match the signature string "LSPK" or [0x4C,0x53,0x50,0x4B]
// 3. if match, seek past the 4 byte signature block.
// 4. read ext block of binary into the Header struct (see code fo data shape)
// 5. confirm version code is = 18, other engine versions are unsupported at this time
// 6. If version matches 18, congrats you've pulled the meta data successfully from the pak file!

func ReadPak(filePath string) *LSPK {
	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)
	fmt.Println(reader.Len())

	var pakResults = new(LSPK)

	sig, err := ReadSignature(reader)
	if err != nil {
		panic(err)
	}

	if string(sig[:]) != "LSPK" {
		panic("WARNING: pak file signature doesn't match \"LSPK\", this is not a valid Larian Studios pak file.")
	}
	pakResults.Signature = sig

	header, err := ReadHeader(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("header: %v\n", header)
	pakResults.Header = header

	ReadFileList(reader, int64(header.FileListOffset))

	return pakResults
}

func ReadSignature(reader *bytes.Reader) ([4]byte, error) {
	var signatureBuffer [4]byte
	err := binary.Read(reader, binary.LittleEndian, &signatureBuffer)
	if err != nil {
		log.Fatal("Failed to read signature of pak file:", err)
		return [4]byte{}, err
	}
	return signatureBuffer, nil
}

func ReadHeader(reader *bytes.Reader) (LSPKHeader, error) {
	var header LSPKHeader
	err := binary.Read(reader, binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("Failed to read header of pak file:", err)
		return LSPKHeader{}, err
	}
	return header, nil
}

func ReadFileList(reader *bytes.Reader, offset int64) {
	// 1. Seek by offset amount (file list offset is found in the header, turns out you only need to calculate size for older pak formats)
	reader.Seek(offset, 0)

	// 2. read a 32 bit integer, this is the number of files in the mod.
	var numFiles int32
	binary.Read(reader, binary.LittleEndian, &numFiles)
	fmt.Println(numFiles)

	// 3. read another 32 bit integer, that is the size in bytes of the compression block (i think theres a shit ton of )
	var compressedSize int32
	binary.Read(reader, binary.LittleEndian, &compressedSize)
	fmt.Println(compressedSize)

	sourceBytes := make([]byte, compressedSize)
	reader.Read(sourceBytes)
	var fileEntry LSPKFileEntry
	fileBufferSize := binary.Size(fileEntry) * int(numFiles)
	destBytes := make([]byte, fileBufferSize)
	lz4.UncompressBlock(sourceBytes, destBytes)

	var fileEntries []LSPKFileEntry

	newReader := bytes.NewReader(destBytes)

	// reader.Seek(offset+8, 0)
	for i := 0; i < 10; i++ {
		var file LSPKFileEntry

		binary.Read(newReader, binary.LittleEndian, &file)
		fileEntries = append(fileEntries, file)
		fmt.Println(string(file.Name[:]))
	}
}
