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


	t.Run("test_find_parallel1", func(t *testing.T) {
		numCombs := 10000
		maxSize := 2
		minSize := 0
		charSet := "abcd"
		wantedHash := "4a8a08f09d37b73795649038408b5f33"
		print := true
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print)
		got, _ := finder.FindParallel()
		want := "c"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	// t.Run("test_find_parallel_not_found", func(t *testing.T) {
	// 	numCombs := 10000000 * 10000000 // a big number
	// 	maxSize := 2
	// 	minSize := 0
	// 	charSet := "def"
	// 	wantedHash := "4a8a08f09d37b73795649038408b5f33"
	// 	print := true
	// 	finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print)
	// 	got, _ := finder.FindParallel()
	// 	want := "c"
	// 	if diff := cmp.Diff(want, got); diff != "" {
	// 		t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
	// 	}
	// })
}
