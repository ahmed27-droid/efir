package proxy

import (
	"net/http/httputil"
	"net/url"
)

func CreateProxy(target string) (*httputil.ReverseProxy, error) { 
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(u), nil
}
