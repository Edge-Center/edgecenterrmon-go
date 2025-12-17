package provider

import (
	"fmt"
	"time"

	"github.com/Edge-Center/edgecenterrmon-go/edgecenter"
)

// ClientOption — конфигуратор Client (как в edgecentercdn-go).
type ClientOption func(*Client)

// WithTimeout — переопределяет timeout HTTP-клиента.
func WithTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.httpc.Timeout = t
	}
}

// WithSigner — устанавливает подписант (Bearer, API key, JWT...).
func WithSigner(s edgecenter.RequestSigner) ClientOption {
	return func(c *Client) {
		c.signer = s
	}
}

// WithUserAgent — задаёт User-Agent для RMON SDK.
func WithUserAgent(ua string) ClientOption {
	return func(c *Client) {
		c.ua = ua
	}
}

func WithSignerFunc(f edgecenter.RequestSignerFunc) ClientOption {
	return func(c *Client) {
		c.signer = f
	}
}

func AuthenticatedHeaders(apiKey string) (m map[string]string) {
	return map[string]string{"Authorization": fmt.Sprintf("APIKey %s", apiKey)}
}
