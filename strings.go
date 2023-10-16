package utils

import "strings"

func TrimUserNamePrefix(v string) string {
	return strings.TrimPrefix(v, "RIFAS\\")
}
