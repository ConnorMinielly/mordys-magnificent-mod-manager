package lsx

// GetNode returns a node by id from the region.
func (region lsxRegion) GetNode(match string) lsxNode {
	return GetById(region.Nodes, match)
}

// Maps attrs of Region lsx element
type lsxRegion struct {
	Id    string    `xml:"id,attr"`
	Nodes []lsxNode `xml:"node"`
}
