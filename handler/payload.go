package handler

type Payload struct {
	Code int `json:"-"`
	Data any `json:"data"`
}
