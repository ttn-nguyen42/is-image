package utils

import (
	"bytes"
	"io"
)

func ToBytes(file io.Reader) ([]byte, error) {
	buf, err := ToBuffer(file)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ToBuffer(file io.Reader) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
