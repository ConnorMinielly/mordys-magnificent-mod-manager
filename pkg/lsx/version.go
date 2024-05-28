package lsx

import "fmt"

// GetVersionString returns a string representation of the version from an lsx file.
func (version lsxVersion) GetVersionString() string {
	return fmt.Sprintf("%s.%s.%s.%s", version.Major, version.Minor, version.Revision, version.Build)
}

// Maps attrs of Version lsx element
type lsxVersion struct {
	Major    string `xml:"major,attr"`
	Minor    string `xml:"minor,attr"`
	Revision string `xml:"revision,attr"`
	Build    string `xml:"build,attr"`
}
