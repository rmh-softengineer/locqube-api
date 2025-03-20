package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rmh-softengineer/locqube/api/model"
)

func NewClient(appID, appSecret string) *client {
	return &client{
		facebookAppID:     appID,
		facebookAppSecret: appSecret,
	}
}

func (c *client) Login(accessToken string) (*string, error) {
	userID, err := c.validateFacebookToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid facebook token: %w", err)
	}

	token := fmt.Sprintf("mock-jwt-token-for-%s", *userID) // Generate a real JWT token

	return &token, nil
}

func (c *client) Post(post model.Post, accessToken string) error {
	postURL := c.buildPostURL(post, accessToken)

	resp, err := http.Post(postURL, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to post to Facebook: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func (c *client) buildPostURL(post model.Post, accessToken string) string {
	if post.Link != nil {
		return fmt.Sprintf("https://graph.facebook.com/me/feed?message=%s&link=%s&access_token=%s",
			post.Message, *post.Link, accessToken)
	}

	return fmt.Sprintf("https://graph.facebook.com/me/feed?message=%s&access_token=%s",
		post.Message, accessToken)
}

func (c *client) validateFacebookToken(accessToken string) (*string, error) {
	validationURL := fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s|%s",
		accessToken, c.facebookAppID, c.facebookAppSecret)

	resp, err := http.Get(validationURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var validationResponse model.ValidationTokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&validationResponse); err != nil {
		return nil, err
	}

	if !validationResponse.Data.IsValid {
		return nil, errors.New("invalid token")
	}

	return &validationResponse.Data.UserID, nil
}
