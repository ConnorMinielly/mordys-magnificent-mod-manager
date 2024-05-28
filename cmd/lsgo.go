package main

import (
	"fmt"
	"lsgo/pkg/lsx"
)

func main() {
	lsx := lsx.ReadLsxFromFile("./test/files/meta.lsx")

	fmt.Println(lsx.Consolidate())
	// pak.Read("./test/files/BlackDye.pak")
}
