package bcs

import (
	"fmt"
)

type Object struct {
}

func (this *Object) String() string {
	return fmt.Sprint(&this)
}
