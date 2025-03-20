package facebook

import (
	"github.com/rmh-softengineer/locqube/api/http/facebook"
	model "github.com/rmh-softengineer/locqube/api/model"
)

func NewService(client facebook.Client) *service {
	return &service{
		client: client,
	}
}

func (s *service) Login(accessToken string) (*string, error) {
	return s.client.Login(accessToken)
}

func (s *service) Post(post model.Post, accessToken string) error {
	return s.client.Post(post, accessToken)
}
