package chatgpt

type ChatGPTRequest struct {
	Model       string          `json:"model"`
	Messages    []RequesterData `json:"messages"`
	Temperature float64         `json:"temperature"`
}

type RequesterData struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTData struct {
	Id      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []Choices    `json:"choices"`
	Usage   ChatGPTUsage `json:"usage"`
}

type Choices struct {
	Index        int64         `json:"index"`
	Message      ResponderData `json:"message"`
	FinishReason string        `json:"finish_reason"`
}

type ResponderData struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTUsage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalCounts      int64 `json:"total_tokens"`
}

// Messenger is the interface for the ChatGPT provider.
type Messenger interface {
	Post(*ChatGPTRequest) (*ChatGPTData, error)
}

func PostMessage(m Messenger, request *ChatGPTRequest) (*ChatGPTData, error) {
	return m.Post(request)
}
