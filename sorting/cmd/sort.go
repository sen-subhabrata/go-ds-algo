package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	s "github.com/go-ds-algo/sorting/pkg/sort"
)

func sort(a []int) {
	fmt.Printf("Unsorted: %v\n", a)
	s.CountingSort(a)
	fmt.Printf("Sorted:   %v\n", a)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int

	for i := 0; i < n; i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int(sItemTemp)
		s = append(s, sItem)
	}

	sort(s)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
