package main

import (
	"encoding/json"
	"fmt"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"io"
	"net/http"
)

type ServerAddr struct {
	Protocol string
	Host     string
	Port     int
}

func ListServices() ([]dynamic.Router, error) {
	var results []dynamic.Router
	resp, err := http.DefaultClient.Get(
		fmt.Sprintf("%s://%s:%d/api/http/routers", server.Protocol, server.Host, server.Port))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
