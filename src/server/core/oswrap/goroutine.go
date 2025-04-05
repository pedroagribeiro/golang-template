package oswrap

import (
	"runtime"
	"strconv"
	"strings"
)

// Code based on
// https://zhimin-wen.medium.com/logging-for-concurrent-go-programs-abe46ef14d58
// and https://github.com/golang/net/blob/master/http2/gotrack.go#L51
func GetGoRoutineId() int {
	var buf [64]byte //in the golang source code, 64 seems to be enough
	var id int = 0
	var err error = nil

	n := runtime.Stack(buf[:], false)
	idFields := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))
	if len(idFields) > 0 {
		id, err = strconv.Atoi(idFields[0])
		if err != nil {
			return 0
		}
	}
	return id
}
