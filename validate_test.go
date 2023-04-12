package main

import (
	"github.com/bmizerany/assert"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestValidateOwner(t *testing.T) {
	testCases := []struct {
		desc     string
		ing      v1alpha1.IngressRoute
		route    Route
		expected bool
	}{
		{
			desc: "simple test",
			ing: v1alpha1.IngressRoute{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
				},
			},
			route:    Route{Owner: "default-test-9dd2d013e4859514de1a@kubernetescrd"},
			expected: true,
		},
		{
			desc: "error test",
			ing: v1alpha1.IngressRoute{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default",
					Namespace: "test",
				},
			},
			route:    Route{Owner: "default-test-9dd2d013e4859514de1a@kubernetescrd"},
			expected: false,
		},
		{
			desc: "error test 2",
			ing: v1alpha1.IngressRoute{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-9",
					Namespace: "default",
				},
			},
			route:    Route{Owner: "default-test-9dd2d013e4859514de1a@kubernetescrd"},
			expected: false,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			output := validateOwner(test.ing, test.route)
			assert.Equal(t, test.expected, output)
		})
	}
}
