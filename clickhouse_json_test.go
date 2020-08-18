package clickhouseJSONEachRow

import (
	"bytes"
	"io"
	"testing"
)

var (
	wantJson = []byte(`[{"SearchPhrase":"дизайн штор","count()":"1064"}
,{"SearchPhrase":"баку","count()":"1000"}
,{"SearchPhrase":"","count":"8267016"}
]`)

	fromJson = []byte(`{"SearchPhrase":"дизайн штор","count()":"1064"}
{"SearchPhrase":"баку","count()":"1000"}
{"SearchPhrase":"","count":"8267016"}
`)
)

func TestCopy(t *testing.T) {
	for i := 2; i < 30; i++ {
		buf1 := bytes.NewBuffer(make([]byte, 0, 145))
		buf2 := bytes.NewBuffer(fromJson)

		count, err := Copy(buf1, buf2, i)
		if err != nil && err != io.EOF {
			t.Fatal(err)
			return
		}

		if count != 145 {
			t.Fatal("wrong copy count elements")
		}
		if bytes.Compare(buf1.Bytes(), wantJson) != 0 {
			t.Fatal("not same")
		}
	}
}

func BenchmarkCopyTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf1 := bytes.NewBuffer(make([]byte, 0, 145))
		buf2 := bytes.NewBuffer(fromJson)

		count, err := Copy(buf1, buf2, 4)
		if err != nil && err != io.EOF {
			b.Fatal(err)
			return
		}
		if count != 145 {
			b.Fatal("wrong copy count elements")
		}

		if bytes.Compare(buf1.Bytes(), wantJson) != 0 {
			b.Fatal("not same")
		}
	}
}
