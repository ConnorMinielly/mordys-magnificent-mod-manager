package lsx

// Receiver function for node elements
func (node lsxNode) GetId() string {
	return node.Id // Node elements have node ids
}

func (node lsxNode) GetAttribute(match string) lsxAttribute {
	return GetById(node.Attributes, match)
}

func (rootNode lsxNode) GetChild(match string) lsxNode {
	return GetById(rootNode.Children.Nodes, match)
}

// Maps attrs of Node lsx element
type lsxNode struct {
	Id         string         `xml:"id,attr"`
	Attributes []lsxAttribute `xml:"attribute"` // A node can have some or no attributes AND
	Children   lsxChildren    `xml:"children"`  // A node can have no children lsx element OR one children lsx element
}
