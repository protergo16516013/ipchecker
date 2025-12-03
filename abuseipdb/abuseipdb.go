package abuseipdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Client struct {
	APIKey         string
	HTTPClient     *http.Client
	baseURL        string
	default_header RequestHeader
}

func NewClient(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		baseURL:    "https://api.abuseipdb.com",
		default_header: RequestHeader{
			Key:    apiKey,
			Accept: "application/json",
		},
	}
}

func (c *Client) PrettyPrint(i interface{}) {
	prettyJSON, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}

func (c *Client) _decodeResponse(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(v)
}

func (c *Client) Check(ip string, maxAgeInDays int, verbose bool) (*CheckResponse, error) {
	body := RequestBody{
		IpAddress: ip,
	}
	if maxAgeInDays > 0 {
		body.MaxAgeInDays = strconv.Itoa(maxAgeInDays)
	}
	if verbose {
		body.Verbose = verbose
	}

	res, err := c._request(GET, CHECK_ENDPOINT, body, c.default_header)
	if err != nil {
		println("Error whilst making request:", err)
		return nil, err
	}
	defer res.Body.Close()

	println("response status:", res.Status)
	println("response body:", res.Body)

	response := CheckResponse{}
	if err := c._decodeResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Reports(ip string, page int, perPage int, maxAgeInDays int) (*ReportsResponse, error) {
	body := RequestBody{
		IpAddress: ip,
		Page:      strconv.Itoa(page),
		PerPage:   strconv.Itoa(perPage),
	}
	if maxAgeInDays > 0 {
		body.MaxAgeInDays = strconv.Itoa(maxAgeInDays)
	}

	res, err := c._request(GET, REPORTS_ENDPOINT, body, c.default_header)
	if err != nil {
		println("Error whilst making request:", err)
		return nil, err
	}
	defer res.Body.Close()

	response := ReportsResponse{}
	if err := c._decodeResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
