package service

import (
	"errors"
	"io"
	"unicode/utf8"
)

const (
	nextLine = 10
	comma    = 44
)

var (
	openSquareBracket  = [1]byte{91}
	closeSquareBracket = [1]byte{93}
)

func Copy(dst io.Writer, src io.Reader, bufSize int) (written int64, err error) {
	if bufSize < 2 {
		return 0, errors.New(" buffer size is small")
	}

	wOpen, wErrOpen := dst.Write(openSquareBracket[:])
	if wErrOpen != nil {
		return int64(wOpen), err
	}
	written += int64(wOpen)
	var isNextLine = false
	buf := make([]byte, bufSize)
	bufTemp := make([]byte, 0, bufSize+bufSize/20)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			for i := 0; i < nr; i++ {
				if isNextLine {
					bufTemp = append(bufTemp, comma)
					isNextLine = false
				}
				if buf[i] == nextLine && utf8.Valid([]byte{buf[i]}) {
					isNextLine = true
				}
				bufTemp = append(bufTemp, buf[i])
			}
			nw, ew := dst.Write(bufTemp)
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if len(bufTemp) != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		bufTemp = bufTemp[:0]
	}
	nw, err := dst.Write(closeSquareBracket[:])
	return written + int64(nw), err
}
