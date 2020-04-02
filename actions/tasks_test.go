package actions

import (
	"net/http"
	"time"

	"github.com/larrymjordan/tasks/models"

	"github.com/gobuffalo/httptest"
	"github.com/gofrs/uuid"
)

func (as *ActionSuite) Test_Tasks_Create() {
	task := models.Task{
		Description: "Do this",
		RequestedBy: "test",
		ExecutedBy:  "larry",
	}

	resp := as.JSON("/tasks").Post(task)
	as.Equal(http.StatusCreated, resp.Code)

	createdTask := models.Task{}
	resp.Bind(&createdTask)

	as.NotEqual(uuid.UUID{}, createdTask.ID)
	as.False(createdTask.CreatedAt.IsZero())
	as.False(createdTask.UpdatedAt.IsZero())

	as.Len(tasksStorage, 1)
	as.Equal(task.Description, createdTask.Description)
	as.Equal(task.RequestedBy, createdTask.RequestedBy)
	as.Equal(task.ExecutedBy, createdTask.ExecutedBy)
}

func (as *ActionSuite) Test_Tasks_List() {
	tasks := as.exercise(func() *httptest.JSONResponse {
		return as.JSON("/tasks").Get()
	})
	as.Len(tasks, 2)
}

func (as *ActionSuite) Test_Tasks_Done() {
	tasks := as.exercise(func() *httptest.JSONResponse {
		return as.JSON("/tasks/done").Get()
	})
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_Tasks_Pending() {
	tasks := as.exercise(func() *httptest.JSONResponse {
		return as.JSON("/tasks/pending").Get()
	})
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_Tasks_Range() {
	tasks := as.exercise(func() *httptest.JSONResponse {
		twoDaysAgo := time.Now().Add(-24 * 2 * time.Hour)
		return as.JSON("/tasks/range/%s/%s", twoDaysAgo.Add(-2*time.Hour).Format(time.RFC3339), twoDaysAgo.Add(2*time.Hour).Format(time.RFC3339)).Get()
	})
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_Tasks_Requester() {
	tasks := as.exercise(func() *httptest.JSONResponse {
		return as.JSON("/tasks/requester/%s", "Bob").Get()
	})
	as.Len(tasks, 1)
}

func (as *ActionSuite) Test_Tasks_Executer() {
	tasks := as.exercise(func() *httptest.JSONResponse {
		return as.JSON("/tasks/executer/%s", "Rob").Get()
	})
	as.Len(tasks, 1)
}

func (as *ActionSuite) exercise(fn func() *httptest.JSONResponse) []models.Task {
	setup()

	resp := fn()
	as.Equal(http.StatusOK, resp.Code)

	tasks := []models.Task{}
	resp.Bind(&tasks)

	return tasks
}

func setup() {
	tasksStorage = models.MemoryTasksStorage{}

	twoDaysAgo := time.Now().Add(-24 * 2 * time.Hour)
	tasksStorage.Add(&models.Task{
		Description: "Do this first!",
		IsDone:      true,
		CompletedOn: twoDaysAgo,
		RequestedBy: "Bob",
		ExecutedBy:  "Rob",
	})

	tasksStorage.Add(&models.Task{
		Description: "Do this next!",
		RequestedBy: "Rob",
		ExecutedBy:  "Bob",
	})
}
