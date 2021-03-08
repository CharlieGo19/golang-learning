package versioncontrol

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

// We create a new semantic version and return the struct.
// This is a simple version label.
func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// So instead of passing SemanticVersion as an arguement we can use a method reciever.
// This is the part straight after func, so it's saying, do something on this 'object'.
// This allows us in the code to do someting like: sv.String(). Similar syntax to a class.
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// This function does not work, because behind the scenes, sv here is just a copy of the calling variable/instance.
// Therefore the changes are not reflected in the calling instance.
func (sv SemanticVersion) IncrementMajor() {
	sv.major++
}

// This function will work as intended, because this is pointer method recievers, where we're telling it where the
// instance is that we would like to operate on; thus, the changes are reflected on the calling instance.
func (sv *SemanticVersion) IncrementMajorPointer() {
	sv.major++
}
