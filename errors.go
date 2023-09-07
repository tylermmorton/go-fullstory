package fullstory

import "fmt"

var (
	ErrTooManyProperties = fmt.Errorf("too many properties. the maximum allowed is 20")
)
