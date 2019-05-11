package camunda

import (
	"encoding/json"

	"github.com/pkg/errors"

	"gopkg.in/resty.v1"
)

// Client ...
type Client struct {
	APIURL string
}

// FetchAndLock ...
func (c *Client) FetchAndLock(param *FetchAndLock) ([]Task, error) {

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(c.APIURL + "/engine-rest/external-task/fetchAndLock")
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	err = json.Unmarshal(resp.Body(), &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, err
}

// Complete ...
func (c *Client) Complete(id string, param *Complete) error {

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(c.APIURL + "/engine-rest/external-task/" + id + "/complete")
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 {
		return errors.New("Complete failed with response " + resp.String())
	}

	return nil
}
