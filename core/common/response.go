package common

import (
	"github.com/gofiber/fiber/v2"
)


type StatusType string

const (
	StatusSuccess StatusType = "Success"
	StatusFailed StatusType = "Failed"
)

func NewResponse(status StatusType, message string, additionalProps map[string]interface{}) fiber.Map {
	response := fiber.Map{
		"status": status,
		"message": message, 
	}

	for key, value := range additionalProps {
		response[key] = value
	}

	return response
}

func NewSuccessResponse(message string, data ...interface{}) fiber.Map {
	if len(data) == 0 {
		return NewResponse(StatusSuccess, message, nil)
	}

	return NewResponse(StatusSuccess, message, map[string]interface{}{
		"data": data[0],
	})
}

func NewFailedResponse(message string, err interface{}) fiber.Map {
	return NewResponse(StatusFailed, message, map[string]interface{}{
		"error": err,
	})
}