package lsgo

import (
	"encoding/xml"
)

// Root struct representing an LSX file
type LSX struct {
	XMLName xml.Name   `xml:"save"` // save seems to be the root tag of all (?) lsx files
	Version lsxVersion `xml:"version"`
	Region  lsxRegion  `xml:"region"`
}

// Maps attrs of Version lsx element
type lsxVersion struct {
	Major    string `xml:"major,attr"`
	Minor    string `xml:"minor,attr"`
	Revision string `xml:"revision,attr"`
	Build    string `xml:"build,attr"`
}

// Maps attrs of Region lsx element
type lsxRegion struct {
	Id    string    `xml:"id,attr"`
	Nodes []lsxNode `xml:"node"`
}

// Maps attrs of Node lsx element
type lsxNode struct {
	Id         string         `xml:"id,attr"`
	Attributes []lsxAttribute `xml:"attribute"` // A node can have some or no attributes AND
	Children   lsxChildren    `xml:"children"`  // A node can have no children lsx element OR one children lsx element
}

// Maps attrs of Children lsx element
type lsxChildren struct {
	Nodes []lsxNode `xml:"node"` // This is a recursive struct, containing nodes
}

// Maps attrs of Attribute lsx element
type lsxAttribute struct {
	Id    string `xml:"id,attr"`
	Type  string `xml:"type,attr"`  // directly maps to the type of the attribute
	Value string `xml:"value,attr"` // directly maps to the value of the attribute
}

// Final struct representing meta.lsx attrs we care about
type MetaLsx struct {
	Version string
	Name    string
	Author  string
	UUID    string
}
