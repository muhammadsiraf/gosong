// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int(nTemp)
// 	var testCase []interface{}
// 	testCase = append(testCase, n)

// 	for i := 0; i < int(n); i++ {
// 		lenTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)
// 		arrLen := int32(lenTemp)

// 		testCase = append(testCase, arrLen)

// 		arrayTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
// 		var arrayFinal []int

// 		for j := 0; j < int(arrLen); j++ {
// 			itemTemp, err := strconv.ParseInt(arrayTemp[j], 10, 64)
// 			checkError(err)
// 			item := int(itemTemp)
// 			arrayFinal = append(arrayFinal, item)
// 		}

// 		testCase = append(testCase, arrayFinal)
// 	}

// 	squareRunTest(testCase)

// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
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
