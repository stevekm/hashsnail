package hash

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("test_find_hash", func(t *testing.T) {
		numCombs := 10000000 * 10000000 // a big number
		maxSize := 4
		minSize := 0
		charSet := "abcdefg"
		wantedHash := "900150983cd24fb0d6963f7d28e17f72"
		print := false
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print)
		got, _ := finder.Find()
		want := "abc"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})
}
