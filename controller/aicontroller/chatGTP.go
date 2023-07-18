package aicontroller

import (
	"context"
	"net/http"
	"net/url"
	cf "open_ai_chat/config"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

var (
	OPEN_AI_KEY = cf.Confs.GetString("chatGTP.ApiKey")
	OPEN_AI_URL = cf.Confs.GetString("chatGTP.URL")
)

func ChatGTPRequest(ctx *gin.Context) {

	config := openai.DefaultConfig(OPEN_AI_KEY)
	proxyUrl, err := url.Parse("http://localhost:7890")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors":  err.Error(),
			"message": "failed!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errors":  nil,
		"message": resp.Choices[0].Message.Content,
	})
}
