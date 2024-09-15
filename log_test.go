package dnutlogger

import (
	"testing"
)

func Test(t *testing.T) {
	Debug("Debug")
	Success("Success")
	SetLevel(LevelDebug)
	Debug("Debug")
	Success("Success")
}
