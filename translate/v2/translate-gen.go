// Package translate provides access to the Translate API.
//
// See http://code.google.com/apis/language/translate/v2/using_rest.html
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/translate/v2"
//   ...
//   translateService, err := translate.New(oauthHttpClient)
package translate

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

const apiId = "translate:v2"
const apiName = "translate"
const apiVersion = "v2"
const basePath = "https://www.googleapis.com/language/translate/"

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Detections = &DetectionsService{s: s}
	s.Languages = &LanguagesService{s: s}
	s.Translations = &TranslationsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Detections *DetectionsService

	Languages *LanguagesService

	Translations *TranslationsService
}

type DetectionsService struct {
	s *Service
}

type LanguagesService struct {
	s *Service
}

type TranslationsService struct {
	s *Service
}

type TranslationsListResponse struct {
	// Translations: Translations contains list of translation results of
	// given text
	Translations []*TranslationsResource `json:"translations,omitempty"`
}

type DetectionsResourceItem struct {
	// IsReliable: A boolean to indicate is the language detection result
	// reliable.
	IsReliable bool `json:"isReliable,omitempty"`

	// Language: The language we detect
	Language string `json:"language,omitempty"`

	// Confidence: The confidence of the detection resul of this language.
	Confidence float64 `json:"confidence,omitempty"`
}

type TranslationsResource struct {
	// TranslatedText: The translation.
	TranslatedText string `json:"translatedText,omitempty"`

	// DetectedSourceLanguage: Detected source language if source parameter
	// is unspecified.
	DetectedSourceLanguage string `json:"detectedSourceLanguage,omitempty"`
}

type LanguagesListResponse struct {
	// Languages: List of source/target languages supported by the
	// translation API. If target parameter is unspecified, the list is
	// sorted by the ASCII code point order of the language code. If target
	// parameter is specified, the list is sorted by the collation order of
	// the language name in the target language.
	Languages []*LanguagesResource `json:"languages,omitempty"`
}

type DetectionsListResponse struct {
	// Detections: A detections contains detection results of several text
	Detections [][]*DetectionsResourceItem `json:"detections,omitempty"`
}

type LanguagesResource struct {
	// Language: The language code.
	Language string `json:"language,omitempty"`

	// Name: The localized name of the language if target parameter is
	// given.
	Name string `json:"name,omitempty"`
}

// method id "language.detections.list":

type DetectionsListCall struct {
	s    *Service
	q    []string
	opt_ map[string]interface{}
}

// List: Detect the language of text.
func (r *DetectionsService) List(q []string) *DetectionsListCall {
	c := &DetectionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.q = q
	return c
}

func (c *DetectionsListCall) Do() (*DetectionsListResponse, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	for _, v := range c.q {
		params.Add("q", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/language/translate/", "v2/detect")
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
	ret := new(DetectionsListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Detect the language of text.",
	//   "httpMethod": "GET",
	//   "id": "language.detections.list",
	//   "parameterOrder": [
	//     "q"
	//   ],
	//   "parameters": {
	//     "q": {
	//       "description": "The text to detect",
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/detect",
	//   "response": {
	//     "$ref": "DetectionsListResponse"
	//   }
	// }

}

// method id "language.languages.list":

type LanguagesListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: List the source/target languages supported by the API
func (r *LanguagesService) List() *LanguagesListCall {
	c := &LanguagesListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Target sets the optional parameter "target": the language and
// collation in which the localized results should be returned
func (c *LanguagesListCall) Target(target string) *LanguagesListCall {
	c.opt_["target"] = target
	return c
}

func (c *LanguagesListCall) Do() (*LanguagesListResponse, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["target"]; ok {
		params.Set("target", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/language/translate/", "v2/languages")
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
	ret := new(LanguagesListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List the source/target languages supported by the API",
	//   "httpMethod": "GET",
	//   "id": "language.languages.list",
	//   "parameters": {
	//     "target": {
	//       "description": "the language and collation in which the localized results should be returned",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/languages",
	//   "response": {
	//     "$ref": "LanguagesListResponse"
	//   }
	// }

}

// method id "language.translations.list":

type TranslationsListCall struct {
	s      *Service
	q      []string
	target string
	opt_   map[string]interface{}
}

// List: Returns text translations from one language to another.
func (r *TranslationsService) List(q []string, target string) *TranslationsListCall {
	c := &TranslationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.q = q
	c.target = target
	return c
}

// Format sets the optional parameter "format": The format of the text
func (c *TranslationsListCall) Format(format string) *TranslationsListCall {
	c.opt_["format"] = format
	return c
}

// Cid sets the optional parameter "cid": The customization id for
// translate
func (c *TranslationsListCall) Cid(cid string) *TranslationsListCall {
	c.opt_["cid"] = cid
	return c
}

// Source sets the optional parameter "source": The source language of
// the text
func (c *TranslationsListCall) Source(source string) *TranslationsListCall {
	c.opt_["source"] = source
	return c
}

func (c *TranslationsListCall) Do() (*TranslationsListResponse, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("target", fmt.Sprintf("%v", c.target))
	for _, v := range c.q {
		params.Add("q", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["format"]; ok {
		params.Set("format", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cid"]; ok {
		params.Set("cid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["source"]; ok {
		params.Set("source", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative("https://www.googleapis.com/language/translate/", "v2")
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
	ret := new(TranslationsListResponse)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns text translations from one language to another.",
	//   "httpMethod": "GET",
	//   "id": "language.translations.list",
	//   "parameterOrder": [
	//     "q",
	//     "target"
	//   ],
	//   "parameters": {
	//     "cid": {
	//       "description": "The customization id for translate",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "format": {
	//       "description": "The format of the text",
	//       "enum": [
	//         "html",
	//         "text"
	//       ],
	//       "enumDescriptions": [
	//         "Specifies the input is in HTML",
	//         "Specifies the input is in plain textual format"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "The text to translate",
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "source": {
	//       "description": "The source language of the text",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "target": {
	//       "description": "The target language into which the text should be translated",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2",
	//   "response": {
	//     "$ref": "TranslationsListResponse"
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
