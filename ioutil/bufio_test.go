package ioutil

import (
	"testing"
	"bufio"
	"os"
	"fmt"
)

func TestBufio(T *testing.T) {
	f, err := os.Open("bufio_test.go")
	if err != nil {
		T.Error(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		T.Error(err)
	}

}