package types

import "net/http"

type IRequestService interface {
	Get(url string) (*http.Request, error)
}
