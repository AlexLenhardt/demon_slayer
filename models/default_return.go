package models

type HttpReturn struct {
	Object any    `json:"object"`
	Error  *error `json:"error"`
}
