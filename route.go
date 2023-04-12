package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Route struct {
	Rule  map[string]string
	Owner string
}

type RouteOption func(route Route) Route

func (r *Route) IsEmpty() bool {
	return r.Rule == nil
}

func SplitMatchPath(match string, options ...RouteOption) []Route {
	segments := splitAndCombine(match)

	var routes []Route
	for _, segment := range segments {
		rules := parseSingleRule(segment)
		r := Route{Rule: rules}
		for _, option := range options {
			r = option(r)
		}
		routes = append(routes, r)
	}

	return routes
}

func splitAndCombine(input string) []string {
	andRegEx := regexp.MustCompile(`\s*&&\s*`)
	orRegEx := regexp.MustCompile(`\s*\|\|\s*`)

	andSegments := andRegEx.Split(input, -1)
	var result []string

	if len(andSegments) > 1 {
		orCombinations := make([][]string, len(andSegments))

		for i, andSegment := range andSegments {
			orCombinations[i] = orRegEx.Split(andSegment, -1)
		}

		for _, a := range orCombinations[0] {
			for _, b := range orCombinations[1] {
				result = append(result, fmt.Sprintf("%s && %s", a, b))
			}
		}

		for i := 2; i < len(orCombinations); i++ {
			var newResult []string
			for _, a := range result {
				for _, b := range orCombinations[i] {
					newResult = append(newResult, fmt.Sprintf("%s && %s", a, b))
				}
			}
			result = newResult
		}

	} else {
		result = orRegEx.Split(input, -1)
	}

	return result
}

func parseSingleRule(input string) map[string]string {
	ruleRegEx := regexp.MustCompile(`([\!A-Za-z]+)\(` + "`" + `([^` + "`" + `]+)` + "`" + `\)`)

	matches := ruleRegEx.FindAllStringSubmatch(input, -1)

	rulesMap := make(map[string]string)
	for _, match := range matches {
		if len(match) > 2 {
			rulesMap[match[1]] = match[2]
		}
	}
	return rulesMap
}

func (r *Route) ToString() string {
	byteData, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(byteData)
}
