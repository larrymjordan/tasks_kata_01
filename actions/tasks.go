package actions

import (
	"net/http"
	"tasks/models"
	"time"

	"github.com/gobuffalo/buffalo"
)

var tasksStorage = models.MemoryTasksStorage{}

// TasksCreate default implementation.
func TasksCreate(c buffalo.Context) error {
	var task models.Task
	c.Bind(&task)

	tasksStorage.Add(&task)
	return c.Render(http.StatusCreated, r.JSON(task))
}

// TasksList default implementation.
func TasksList(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(tasksStorage))
}

// TasksDone default implementation.
func TasksDone(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(tasksStorage.TasksDone()))
}

// TasksPending default implementation.
func TasksPending(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(tasksStorage.TasksPending()))
}

// TasksRange default implementation.
func TasksRange(c buffalo.Context) error {
	from, _ := time.Parse(time.RFC3339, c.Param("from"))
	to, _ := time.Parse(time.RFC3339, c.Param("to"))
	return c.Render(http.StatusOK, r.JSON(tasksStorage.TasksInRange(from, to)))
}

// TasksRequester default implementation.
func TasksRequester(c buffalo.Context) error {
	return c.Render(200, r.JSON(tasksStorage.TasksRequestedBy(c.Param("requester"))))
}

// TasksExecuter default implementation.
func TasksExecuter(c buffalo.Context) error {
	return c.Render(200, r.JSON(tasksStorage.TasksExecutedBy(c.Param("executer"))))
}
