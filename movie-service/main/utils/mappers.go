package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func MarshalJSON(content interface{}) []byte {
	result, _ := json.Marshal(content)
	return result
}

func WriteBody(content interface{}) io.Reader {
	return ioutil.NopCloser(bytes.NewBuffer(MarshalJSON(content)))
}
