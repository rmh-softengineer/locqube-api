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

func (s *Service) GetLoginUrl() string {
	return s.client.GetLoginURL()
}

func (s *Service) Login(code string) (string, error) {
	return s.client.Login(code)
}

func (s *Service) Post(post model.Post, accessToken string) error {
	return s.client.Post(post, accessToken)
}
