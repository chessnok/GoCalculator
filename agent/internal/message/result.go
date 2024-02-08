package message

import "encoding/json"

type Result struct {
	Id     string  `json:"id"`
	Result float64 `json:"result"`
	IsErr  bool    `json:"is_err"`
	Error  string  `json:"error"`
}

func NewResult(id string, result float64, isErr bool, err string) *Result {
	return &Result{
		Id:     id,
		Result: result,
		IsErr:  isErr,
		Error:  err,
	}
}
func (r *Result) ToJSON() []byte {
	res, _ := json.Marshal(r)
	return res
}
