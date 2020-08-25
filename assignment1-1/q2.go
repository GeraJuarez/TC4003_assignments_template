package cos418_hw1_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	total := 0
	for num := range nums {
		total += num
	}
	out <- total
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	integers, err := readInts(file)
	checkError(err)

	chNums := make(chan int, len(integers)/num)
	chOut := make(chan int, num)

	for i := 0; i < num; i++ {
		go sumWorker(chNums, chOut)
	}

	for _, val := range integers {
		chNums <- val
	}
	close(chNums)

	total := 0
	for i := 0; i < num; i++ {
		total += <-chOut
	}
	close(chOut)

	return total
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
