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

// func GetGeneric[T lsxNode | lsxAttribute](data []T, match string) T {
// 	for _, node := range data {
// 		if node.Id == match {
// 			return node
// 		}
// 	}
// 	return lsxNode{}
// }

func (region lsxRegion) Get(match string) lsxNode {
	for _, node := range region.Nodes {
		if node.Id == match {
			return node
		}
	}
	return lsxNode{}
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
