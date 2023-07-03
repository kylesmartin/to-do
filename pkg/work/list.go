package work

// List contains a list of tasks
type List struct {
	Tasks []Task
}

// Add adds a task to the list
func (l *List) Add(task Task) {
	l.Tasks = append(l.Tasks, task)
}

// Remove removes the task at a given index from the task list
func (l *List) Remove(ind int) {
	l.Tasks = append(l.Tasks[:ind], l.Tasks[ind+1:]...)
}
