package facebook

import (
	"github.com/rmh-softengineer/locqube/api/http/facebook"
	model "github.com/rmh-softengineer/locqube/api/model"
)

func NewService(client *facebook.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Login() string {
	return s.client.Login()
}

func (s *Service) Post(post model.Post, accessToken string) error {
	return s.client.Post(post, accessToken)
}
