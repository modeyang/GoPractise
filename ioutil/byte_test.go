package ioutil

import (
	"bytes"
	"encoding/binary"
	"errors"
	"testing"
	"fmt"
	//"os"
)

func readString(buf *bytes.Buffer) (string, error) {
	var strlen int16
	err := binary.Read(buf, binary.BigEndian, &strlen)
	if err != nil {
		return "", err
	}
	if strlen == -1 {
		return "", nil
	}
	strbytes := make([]byte, strlen)
	n, err := buf.Read(strbytes)
	if (err != nil) || (n != int(strlen)) {
		return "", errors.New("string underflow")
	}
	return string(strbytes), nil
}

func TestPutVarint64(T *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
	}
	fmt.Printf("%x\n", buf[:len(buf)])
}

func TestReadBytes(T *testing.T) {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}

	r := bytes.NewReader(b)

	var data struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}

	err := binary.Read(r, binary.LittleEndian, &data)
	if err != nil {
		T.Error(err)
	}
	fmt.Println(data)
	fmt.Printf("%x \n", data.Mine)
}

func TestReadString(T *testing.T) {
	var b bytes.Buffer

	b.Write([]byte("hello "))
	//pb := bytes.NewBuffer(b.Bytes())
	//fmt.Fprintf(&b, "world!")
	//b.WriteTo(os.Stdout)

	fmt.Printf("%x, %d\n", b.Bytes()[0], len(b.Bytes()))
	//sb, err := readString(pb)
	//if err != nil {
	//	T.Error(err)
	//}
	//fmt.Println(sb)
}