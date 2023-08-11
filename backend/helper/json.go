package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MasLazu/CheatChatV2/model"
)

func ReadRequestBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func WriteResponse(writer http.ResponseWriter, code int, stauts string, responseBody interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	encoder := json.NewEncoder(writer)
	response := model.JsonResponse{
		Code:   code,
		Status: stauts,
		Data:   responseBody,
	}
	encoder.Encode(response)
}
