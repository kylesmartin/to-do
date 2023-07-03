package work_test

import (
	"os"
	"path"
	"testing"

	"github.com/kylesmartin/to-do/pkg/work"
	"github.com/stretchr/testify/assert"
)

// TestLoadEmptyList verifies that LoadList returns an empty struct if no file exists
func TestLoadEmptyList(t *testing.T) {
	fileNameAndPath := path.Join(os.TempDir(), "to-do-1.json")

	list, err := work.LoadList(fileNameAndPath)

	assert.Nil(t, err)
	assert.Equal(t, list, &work.List{})
}

// TestSaveAndLoadList verifies that SaveList stores the struct properly and that LoadList returns the correct struct
func TestSaveAndLoadList(t *testing.T) {
	fileNameAndPath := path.Join(os.TempDir(), "to-do-2.json")

	list := work.List{
		Tasks: []work.Task{
			{
				Description: dummyTask,
			},
		},
	}

	err := work.SaveList(fileNameAndPath, &list)
	assert.Nil(t, err)

	newList, err := work.LoadList(fileNameAndPath)
	assert.Nil(t, err)

	assert.Equal(t, list, *newList)
}
