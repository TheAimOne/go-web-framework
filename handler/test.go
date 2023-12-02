package handler

import "github.com/TheAimOne/go-web-framework/model"

func CreateTestHandler(request interface{}) (*model.Response, error) {
	response := model.Response{
		Body: nil,
	}
	return &response, nil
}
