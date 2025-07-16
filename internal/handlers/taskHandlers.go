package handlers

import (
	"Tasks/internal/taskService"
	"Tasks/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	if err := h.service.DeleteTask(taskID); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:   &tsk.ID,
			Name: &tsk.Name,
		}
		response = append(response, task)
	}
	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	taskToUpdate := taskService.TaskRequest{
		Name: *taskRequest.Name,
	}
	updatedTask, err := h.service.UpdateTask(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:   &updatedTask.ID,
		Name: &updatedTask.Name,
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

	taskRequest := request.Body
	taskToCreate := taskService.TaskRequest{
		Name: *taskRequest.Name,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:   &createdTask.ID,
		Name: &createdTask.Name,
	}

	return response, nil
}
