package tenancies

import (
	"errors"

	"github.com/rkvst/go-rkvstcommon-api-gen/quorumnetwork/v1/quorumnetwork"
)

var (
	ErrNetworkNotFound = errors.New("quorum network not found")
)

// TesseraPubKeys returns the tessera public key for each configured chainID.
// Note that we only support 1 chainID per tenant today. Callers may assume or require len(pubs) == 1
func (r *TenantResponse) NetworkTesseraPubKeys() []string {

	var pubs []string
	for _, net := range r.QuorumNetworks {
		if net.TesseraKey == "" {
			continue
		}
		pubs = append(pubs, net.TesseraKey)
	}
	return pubs
}

// GetNetwork gets the corresponding quorum network, given the chainID.
func (r *TenantResponse) GetNetwork(chainID string) (*quorumnetwork.QuorumNetwork, error) {

	for id, network := range r.QuorumNetworks {

		if id != chainID {
			continue
		}

		return network, nil
	}

	return nil, ErrNetworkNotFound

}

// GetNetworkTopic gets a quorum network topic
func (r *TenantResponse) GetNetworkTopic(chainID string) (string, error) {
	network, err := r.GetNetwork(chainID)
	if err != nil {
		return "", err
	}

	return network.Topic, nil
}
