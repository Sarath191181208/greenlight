package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	runTimeStr := fmt.Sprintf("%d mins", r)
	quotedStr := strconv.Quote(runTimeStr)
	return []byte(quotedStr), nil
}
