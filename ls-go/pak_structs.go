package lsgo

type LSPK struct {
	Header LSPKHeader
}

// Larian Studios Pak Header
type LSPKHeader struct {
	Version        uint32
	FileListOffset uint64
	FileListSize   uint32
	Flags          byte
	Priority       byte
	Md5            [16]byte
	NumParts       uint16
}

// Larian Studios Pak Sub-File Entry
type LSPKFileEntry struct {
	Name             [256]byte
	OffsetInFile1    uint32
	OffsetInFile2    uint16
	ArchivePart      byte
	Flags            byte
	SizeOnDisk       uint32
	UncompressedSize uint32
}
