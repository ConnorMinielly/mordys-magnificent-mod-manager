package pak

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"lsgo/pkg/lsx"
	"os"

	"github.com/pierrec/lz4"
)

// Binary byte order for reading pak files.
var ByteOrder = binary.LittleEndian

// Reads a pak file and returns a Pak struct.
func Read(filePath string) Pak {
	// Read the pak file.
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read pak file:", err)
	}

	// Create a reader from the data.
	reader := bytes.NewReader(data)

	// Initialize a new Pak struct.
	pak := new(Pak)

	// Read the signature of the pak file.
	signature, err := ReadSignature(reader)
	if err != nil {
		fmt.Println("Failed to read signature of pak file:", err)
	}

	// Check if the signature is valid.
	if string(signature[:]) != "LSPK" {
		panic("WARNING: pak file signature doesn't match \"LSPK\", this is not a valid Larian Studios pak file.")
	}
	// Assign the signature to the pak struct.
	pak.Signature = signature

	// Read the header of the pak file.
	header, err := ReadHeader(reader)
	if err != nil {
		fmt.Println("Failed to read header of pak file:", err)
	}
	// Assign the header to the pak struct.
	pak.Header = header

	// Read the files of the pak file.
	files := ReadFiles(reader, int64(header.FileListOffset))

	// Read the meta file of the pak file.
	metaFile := files[2] // meta.lsx is always the third file in the list.

	// Seek to the offset of the meta file in the pak file.
	reader.Seek(int64(metaFile.OffsetInFile1), io.SeekStart)
	// Create a byte slice for the source bytes of the meta file.
	sourceBytes := make([]byte, metaFile.SizeOnDisk)
	// Read the source bytes of the meta file.
	reader.Read(sourceBytes)

	// Create a byte slice for the destination bytes of the meta file.
	destBytes := make([]byte, metaFile.UncompressedSize)
	// Decompress the source bytes of the meta file.
	lz4.UncompressBlock(sourceBytes, destBytes)

	// Read the meta data of the pak file.
	metaData := lsx.ReadLsx(destBytes)

	// Print the consolidated meta data.
	fmt.Println(metaData.Consolidate())

	// Assign the files to the pak struct.
	return *pak
}

func ReadSignature(reader *bytes.Reader) (signatureBuffer [4]byte, err error) {
	err = binary.Read(reader, ByteOrder, &signatureBuffer)
	if err != nil {
		log.Fatal("Failed to read signature of pak file:", err)
		return [4]byte{}, err
	}

	return
}

// Represents a Larian Studios Pak file
type Pak struct {
	Files     []File
	Header    Header
	Signature [4]byte
}
