package tenancies

import (
	"testing"

	"github.com/rkvst/go-rkvstcommon-api-gen/quorumnetwork/v1/quorumnetwork"
	"github.com/stretchr/testify/assert"
)

// TestGetNetwork
func TestGetNetwork(t *testing.T) {

	table := []struct {
		name   string
		tenant *TenantResponse

		chainID string

		expected *quorumnetwork.QuorumNetwork
		err      error
	}{
		{
			name: "positive",
			tenant: &TenantResponse{
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						TesseraKey: "key1",
						Topic:      "topic1",
					},
					"2": {
						TesseraKey: "key2",
						Topic:      "topic2",
					},
				},
			},

			chainID: "2",

			expected: &quorumnetwork.QuorumNetwork{
				TesseraKey: "key2",
				Topic:      "topic2",
			},

			err: nil,
		},
		{
			name: "can't find network",
			tenant: &TenantResponse{
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						TesseraKey: "key1",
						Topic:      "topic1",
					},
					"2": {
						TesseraKey: "key2",
						Topic:      "topic2",
					},
				},
			},

			chainID: "3",

			expected: nil,

			err: ErrNetworkNotFound,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.tenant.GetNetwork(test.chainID)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

// TestGetNetworkTopic
func TestGetNetworkTopic(t *testing.T) {

	table := []struct {
		name   string
		tenant *TenantResponse

		chainID string

		expected string
		err      error
	}{
		{
			name: "positive",
			tenant: &TenantResponse{
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						TesseraKey: "key1",
						Topic:      "topic1",
					},
					"2": {
						TesseraKey: "key2",
						Topic:      "topic2",
					},
				},
			},

			chainID: "2",

			expected: "topic2",

			err: nil,
		},
		{
			name: "can't find network",
			tenant: &TenantResponse{
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						TesseraKey: "key1",
						Topic:      "topic1",
					},
					"2": {
						TesseraKey: "key2",
						Topic:      "topic2",
					},
				},
			},

			chainID: "3",

			expected: "",

			err: ErrNetworkNotFound,
		},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.tenant.GetNetworkTopic(test.chainID)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
