package abuseipdb

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	CHECK_ENDPOINT        = "/api/v2/check"
	REPORTS_ENDPOINT      = "/api/v2/reports"
	BLACKLIST_ENDPOINT    = "/api/v2/blacklist"
	REPORT_ENDPOINT       = "/api/v2/report"
	CHECKBLOCK_ENDPOINT   = "/api/v2/check-block"
	BULKREPORT_ENDPOINT   = "/api/v2/bulk-report"
	CLEARADDRESS_ENDPOINT = "/api/v2/clear-address"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (c *Client) _request(method, endpoint string, body RequestBody, header RequestHeader) (*http.Response, error) {
	data := url.Values{}
	bodyType := reflect.TypeOf(body)
	bodyValue := reflect.ValueOf(body)
	for i := 0; i < bodyType.NumField(); i++ {
		field := bodyType.Field(i)
		value := bodyValue.Field(i)

		if value.IsZero() {
			continue
		}

		tagName := field.Tag.Get("url")
		if tagName == "" {
			tagName = field.Name
		}

		data.Set(tagName, fmt.Sprintf("%v", value.Interface()))
	}

	url := c.baseURL + endpoint
	payload := strings.NewReader(data.Encode())
	if method == GET {
		url += "?" + data.Encode()
		payload = strings.NewReader("")
	}
	request, _ := http.NewRequest(method, url, payload)

	headerType := reflect.TypeOf(header)
	headerValue := reflect.ValueOf(header)
	for i := 0; i < headerType.NumField(); i++ {
		field := headerType.Field(i).Name
		value := headerValue.Field(i).String()
		request.Header.Set(field, value)
	}
	result, err := c.HTTPClient.Do(request)

	return result, err
}

type RequestHeader struct {
	Key    string
	Accept string
}

type RequestBody struct {
	MaxAgeInDays      string `url:"maxAgeInDays"`
	Page              string `url:"page"`
	PerPage           string `url:"perPage"`
	ConfidenceMinimum string `url:"confidenceMinimum"`
	Comment           string `url:"comment"`
	Ip                string `url:"ip"`
	IpAddress         string `url:"ipAddress"`
	Plaintext         bool   `url:"plaintext"`
	Verbose           bool   `url:"verbose"`
}

type CheckResponse struct {
	Data struct {
		IPAddress            string        `json:"ipAddress"`
		IsPublic             bool          `json:"isPublic"`
		IPVersion            int           `json:"ipVersion"`
		IsWhitelisted        bool          `json:"isWhitelisted"`
		AbuseConfidenceScore int           `json:"abuseConfidenceScore"`
		CountryCode          string        `json:"countryCode"`
		CountryName          string        `json:"countryName"`
		UsageType            string        `json:"usageType"`
		Isp                  string        `json:"isp"`
		Domain               string        `json:"domain"`
		Hostnames            []interface{} `json:"hostnames"`
		IsTor                bool          `json:"isTor"`
		TotalReports         int           `json:"totalReports"`
		NumDistinctUsers     int           `json:"numDistinctUsers"`
		LastReportedAt       time.Time     `json:"lastReportedAt"`
		Reports              []struct {
			ReportedAt          time.Time `json:"reportedAt"`
			Comment             string    `json:"comment"`
			Categories          []int     `json:"categories"`
			ReporterID          int       `json:"reporterId"`
			ReporterCountryCode string    `json:"reporterCountryCode"`
			ReporterCountryName string    `json:"reporterCountryName"`
		} `json:"reports,omitempty"`
	} `json:"data"`
}

type ReportsResponse struct {
	Data struct {
		Total           int    `json:"total"`
		Page            int    `json:"page"`
		Count           int    `json:"count"`
		PerPage         int    `json:"perPage"`
		LastPage        int    `json:"lastPage"`
		NextPageURL     string `json:"nextPageUrl"`
		PreviousPageURL string `json:"previousPageUrl"`
		Results         []struct {
			ReportedAt          time.Time `json:"reportedAt"`
			Comment             string    `json:"comment"`
			Categories          []int     `json:"categories"`
			ReporterID          int       `json:"reporterId"`
			ReporterCountryCode string    `json:"reporterCountryCode"`
			ReporterCountryName string    `json:"reporterCountryName"`
		} `json:"results,omitempty"`
	} `json:"data"`
}

type BlacklistResponse struct {
	Meta struct {
		GeneratedAt time.Time `json:"generatedAt"`
	} `json:"meta"`
	Data []struct {
		IPAddress            string    `json:"ipAddress"`
		AbuseConfidenceScore int       `json:"abuseConfidenceScore"`
		LastReportedAt       time.Time `json:"lastReportedAt"`
	} `json:"data"`
}

type ReportResponse struct {
	Data struct {
		IPAddress            string `json:"ipAddress"`
		AbuseConfidenceScore int    `json:"abuseConfidenceScore"`
	} `json:"data"`
}

type CheckBlockResponse struct {
	Data struct {
		NetworkAddress   string `json:"networkAddress"`
		Netmask          string `json:"netmask"`
		MinAddress       string `json:"minAddress"`
		MaxAddress       string `json:"maxAddress"`
		NumPossibleHosts int    `json:"numPossibleHosts"`
		AddressSpaceDesc string `json:"addressSpaceDesc"`
		ReportedAddress  []struct {
			IPAddress            string      `json:"ipAddress"`
			NumReports           int         `json:"numReports"`
			MostRecentReport     time.Time   `json:"mostRecentReport"`
			AbuseConfidenceScore int         `json:"abuseConfidenceScore"`
			CountryCode          interface{} `json:"countryCode"`
		} `json:"reportedAddress"`
	} `json:"data"`
}

type BulkReportResponse struct {
	Data struct {
		SavedReports   int `json:"savedReports"`
		InvalidReports []struct {
			Error     string `json:"error"`
			Input     string `json:"input"`
			RowNumber int    `json:"rowNumber"`
		} `json:"invalidReports"`
	} `json:"data"`
}

type ClearAddressResponse struct {
	Data struct {
		NumReportsDeleted int `json:"numReportsDeleted"`
	} `json:"data"`
}
