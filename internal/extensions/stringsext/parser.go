package stringsext

import "strings"

func ParseKeyValuePair(pair string, separator byte) (key string, value string) {
	if index := strings.IndexByte(pair, separator); index >= 0 {
		return pair[:index], pair[index+1:]
	}

	return
}
