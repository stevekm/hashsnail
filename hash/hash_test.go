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
		numThreads := 2
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
		got, _ := finder.Find()
		want := "abc"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("0test_find_parallel", func(t *testing.T) {
		numCombs := 10000
		maxSize := 2
		minSize := 0
		charSet := "abcd"
		wantedHash := "4a8a08f09d37b73795649038408b5f33"
		print := false
		numThreads := 2
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
		got, _ := finder.FindParallel()
		want := "c"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("test_find_parallel_neg_numCombs", func(t *testing.T) {
		numCombs := -1 // this should be unlimited number of searches
		maxSize := 2
		minSize := 0
		charSet := "abcd"
		wantedHash := "4a8a08f09d37b73795649038408b5f33"
		print := false
		numThreads := 2
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
		got, _ := finder.FindParallel()
		want := "c"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("1test_find_parallel_max_size_exceeded", func(t *testing.T) {
		// the desired combination is larger than the allowed maxSize and wont be found
		numCombs := 10000
		maxSize := 2
		minSize := 0
		charSet := "abc"
		wantedHash := "26ca5bfe74f8de88ccaac5c0f44b349d"
		print := false
		numThreads := 2
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
		got, _ := finder.FindParallel() // TODO: check the error instead!
		want := ""                      // "abcc"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("1test_find_parallel_unlimited_max_size", func(t *testing.T) {
		// find the combination with an unlimited max size
		numCombs := 10000
		maxSize := -1
		minSize := 0
		charSet := "abc"
		wantedHash := "26ca5bfe74f8de88ccaac5c0f44b349d"
		print := false
		numThreads := 2
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
		got, _ := finder.FindParallel()
		want := "abcc"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("1test_find_parallel_not_found", func(t *testing.T) {
		// the desired combination is not possible from the charSet values
		numCombs := 10000000
		maxSize := 2
		minSize := 0
		charSet := "def"
		wantedHash := "4a8a08f09d37b73795649038408b5f33"
		print := false
		numThreads := 2
		finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
		got, _ := finder.FindParallel() // TODO: check the error instead!
		want := ""                      // "c"
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})

	// NOTE: This one is deadlocking, need to fix this !!
	// t.Run("1test_find_not_found2", func(t *testing.T) {
	// 	// the desired combination is not possible from the charSet values
	// 	numCombs := -1
	// 	maxSize := -1
	// 	minSize := 0
	// 	charSet := "def"
	// 	wantedHash := "e2fc714c4727ee9395f324cd2e7f331f" // "abcd"
	// 	print := false
	// 	numThreads := 2
	// 	finder := NewHashFinder(numCombs, maxSize, minSize, charSet, wantedHash, print, numThreads)
	// 	got, _ := finder.FindParallel() // TODO: check the error instead!
	// 	want := "" // "abcd"
	// 	if diff := cmp.Diff(want, got); diff != "" {
	// 		t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
	// 	}
	// })
}
