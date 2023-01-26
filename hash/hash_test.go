package hash

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRunHash(t *testing.T) {
	tests := map[string]struct {
		finder HashFinder
		want   string
		got    string
	}{
		"hash_abc": {
			finder: NewHashFinder(
				10000000*10000000,                  // numCombs
				4,                                  // maxSize
				0,                                  // minSize
				"abcdefg",                          // charSet
				"900150983cd24fb0d6963f7d28e17f72", // wantedHash
				false,                              // print
				2),
			want: "abc",
		},
		"hash_c": {
			finder: NewHashFinder(
				10000,                              // numCombs
				2,                                  // maxSize
				0,                                  // minSize
				"abcd",                             // charSet
				"4a8a08f09d37b73795649038408b5f33", // wantedHash
				false,                              // print
				2),
			want: "c",
		},
		"unlimited_combs": {
			finder: NewHashFinder(
				-1,                                 // numCombs
				2,                                  // maxSize
				0,                                  // minSize
				"abcd",                             // charSet
				"4a8a08f09d37b73795649038408b5f33", // wantedHash
				false,                              // print
				2),
			want: "c",
		},
		"max_size_exceeded": {
			// the desired combination is larger than the allowed maxSize and wont be found
			// TODO: check the error instead!
			finder: NewHashFinder(
				10000,                              // numCombs
				2,                                  // maxSize
				0,                                  // minSize
				"abcd",                             // charSet
				"26ca5bfe74f8de88ccaac5c0f44b349d", // wantedHash
				false,                              // print
				2),
			want: "", // "abcc"
		},
		"unlimited_max_size": {
			// find the combination with an unlimited max size
			finder: NewHashFinder(
				10000,                              // numCombs
				-1,                                 // maxSize
				0,                                  // minSize
				"abcd",                             // charSet
				"26ca5bfe74f8de88ccaac5c0f44b349d", // wantedHash
				false,                              // print
				2),
			want: "abcc", // "abcc"
		},
		"not_found_in_charset": {
			// the desired combination is not possible from the charSet values
			finder: NewHashFinder(
				10000000,                           // numCombs
				2,                                  // maxSize
				0,                                  // minSize
				"def",                              // charSet
				"4a8a08f09d37b73795649038408b5f33", // wantedHash
				false,                              // print
				2),
			want: "", // "c"
		},
		// "not_found_maxsize_unlimited": {
		// 	// NOTE: This one is deadlocking, need to fix this !!
		// 	finder: NewHashFinder(
		// 		-1, // numCombs
		// 		-1, // maxSize
		// 		0, // minSize
		// 		"def", // charSet
		// 		"e2fc714c4727ee9395f324cd2e7f331f", // wantedHash
		// 		false, // print
		// 		2),
		// 	want: "", // "abcd"
		// },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := tc.finder.FindParallel()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
