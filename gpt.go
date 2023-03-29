// Package gpt provides a client for the OpenAI GPT-3 API
package gpt

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Define GPT-3 Engine Types
const (
	TextAda001Engine     = "text-ada-001"     // TextAda001Engine Text Ada 001
	TextBabbage001Engine = "text-babbage-001" // TextBabbage001Engine Text Babbage 001
	TextCurie001Engine   = "text-curie-001"   // TextCurie001Engine Text Curie 001
	TextDavinci001Engine = "text-davinci-001" // TextDavinci001Engine Text Davinci 001
	TextDavinci002Engine = "text-davinci-002" // TextDavinci002Engine Text Davinci 002
	TextDavinci003Engine = "text-davinci-003" // TextDavinci003Engine Text Davinci 003
	AdaEngine            = "ada"              // AdaEngine Ada
	BabbageEngine        = "babbage"          // BabbageEngine Babbage
	CurieEngine          = "curie"            // CurieEngine Curie
	DavinciEngine        = "davinci"          // DavinciEngine Davinci
	DefaultEngine        = DavinciEngine      // DefaultEngine Default Engine
)

const (
	GPT4                      = "gpt4"                          // GPT4 GPT-4
	GPT3Dot5Turbo             = "gpt-3.5-turbo"                 // GPT3Dot5Turbo GPT-3.5 Turbo
	GPT3Dot5Turbo0301         = "gpt-3.5-turbo-0301"            // GPT3Dot5Turbo0301 GPT-3.5 Turbo 0301
	TextSimilarityAda001      = "text-similarity-ada-001"       // TextSimilarityAda001 Text Similarity Ada 001
	TextSimilarityBabbage001  = "text-similarity-babbage-001"   // TextSimilarityBabbage001 Text Similarity Babbage 001
	TextSimilarityCurie001    = "text-similarity-curie-001"     // TextSimilarityCurie001 Text Similarity Curie 001
	TextSimilarityDavinci001  = "text-similarity-davinci-001"   // TextSimilarityDavinci001 Text Similarity Davinci 001
	TextSearchAdaDoc001       = "text-search-ada-doc-001"       // TextSearchAdaDoc001 Text Search Ada Doc 001
	TextSearchAdaQuery001     = "text-search-ada-query-001"     // TextSearchAdaQuery001 Text Search Ada Query 001
	TextSearchBabbageDoc001   = "text-search-babbage-doc-001"   // TextSearchBabbageDoc001 Text Search Babbage Doc 001
	TextSearchBabbageQuery001 = "text-search-babbage-query-001" // TextSearchBabbageQuery001 Text Search Babbage Query 001
	TextSearchCurieDoc001     = "text-search-curie-doc-001"     // TextSearchCurieDoc001 Text Search Curie Doc 001
	TextSearchCurieQuery001   = "text-search-curie-query-001"   // TextSearchCurieQuery001 Text Search Curie Query 001
	TextSearchDavinciDoc001   = "text-search-davinci-doc-001"   // TextSearchDavinciDoc001 Text Search Davinci Doc 001
	TextSearchDavinciQuery001 = "text-search-davinci-query-001" // TextSearchDavinciQuery001 Text Search Davinci Query 001
	CodeSearchAdaCode001      = "code-search-ada-code-001"      // CodeSearchAdaCode001 Code Search Ada Code 001
	CodeSearchAdaText001      = "code-search-ada-text-001"      // CodeSearchAdaText001 Code Search Ada Text 001
	CodeSearchBabbageCode001  = "code-search-babbage-code-001"  // CodeSearchBabbageCode001 Code Search Babbage Code 001
	CodeSearchBabbageText001  = "code-search-babbage-text-001"  // CodeSearchBabbageText001 Code Search Babbage Text 001
	TextEmbeddingAda002       = "text-embedding-ada-002"        // TextEmbeddingAda002 Text Embedding Ada 002
)

const (
	defaultBaseURL        = "https://api.openai.com/v1"
	defaultUserAgent      = "go-gpt3"
	defaultTimeoutSeconds = 30
)

// Image sizes defined by the OpenAI API.
const (
	CreateImageSize256x256   = "256x256"   // CreateImageSize256x256 256x256
	CreateImageSize512x512   = "512x512"   // CreateImageSize512x512 512x512
	CreateImageSize1024x1024 = "1024x1024" // CreateImageSize1024x1024 1024x1024

	CreateImageResponseFormatURL     = "url"      // CreateImageResponseFormatURL URL
	CreateImageResponseFormatB64JSON = "b64_json" // CreateImageResponseFormatB64JSON B64 JSON
)

// Client is an API client to communicate with the OpenAI gpt-3 APIs
type Client interface {
	// Engines lists the currently available engines, and provides basic information about each
	// option such as the owner and availability.
	Engines(ctx context.Context) (*EnginesResponse, error)

	// Engine retrieves an engine instance, providing basic information about the engine such
	// as the owner and availability.
	Engine(ctx context.Context, engine string) (*EngineObject, error)

	// ChatCompletion creates a completion with the Chat completion endpoint which
	// is what powers the ChatGPT experience.
	ChatCompletion(ctx context.Context, request *ChatCompletionRequest) (*ChatCompletionResponse, error)

	// ChatCompletionStream creates a completion with the Chat completion endpoint which
	// is what powers the ChatGPT experience.
	ChatCompletionStream(ctx context.Context, request *ChatCompletionRequest, onData func(*ChatCompletionStreamResponse)) error

	// Completion creates a completion with the default engine. This is the main endpoint of the API
	// which auto-completes based on the given prompt.
	Completion(ctx context.Context, request *CompletionRequest) (*CompletionResponse, error)

	// CompletionStream creates a completion with the default engine and streams the results through
	// multiple calls to onData.
	CompletionStream(ctx context.Context, request *CompletionRequest, onData func(*CompletionResponse)) error

	// CompletionWithEngine is the same as Completion except allows overriding the default engine on the client
	CompletionWithEngine(ctx context.Context, request *CompletionRequest) (*CompletionResponse, error)

	// CompletionStreamWithEngine is the same as CompletionStream allows overriding the default engine on the client
	CompletionStreamWithEngine(ctx context.Context, request *CompletionRequest, onData func(*CompletionResponse)) error

	// Edits is given a prompt and an instruction, the model will return an edited version of the prompt.
	Edits(ctx context.Context, request *EditsRequest) (*EditsResponse, error)

	// Search performs a semantic search over a list of documents with the default engine.
	Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error)

	// SearchWithEngine performs a semantic search over a list of documents with the specified engine.
	SearchWithEngine(ctx context.Context, engine string, request *SearchRequest) (*SearchResponse, error)

	// Embeddings Returns an embedding using the provided request.
	Embeddings(ctx context.Context, request *EmbeddingsRequest) (*EmbeddingsResponse, error)

	// Image returns an image using the provided request.
	Image(ctx context.Context, request *ImageRequest) (*ImageResponse, error)
}

type client struct {
	baseURL       string
	apiKey        string
	userAgent     string
	httpClient    *http.Client
	defaultEngine string
	idOrg         string
}

// NewClient returns a new OpenAI GPT-3 API client. An APIKey is required to use the client
func NewClient(apiKey string, options ...ClientOption) Client {
	httpClient := &http.Client{
		Timeout: defaultTimeoutSeconds * time.Second,
	}
	cli := &client{
		userAgent:     defaultUserAgent,
		apiKey:        apiKey,
		baseURL:       defaultBaseURL,
		httpClient:    httpClient,
		defaultEngine: DefaultEngine,
		idOrg:         "",
	}
	for _, opt := range options {
		cli = opt.apply(cli)
	}
	return cli
}

// Engines lists the currently available engines, and provides basic information about each
// option such as the owner and availability.
func (c *client) Engines(ctx context.Context) (*EnginesResponse, error) {
	req, err := c.newRequest(ctx, "GET", "/engines", nil)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(EnginesResponse)
	if err := getResponseObject(rsp, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Engine retrieves an engine instance, providing basic information about the engine such
// as the owner and availability.
func (c *client) Engine(ctx context.Context, engine string) (*EngineObject, error) {
	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/engines/%s", engine), nil)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(EngineObject)
	if err := getResponseObject(rsp, output); err != nil {
		return nil, err
	}
	return output, nil
}

// ChatCompletion creates a completion with the Chat completion endpoint which
// is what powers the ChatGPT experience.
func (c *client) ChatCompletion(ctx context.Context, request *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	if request.Model == "" {
		request.Model = GPT3Dot5Turbo
	}
	request.Stream = false
	req, err := c.newRequest(ctx, "POST", "/chat/completions", &request)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(ChatCompletionResponse)
	if err := getResponseObject(rsp, output); err != nil {
		return nil, err
	}
	return output, nil
}

// ChatCompletionStream creates a completion with the Chat completion endpoint which
// is what powers the ChatGPT experience.
func (c *client) ChatCompletionStream(ctx context.Context, request *ChatCompletionRequest, onData func(*ChatCompletionStreamResponse)) error {
	if request.Model == "" {
		request.Model = GPT3Dot5Turbo
	}
	request.Stream = true
	req, err := c.newRequest(ctx, "POST", "/chat/completions", request)
	if err != nil {
		return err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(rsp.Body)
	defer rsp.Body.Close()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}
		// make sure there isn't any extra whitespace before or after
		line = bytes.TrimSpace(line)
		// the completion API only returns data events
		if !bytes.HasPrefix(line, dataPrefix) {
			continue
		}
		line = bytes.TrimPrefix(line, dataPrefix)
		// the stream is completed when terminated by [DONE]
		if bytes.HasPrefix(line, doneSequence) {
			break
		}
		output := new(ChatCompletionStreamResponse)
		if err := json.Unmarshal(line, output); err != nil {
			return fmt.Errorf("invalid json stream data: %v", err)
		}
		onData(output)
	}
	return nil
}

// Completion creates a completion with the default engine.
func (c *client) Completion(ctx context.Context, request *CompletionRequest) (*CompletionResponse, error) {
	return c.CompletionWithEngine(ctx, request)
}

// CompletionWithEngine creates a completion with the specified engine.
func (c *client) CompletionWithEngine(ctx context.Context, request *CompletionRequest) (*CompletionResponse, error) {
	request.Stream = false
	req, err := c.newRequest(ctx, "POST", "/completions", &request)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(CompletionResponse)
	if err := getResponseObject(rsp, output); err != nil {
		return nil, err
	}
	return output, nil
}

// CompletionStream creates a completion with the default engine.
func (c *client) CompletionStream(ctx context.Context, request *CompletionRequest,
	onData func(*CompletionResponse)) error {
	return c.CompletionStreamWithEngine(ctx, request, onData)
}

var (
	dataPrefix   = []byte("data: ")
	doneSequence = []byte("[DONE]")
)

// CompletionStreamWithEngine creates a completion with the specified engine.
func (c *client) CompletionStreamWithEngine(ctx context.Context, request *CompletionRequest,
	onData func(*CompletionResponse)) error {
	request.Stream = true
	req, err := c.newRequest(ctx, "POST", "/completions", &request)
	if err != nil {
		return err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(rsp.Body)
	defer rsp.Body.Close()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}
		// make sure there isn't any extra whitespace before or after
		line = bytes.TrimSpace(line)
		// the completion API only returns data events
		if !bytes.HasPrefix(line, dataPrefix) {
			continue
		}
		line = bytes.TrimPrefix(line, dataPrefix)
		// the stream is completed when terminated by [DONE]
		if bytes.HasPrefix(line, doneSequence) {
			break
		}
		output := new(CompletionResponse)
		if err := json.Unmarshal(line, output); err != nil {
			return fmt.Errorf("invalid json stream data: %v", err)
		}
		onData(output)
	}
	return nil
}

// Edits is given a prompt and an instruction, the model will return an edited version of the prompt.
func (c *client) Edits(ctx context.Context, request *EditsRequest) (*EditsResponse, error) {
	req, err := c.newRequest(ctx, "POST", "/edits", &request)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(EditsResponse)
	if err := getResponseObject(rsp, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Search creates a search with the default engine.
func (c *client) Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	return c.SearchWithEngine(ctx, c.defaultEngine, request)
}

// SearchWithEngine performs a semantic search over a list of documents with the specified engine.
func (c *client) SearchWithEngine(ctx context.Context, engine string, request *SearchRequest) (*SearchResponse, error) {
	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("/engines/%s/search", engine), &request)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := new(SearchResponse)
	if err := getResponseObject(rsp, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Embeddings creates text embeddings for a supplied slice of inputs with a provided model.
// See: https://beta.openai.com/docs/api-reference/embeddings
func (c *client) Embeddings(ctx context.Context, request *EmbeddingsRequest) (*EmbeddingsResponse, error) {
	req, err := c.newRequest(ctx, "POST", "/embeddings", &request)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := EmbeddingsResponse{}
	if err := getResponseObject(rsp, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

// Image creates an image
func (c *client) Image(ctx context.Context, request *ImageRequest) (*ImageResponse, error) {
	req, err := c.newRequest(ctx, "POST", "/images/generations", &request)
	if err != nil {
		return nil, err
	}
	rsp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}
	output := ImageResponse{}
	if err := getResponseObject(rsp, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (c *client) performRequest(req *http.Request) (*http.Response, error) {
	rsp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := checkForSuccess(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

// checkForSuccess returns an error if this response includes an error.
func checkForSuccess(rsp *http.Response) error {
	if rsp.StatusCode >= 200 && rsp.StatusCode < 300 {
		return nil
	}
	defer rsp.Body.Close()
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("failed to read from body: %w", err)
	}
	var result APIErrorResponse
	if err := json.Unmarshal(data, &result); err != nil {
		// if we can't decode the json error then create an unexpected error
		apiError := APIError{
			StatusCode: rsp.StatusCode,
			Type:       "Unexpected",
			Message:    string(data),
		}
		return apiError
	}
	result.Error.StatusCode = rsp.StatusCode
	return result.Error
}

func getResponseObject(rsp *http.Response, v interface{}) error {
	defer rsp.Body.Close()
	if err := json.NewDecoder(rsp.Body).Decode(v); err != nil {
		return fmt.Errorf("invalid json response: %w", err)
	}
	return nil
}

func jsonBodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return bytes.NewBuffer(nil), nil
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed encoding json: %w", err)
	}
	return bytes.NewBuffer(raw), nil
}

func (c *client) newRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
	bodyReader, err := jsonBodyReader(payload)
	if err != nil {
		return nil, err
	}
	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	if len(c.idOrg) > 0 {
		req.Header.Set("OpenAI-Organization", c.idOrg)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	return req, nil
}
