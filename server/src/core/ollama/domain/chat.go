package domain

import "time"

type RequestBody struct {
	Model    string           `json:"model"`
	Messages []RequestMessage `json:"messages"`
	Stream   bool             `json:"stream"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseBody struct {
	Model           string          `json:"model"`
	CreateAt        time.Time       `json:"created_at"`
	Message         ResponseMessage `json:"message"`
	Done            bool            `json:"done"`
	DoneReason      string          `json:"done_reason"`
	PromptEvalCount uint            `json:"prompt_eval_count"`
	EvalCount       uint            `json:"eval_count"`
}

type ResponseMessage struct {
	Role     string `json:"role"`
	Content  string `json:"content"`
	Thinking string `json:"thinking"`
}
