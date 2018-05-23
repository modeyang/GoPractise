package strings

import (
	"testing"
	"time"
)


func TestJoin(T *testing.T) {
	way := make(map[int]string, 5)
	way[0] = "fmt.Sprintf"
	way[1] = "+"
	way[2] = "strings.Join"
	way[3] = "bytes.Buffer"

	k := 4
	d := [5]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(way,10000, i)
	}
}
