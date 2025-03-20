package facebook

import (
	"net/http"

	"github.com/rmh-softengineer/locqube/api/model"
)

type Client interface {
	Login(accessToken string) (*string, error)
	Post(post model.Post, accessToken string) error
}
type client struct {
	facebookAppID     string
	facebookAppSecret string
	httpClient        http.Client
}
