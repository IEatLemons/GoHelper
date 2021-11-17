package file

import (
	"fmt"
	"testing"
)

func TestFile_Read(t *testing.T) {
	f, err := Exc.ReadExc("file.xlsx")
	if err != nil {
		panic(err)
	}
	var Re = []Result{}
	Re = ReadFacebook(f, Re)
	for _, r := range Re {
		fmt.Printf(r.Address, r.Platform)
	}
}
