package facebook

import (
	"github.com/rmh-softengineer/locqube/api/http/facebook"
	model "github.com/rmh-softengineer/locqube/api/model"
)

type Service interface {
	Login(accessToken string) (*string, error)
	Post(post model.Post, accessToken string) error
}

type service struct {
	client facebook.Client
}
