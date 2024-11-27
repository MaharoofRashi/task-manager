package core

import "errors"

type Task struct {
	ID          string
	Title       string
	Description string
	Status      string
	UserID      string
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
