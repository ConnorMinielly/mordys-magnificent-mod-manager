package pak

import (
	"bytes"
)

func ReadFiles(reader *bytes.Reader, offset int64) (files []File) {
	fr := FileReader{reader, offset}

	// 1. Seek by offset amount (file list offset is found in the header,
	//    turns out you only need to calculate size for older pak formats).
	reader.Seek(offset, 0)

	// 2. Read a 32 bit integer, this is the number of files in the mod.
	numFiles := fr.GetNumberOfFiles()

	destBytes := fr.ExtractSourceBytes()

	newReader := bytes.NewReader(destBytes)

	newFileReader := FileReader{newReader, 0}

	for i := 0; i < int(numFiles); i++ {
		file := newFileReader.ReadFile()

		files = append(files, file)
		// fmt.Println(string(file.Name[:]))
	}
	return
}

// Larian Studios Pak File
type File struct {
	Name             [256]byte
	OffsetInFile1    uint32
	OffsetInFile2    uint16
	ArchivePart      byte
	Flags            byte
	SizeOnDisk       uint32
	UncompressedSize uint32
}
