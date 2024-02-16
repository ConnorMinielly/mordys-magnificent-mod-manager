package lsgo

import (
	"encoding/xml"
	"fmt"
)

// Root struct representing an LSX file
type LSX struct {
	XMLName xml.Name   `xml:"save"` // save seems to be the root tag of all (?) lsx files
	Version lsxVersion `xml:"version"`
	Region  lsxRegion  `xml:"region"`
}

// This is a helper function that will likely get scraped or rewritten, but for now can be used to consolidate complex nested lsx data into known data structures
func (lsx LSX) Simplify() any {
	switch lsx.Region.Id {
	case "Config": // meta.lsx
		metaLsx := MetaLsx{
			Version: lsx.Version.GetVersionString(),
			Name:    lsx.Region.Get("root").GetChild("ModuleInfo").GetAttribute("Name").Value,
			Author:  lsx.Region.Get("root").GetChild("ModuleInfo").GetAttribute("Author").Value,
			UUID:    lsx.Region.Get("root").GetChild("ModuleInfo").GetAttribute("UUID").Value,
		}
		return metaLsx
	case "ModuleSettings": // modsettings.lsx
		return lsx
	default:
		fmt.Println("WARNING: Can't simplify this lsx data into a known structure.")
		return lsx
	}
}

type lsxVersion struct {
	Major    string `xml:"major,attr"`
	Minor    string `xml:"minor,attr"`
	Revision string `xml:"revision,attr"`
	Build    string `xml:"build,attr"`
}

func (version lsxVersion) GetVersionString() string {
	return fmt.Sprintf("%s.%s.%s.%s", version.Major, version.Minor, version.Revision, version.Build)
}

type lsxRegion struct {
	Id    string    `xml:"id,attr"`
	Nodes []lsxNode `xml:"node"`
}

func (region lsxRegion) Get(match string) lsxNode {
	for _, node := range region.Nodes {
		if node.Id == match {
			return node
		}
	}
	return lsxNode{}
}

type lsxNode struct {
	Id         string         `xml:"id,attr"`
	Attributes []lsxAttribute `xml:"attribute"`
	Children   lsxChildren    `xml:"children"`
}

func (node lsxNode) GetAttribute(match string) lsxAttribute {
	for _, attr := range node.Attributes {
		if attr.Id == match {
			return attr
		}
	}
	return lsxAttribute{}
}

func (rootNode lsxNode) GetChild(match string) lsxNode {
	for _, node := range rootNode.Children.Nodes {
		if node.Id == match {
			return node
		}
	}
	return lsxNode{}
}

type lsxChildren struct {
	Nodes []lsxNode `xml:"node"`
}

type lsxAttribute struct {
	Id    string `xml:"id,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type MetaLsx struct {
	Version string
	Name    string
	Author  string
	UUID    string
}
