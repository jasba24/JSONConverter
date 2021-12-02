package JSONConverter

import (
	"io/ioutil"
	"encoding/json"
)

func Converter(class interface{}, path string) interface{} {
	file, err := ioutil.ReadFile(path)
	file_struct := class
	json.Unmarshal(file, &file_struct)

	return file_struct, err
}