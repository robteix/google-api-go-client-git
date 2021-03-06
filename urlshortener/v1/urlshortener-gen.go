// Package urlshortener provides access to the URL Shortener API.
//
// See http://code.google.com/apis/urlshortener/v1/getting_started.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/urlshortener/v1"
//   ...
//   urlshortenerService, err := urlshortener.New(oauthHttpClient)
package urlshortener

import (
	"bytes"
	"fmt"
	"http"
	"io"
	"json"
	"os"
	"strings"
	"strconv"
	"url"
	"google-api-go-client.googlecode.com/hg/google-api"
)

var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version

const apiId = "urlshortener:v1"
const apiName = "urlshortener"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/urlshortener/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage your goo.gl short URLs
	UrlshortenerScope = "https://www.googleapis.com/auth/urlshortener"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Url = &UrlService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Url *UrlService
}

type UrlService struct {
	s *Service
}

type AnalyticsSnapshot struct {
	// Browsers: Top browsers, e.g. "Chrome"; sorted by (descending) click
	// counts. Only present if this data is available.
	Browsers []*StringCount `json:"browsers,omitempty"`

	// Platforms: Top platforms or OSes, e.g. "Windows"; sorted by
	// (descending) click counts. Only present if this data is available.
	Platforms []*StringCount `json:"platforms,omitempty"`

	// Referrers: Top referring hosts, e.g. "www.google.com"; sorted by
	// (descending) click counts. Only present if this data is available.
	Referrers []*StringCount `json:"referrers,omitempty"`

	// LongUrlClicks: Number of clicks on all goo.gl short URLs pointing to
	// this long URL.
	LongUrlClicks int64 `json:"longUrlClicks,omitempty,string"`

	// ShortUrlClicks: Number of clicks on this short URL.
	ShortUrlClicks int64 `json:"shortUrlClicks,omitempty,string"`

	// Countries: Top countries (expressed as country codes), e.g. "US" or
	// "DE"; sorted by (descending) click counts. Only present if this data
	// is available.
	Countries []*StringCount `json:"countries,omitempty"`
}

type AnalyticsSummary struct {
	// Week: Click analytics over the last week.
	Week *AnalyticsSnapshot `json:"week,omitempty"`

	// AllTime: Click analytics over all time.
	AllTime *AnalyticsSnapshot `json:"allTime,omitempty"`

	// Month: Click analytics over the last month.
	Month *AnalyticsSnapshot `json:"month,omitempty"`

	// TwoHours: Click analytics over the last two hours.
	TwoHours *AnalyticsSnapshot `json:"twoHours,omitempty"`

	// Day: Click analytics over the last day.
	Day *AnalyticsSnapshot `json:"day,omitempty"`
}

type UrlHistory struct {
	// Items: A list of URL resources.
	Items []*Url `json:"items,omitempty"`

	// NextPageToken: A token to provide to get the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of short URLs associated with this user (may
	// be approximate).
	TotalItems int64 `json:"totalItems,omitempty"`

	// Kind: The fixed string "urlshortener#urlHistory".
	Kind string `json:"kind,omitempty"`

	// ItemsPerPage: Number of items returned with each full "page" of
	// results. Note that the last page could have fewer items than the
	// "itemsPerPage" value.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
}

type StringCount struct {
	// Count: Number of clicks for this top entry, e.g. for this particular
	// country or browser.
	Count int64 `json:"count,omitempty,string"`

	// Id: Label assigned to this top entry, e.g. "US" or "Chrome".
	Id string `json:"id,omitempty"`
}

type Url struct {
	// Kind: The fixed string "urlshortener#url".
	Kind string `json:"kind,omitempty"`

	// Id: Short URL, e.g. "http://goo.gl/l6MS".
	Id string `json:"id,omitempty"`

	// Created: Time the short URL was created; ISO 8601 representation
	// using the yyyy-MM-dd'T'HH:mm:ss.SSSZZ format, e.g.
	// "2010-10-14T19:01:24.944+00:00".
	Created string `json:"created,omitempty"`

	// Status: Status of the target URL. Possible values: "OK", "MALWARE",
	// "PHISHING", or "REMOVED". A URL might be marked "REMOVED" if it was
	// flagged as spam, for example.
	Status string `json:"status,omitempty"`

	// LongUrl: Long URL, e.g. "http://www.google.com/". Might not be
	// present if the status is "REMOVED".
	LongUrl string `json:"longUrl,omitempty"`

	// Analytics: A summary of the click analytics for the short and long
	// URL. Might not be present if not requested or currently unavailable.
	Analytics *AnalyticsSummary `json:"analytics,omitempty"`
}

// method id "urlshortener.url.list":

type UrlListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieves a list of URLs shortened by a user.
func (r *UrlService) List() *UrlListCall {
	c := &UrlListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Projection sets the optional parameter "projection": Additional
// information to return.
func (c *UrlListCall) Projection(projection string) *UrlListCall {
	c.opt_["projection"] = projection
	return c
}

// StartToken sets the optional parameter "start-token": Token for
// requesting successive pages of results.
func (c *UrlListCall) StartToken(startToken string) *UrlListCall {
	c.opt_["start-token"] = startToken
	return c
}

func (c *UrlListCall) Do() (*UrlHistory, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start-token"]; ok {
		params.Set("start-token", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/urlshortener/v1/", "url/history")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(UrlHistory)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of URLs shortened by a user.",
	//   "httpMethod": "GET",
	//   "id": "urlshortener.url.list",
	//   "parameters": {
	//     "projection": {
	//       "description": "Additional information to return.",
	//       "enum": [
	//         "ANALYTICS_CLICKS",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Returns short URL click counts.",
	//         "Returns short URL click counts."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "start-token": {
	//       "description": "Token for requesting successive pages of results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "url/history",
	//   "response": {
	//     "$ref": "UrlHistory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/urlshortener"
	//   ]
	// }

}

// method id "urlshortener.url.insert":

type UrlInsertCall struct {
	s    *Service
	url  *Url
	opt_ map[string]interface{}
}

// Insert: Creates a new short URL.
func (r *UrlService) Insert(url *Url) *UrlInsertCall {
	c := &UrlInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.url = url
	return c
}

func (c *UrlInsertCall) Do() (*Url, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.url)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/urlshortener/v1/", "url")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Url)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new short URL.",
	//   "httpMethod": "POST",
	//   "id": "urlshortener.url.insert",
	//   "path": "url",
	//   "request": {
	//     "$ref": "Url"
	//   },
	//   "response": {
	//     "$ref": "Url"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/urlshortener"
	//   ]
	// }

}

// method id "urlshortener.url.get":

type UrlGetCall struct {
	s        *Service
	shortUrl string
	opt_     map[string]interface{}
}

// Get: Expands a short URL or gets creation time and analytics.
func (r *UrlService) Get(shortUrl string) *UrlGetCall {
	c := &UrlGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.shortUrl = shortUrl
	return c
}

// Projection sets the optional parameter "projection": Additional
// information to return.
func (c *UrlGetCall) Projection(projection string) *UrlGetCall {
	c.opt_["projection"] = projection
	return c
}

func (c *UrlGetCall) Do() (*Url, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("shortUrl", fmt.Sprintf("%v", c.shortUrl))
	if v, ok := c.opt_["projection"]; ok {
		params.Set("projection", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/urlshortener/v1/", "url")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Url)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Expands a short URL or gets creation time and analytics.",
	//   "httpMethod": "GET",
	//   "id": "urlshortener.url.get",
	//   "parameterOrder": [
	//     "shortUrl"
	//   ],
	//   "parameters": {
	//     "projection": {
	//       "description": "Additional information to return.",
	//       "enum": [
	//         "ANALYTICS_CLICKS",
	//         "ANALYTICS_TOP_STRINGS",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Returns only click counts.",
	//         "Returns only top string counts.",
	//         "Returns the creation timestamp and all available analytics."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "shortUrl": {
	//       "description": "The short URL, including the protocol.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "url",
	//   "response": {
	//     "$ref": "Url"
	//   }
	// }

}

func cleanPathString(s string) string {
	return strings.Map(func(r int) int {
		if r >= 0x30 && r <= 0x7a {
			return r
		}
		return -1
	}, s)
}
