package main

import (
	"encoding/json"
	"fmt"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/config/static"
	"io"
	"net/http"
)

type ServerAddr struct {
	Protocol string
	Host     string
	Port     int
}

type RouteInfo struct {
	dynamic.Router
	Name string `json:"name,omitempty"`
}

type EntryPoint struct {
	static.EntryPoint
	Name string `json:"name,omitempty"`
}

func ListServices() ([]RouteInfo, error) {
	var results []RouteInfo
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

func ListEntryPoints() ([]EntryPoint, error) {
	var results []EntryPoint
	resp, err := http.DefaultClient.Get(
		fmt.Sprintf("%s://%s:%d/api/entrypoints", server.Protocol, server.Host, server.Port))
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
