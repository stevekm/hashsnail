package hash

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// wantedHash := "0cc175b9c0f1b6a831c399e269772661" // a 0.225s
// wantedHash := "900150983cd24fb0d6963f7d28e17f72" // abc 2s
// wantedHash := "e2fc714c4727ee9395f324cd2e7f331f" // abcd 2:36
// wantedHash := "ab56b4d92b40713acc5af89985d4b786" // abcde 1:40:50

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
