package catsapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseUrl string
	userID  string
	client  *http.Client
}

func NewClient(baseUrl string, apiKEY string, userID string) *Client {
	client := http.DefaultClient
	client.Transport = NewAuthHeaderRoundTripper(apiKEY, http.DefaultTransport)

	return &Client{
		baseUrl: baseUrl,
		userID:  userID,
		client:  http.DefaultClient,
	}
}

func (c Client) List(ctx context.Context, params *ListParams) (ListResponse, error) {
	limit := 3
	page := 1
	order := "asc"

	if params != nil {
		limit = params.Limit
		page = params.Page
		if params.DescOrder {
			order = "desc"
		}
	}

	url := fmt.Sprintf("%s/images/search?limit=%d&page=%d&order=%s", c.baseUrl, limit, page, order)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ListResponse{}, err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return ListResponse{}, err
	}

	defer resp.Body.Close()

	if err = c.handleStatusCode(resp); err != nil {
		return ListResponse{}, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListResponse{}, err
	}

	var r ListResponse
	if err := json.Unmarshal(b, &r); err != nil {
		return ListResponse{}, err
	}

	return r, nil
}

func (c Client) GetImage(ctx context.Context, imageId string) (CatImage, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/%s", imageId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return CatImage{}, err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return CatImage{}, err
	}

	defer resp.Body.Close()

	if err = c.handleStatusCode(resp); err != nil {
		return CatImage{}, nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return CatImage{}, err
	}

	var r CatImage
	if err = json.Unmarshal(b, &r); err != nil {
		return CatImage{}, err
	}

	return r, nil
}

func (c Client) CreateVote(ctx context.Context, imageId string, iLikeIt bool) error {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/votes")

	vb := VoteBody{
		ImageId: imageId,
		SubId:   c.userID,
	}

	if iLikeIt {
		vb.Value = 1
	} else {
		vb.Value = 0
	}

	bb, err := json.Marshal(vb)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bb))
	if err != nil {
		return err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)

	return c.handleStatusCode(resp)
}

func (c Client) GetVotes(ctx context.Context) (ListRespVote, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/votes")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ListRespVote{}, err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return ListRespVote{}, err
	}

	defer resp.Body.Close()

	if err := c.handleStatusCode(resp); err != nil {
		return ListRespVote{}, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListRespVote{}, err
	}

	var r ListRespVote
	if err = json.Unmarshal(b, &r); err != nil {
		return ListRespVote{}, err
	}

	return r, nil
}

func (c Client) DeleteVote(ctx context.Context, voteId string) (Message, error) {
	url := fmt.Sprintf("https://api.thecatapi.com/v1/votes/%s", voteId)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return Message{}, err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return Message{}, err
	}

	defer resp.Body.Close()

	if err = c.handleStatusCode(resp); err != nil {
		return Message{}, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	var m Message
	if err = json.Unmarshal(b, &m); err != nil {
		return Message{}, err
	}

	return m, nil
}

func (c Client) handleStatusCode(response *http.Response) error {
	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return nil
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusInternalServerError:
		return ErrInternal
	case http.StatusMethodNotAllowed:
		return ErrMethodNotAllowed
	}

	return ErrInternal
}
