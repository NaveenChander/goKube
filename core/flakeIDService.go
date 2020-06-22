package core

import flake "github.com/davidnarayan/go-flake"

// GoFlake ... GoFlake
var GoFlake *flake.Flake

// InitFlakeIDGenerator ... Init Global variables
func InitFlakeIDGenerator() {
	GoFlake, _ = flake.New()
}

// GetNextFlakeID ... Returns new Flake ID
func GetNextFlakeID() uint64 {
	return uint64(GoFlake.NextId())
}
