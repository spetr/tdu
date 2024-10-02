package units

import (
	"math"
	"testing"
)

func TestHumanSize(t *testing.T) {
	tEquals(t, "1KB", DataSizeHuman(1024, 0))
	tEquals(t, "1MB", DataSizeHuman(1024*1024, 0))
	tEquals(t, "1MB", DataSizeHuman(1048576, 0))
	tEquals(t, "2MB", DataSizeHuman(2*MB, 0))
	tEquals(t, "3.42GB", DataSizeHuman(int64(math.Round(3.42*GB)), 2))
	tEquals(t, "5.372TB", DataSizeHuman(int64(math.Round(5.372*TB)), 3))
	tEquals(t, "2.22PB", DataSizeHuman(int64(math.Round(2.22*PB)), 2))
}

func tEquals(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}
