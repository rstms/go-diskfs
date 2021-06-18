package diskfs

// this constants should be part of "golang.org/x/sys/unix", but aren't, yet
const (
	DKIOCGETBLOCKSIZE         = 0x40046418
	DKIOCGETPHYSICALBLOCKSIZE = 0x4004644D
	DKIOCGETBLOCKCOUNT        = 0x40086419
)
