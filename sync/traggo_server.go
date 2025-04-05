package sync

import (
	"fmt"
	"net/url"
	"strings"
)

type TraggoServer struct {
	Host string
	User string
	Pass string
}

func (s *TraggoServer) CheckValid() error {

	url, err := url.ParseRequestURI(s.Host)
	if err != nil {
		return err
	}
	switch strings.ToLower(url.Scheme) {
	case "http":
	case "https":
		return nil
	default:
		return fmt.Errorf("invalid url scheme: %s", url.Scheme)
	}
	return nil
}
