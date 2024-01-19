package main

// PAK READER STRUCTURE

// 1. read .pak file into go as binary.

// 2. grab first 4 bytes.

// 3. check to see if first 4 bytes equal the signature array of [0x4C,0x53,0x50,0x4B].

// 4. convert next(?) 4 bytes into a 32 bit integer to get version number

// 5. confirm that integer = 18 (version 18 of the pak toolchain is what's used for BG3 release).

// 6. read the next 16 bytes (the file header) into a header struct (??)

/// NOTE: Binary reader is little endian by default.

// Struct to marshal the pak file header into.
// type LarianStudioPakHeader struct {
// 	Version uint32
// 	FileListOffset uint64
// 	FileListSize uint32
// 	Flags byte
// 	Priority byte
// 	Md5 []byte // 16 bytes?
// 	NumParts uint16
// }

// func main() {

// }