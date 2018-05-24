package strings

import "testing"

func TestEqual(T *testing.T) {
	bs := []byte{'a', 'b'}
	ss := "ab"
	if ! unsafeEqual(ss, bs) {
		T.Error("must equal !!")
	}

	s2 := "abc"
	if unsafeEqual(s2, bs) {
		T.Error("equal??")
	}
}
