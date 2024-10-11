package client

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/minnowo/traggo-sync/graph"
	"github.com/rs/zerolog/log"

	"github.com/Khan/genqlient/graphql"
)

type TraggoHttpClient struct {
	httpClient    http.Client
	graphqlClient graphql.Client
	traggoHost    *url.URL
}

func NewTraggoClient(traggoUrl string) (*TraggoHttpClient, error) {

	url, err := url.Parse(traggoUrl)

	if err != nil {
		return nil, err
	}

	// TODO: suffix list
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})

	if err != nil {
		return nil, err
	}

	httpClient := http.Client{Jar: jar}

	graphClient := graphql.NewClient(url.JoinPath(GraphqlPath).String(), &httpClient)

	log.Info().
		Str("traggo", url.String()).
		Msg("Creating new TraggoHttpClient")

	c := &TraggoHttpClient{
		httpClient:    httpClient,
		graphqlClient: graphClient,
		traggoHost:    url,
	}

	return c, nil
}

func (c *TraggoHttpClient) Ql() graphql.Client {
	return c.graphqlClient
}

func (c *TraggoHttpClient) Url() string {
	return c.traggoHost.String()
}

func (c *TraggoHttpClient) Login(username, password string) error {

	var res *graph.LoginResponse

	res, err := graph.Login(context.Background(), c.Ql(), username, password, ClientName, graph.DeviceTypeLongexpiry)

	if err != nil {

		return err
	}

	u := res.Login.User

	log.Info().
		Int("id", u.Id).
		Str("user", u.Name).
		Str("traggo", c.Url()).
		Bool("admin", u.Admin).
		Msg("Logged into Traggo instance")

	return nil
}

func (c *TraggoHttpClient) LogVersion() error {

	var version *graph.GetVersionResponse

	version, err := graph.GetVersion(context.Background(), c.Ql())

	if err != nil {
		return err
	}

	log.Info().
		Str("version", version.Version.Name).
		Str("commit", version.Version.Commit).
		Str("built", version.Version.BuildDate).
		Str("traggo", c.Url()).
		Msg("Traggo information")

	return nil
}
