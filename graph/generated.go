// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graph

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// CreateTagCreateTagTagDefinition includes the requested fields of the GraphQL type TagDefinition.
type CreateTagCreateTagTagDefinition struct {
	Key string `json:"key"`
}

// GetKey returns CreateTagCreateTagTagDefinition.Key, and is useful for accessing the field via an interface.
func (v *CreateTagCreateTagTagDefinition) GetKey() string { return v.Key }

// CreateTagResponse is returned by CreateTag on success.
type CreateTagResponse struct {
	CreateTag CreateTagCreateTagTagDefinition `json:"createTag"`
}

// GetCreateTag returns CreateTagResponse.CreateTag, and is useful for accessing the field via an interface.
func (v *CreateTagResponse) GetCreateTag() CreateTagCreateTagTagDefinition { return v.CreateTag }

// CreateTimeSpanResponse is returned by CreateTimeSpan on success.
type CreateTimeSpanResponse struct {
	CreateTimeSpan TimeSpan `json:"createTimeSpan"`
}

// GetCreateTimeSpan returns CreateTimeSpanResponse.CreateTimeSpan, and is useful for accessing the field via an interface.
func (v *CreateTimeSpanResponse) GetCreateTimeSpan() TimeSpan { return v.CreateTimeSpan }

type DeviceType string

const (
	DeviceTypeLongexpiry  DeviceType = "LongExpiry"
	DeviceTypeShortexpiry DeviceType = "ShortExpiry"
	DeviceTypeNoexpiry    DeviceType = "NoExpiry"
)

// GetAllTagsResponse is returned by GetAllTags on success.
type GetAllTagsResponse struct {
	Tags []GetAllTagsTagsTagDefinition `json:"tags"`
}

// GetTags returns GetAllTagsResponse.Tags, and is useful for accessing the field via an interface.
func (v *GetAllTagsResponse) GetTags() []GetAllTagsTagsTagDefinition { return v.Tags }

// GetAllTagsTagsTagDefinition includes the requested fields of the GraphQL type TagDefinition.
type GetAllTagsTagsTagDefinition struct {
	Color  string `json:"color"`
	Key    string `json:"key"`
	Usages int    `json:"usages"`
}

// GetColor returns GetAllTagsTagsTagDefinition.Color, and is useful for accessing the field via an interface.
func (v *GetAllTagsTagsTagDefinition) GetColor() string { return v.Color }

// GetKey returns GetAllTagsTagsTagDefinition.Key, and is useful for accessing the field via an interface.
func (v *GetAllTagsTagsTagDefinition) GetKey() string { return v.Key }

// GetUsages returns GetAllTagsTagsTagDefinition.Usages, and is useful for accessing the field via an interface.
func (v *GetAllTagsTagsTagDefinition) GetUsages() int { return v.Usages }

// GetCurrentUserCurrentUser includes the requested fields of the GraphQL type User.
type GetCurrentUserCurrentUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

// GetId returns GetCurrentUserCurrentUser.Id, and is useful for accessing the field via an interface.
func (v *GetCurrentUserCurrentUser) GetId() int { return v.Id }

// GetName returns GetCurrentUserCurrentUser.Name, and is useful for accessing the field via an interface.
func (v *GetCurrentUserCurrentUser) GetName() string { return v.Name }

// GetAdmin returns GetCurrentUserCurrentUser.Admin, and is useful for accessing the field via an interface.
func (v *GetCurrentUserCurrentUser) GetAdmin() bool { return v.Admin }

// GetCurrentUserDevicesDevice includes the requested fields of the GraphQL type Device.
type GetCurrentUserDevicesDevice struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Type      DeviceType `json:"type"`
	CreatedAt time.Time  `json:"createdAt"`
	ActiveAt  time.Time  `json:"activeAt"`
}

// GetId returns GetCurrentUserDevicesDevice.Id, and is useful for accessing the field via an interface.
func (v *GetCurrentUserDevicesDevice) GetId() int { return v.Id }

// GetName returns GetCurrentUserDevicesDevice.Name, and is useful for accessing the field via an interface.
func (v *GetCurrentUserDevicesDevice) GetName() string { return v.Name }

// GetType returns GetCurrentUserDevicesDevice.Type, and is useful for accessing the field via an interface.
func (v *GetCurrentUserDevicesDevice) GetType() DeviceType { return v.Type }

// GetCreatedAt returns GetCurrentUserDevicesDevice.CreatedAt, and is useful for accessing the field via an interface.
func (v *GetCurrentUserDevicesDevice) GetCreatedAt() time.Time { return v.CreatedAt }

// GetActiveAt returns GetCurrentUserDevicesDevice.ActiveAt, and is useful for accessing the field via an interface.
func (v *GetCurrentUserDevicesDevice) GetActiveAt() time.Time { return v.ActiveAt }

// GetCurrentUserResponse is returned by GetCurrentUser on success.
type GetCurrentUserResponse struct {
	CurrentUser GetCurrentUserCurrentUser     `json:"currentUser"`
	Devices     []GetCurrentUserDevicesDevice `json:"devices"`
}

// GetCurrentUser returns GetCurrentUserResponse.CurrentUser, and is useful for accessing the field via an interface.
func (v *GetCurrentUserResponse) GetCurrentUser() GetCurrentUserCurrentUser { return v.CurrentUser }

// GetDevices returns GetCurrentUserResponse.Devices, and is useful for accessing the field via an interface.
func (v *GetCurrentUserResponse) GetDevices() []GetCurrentUserDevicesDevice { return v.Devices }

// GetTimeSpansResponse is returned by GetTimeSpans on success.
type GetTimeSpansResponse struct {
	TimeSpans GetTimeSpansTimeSpansPagedTimeSpans `json:"timeSpans"`
}

// GetTimeSpans returns GetTimeSpansResponse.TimeSpans, and is useful for accessing the field via an interface.
func (v *GetTimeSpansResponse) GetTimeSpans() GetTimeSpansTimeSpansPagedTimeSpans { return v.TimeSpans }

// GetTimeSpansTimeSpansPagedTimeSpans includes the requested fields of the GraphQL type PagedTimeSpans.
type GetTimeSpansTimeSpansPagedTimeSpans struct {
	TimeSpans []TimeSpan                                `json:"timeSpans"`
	Cursor    GetTimeSpansTimeSpansPagedTimeSpansCursor `json:"cursor"`
}

// GetTimeSpans returns GetTimeSpansTimeSpansPagedTimeSpans.TimeSpans, and is useful for accessing the field via an interface.
func (v *GetTimeSpansTimeSpansPagedTimeSpans) GetTimeSpans() []TimeSpan { return v.TimeSpans }

// GetCursor returns GetTimeSpansTimeSpansPagedTimeSpans.Cursor, and is useful for accessing the field via an interface.
func (v *GetTimeSpansTimeSpansPagedTimeSpans) GetCursor() GetTimeSpansTimeSpansPagedTimeSpansCursor {
	return v.Cursor
}

// GetTimeSpansTimeSpansPagedTimeSpansCursor includes the requested fields of the GraphQL type Cursor.
type GetTimeSpansTimeSpansPagedTimeSpansCursor struct {
	HasMore  bool `json:"hasMore"`
	Offset   int  `json:"offset"`
	StartId  int  `json:"startId"`
	PageSize int  `json:"pageSize"`
}

// GetHasMore returns GetTimeSpansTimeSpansPagedTimeSpansCursor.HasMore, and is useful for accessing the field via an interface.
func (v *GetTimeSpansTimeSpansPagedTimeSpansCursor) GetHasMore() bool { return v.HasMore }

// GetOffset returns GetTimeSpansTimeSpansPagedTimeSpansCursor.Offset, and is useful for accessing the field via an interface.
func (v *GetTimeSpansTimeSpansPagedTimeSpansCursor) GetOffset() int { return v.Offset }

// GetStartId returns GetTimeSpansTimeSpansPagedTimeSpansCursor.StartId, and is useful for accessing the field via an interface.
func (v *GetTimeSpansTimeSpansPagedTimeSpansCursor) GetStartId() int { return v.StartId }

// GetPageSize returns GetTimeSpansTimeSpansPagedTimeSpansCursor.PageSize, and is useful for accessing the field via an interface.
func (v *GetTimeSpansTimeSpansPagedTimeSpansCursor) GetPageSize() int { return v.PageSize }

// GetVersionResponse is returned by GetVersion on success.
type GetVersionResponse struct {
	Version GetVersionVersion `json:"version"`
}

// GetVersion returns GetVersionResponse.Version, and is useful for accessing the field via an interface.
func (v *GetVersionResponse) GetVersion() GetVersionVersion { return v.Version }

// GetVersionVersion includes the requested fields of the GraphQL type Version.
type GetVersionVersion struct {
	Name      string `json:"name"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
}

// GetName returns GetVersionVersion.Name, and is useful for accessing the field via an interface.
func (v *GetVersionVersion) GetName() string { return v.Name }

// GetCommit returns GetVersionVersion.Commit, and is useful for accessing the field via an interface.
func (v *GetVersionVersion) GetCommit() string { return v.Commit }

// GetBuildDate returns GetVersionVersion.BuildDate, and is useful for accessing the field via an interface.
func (v *GetVersionVersion) GetBuildDate() string { return v.BuildDate }

// LoginLogin includes the requested fields of the GraphQL type Login.
type LoginLogin struct {
	User     LoginLoginUser `json:"user"`
	Typename string         `json:"__typename"`
}

// GetUser returns LoginLogin.User, and is useful for accessing the field via an interface.
func (v *LoginLogin) GetUser() LoginLoginUser { return v.User }

// GetTypename returns LoginLogin.Typename, and is useful for accessing the field via an interface.
func (v *LoginLogin) GetTypename() string { return v.Typename }

// LoginLoginUser includes the requested fields of the GraphQL type User.
type LoginLoginUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Admin    bool   `json:"admin"`
	Typename string `json:"__typename"`
}

// GetId returns LoginLoginUser.Id, and is useful for accessing the field via an interface.
func (v *LoginLoginUser) GetId() int { return v.Id }

// GetName returns LoginLoginUser.Name, and is useful for accessing the field via an interface.
func (v *LoginLoginUser) GetName() string { return v.Name }

// GetAdmin returns LoginLoginUser.Admin, and is useful for accessing the field via an interface.
func (v *LoginLoginUser) GetAdmin() bool { return v.Admin }

// GetTypename returns LoginLoginUser.Typename, and is useful for accessing the field via an interface.
func (v *LoginLoginUser) GetTypename() string { return v.Typename }

// LoginResponse is returned by Login on success.
type LoginResponse struct {
	Login LoginLogin `json:"login"`
}

// GetLogin returns LoginResponse.Login, and is useful for accessing the field via an interface.
func (v *LoginResponse) GetLogin() LoginLogin { return v.Login }

// __CreateTagInput is used internally by genqlient
type __CreateTagInput struct {
	Key   string `json:"key"`
	Color string `json:"color"`
}

// GetKey returns __CreateTagInput.Key, and is useful for accessing the field via an interface.
func (v *__CreateTagInput) GetKey() string { return v.Key }

// GetColor returns __CreateTagInput.Color, and is useful for accessing the field via an interface.
func (v *__CreateTagInput) GetColor() string { return v.Color }

// __CreateTimeSpanInput is used internally by genqlient
type __CreateTimeSpanInput struct {
	Start time.Time     `json:"start"`
	End   time.Time     `json:"end"`
	Tags  []TimeSpanTag `json:"tags"`
	Note  string        `json:"note"`
}

// GetStart returns __CreateTimeSpanInput.Start, and is useful for accessing the field via an interface.
func (v *__CreateTimeSpanInput) GetStart() time.Time { return v.Start }

// GetEnd returns __CreateTimeSpanInput.End, and is useful for accessing the field via an interface.
func (v *__CreateTimeSpanInput) GetEnd() time.Time { return v.End }

// GetTags returns __CreateTimeSpanInput.Tags, and is useful for accessing the field via an interface.
func (v *__CreateTimeSpanInput) GetTags() []TimeSpanTag { return v.Tags }

// GetNote returns __CreateTimeSpanInput.Note, and is useful for accessing the field via an interface.
func (v *__CreateTimeSpanInput) GetNote() string { return v.Note }

// __GetTimeSpansInput is used internally by genqlient
type __GetTimeSpansInput struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

// GetFrom returns __GetTimeSpansInput.From, and is useful for accessing the field via an interface.
func (v *__GetTimeSpansInput) GetFrom() time.Time { return v.From }

// GetTo returns __GetTimeSpansInput.To, and is useful for accessing the field via an interface.
func (v *__GetTimeSpansInput) GetTo() time.Time { return v.To }

// __LoginInput is used internally by genqlient
type __LoginInput struct {
	Name       string     `json:"name"`
	Pass       string     `json:"pass"`
	DeviceName string     `json:"deviceName"`
	DeviceType DeviceType `json:"deviceType"`
}

// GetName returns __LoginInput.Name, and is useful for accessing the field via an interface.
func (v *__LoginInput) GetName() string { return v.Name }

// GetPass returns __LoginInput.Pass, and is useful for accessing the field via an interface.
func (v *__LoginInput) GetPass() string { return v.Pass }

// GetDeviceName returns __LoginInput.DeviceName, and is useful for accessing the field via an interface.
func (v *__LoginInput) GetDeviceName() string { return v.DeviceName }

// GetDeviceType returns __LoginInput.DeviceType, and is useful for accessing the field via an interface.
func (v *__LoginInput) GetDeviceType() DeviceType { return v.DeviceType }

// The query or mutation executed by CreateTag.
const CreateTag_Operation = `
mutation CreateTag ($key: String!, $color: String!) {
	createTag(key: $key, color: $color) {
		key
	}
}
`

func CreateTag(
	ctx_ context.Context,
	client_ graphql.Client,
	key string,
	color string,
) (*CreateTagResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateTag",
		Query:  CreateTag_Operation,
		Variables: &__CreateTagInput{
			Key:   key,
			Color: color,
		},
	}
	var err_ error

	var data_ CreateTagResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by CreateTimeSpan.
const CreateTimeSpan_Operation = `
mutation CreateTimeSpan ($start: Time!, $end: Time, $tags: [InputTimeSpanTag!]!, $note: String!) {
	createTimeSpan(start: $start, end: $end, tags: $tags, note: $note) {
		id
		start
		end
		note
		tags {
			key
			value
		}
	}
}
`

func CreateTimeSpan(
	ctx_ context.Context,
	client_ graphql.Client,
	start time.Time,
	end time.Time,
	tags []TimeSpanTag,
	note string,
) (*CreateTimeSpanResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateTimeSpan",
		Query:  CreateTimeSpan_Operation,
		Variables: &__CreateTimeSpanInput{
			Start: start,
			End:   end,
			Tags:  tags,
			Note:  note,
		},
	}
	var err_ error

	var data_ CreateTimeSpanResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetAllTags.
const GetAllTags_Operation = `
query GetAllTags {
	tags {
		color
		key
		usages
	}
}
`

func GetAllTags(
	ctx_ context.Context,
	client_ graphql.Client,
) (*GetAllTagsResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetAllTags",
		Query:  GetAllTags_Operation,
	}
	var err_ error

	var data_ GetAllTagsResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetCurrentUser.
const GetCurrentUser_Operation = `
query GetCurrentUser {
	currentUser {
		id
		name
		admin
	}
	devices {
		id
		name
		type
		createdAt
		activeAt
	}
}
`

func GetCurrentUser(
	ctx_ context.Context,
	client_ graphql.Client,
) (*GetCurrentUserResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetCurrentUser",
		Query:  GetCurrentUser_Operation,
	}
	var err_ error

	var data_ GetCurrentUserResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetTimeSpans.
const GetTimeSpans_Operation = `
query GetTimeSpans ($from: Time, $to: Time) {
	timeSpans(fromInclusive: $from, toInclusive: $to) {
		timeSpans {
			id
			start
			end
			oldStart
			note
			tags {
				key
				value
			}
		}
		cursor {
			hasMore
			offset
			startId
			pageSize
		}
	}
}
`

func GetTimeSpans(
	ctx_ context.Context,
	client_ graphql.Client,
	from time.Time,
	to time.Time,
) (*GetTimeSpansResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetTimeSpans",
		Query:  GetTimeSpans_Operation,
		Variables: &__GetTimeSpansInput{
			From: from,
			To:   to,
		},
	}
	var err_ error

	var data_ GetTimeSpansResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetVersion.
const GetVersion_Operation = `
query GetVersion {
	version {
		name
		commit
		buildDate
	}
}
`

func GetVersion(
	ctx_ context.Context,
	client_ graphql.Client,
) (*GetVersionResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetVersion",
		Query:  GetVersion_Operation,
	}
	var err_ error

	var data_ GetVersionResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by Login.
const Login_Operation = `
mutation Login ($name: String!, $pass: String!, $deviceName: String!, $deviceType: DeviceType!) {
	login(username: $name, pass: $pass, deviceName: $deviceName, type: $deviceType, cookie: true) {
		user {
			id
			name
			admin
			__typename
		}
		__typename
	}
}
`

func Login(
	ctx_ context.Context,
	client_ graphql.Client,
	name string,
	pass string,
	deviceName string,
	deviceType DeviceType,
) (*LoginResponse, error) {
	req_ := &graphql.Request{
		OpName: "Login",
		Query:  Login_Operation,
		Variables: &__LoginInput{
			Name:       name,
			Pass:       pass,
			DeviceName: deviceName,
			DeviceType: deviceType,
		},
	}
	var err_ error

	var data_ LoginResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
