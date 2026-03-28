package domain

import (
	"time"

	chatDomain "chat-bot/src/common-service/chat/domain"
	core_values "chat-bot/src/core/values"
	"chat-bot/src/initial/default_data"
)

type RequestBody struct {
	Model    string           `json:"model"`
	Options  RequestOptions   `json:"options"`
	Messages []RequestMessage `json:"messages"`
	Stream   bool             `json:"stream"`
}

func (r *RequestBody) SetBody(roomHistory chatDomain.Room, reqChat RequestChat) {
	r.Model = core_values.OllamaModel
	r.Options = RequestOptions{
		Temperature: 0.7,
		TopP:        0.8,
		TopK:        60,
	}
	r.Stream = false
	r.Messages = []RequestMessage{}

	// system
	r.Messages = append(r.Messages, RequestMessage{
		Role:    core_values.OllamaRoleSystem,
		Content: roomHistory.AiSystem,
	})

	// user, assistant history
	for _, chat := range roomHistory.ChatList {
		role := core_values.OllamaRoleUser
		if chat.CreatorID == default_data.GetAIWorker().ID {
			role = core_values.OllamaRoleAssistant
		}
		reqMsg := RequestMessage{
			Role:    role,
			Content: chat.Content,
		}
		r.Messages = append(r.Messages, reqMsg)
	}

	// 가장 최근 user 채팅
	r.Messages = append(r.Messages, RequestMessage{
		Role:    reqChat.Role,
		Content: reqChat.Content,
	})

}

type RequestOptions struct {
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	TopK        uint    `json:"top_k"`
}

// Role 종류
// "system" — 최상위 규칙 (컨텍스트 설정)
// "user" — 사용자의 입력 (모델 답변 생성 기준)
// "assistant" — 모델의 이전 응답 (맥락 유지용)
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

type RequestChat struct {
	Role    string
	Content string
}
