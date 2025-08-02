package partition

import (
	"github.com/rstms/go-diskfs/partition/part"
	"github.com/rstms/go-diskfs/util"
)

// Table reference to a partitioning table on disk
type Table interface {
	Type() string
	Write(util.File, int64) error
	GetPartitions() []part.Partition
}
