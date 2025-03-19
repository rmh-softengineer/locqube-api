package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rmh-softengineer/locqube/api/model"
)

func NewClient(appID, appSecret, redirectURL string) *Client {
	return &Client{
		facebookAppID:     appID,
		facebookAppSecret: appSecret,
		redirectURL:       redirectURL,
	}
}

func (c *Client) GetLoginURL() string {
	return fmt.Sprintf("https://www.facebook.com/v17.0/dialog/oauth?client_id=%s&redirect_uri=%s&scope=public_profile,email,pages_manage_posts,publish_video",
		c.facebookAppID, c.redirectURL)
}

func (c *Client) Login(code string) (string, error) {
	tokenURL := fmt.Sprintf("https://graph.facebook.com/v17.0/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s",
		c.facebookAppID, c.redirectURL, c.facebookAppSecret, code)

	resp, err := http.Get(tokenURL)
	if err != nil {
		return "", fmt.Errorf("failed to get access token: %w", err)
	}
	defer resp.Body.Close()

	var fbToken map[string]interface{}

	if err = json.NewDecoder(resp.Body).Decode(&fbToken); err != nil {
		return "", fmt.Errorf("failed to decode access token response: %w", err)
	}

	accessToken, ok := fbToken["access_token"].(string)
	if !ok {
		return "", errors.New("invalid access token response")
	}

	return accessToken, nil
}

func (c *Client) Post(post model.Post, accessToken string) error {
	postURL := c.buildPostURL(post, accessToken)

	resp, err := http.Post(postURL, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to post to Facebook: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) buildPostURL(post model.Post, accessToken string) string {
	if post.Link != nil {
		return fmt.Sprintf("https://graph.facebook.com/me/feed?message=%s&link=%s&access_token=%s",
			post.Message, *post.Link, accessToken)
	}

	return fmt.Sprintf("https://graph.facebook.com/me/feed?message=%s&access_token=%s",
		post.Message, accessToken)
}
