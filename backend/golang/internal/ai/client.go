package ai

import (
    "bytes"
	"encoding/json"
    "fmt"
	"io"
	"net/http"
	"strings"
)


type client struct {
   apiKey string
   BaseURL string
   model string
}

func NewClient(apiKey, baseURL,model string) *client {
   retun &Client{
    apiKey: apiKey,
    BaseURL: baseURL,
    model: model,
   }
}

type chatMessage struct {
   Role string `json:"role"`
   Content string `json:"content`
}

type chatRequest struct {
	Model string `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}