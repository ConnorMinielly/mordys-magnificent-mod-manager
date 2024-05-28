package pak

import (
	"bytes"
	"encoding/binary"

	"github.com/pierrec/lz4"
)

type FileReader struct {
	reader *bytes.Reader
	offset int64
}

// Reads a file entry from the reader, and returns a File struct.
func (fr FileReader) ReadFile() (file File) {
	binary.Read(fr.reader, ByteOrder, file)
	return
}

func (fr FileReader) GetNumberOfFiles() (numFiles int32) {
	binary.Read(fr.reader, ByteOrder, &numFiles)
	return
}

func (fr FileReader) GetCompressedSize() (compressedSize int32) {
	binary.Read(fr.reader, ByteOrder, &compressedSize)
	return
}

func (fr FileReader) GetSourceBytes() []byte {
	sourceBytes := make([]byte, fr.GetCompressedSize())
	fr.reader.Read(sourceBytes)
	return sourceBytes
}

func (fr FileReader) GetFileEntrySize() int {
	var fileEntry File
	return binary.Size(fileEntry)
}

func (fr FileReader) GetFileBufferSize() int {
	return fr.GetFileEntrySize() * int(fr.GetNumberOfFiles())
}

func (fr FileReader) ExtractSourceBytes() []byte {
	destBytes := make([]byte, fr.GetCompressedSize())
	lz4.UncompressBlock(fr.GetSourceBytes(), destBytes)
	return destBytes
}
