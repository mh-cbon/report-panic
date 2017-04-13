package ga

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var trackingIDMatcher = regexp.MustCompile(`^UA-\d+-\d+$`)

// NewClient prepares a new client for ga service
func NewClient(clientID, trackingID string) (*Client, error) {
	if !trackingIDMatcher.MatchString(trackingID) {
		return nil, fmt.Errorf("Invalid Tracking ID: %s", trackingID)
	}
	return &Client{
		UseTLS:     true,
		HTTPClient: http.DefaultClient,
		trackingID: trackingID,
		clientID:   clientID,
	}, nil
}

// Client for ga
type Client struct {
	UseTLS     bool
	trackingID string
	clientID   string
	HTTPClient *http.Client
}

// PageView triggers hits
func (c *Client) PageView(pvurl string) error {

	v := url.Values{}
	v.Add("cid", c.clientID)
	v.Add("dl", "http://mh-cbon.github.io/report-panic/"+pvurl)
	v.Add("t", "pageview")
	v.Add("tid", c.trackingID)
	v.Add("v", "1")

	rurl := "http://ssl.google-analytics.com/collect"
	if c.UseTLS {
		rurl = "https://www.google-analytics.com/collect"
	}

	str := v.Encode()
	buf := bytes.NewBufferString(str)

	resp, err := c.HTTPClient.Post(rurl, "application/x-www-form-urlencoded", buf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("Rejected by Google with code %d", resp.StatusCode)
	}

	// fmt.Printf("POST %s => %d\n", str, resp.StatusCode)

	return nil
}
