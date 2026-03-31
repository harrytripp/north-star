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

// Generate a response
// Returns a response ID
func Response() {
	client := openai.NewClient(
		option.WithBaseURL("http://host.internal:8080/v1"),
		option.WithAPIKey("not-needed"),
	)
	response, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "Ministral-3-8B-Instruct-2512-Q8_0.gguf",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String("Say this is a test")},
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", response.ID)
}

// TODO Retreive resonse content
func Retrieve() {
	client := openai.NewClient(
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Responses.Get(
		context.TODO(),
		"resp_677efb5139a88190b512bc3fef8e535d",
		responses.ResponseGetParams{},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", response.ID)
}
