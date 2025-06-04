package types

type IHttpService interface {
	Do(url string) ([]byte, error)
}
