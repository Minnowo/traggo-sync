// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graph

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

type DeviceType string

const (
	DeviceTypeLongexpiry  DeviceType = "LongExpiry"
	DeviceTypeShortexpiry DeviceType = "ShortExpiry"
	DeviceTypeNoexpiry    DeviceType = "NoExpiry"
)

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
