package work

import "fmt"

// List contains a list of tasks
type List struct {
	Tasks []Task
}

// Add adds a task to the list
func (l *List) Add(task Task) {
	l.Tasks = append(l.Tasks, task)
}

// Remove removes the task at a given index from the task list
func (l *List) Remove(ind int) error {
	if ind < 0 || ind >= len(l.Tasks) {
		return fmt.Errorf("error removing task from list: index must be greater than 0 and less than %d", len(l.Tasks))
	}
	l.Tasks = append(l.Tasks[:ind], l.Tasks[ind+1:]...)
	return nil
}
