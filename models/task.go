package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Task is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	RequestedBy string    `json:"requested_by" db:"requested_by"`
	ExecutedBy  string    `json:"executed_by" db:"executed_by"`
	IsDone      bool      `json:"is_done" db:"is_done"`
	CompletedOn time.Time `json:"completed_on" db:"completed_on"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Task) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tasks is not required by pop and may be deleted
type MemoryTasksStorage []Task

func (ts *MemoryTasksStorage) Add(task *Task) {
	task.ID, _ = uuid.NewV4()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	*ts = append(*ts, *task)
}

func (ts *MemoryTasksStorage) TasksDone() []Task {
	return ts.filter(func(task Task) bool {
		return task.IsDone
	})
}

func (ts *MemoryTasksStorage) TasksPending() []Task {
	return ts.filter(func(task Task) bool {
		return !task.IsDone
	})
}

func (ts *MemoryTasksStorage) TasksInRange(from time.Time, to time.Time) []Task {
	return ts.filter(func(task Task) bool {
		return task.CompletedOn.After(from) && task.CompletedOn.Before(to)
	})
}

func (ts *MemoryTasksStorage) TasksRequestedBy(requester string) []Task {
	return ts.filter(func(task Task) bool {
		return task.RequestedBy == requester
	})
}

func (ts *MemoryTasksStorage) TasksExecutedBy(executer string) []Task {
	return ts.filter(func(task Task) bool {
		return task.ExecutedBy == executer
	})
}

func (ts MemoryTasksStorage) filter(fn func(task Task) bool) []Task {
	var tasks []Task
	for _, t := range ts {
		if fn(t) {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// String is not required by pop and may be deleted
func (t MemoryTasksStorage) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
		&validators.StringIsPresent{Field: t.RequestedBy, Name: "RequestedBy"},
		&validators.StringIsPresent{Field: t.ExecutedBy, Name: "ExecutedBy"},
		&validators.TimeIsPresent{Field: t.CompletedOn, Name: "CompletedOn"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Task) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
