package lsgo

import "fmt"

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

func (node lsxNode) GetId() string {
	return node.Id
}

func (attr lsxAttribute) GetId() string {
	return attr.Id
}

func GetGeneric[T ILsx](data []T, match string) T {
	for _, node := range data {
		if node.GetId() == match {
			return node
		}
	}
	var emptyReturn T
	return emptyReturn // TODO: Find better way of doing this
}

func (region lsxRegion) Get(match string) lsxNode {
	return GetGeneric[lsxNode](region.Nodes, match)
}

func (node lsxNode) GetAttribute(match string) lsxAttribute {
	return GetGeneric[lsxAttribute](node.Attributes, match)
}

func (rootNode lsxNode) GetChild(match string) lsxNode {
	return GetGeneric[lsxNode](rootNode.Children.Nodes, match)
}
