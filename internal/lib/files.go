package lib

import (
	"io"
	"os"
	"strings"
)

func ReadStringFile(fn string) ([]string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	strs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	res := strings.Split(string(strs), "\n")
	return res, nil
}
