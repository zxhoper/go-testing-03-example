package sum_test

import (
	"testing"

	"github.com/zxhoper/sum"
)

// github.com/zxhoper/sum/sum.go
// package sum
// // Ints returns the sum of a list of intergers.
// func Ints(vs ...int) int {
// 	return ints(vs)
// }
// func ints(vs []int) int {
// 	if len(vs) == 0 {
// 		return 0
// 	}
// 	return ints(vs[1:]) + vs[0]
// }

// Use sub tests: t.Run(name, func)
// Now we can test parts of the table of tests
// go test -run Ints/one   => only run sub-tests which name matches *one*
func TestInts(t *testing.T) {
	// tt => Table of Test case
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	// range Test Cases
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			s := sum.Ints(tc.numbers...)
			if s != tc.sum {
				t.Fatalf("sum of %v should be %v; got %v", tc.numbers, tc.sum, s)
			}
		})
	}

}
