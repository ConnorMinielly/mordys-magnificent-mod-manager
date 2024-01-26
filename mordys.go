package main

import (
	"fmt"
	lsgo "mordys/lsgo"
)

func main() {
	pak := lsgo.ReadPak("./test_files/5eSpells/5eSpells.pak")
	fmt.Printf("header: %v\n", pak.Header)
}
