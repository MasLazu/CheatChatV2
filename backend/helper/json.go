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

func WriteInternalServerError(writer http.ResponseWriter) {
	WriteResponse(writer, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", model.MessageResponse{Message: "someting went wrong"})
}

func WriteNotFoundError(writer http.ResponseWriter) {
	WriteResponse(writer, http.StatusNotFound, "NOT_FOUND", model.MessageResponse{Message: "not found"})
}

func WriteUnauthorizedError(writer http.ResponseWriter) {
	WriteResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", model.MessageResponse{Message: "unauthorized"})
}

func WriteBadRequestError(writer http.ResponseWriter) {
	WriteResponse(writer, http.StatusBadRequest, "BAD_REQUEST", model.MessageResponse{Message: "bad request"})
}

func WriteOk(writer http.ResponseWriter, responseBody any) {
	WriteResponse(writer, http.StatusOK, "OK", responseBody)
}
