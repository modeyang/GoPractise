package ioutil

import (
	"testing"
	"crypto/sha1"
	"fmt"
)

func TestSha(t *testing.T){
	var s = "fdafda"
	var h = sha1.New()
	h.Write([]byte(s))

	var bs = h.Sum(nil)
	if bs == nil {
		t.Error("error sha")
	}
	fmt.Printf("%x, %d", bs, len(bs))
}