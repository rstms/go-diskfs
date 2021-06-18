// Package partition provides ability to work with individual partitions.
// All useful implementations are subpackages of this package, e.g. github.com/cusspvz/go-diskfs/gpt
package partition

import (
	"fmt"

	"github.com/cusspvz/go-diskfs/partition/gpt"
	"github.com/cusspvz/go-diskfs/partition/mbr"
	"github.com/cusspvz/go-diskfs/util"
)

// Read read a partition table from a disk
func Read(f util.File, logicalBlocksize, physicalBlocksize int) (Table, error) {
	// just try each type
	gptTable, err := gpt.Read(f, logicalBlocksize, physicalBlocksize)
	if err == nil {
		return gptTable, nil
	}
	mbrTable, err := mbr.Read(f, logicalBlocksize, physicalBlocksize)
	if err == nil {
		return mbrTable, nil
	}
	// we are out
	return nil, fmt.Errorf("Unknown disk partition type")
}
