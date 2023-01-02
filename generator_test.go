package uuid7

import (
	"bytes"
	"testing"
	"time"
)

func TestUuid(t *testing.T) {
	u := New()
	now := time.Now().UnixMilli()

	v := u.Next()

	if v.Timestamp() < uint64(now)-1 || v.Timestamp() > uint64(now)+1 {
		t.Fatalf("Timestamp: expected %d, got %d", now, v.Timestamp())
	}

	vString := v.String()

	if len(vString) != 36 {
		t.Fatal("bad uuid string length")
	}

	if vString[8] != '-' || vString[13] != '-' || vString[18] != '-' || vString[23] != '-' {
		t.Fatal("bad uuid string format")
	}

	v2, err := Parse(vString)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(v[:], v2[:]) {
		t.Fatalf("uuids not equal: %s != %s", v, v2)
	}
}

func BenchmarkNext(b *testing.B) {
	u := New()

	for i := 0; i < b.N; i++ {
		_ = u.Next()
	}
}

func BenchmarkString(b *testing.B) {
	u := New().Next()

	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}

func BenchmarkParse(b *testing.B) {
	v := "017F21CF-D130-7CC3-98C4-DC0C0C07398F"

	for i := 0; i < b.N; i++ {
		_, err := Parse(v)
		if err != nil {
			b.Fatal(err)
		}
	}
}
