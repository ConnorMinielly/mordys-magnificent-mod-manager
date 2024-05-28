package lsx

import (
	"encoding/xml"
	"fmt"
	"os"
)

// ReadLsx unmarshals an lsx byte slice into an LSX struct.
func ReadLsx(lsxData []byte) (lsx LSX) {
	err := xml.Unmarshal(lsxData, &lsx)
	if err != nil {
		fmt.Println("Failed to unmarshal lsx data:", err)
	}
	return lsx
}

// ReadLsxFromFile reads an lsx file and unmarshals it into an LSX struct.
func ReadLsxFromFile(filePath string) LSX {
	lsxData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read lsx file:", err)
	}
	return ReadLsx(lsxData)
}

// Generic type for accessing lsx elements
type lsxElement interface {
	lsxAttribute | lsxNode // Type constraint for lsx elements
	GetId() string         // All lsx elements have an id attribute, which must be accessible from within the type constraint
}

// Generic function for getting a generic lsx element by id
func GetById[T lsxElement](data []T, match string) T {
	for _, node := range data {
		if node.GetId() == match {
			return node
		}
	}

	// If no match is found, return a new instance of the type as a default value:
	// The `new` built-in allocates storage for a variable of any type and returns
	// a pointer to it, so dereferencing *new(T) effectively yields the zero value for T.
	// https://stackoverflow.com/a/70589302
	return *new(T) // https://go.dev/ref/spec#Allocation
}

// Root struct representing an LSX file
type LSX struct {
	XMLName xml.Name   `xml:"save"` // save seems to be the root tag of all (?) lsx files
	Version lsxVersion `xml:"version"`
	Region  lsxRegion  `xml:"region"`
}

// Maps attrs of Children lsx element
type lsxChildren struct {
	Nodes []lsxNode `xml:"node"` // This is a recursive struct, containing nodes
}
