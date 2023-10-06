package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int16

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quoytedJSONValue := strconv.Quote(jsonValue)

	return []byte(quoytedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(value []byte) error {
	unqoutedValue, err := strconv.Unquote(string(value))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unqoutedValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 16)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}
