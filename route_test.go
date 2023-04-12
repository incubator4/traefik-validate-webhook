package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSplitMatchPath(t *testing.T) {
	testCases := []struct {
		desc     string
		match    string
		expected []Route
	}{
		{
			desc:  "Simple Route",
			match: "Host(`foo.com`)",
			expected: []Route{
				{
					"Host": "foo.com",
				},
			},
		},
		{
			desc:  "Host with multi PathPrefix routes",
			match: "Host(`foo.com`) && (PathPrefix(`/a`) || PathPrefix(`/b`))",
			expected: []Route{
				{
					"Host":       "foo.com",
					"PathPrefix": "/a",
				},
				{
					"Host":       "foo.com",
					"PathPrefix": "/b",
				},
			},
		},
		{
			desc:  "Multi Host routes",
			match: "(Host(`foo.com`) || Host(`bar.com`))",
			expected: []Route{
				{
					"Host": "foo.com",
				},
				{
					"Host": "bar.com",
				},
			},
		},
		{
			desc:  "Multi Host routes with Prefix",
			match: "(Host(`foo.com`) || !Host(`bar.com`)) && PathPrefix(`/a`)",
			expected: []Route{
				{
					"Host":       "foo.com",
					"PathPrefix": "/a",
				},
				{
					"!Host":      "bar.com",
					"PathPrefix": "/a",
				},
			},
		},
	}
	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			output := SplitMatchPath(test.match)
			assert.Equal(t, test.expected, output)
		})
	}
}
