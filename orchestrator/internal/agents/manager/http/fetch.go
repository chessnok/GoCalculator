package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"io"
	"net/http"
	"time"
)

var timeout = 3 * time.Second

type Response struct {
	Task string `json:"Task"`
}

func fetch(ctx context.Context, url, method, contentType string, timeout time.Duration, data []byte) (int, string) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(data))
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if err != nil {
		return -1, ""
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, ""
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return -1, ""
	}
	return resp.StatusCode, response.Task

}
func Ping(ctx context.Context, url string) (bool, string) {
	s, res := fetch(ctx, url+"/currentOperation", "GET", "", timeout, nil)
	return s >= 200 && s < 300, res
}

func NewConfig(ctx context.Context, url string, config *calculator.Config) (bool, string) {
	c, _ := json.Marshal(config)
	s, res := fetch(ctx, url+"/updateConfig", "POST", "application/json", timeout, c)
	return s >= 200 && s < 300, res
}
