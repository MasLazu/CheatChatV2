package model

type MakeGroupRequest struct {
	Name string `json:"name" validate:"required"`
}
