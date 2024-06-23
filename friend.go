package hashicups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetFriends() ([]Friend, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/friends", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	friends := []Friend{}
	err = json.Unmarshal(body, &friends)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (c *Client) GetFriend(friendID string) ([]Friend, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/friends/%s", c.HostURL, friendID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	friend := []Friend{}
	err = json.Unmarshal(body, &friend)
	if err != nil {
		return nil, err
	}

	return friend, nil
}

func (c *Client) CreateFriend(friends []Friend) (*Friend, error) {
	rb, err := json.Marshal(friends)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/friends", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	friend := Friend{}
	err = json.Unmarshal(body, &friend)
	if err != nil {
		return nil, err
	}

	return &friend, nil
}

func (c *Client) UpdateFriend(friendID string, friends []Friend) (*Friend, error) {
	rb, err := json.Marshal(friends)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/friends/%s", c.HostURL, friendID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	friend := Friend{}
	err = json.Unmarshal(body, &friend)
	if err != nil {
		return nil, err
	}

	return &friend, nil
}

func (c *Client) DeleteFriend(friendID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/friends/%s", c.HostURL, friendID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return err
	}

	if string(body) != "Deleted friend" {
		return errors.New(string(body))
	}

	return nil
}