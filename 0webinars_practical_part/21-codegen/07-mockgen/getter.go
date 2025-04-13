package mockgen

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -destination ./mocks/mock_getter.go -package mocks Getter
type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

func GetPage(url string, opts ...Option) ([]byte, error) {
	options := GetPageOptions{
		Getter: http.DefaultClient,
	}
	for _, o := range opts {
		o(&options)
	}

	resp, err := options.Getter.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to GET %q: %w", url, err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	return result, nil
}
