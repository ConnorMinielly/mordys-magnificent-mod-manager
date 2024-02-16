package main

import (
	lsgo "mordys/lsgo"
)

func main() {
	// lsx := lsgo.ReadLsxFromFile("./test_files/meta.lsx")

	// fmt.Println(lsx.Simplify())
	lsgo.ReadPak("./test_files/BlackDye.pak")
}
