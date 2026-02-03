package checkgroup

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Edge-Center/edgecenteredgemon-go/edgecenter"
)

type Service interface {
	Create(ctx context.Context, req *Request) (*Response, error)
	Get(ctx context.Context, id int) (*Response, error)
	Update(ctx context.Context, id int, req *Request) (*Response, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	r edgecenter.Requester
}

func New(r edgecenter.Requester) Service {
	return &service{r: r}
}

func (s *service) Create(ctx context.Context, req *Request) (*Response, error) {
	var resp Response
	u := url.URL{Path: "/rmon/check-group"}

	if err := s.r.Request(ctx, http.MethodPost, u.String(), req, &resp); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resp, nil
}

func (s *service) Get(ctx context.Context, id int) (*Response, error) {
	var resp Response

	u := url.URL{
		Path: fmt.Sprintf("/rmon/check-group/%d", id),
	}

	if err := s.r.Request(ctx, http.MethodGet, u.String(), nil, &resp); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resp, nil
}

func (s *service) Update(ctx context.Context, id int, req *Request) (*Response, error) {
	var resp Response

	u := url.URL{
		Path: fmt.Sprintf("/rmon/check-group/%d", id),
	}

	if err := s.r.Request(ctx, http.MethodPut, u.String(), req, &resp); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &resp, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	u := url.URL{
		Path: fmt.Sprintf("/rmon/check-group/%d", id),
	}

	if err := s.r.Request(ctx, http.MethodDelete, u.String(), nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
