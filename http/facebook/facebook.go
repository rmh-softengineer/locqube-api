package facebook

import (
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

func (c *Client) Login() string {
	return fmt.Sprintf("https://www.facebook.com/v17.0/dialog/oauth?client_id=%s&redirect_uri=%s&scope=public_profile,email,pages_manage_posts,publish_video",
		c.facebookAppID, c.redirectURL)
}

func (c *Client) Post(post facebook.Post, accessToken string) error {
	postURL := c.buildPostURL(post, accessToken)

	_, err := http.Post(postURL, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to post to Facebook: %w", err)
	}

	return nil
}

func (c *Client) buildPostURL(post facebook.Post, accessToken string) string {
	if post.Link != nil {
		return fmt.Sprintf("https://graph.facebook.com/me/feed?message=%s&link=%s&access_token=%s",
			post.Message, *post.Link, accessToken)
	}

	return fmt.Sprintf("https://graph.facebook.com/me/feed?message=%s&access_token=%s",
		post.Message, accessToken)
}
