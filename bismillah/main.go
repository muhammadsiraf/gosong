// package main

// import "fmt"

// func main() {
// 	test := 2
// 	len1 := 5
// 	len2 := 4
// 	arrayInput := []int{9, 6, -53, 32, 16}
// 	arrayInput2 := []int{3, -1, 1, 14}
// 	testCase := []interface{}{test, len1, arrayInput, len2, arrayInput2}
// 	squareRunTest(testCase)
// }

// func squareRunTest(testCase []interface{}) {
// 	lenTest := testCase[0].(int)
// 	limitTest := int(lenTest) * 2
// 	run := 2
// 	runEach(run, limitTest, testCase)
// }

// func runEach(run int, limit int, testCase []interface{}) {
// 	if run > limit {
// 		return
// 	}
// 	hasil := square(testCase[run].([]int))
// 	fmt.Println(hasil)
// 	run += 2
// 	runEach(run, limit, testCase)
// }

// func square(input []int) int {
// 	i := 0
// 	hasil := count_square(i, input)
// 	return hasil
// }

// func count_square(i int, input []int) int {
// 	var hasil int
// 	if i == len(input) {
// 		return hasil
// 	}
// 	if input[i] > 0 {
// 		hasil = input[i] * input[i]
// 	}
// 	i++
// 	hasil = hasil + count_square(i, input)
// 	return hasil
// }
