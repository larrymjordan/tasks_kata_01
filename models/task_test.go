package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_MemoryTasksStorageWrite(t *testing.T) {
	ts := MemoryTasksStorage{}
	ts.Add(&Task{})
	a := require.New(t)
	a.Equal(1, len(ts))
}

func Test_MemoryTasksStorageTasksDone(t *testing.T) {
	ts := MemoryTasksStorage{}
	ts.Add(&Task{Description: "Task to be done"})
	ts.Add(&Task{Description: "Task that was done", IsDone: true})

	tasks := ts.TasksDone()
	a := require.New(t)
	a.Equal(1, len(tasks))
	a.Equal(ts[1].Description, tasks[0].Description)
}

func Test_MemoryTasksStorageTasksPending(t *testing.T) {
	ts := MemoryTasksStorage{}
	ts.Add(&Task{Description: "Task to be done"})
	ts.Add(&Task{Description: "Task that was done", IsDone: true})

	tasks := ts.TasksPending()
	a := require.New(t)
	a.Equal(1, len(tasks))
	a.Equal(ts[0].Description, tasks[0].Description)
}

func Test_MemoryTasksStorageTaksInRange(t *testing.T) {
	ts := MemoryTasksStorage{}
	twoDaysAgo := time.Now().Add(-48 * time.Hour)
	ts.Add(&Task{Description: "Task to be done", CompletedOn: twoDaysAgo})
	ts.Add(&Task{Description: "Task that was done", IsDone: true, CompletedOn: time.Now()})

	tasks := ts.TasksInRange(twoDaysAgo.Add(-24*time.Hour), twoDaysAgo.Add(24*time.Hour))
	a := require.New(t)
	a.Len(tasks, 1)
	a.Equal(ts[0].Description, tasks[0].Description)
}

func Test_MemoryTasksStorageTasksRequestedBy(t *testing.T) {
	ts := MemoryTasksStorage{}
	ts.Add(&Task{RequestedBy: "Bob"})
	ts.Add(&Task{RequestedBy: "Rob"})

	tasks := ts.TasksRequestedBy("Bob")
	a := require.New(t)
	a.Len(tasks, 1)
	a.Equal("Bob", tasks[0].RequestedBy)
}

func Test_MemoryTasksStorageTasksExecutedBy(t *testing.T) {
	ts := MemoryTasksStorage{}
	ts.Add(&Task{ExecutedBy: "Bob"})
	ts.Add(&Task{ExecutedBy: "Rob"})

	tasks := ts.TasksExecutedBy("Rob")
	a := require.New(t)
	a.Len(tasks, 1)
}
