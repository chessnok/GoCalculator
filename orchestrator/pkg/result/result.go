package result

type Result struct {
	// Id - unique identifier
	Id string `json:"id" xml:"id"`
	//	Result - result of the task
	Result float64 `json:"result" xml:"result"`
}
