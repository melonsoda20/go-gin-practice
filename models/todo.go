package models

type CreateToDoReqDTO struct {
	Name       string `json:"name"`
	IsTaskDone bool   `json:"isTaskDone"`
}

type UpdateToDoReqDTO struct {
	Name       string `json:"name"`
	IsTaskDone bool   `json:"isTaskDone"`
}
