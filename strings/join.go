package strings

import (
	"strings"
	"time"
	"fmt"
	"bytes"
)

// https://gocn.io/question/265
// https://gocn.io/article/704
func benchmarkStringFunction(way map[int]string, n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	jv := []string{s, "[", v, "]"}

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join(jv,  "")
		case 3: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

	}
	d = time.Since(t0)
	if index == 3 {
		s = buf.String()
	}
	fmt.Printf("string len: %d\t", len(s))
	fmt.Printf("time of [%s]=\t %v\n", way[index], d)
	return d
}


