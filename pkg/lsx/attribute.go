package lsx

// Receiver function for attribute elements
func (attr lsxAttribute) GetId() string {
	return attr.Id // Attribute elements have attribute ids
}

// Maps attrs of Attribute lsx element
type lsxAttribute struct {
	Id    string `xml:"id,attr"`
	Type  string `xml:"type,attr"`  // directly maps to the type of the attribute
	Value string `xml:"value,attr"` // directly maps to the value of the attribute
}
