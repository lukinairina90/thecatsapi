package catsapi

import "net/http"

const authorizationHeaderName = "x-api-key"

type AuthHeaderRoundTripper struct {
	apiKEY string
	next   http.RoundTripper
}

func NewAuthHeaderRoundTripper(apiKEY string, next http.RoundTripper) *AuthHeaderRoundTripper {
	return &AuthHeaderRoundTripper{
		apiKEY: apiKEY,
		next:   next,
	}
}

func (a AuthHeaderRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add(authorizationHeaderName, a.apiKEY)
	request.Header.Add("Content-Type", "application/json")

	return a.next.RoundTrip(request)
}
