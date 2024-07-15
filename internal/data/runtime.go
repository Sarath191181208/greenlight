package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshalJSON() ([]byte, error) {
	runTimeStr := fmt.Sprintf("%d mins", r)
	quotedStr := strconv.Quote(runTimeStr)
	return []byte(quotedStr), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// incoming json = "<runtime> mins"
	// - remove the quotes
	val, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

  // Split the string to isolate the part containing the number.
	parts := strings.Split(val, " ")

  // check if the string is in expected format
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

  // str -> int32
	runtimeVal, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

  // Convert the int32 to a Runtime type and assign this to the receiver. Note that we use
  // use the * operator to deference the receiver (which is a pointer to a Runtime
  // type) in order to set the underlying value of the pointer.
	*r = Runtime(runtimeVal)
	return nil
}
