package lsgo

import (
	"encoding/xml"
	"os"
)

func ReadLsx(lsxContent []byte) LSX {
	var lsxSave LSX
	err := xml.Unmarshal(lsxContent, &lsxSave)
	if err != nil {
		panic(err)
	}
	return lsxSave
}

func ReadLsxFromFile(filePath string) LSX {
	lsxData, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return ReadLsx(lsxData)
}
