package http

import (
	"net/url"
)

func BuildURL(scheme, host, path string) string {
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	return url.String()
}
