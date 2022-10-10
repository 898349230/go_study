package main

// flag 包用于解析命令行参数
import (
	"bufio"
	"flag"
	"fmt"
	"hello/chapter2/sorter/sorter/algorithms/bubblesort"
	"hello/chapter2/sorter/sorter/algorithms/qsort"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("failed to open the input file", infile)
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		// 转换字符数组为字符串
		str := string(line)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed to create file", outfile)
		return err
	}

	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err != nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("no algorithm", *algorithm)
		}
		t2 := time.Now()
		fmt.Println("Sorting algorithm", *algorithm, " costs ", t2.Sub(t1))
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
