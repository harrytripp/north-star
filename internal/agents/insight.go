package agents

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
)

// Test basic server connection
func ConnectionDebug() {
	resp, err := http.Post(
		"http://host.internal:8080/v1/responses", // host.internal is OrbStack specific for reaching local host
		"application/json",
		strings.NewReader(`{"model":"test","input":[]}`),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// Generate a response. Returns an unformatted string or an error.
func Response() (string, error) {
	ctx := context.Background()
	client := openai.NewClient(
		option.WithBaseURL("http://host.internal:8080/v1"),
		option.WithAPIKey("not-needed"),
	)

	input := "Say this is a test"
	sysprompt := "You must respond only in a sarcastic haiku."

	fmt.Printf("\nGenerating response...")
	response, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Model:        "Ministral-3-8B-Instruct-2512-Q8_0.gguf",
		Instructions: openai.String(sysprompt),
		Input:        responses.ResponseNewParamsInputUnion{OfString: openai.String(input)},
	})

	if err != nil {
		return "", err // Return error instead of panicking
	}

	return response.OutputText(), nil // Return output unformatted output text on success
}
