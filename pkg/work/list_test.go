package work_test

import (
	"testing"

	"github.com/kylesmartin/to-do/pkg/work"
	"github.com/stretchr/testify/assert"
)

const dummyTask = "Take out the trash"

// TestAdd verifies that we can add an item to a list
func TestAdd(t *testing.T) {
	list := work.List{
		Tasks: []work.Task{},
	}

	list.Add(
		work.Task{
			Description: dummyTask,
		},
	)

	assert.Equal(t, len(list.Tasks), 1)
	assert.Equal(t, list.Tasks[0].Description, dummyTask)
}

// TestValidRemove verifies that we can remove an item from a list if the index exists
func TestValidRemove(t *testing.T) {
	list := work.List{
		Tasks: []work.Task{
			{
				Description: dummyTask,
			},
		},
	}

	err := list.Remove(0)

	assert.Nil(t, err)
	assert.Equal(t, len(list.Tasks), 0)
}

// TestInvalidRemove verifies that we get an error if we input an invalid index
func TestInvalidRemove(t *testing.T) {
	list := work.List{
		Tasks: []work.Task{
			{
				Description: dummyTask,
			},
		},
	}

	err := list.Remove(1)
	assert.EqualError(t, err, "error removing task from list: index must be greater than 0 and less than 1")
	assert.Equal(t, len(list.Tasks), 1)

	err = list.Remove(-1)
	assert.EqualError(t, err, "error removing task from list: index must be greater than 0 and less than 1")
	assert.Equal(t, len(list.Tasks), 1)
}
