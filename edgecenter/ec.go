package edgecenter

import (
	"context"
	"net/http"
)

// Requester — интерфейс, который должен реализовать HTTP-клиент SDK.
// Все сервисы SDK вызывают только Request(...), как в edgecentercdn-go.
type Requester interface {
	Request(ctx context.Context, method, path string, payload interface{}, result interface{}) error
}

// RequestSigner — механизм подписания HTTP-запросов.
// В провайдере EdgeCenter используется signer, который ставит заголовки авторизации.
type RequestSigner interface {
	Sign(req *http.Request) error
}

// RequestSignerFunc — адаптер, чтобы любая функция могла быть Signer-ом.
type RequestSignerFunc func(req *http.Request) error

func (f RequestSignerFunc) Sign(req *http.Request) error {
	return f(req)
}
