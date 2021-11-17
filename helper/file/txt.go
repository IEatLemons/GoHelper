package file

import (
	"bufio"
	"os"
)

type TXT struct {
	File
}

func (C *TXT) ReadTXT(filename string) (br *bufio.Reader, err error) {
	fi, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fi.Close()
	br = bufio.NewReader(fi)
	return
}
