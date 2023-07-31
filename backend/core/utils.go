package core

import (
	"fmt"
)

func makeError(e error, tag string) error {
	if e == nil {
		return nil
	}

	return fmt.Errorf("%s: %v", tag, e)
}
