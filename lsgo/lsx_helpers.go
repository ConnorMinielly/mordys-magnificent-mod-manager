package lsgo

import "fmt"

// This is a helper function that will likely get rewritten, but for now can be used to consolidate complex nested lsx data into known data structures
func (lsx LSX) Consolidate() any {
	switch lsx.Region.Id {
	case "Config": // meta.lsx
		metaLsx := MetaLsx{
			Version: lsx.Version.GetVersionString(),
			Name:    lsx.Region.GetNode("root").GetChild("ModuleInfo").GetAttribute("Name").Value,
			Author:  lsx.Region.GetNode("root").GetChild("ModuleInfo").GetAttribute("Author").Value,
			UUID:    lsx.Region.GetNode("root").GetChild("ModuleInfo").GetAttribute("UUID").Value,
		}
		return metaLsx
	case "ModuleSettings": // modsettings.lsx
		//  TODO: Implement modsettings.lsx reducer
		return lsx
	default:
		fmt.Println("WARNING: Can't simplify this lsx data into a known structure.")
		return lsx
	}
}

func (version lsxVersion) GetVersionString() string {
	return fmt.Sprintf("%s.%s.%s.%s", version.Major, version.Minor, version.Revision, version.Build)
}

// Generic type for accessing lsx elements
type lsxElement interface {
	lsxAttribute | lsxNode // Type constraint for lsx elements
	GetId() string         // All lsx elements have an id attribute, which must be accessible from within the type constraint
}

func (node lsxNode) GetId() string {
	return node.Id
}

func (attr lsxAttribute) GetId() string {
	return attr.Id
}

func GetGeneric[T lsxElement](data []T, match string) T {
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

func (region lsxRegion) GetNode(match string) lsxNode {
	return GetGeneric[lsxNode](region.Nodes, match)
}

func (node lsxNode) GetAttribute(match string) lsxAttribute {
	return GetGeneric[lsxAttribute](node.Attributes, match)
}

func (rootNode lsxNode) GetChild(match string) lsxNode {
	return GetGeneric[lsxNode](rootNode.Children.Nodes, match)
}
