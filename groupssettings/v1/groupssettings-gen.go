// Package groupssettings provides access to the .
//
// Usage example:
//
//   import "google-api-go-client.googlecode.com/hg/groupssettings/v1"
//   ...
//   groupssettingsService, err := groupssettings.New(oauthHttpClient)
package groupssettings

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

const apiId = "groupssettings:v1"
const apiName = "groupssettings"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/groups/v1/groups/"

// OAuth2 scopes used by this API.
const (
	// View and manage the settings of a Google Apps Group
	AppsGroupsSettingsScope = "https://www.googleapis.com/auth/apps.groups.settings"
)

func New(client *http.Client) (*Service, os.Error) {
	if client == nil {
		return nil, os.NewError("client is nil")
	}
	s := &Service{client: client}
	s.Groups = &GroupsService{s: s}
	return s, nil
}

type Service struct {
	client *http.Client

	Groups *GroupsService
}

type GroupsService struct {
	s *Service
}

type Groups struct {
	// Email: Email id of the group
	Email string `json:"email,omitempty"`

	// DefaultMessageDenyNotificationText: Default message deny notification
	// message
	DefaultMessageDenyNotificationText string `json:"defaultMessageDenyNotificationText,omitempty"`

	// IsArchived: If the contents of the group are archived.
	IsArchived string `json:"isArchived,omitempty"`

	// MessageDisplayFont: Default message display font. Possible values
	// are: DEFAULT_FONT FIXED_WIDTH_FONT
	MessageDisplayFont string `json:"messageDisplayFont,omitempty"`

	// WhoCanJoin: Permissions to join the group. Possible values are:
	// ANYONE_CAN_JOIN ALL_IN_DOMAIN_CAN_JOIN INVITED_CAN_JOIN
	// CAN_REQUEST_TO_JOIN
	WhoCanJoin string `json:"whoCanJoin,omitempty"`

	// Name: Name of the Group
	Name string `json:"name,omitempty"`

	// ReplyTo: Whome should the default reply to a message go to. Possible
	// values are: REPLY_TO_CUSTOM REPLY_TO_SENDER REPLY_TO_LIST
	// REPLY_TO_OWNER REPLY_TO_IGNORE REPLY_TO_MANAGERS
	ReplyTo string `json:"replyTo,omitempty"`

	// MaxMessageBytes: Maximum message size allowed.
	MaxMessageBytes int64 `json:"maxMessageBytes,omitempty"`

	// CustomReplyTo: Default email to which reply to any message should go.
	CustomReplyTo string `json:"customReplyTo,omitempty"`

	// ArchiveOnly: If the group is archive only
	ArchiveOnly string `json:"archiveOnly,omitempty"`

	// AllowExternalMembers: Are external members allowed to join the group.
	AllowExternalMembers string `json:"allowExternalMembers,omitempty"`

	// Kind: The type of the resource.
	Kind string `json:"kind,omitempty"`

	// PrimaryLanguage: Primary language for the group.
	PrimaryLanguage string `json:"primaryLanguage,omitempty"`

	// WhoCanPostMessage: Permissions to post messages to the group.
	// Possible values are: NONE_CAN_POST ALL_MANAGERS_CAN_POST
	// ALL_MEMBERS_CAN_POST ALL_IN_DOMAIN_CAN_POST ANYONE_CAN_POST
	WhoCanPostMessage string `json:"whoCanPostMessage,omitempty"`

	// ShowInGroupDirectory: Is the group listed in groups directory
	ShowInGroupDirectory string `json:"showInGroupDirectory,omitempty"`

	// AllowWebPosting: If posting from web is allowed.
	AllowWebPosting string `json:"allowWebPosting,omitempty"`

	// MembersCanPostAsTheGroup: Can members post using the group email
	// address.
	MembersCanPostAsTheGroup string `json:"membersCanPostAsTheGroup,omitempty"`

	// WhoCanViewGroup: Permissions to view group. Possbile values are:
	// ANYONE_CAN_VIEW ALL_IN_DOMAIN_CAN_VIEW ALL_MEMBERS_CAN_VIEW
	// ALL_MANAGERS_CAN_VIEW
	WhoCanViewGroup string `json:"whoCanViewGroup,omitempty"`

	// AllowGoogleCommunication: Is google allowed to contact admins.
	AllowGoogleCommunication string `json:"allowGoogleCommunication,omitempty"`

	// MessageModerationLevel: Moderation level for messages. Possible
	// values are: MODERATE_ALL_MESSAGES MODERATE_NON_MEMBERS
	// MODERATE_NEW_MEMBERS MODERATE_NONE
	MessageModerationLevel string `json:"messageModerationLevel,omitempty"`

	// SendMessageDenyNotification: Should the member be notified if his
	// message is denied by owner.
	SendMessageDenyNotification string `json:"sendMessageDenyNotification,omitempty"`

	// WhoCanViewMembership: Permissions to view membership. Possbile values
	// are: ALL_IN_DOMAIN_CAN_VIEW ALL_MEMBERS_CAN_VIEW
	// ALL_MANAGERS_CAN_VIEW
	WhoCanViewMembership string `json:"whoCanViewMembership,omitempty"`

	// WhoCanInvite: Permissions to invite members. Possbile values are:
	// ALL_MEMBERS_CAN_INVITE ALL_MANAGERS_CAN_INVITE
	WhoCanInvite string `json:"whoCanInvite,omitempty"`

	// Description: Description of the group
	Description string `json:"description,omitempty"`
}

// method id "groupsSettings.groups.update":

type GroupsUpdateCall struct {
	s             *Service
	groupUniqueId string
	groups        *Groups
	opt_          map[string]interface{}
}

// Update: Updates an existing resource.
func (r *GroupsService) Update(groupUniqueId string, groups *Groups) *GroupsUpdateCall {
	c := &GroupsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.groupUniqueId = groupUniqueId
	c.groups = groups
	return c
}

func (c *GroupsUpdateCall) Do() (*Groups, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.groups)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/groups/v1/groups/", "{groupUniqueId}")
	urls = strings.Replace(urls, "{groupUniqueId}", cleanPathString(c.groupUniqueId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Groups)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing resource.",
	//   "httpMethod": "PUT",
	//   "id": "groupsSettings.groups.update",
	//   "parameterOrder": [
	//     "groupUniqueId"
	//   ],
	//   "parameters": {
	//     "groupUniqueId": {
	//       "description": "The resource ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{groupUniqueId}",
	//   "request": {
	//     "$ref": "Groups"
	//   },
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.groups.settings"
	//   ]
	// }

}

// method id "groupsSettings.groups.get":

type GroupsGetCall struct {
	s             *Service
	groupUniqueId string
	opt_          map[string]interface{}
}

// Get: Gets one resource by id.
func (r *GroupsService) Get(groupUniqueId string) *GroupsGetCall {
	c := &GroupsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.groupUniqueId = groupUniqueId
	return c
}

func (c *GroupsGetCall) Do() (*Groups, os.Error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/groups/v1/groups/", "{groupUniqueId}")
	urls = strings.Replace(urls, "{groupUniqueId}", cleanPathString(c.groupUniqueId), 1)
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
	ret := new(Groups)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets one resource by id.",
	//   "httpMethod": "GET",
	//   "id": "groupsSettings.groups.get",
	//   "parameterOrder": [
	//     "groupUniqueId"
	//   ],
	//   "parameters": {
	//     "groupUniqueId": {
	//       "description": "The resource ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{groupUniqueId}",
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.groups.settings"
	//   ]
	// }

}

// method id "groupsSettings.groups.patch":

type GroupsPatchCall struct {
	s             *Service
	groupUniqueId string
	groups        *Groups
	opt_          map[string]interface{}
}

// Patch: Updates an existing resource. This method supports patch
// semantics.
func (r *GroupsService) Patch(groupUniqueId string, groups *Groups) *GroupsPatchCall {
	c := &GroupsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.groupUniqueId = groupUniqueId
	c.groups = groups
	return c
}

func (c *GroupsPatchCall) Do() (*Groups, os.Error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.groups)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative("https://www.googleapis.com/groups/v1/groups/", "{groupUniqueId}")
	urls = strings.Replace(urls, "{groupUniqueId}", cleanPathString(c.groupUniqueId), 1)
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Groups)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "groupsSettings.groups.patch",
	//   "parameterOrder": [
	//     "groupUniqueId"
	//   ],
	//   "parameters": {
	//     "groupUniqueId": {
	//       "description": "The resource ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{groupUniqueId}",
	//   "request": {
	//     "$ref": "Groups"
	//   },
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.groups.settings"
	//   ]
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
