package services

import "gin-todo-app/models"

func ValidateCreateToDo(req models.CreateToDoReqDTO) models.ValidationResults {
	validationResult := models.ValidationResults{
		IsSuccess: true,
	}
	if len(req.Name) == 0 {
		validationResult.Message = append(validationResult.Message, "Please provide a name for the todo")
	}

	if len(validationResult.Message) > 0 {
		validationResult.IsSuccess = false
	}

	return validationResult
}
