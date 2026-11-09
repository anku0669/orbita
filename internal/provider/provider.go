// Package provider defines provider-neutral model interfaces.
package provider

import "context"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Text       string `json:"text"`
	ProviderID string `json:"provider_id"`
}

type Model interface {
	Generate(context.Context, []Message) (Response, error)
}
