package hzk

import (
	"reflect"
	"testing"
)

var str = "我"

func TestTranform(t *testing.T) {
	if b, err := transform([]byte(str)); err != nil {
		t.Fatalf("Transform err: %v", err)
	} else {
		// GBK CED2
		want := []byte{0xCE, 0xD2}
		if !reflect.DeepEqual(b, want) {
			t.Fatalf("Got %X, want %X", b, want)
		}
	}

}

func TestQuweima(t *testing.T) {
	gbk, err := transform([]byte(str))
	if err != nil {
		t.Fatalf("quweima error: %s", err)
	}
	qwm := quweima(gbk)
	want := []byte{46, 50}
	if !reflect.DeepEqual(qwm, want) {
		t.Fatalf("Got %X, want %X", qwm, want)
	}
}

func TestComputeOffset(t *testing.T) {
	gbk, err := transform([]byte(str))
	if err != nil {
		t.Fatalf("Quweima error: %s", err)
	}
	qwm := quweima(gbk)
	offset := computeOffset(qwm)
	var want int64 = 136928
	if offset != want {
		t.Fatalf("Got %X, want %X", qwm, want)
	}
}

func TestMartix(t *testing.T) {
	got, err := Matrix([]byte(str))
	if err != nil {
		t.Fatalf("Martix err: %v", err)
	}
	want := []byte{
		0x04, 0x80, 0x0E, 0xA0, 0x78, 0x90, 0x08, 0x90,
		0x08, 0x84, 0xFF, 0xFE, 0x08, 0x80, 0x08, 0x90,
		0x0A, 0x90, 0x0C, 0x60, 0x18, 0x40, 0x68, 0xA0,
		0x09, 0x20, 0x0A, 0x14, 0x28, 0x14, 0x10, 0x0C,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Martix got %v, want %v", got, want)
	}
}
