package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	chatDomain "chat-bot/src/common-service/chat/domain"
	"chat-bot/src/core/ollama/domain"
	core_values "chat-bot/src/core/values"
)

type ollamaService struct{}

func NewOllamaService() OllamaService {
	return &ollamaService{}
}

func (s *ollamaService) Chat(reqChat domain.RequestChat, roomHistory chatDomain.Room) *domain.ResponseBody {
	chatUrl := fmt.Sprintf("%s/chat", core_values.OllamaApiURL)

	reqBody := domain.RequestBody{}
	reqBody.SetBody(roomHistory, reqChat)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(reqBody); err != nil {
		log.Printf("error: %s\n", err)
		return nil

	}
	request, err := http.NewRequest(http.MethodPost, chatUrl, &body)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv(core_values.EnvOllamaApiKey)))

	client := &http.Client{
		Timeout: 3 * time.Minute,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil
	}

	var result *domain.ResponseBody
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Printf("error: %s\n", err)
	}
	return result
}

func (s *ollamaService) Summarize(roomHistory chatDomain.Room) *domain.ResponseBody {
	chatUrl := fmt.Sprintf("%s/chat", core_values.OllamaApiURL)

	reqBody := domain.RequestBody{}
	reqBody.SetSummaryBody(roomHistory)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(reqBody); err != nil {
		log.Printf("error: %s\n", err)
		return nil

	}
	request, err := http.NewRequest(http.MethodPost, chatUrl, &body)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv(core_values.EnvOllamaApiKey)))

	client := &http.Client{
		Timeout: 3 * time.Minute,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil
	}

	var result *domain.ResponseBody
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Printf("error: %s\n", err)
	}
	return result
}
