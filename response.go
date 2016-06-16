package main

import (
	"encoding/json"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func encodeResponse(status bool, message string, data ...interface{}) []byte {
	res := Response{
		Status:  status,
		Message: message,
	}

	if len(data) > 0 {
		res.Data = data[0]
	}

	result, _ := json.Marshal(res)
	return result
}
