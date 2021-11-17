package file

import (
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	File
}

func (C *Excel) ReadExc(filename string) (*excelize.File, error) {
	return excelize.OpenFile(filename)
}
