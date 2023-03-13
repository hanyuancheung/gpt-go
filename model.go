// Package gpt provides a client for the OpenAI GPT-3 API
package gpt

import "fmt"

// APIError represents an error that occurred on an API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

// Error returns a string representation of the error
func (e APIError) Error() string {
	return fmt.Sprintf("[%d:%s] %s", e.StatusCode, e.Type, e.Message)
}

// APIErrorResponse is the full error response that has been returned by an API.
type APIErrorResponse struct {
	Error APIError `json:"error"`
}

// EngineObject contained in an engine repose
type EngineObject struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

// EnginesResponse is returned from the Engines API
type EnginesResponse struct {
	Data   []EngineObject `json:"data"`
	Object string         `json:"object"`
}

// ChatCompletionRequestMessage is a message to use as the context for the chat completion API
type ChatCompletionRequestMessage struct {
	// Role is the role is the role of the message. Can be "system", "user", or "assistant"
	Role string `json:"role"`
	// Content is the content of the message
	Content string `json:"content"`
}

// ChatCompletionRequest is a request for the chat completion API
type ChatCompletionRequest struct {
	// Model is the name of the model to use. If not specified, will default to gpt-3.5-turbo.
	Model string `json:"model"`
	// Messages is a list of messages to use as the context for the chat completion.
	Messages []ChatCompletionRequestMessage `json:"messages"`
	// Temperature is sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic
	Temperature float32 `json:"temperature,omitempty"`
	// TopP is an alternative to sampling with temperature, called nucleus sampling, where the model considers the results of
	// the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	TopP float32 `json:"top_p,omitempty"`
	// N is number of responses to generate
	N int `json:"n,omitempty"`
	// Stream is whether to stream responses back as they are generated
	Stream bool `json:"stream,omitempty"`
	// Stop is up to 4 sequences where the API will stop generating further tokens.
	Stop []string `json:"stop,omitempty"`
	// MaxTokens is the maximum number of tokens to return.
	MaxTokens int `json:"max_tokens,omitempty"`
	// PresencePenalty (-2, 2) penalize tokens that haven't appeared yet in the history.
	PresencePenalty float32 `json:"presence_penalty,omitempty"`
	// FrequencyPenalty (-2, 2) penalize tokens that appear too frequently in the history.
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`
	// LogitBias modify the probability of specific tokens appearing in the completion.
	LogitBias map[string]float32 `json:"logit_bias,omitempty"`
	// User can be used to identify an end-user
	User string `json:"user,omitempty"`
}

// CompletionRequest is a request for the completions API
type CompletionRequest struct {
	Model string `json:"model"`
	// Prompt sets a list of string prompts to use.
	Prompt []string `json:"prompt,omitempty"`
	// Suffix comes after a completion of inserted text.
	Suffix string `json:"suffix,omitempty"`
	// MaxTokens sets how many tokens to complete up to. Max of 512
	MaxTokens int `json:"max_tokens,omitempty"`
	// Temperature sets sampling temperature to use
	Temperature float32 `json:"temperature,omitempty"`
	// TopP sets alternative to temperature for nucleus sampling
	TopP *float32 `json:"top_p,omitempty"`
	// N sets how many choice to create for each prompt
	N *int `json:"n"`
	// Stream sets whether to stream back results or not. Don't set this value in the request yourself
	// as it will be overridden depending on if you use CompletionStream or Completion methods.
	Stream bool `json:"stream,omitempty"`
	// LogProbs sets include the probabilities of most likely tokens
	LogProbs *int `json:"logprobs"`
	// Echo sets back the prompt in addition to the completion
	Echo bool `json:"echo"`
	// Stop sets up to 4 sequences where the API will stop generating tokens. Response will not contain the stop sequence.
	Stop []string `json:"stop,omitempty"`
	// PresencePenalty sets number between 0 and 1 that penalizes tokens that have already appeared in the text so far.
	PresencePenalty float32 `json:"presence_penalty"`
	// FrequencyPenalty number between 0 and 1 that penalizes tokens on existing frequency in the text so far.
	FrequencyPenalty float32 `json:"frequency_penalty"`
	// BestOf sets how many of the n best completions to return. Defaults to 1.
	BestOf int `json:"best_of,omitempty"`
	// LogitBias sets modify the probability of specific tokens appearing in the completion.
	LogitBias map[string]float32 `json:"logit_bias,omitempty"`
	// User sets an end-user identifier. Can be used to associate completions generated by a specific user.
	User string `json:"user,omitempty"`
}

// EditsRequest is a request for the edits API
type EditsRequest struct {
	// Model is ID of the model to use. You can use the List models API to see all of your available models,
	// or see our Model overview for descriptions of them.
	Model string `json:"model"`
	// Input is the input text to use as a starting point for the edit.
	Input string `json:"input,omitempty"`
	// Instruction is the instruction that tells the model how to edit the prompt.
	Instruction string `json:"instruction"`
	// N is how many edits to generate for the input and instruction. Defaults to 1
	N *int `json:"n,omitempty"`
	// Temperature is sampling temperature to use
	Temperature *float32 `json:"temperature,omitempty"`
	// TopP is alternative to temperature for nucleus sampling
	TopP *float32 `json:"top_p,omitempty"`
}

// EmbeddingsRequest is a request for the Embeddings API
type EmbeddingsRequest struct {
	// Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings
	// for multiple inputs in a single request, pass an array of strings or array of token arrays.
	// Each input must not exceed 2048 tokens in length.
	Input []string `json:"input"`
	// Model is ID of the model to use
	Model string `json:"model"`
	// User is the request user is an optional parameter meant to be used to trace abusive requests
	// back to the originating user. OpenAI states:
	// "The [user] IDs should be a string that uniquely identifies each user. We recommend hashing
	// their username or email address, in order to avoid sending us any identifying information.
	// If you offer a preview of your product to non-logged in users, you can send a session ID
	// instead."
	User string `json:"user,omitempty"`
}

// LogProbResult represents logprob result of Choice
type LogProbResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogProbs []float32            `json:"token_logprobs"`
	TopLogProbs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

// ChatCompletionResponseMessage is a message returned in the response to the Chat Completions API
type ChatCompletionResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionResponseChoice is one of the choices returned in the response to the Chat Completions API
type ChatCompletionResponseChoice struct {
	Index        int                           `json:"index"`
	FinishReason string                        `json:"finish_reason"`
	Message      ChatCompletionResponseMessage `json:"message"`
}

// ChatCompletionStreamResponseChoice is one of the choices returned in the response to the Chat Completions API
type ChatCompletionStreamResponseChoice struct {
	Index        int                           `json:"index"`
	FinishReason string                        `json:"finish_reason"`
	Delta        ChatCompletionResponseMessage `json:"delta"`
}

// ChatCompletionsResponseUsage is the object that returns how many tokens the completion's request used
type ChatCompletionsResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionResponse is the full response from a request to the Chat Completions API
type ChatCompletionResponse struct {
	ID      string                         `json:"id"`
	Object  string                         `json:"object"`
	Created int                            `json:"created"`
	Model   string                         `json:"model"`
	Choices []ChatCompletionResponseChoice `json:"choices"`
	Usage   ChatCompletionsResponseUsage   `json:"usage"`
}

type ChatCompletionStreamResponse struct {
	ID      string                               `json:"id"`
	Object  string                               `json:"object"`
	Created int                                  `json:"created"`
	Model   string                               `json:"model"`
	Choices []ChatCompletionStreamResponseChoice `json:"choices"`
	Usage   ChatCompletionsResponseUsage         `json:"usage"`
}

// CompletionResponseChoice is one of the choices returned in the response to the Completions API
type CompletionResponseChoice struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	LogProbs     LogProbResult `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

// CompletionResponse is the full response from a request to the completions API
type CompletionResponse struct {
	ID      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created int                        `json:"created"`
	Model   string                     `json:"model"`
	Choices []CompletionResponseChoice `json:"choices"`
	Usage   CompletionResponseUsage    `json:"usage"`
}

// CompletionResponseUsage is the object that returns how many tokens the completion's request used
type CompletionResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// EditsResponse is the full response from a request to the edits API
type EditsResponse struct {
	Object  string                `json:"object"`
	Created int                   `json:"created"`
	Choices []EditsResponseChoice `json:"choices"`
	Usage   EditsResponseUsage    `json:"usage"`
}

// EmbeddingsResult The inner result of a create embeddings request, containing the embeddings for a single input.
type EmbeddingsResult struct {
	// The type of object returned (e.g., "list", "object")
	Object string `json:"object"`
	// The embedding data for the input
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// EmbeddingsUsage The usage stats for an embeddings response
type EmbeddingsUsage struct {
	// The number of tokens used by the prompt
	PromptTokens int `json:"prompt_tokens"`
	// The total tokens used
	TotalTokens int `json:"total_tokens"`
}

// EmbeddingsResponse is the response from a create embeddings request.
// See: https://beta.openai.com/docs/api-reference/embeddings/create
type EmbeddingsResponse struct {
	Object string             `json:"object"`
	Data   []EmbeddingsResult `json:"data"`
	Usage  EmbeddingsUsage    `json:"usage"`
}

// EditsResponseChoice is one of the choices returned in the response to the Edits API
type EditsResponseChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

// EditsResponseUsage is a structure used in the response from a request to the edits API
type EditsResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// SearchRequest is a request for the document search API
type SearchRequest struct {
	Documents []string `json:"documents"`
	Query     string   `json:"query"`
}

// SearchData is a single search result from the document search API
type SearchData struct {
	Document int     `json:"document"`
	Object   string  `json:"object"`
	Score    float64 `json:"score"`
}

// SearchResponse is the full response from a request to the document search API
type SearchResponse struct {
	Data   []SearchData `json:"data"`
	Object string       `json:"object"`
}
