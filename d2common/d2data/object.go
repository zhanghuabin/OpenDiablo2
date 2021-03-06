package d2data

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// Object is a game world object
type Object struct {
	Type       int
	Id         int //nolint:golint Id is the right key
	X          int
	Y          int
	Flags      int
	Paths      []d2common.Path
}
