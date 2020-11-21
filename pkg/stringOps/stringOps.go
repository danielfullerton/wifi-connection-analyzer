package stringOps

import (
	"strings"
)

func GetPrefixedFileLocation (filePath, prefix string) string {
	filePathSlice := strings.Split(filePath, "/")
	path := filePathSlice[:len(filePathSlice)-1]
	csvFile := prefix + "-" + filePathSlice[len(filePathSlice)-1]
	path = append(path, csvFile)
	return strings.Join(path, "/")
}
