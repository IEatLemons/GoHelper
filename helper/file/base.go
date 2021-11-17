package file

import "io/ioutil"

type File struct{}

func (f *File) WriteFile(filename string, content []byte) error {
	err := ioutil.WriteFile(filename, content, 0644)
	return err
}
