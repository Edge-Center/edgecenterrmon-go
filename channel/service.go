package channel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Edge-Center/edgecenteredgemon-go/edgecenter"
)

type Service interface {
	Get(ctx context.Context, receiver string, channelId int) (*Response, error)
	Create(ctx context.Context, receiver string, req *Request) (*Response, error)
	Update(ctx context.Context, receiver string, channelId int, req *Request) error
	Delete(ctx context.Context, receiver string, channelId int) error
}

type service struct {
	r edgecenter.Requester
}

func New(r edgecenter.Requester) Service {
	return &service{r: r}
}

func (s *service) Create(ctx context.Context, receiver string, req *Request) (*Response, error) {
	var resp Response
	u := url.URL{
		Path: fmt.Sprintf("/rmon/channel/%s", receiver),
	}

	if err := s.r.Request(ctx, http.MethodPost, u.String(), req, &resp); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resp, nil
}

func (s *service) Get(ctx context.Context, receiver string, channelId int) (*Response, error) {
	var resp Response
	u := url.URL{
		Path: fmt.Sprintf("/rmon/channel/%s/%d", receiver, channelId),
	}

	if err := s.r.Request(ctx, http.MethodGet, u.String(), nil, &resp); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resp, nil
}

func (s *service) Update(ctx context.Context, receiver string, channelId int, req *Request) error {
	u := url.URL{Path: fmt.Sprintf("/rmon/channel/%s/%d", receiver, channelId)}

	if err := s.r.Request(ctx, http.MethodPut, u.String(), req, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}

func (s *service) Delete(ctx context.Context, receiver string, channelId int) error {
	u := url.URL{
		Path: fmt.Sprintf("/rmon/channel/%s/%d", receiver, channelId),
	}

	if err := s.r.Request(ctx, http.MethodDelete, u.String(), nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
