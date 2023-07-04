package cmd

import (
	"fmt"

	"github.com/kylesmartin/to-do/pkg/work"
)

// Validation defines the function that validates a list
type Validation func(list *work.List) error

// IsNonEmpty requires the work list to have elements
func IsNonEmpty(list *work.List) error {
	if len(list.Tasks) == 0 {
		return fmt.Errorf("to-do list is empty")
	}
	return nil
}
