package lsgo

type LSPK struct {
	Signature [4]byte
	Header    LSPKHeader
}

// Larian Studios Pak Header
type LSPKHeader struct {
	Version        uint32   // 4
	FileListOffset uint64   // 8
	FileListSize   uint32   // 4
	Flags          byte     // 1
	Priority       byte     // 1
	Md5            [16]byte // 16
	NumParts       uint16   // 2
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
