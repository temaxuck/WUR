package utils

import (
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ContainsStr(xs []string, x string) bool {
	for _, v := range xs {
		if strings.ToUpper(v) == strings.ToUpper(x) {
			return true
		}
	}

	return false
}
