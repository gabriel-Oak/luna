package thinking

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func ThinkAndRespond(message string) {
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Fatalf("Falha ao inicia o ollama: %v", err)
	}

	ctx := context.Background()
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a AI named Luna."),
		llms.TextParts(llms.ChatMessageTypeSystem, "You ve been developed by Gabriel Oak."),
		llms.TextParts(llms.ChatMessageTypeSystem, "Você responde em português e é prestativa e gentil."),
		llms.TextParts(llms.ChatMessageTypeHuman, message),
	}
	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))

	if err != nil {
		log.Fatal(err)
	}
	_ = completion
	fmt.Printf("\n")
}
