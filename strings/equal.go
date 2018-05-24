package strings

import "unsafe"

func unsafeEqual(a string, b []byte) bool {
	bbp := *(*string)(unsafe.Pointer(&b))
	return a == bbp
}

