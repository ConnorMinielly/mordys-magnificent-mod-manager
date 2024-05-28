package lsx

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

// Representings meta.lsx attrs
type MetaLsx struct {
	Version string
	Name    string
	Author  string
	UUID    string
}
