package chanutil

//import "testing"

type readOps struct {
	key 	int
	resp 	chan int
}

type writeOps struct {
	key 	int
	value 	int
	resp 	chan bool
}
//
//func TestOps(T *testing.T) {
//	var ops int64 = 0
//
//	var reads chan *readOps = make(chan *readOps, 1)
//	var writes chan *writeOps = make(chan *writeOps, 1)
//
//	go func() {
//		var stats = map[int]int
//	}()
//}
