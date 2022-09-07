package sum_test

import (
	"fmt"

	"github.com/zxhoper/sum"
)

func ExampleInts() {

	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Println("sum of one to five is", s)
	// Output:
	// sum of one to five is 15
}

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

/*
go-testing-justforfunc-#16-unit-testing-HTTP-servers-RZ20561-2022-09-07_083801/go-testing-03-example/
[20220907T123927]$ go test -v
=== RUN   TestInts
=== RUN   TestInts/one_to_five
=== RUN   TestInts/no_numbers
=== RUN   TestInts/one_and_minus_one
--- PASS: TestInts (0.00s)
    --- PASS: TestInts/one_to_five (0.00s)
    --- PASS: TestInts/no_numbers (0.00s)
    --- PASS: TestInts/one_and_minus_one (0.00s)
=== RUN   ExampleInts
--- PASS: ExampleInts (0.00s)
PASS
ok  	sum_test	0.002s


*/
