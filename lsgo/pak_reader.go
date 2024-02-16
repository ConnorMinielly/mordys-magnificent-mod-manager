package lsgo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pierrec/lz4"
)

func ReadPak(filePath string) *LSPK {
	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)

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
	pakResults.Header = header

	files := ReadFileList(reader, int64(header.FileListOffset))

	metaFile := files[2]

	reader.Seek(int64(metaFile.OffsetInFile1), io.SeekStart)
	sourceBytes := make([]byte, metaFile.SizeOnDisk)
	reader.Read(sourceBytes)

	destBytes := make([]byte, metaFile.UncompressedSize)
	lz4.UncompressBlock(sourceBytes, destBytes)

	metaData := ReadLsx(destBytes)

	fmt.Println(metaData.Simplify())

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

func ReadFileList(reader *bytes.Reader, offset int64) []LSPKFileEntry {
	// 1. Seek by offset amount (file list offset is found in the header, turns out you only need to calculate size for older pak formats)
	reader.Seek(offset, 0)

	// 2. read a 32 bit integer, this is the number of files in the mod.
	var numFiles int32
	binary.Read(reader, binary.LittleEndian, &numFiles)

	// 3. read another 32 bit integer, that is the size in bytes of the compression block (i think theres a shit ton of )
	var compressedSize int32
	binary.Read(reader, binary.LittleEndian, &compressedSize)

	sourceBytes := make([]byte, compressedSize)
	reader.Read(sourceBytes)
	var fileEntry LSPKFileEntry
	fileBufferSize := binary.Size(fileEntry) * int(numFiles)
	destBytes := make([]byte, fileBufferSize)
	lz4.UncompressBlock(sourceBytes, destBytes)

	var fileEntries []LSPKFileEntry

	newReader := bytes.NewReader(destBytes)

	// TODO: Fix this to iterate a number of times = numFiles
	for i := 0; i < 10; i++ {
		var file LSPKFileEntry

		binary.Read(newReader, binary.LittleEndian, &file)
		fileEntries = append(fileEntries, file)
		// fmt.Println(string(file.Name[:]))
	}

	return fileEntries
}
