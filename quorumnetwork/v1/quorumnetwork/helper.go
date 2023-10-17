package quorumnetwork

import "fmt"

func GetTenantDefaultNetwork(networks map[string]*QuorumNetwork) (string, *QuorumNetwork, error) {

	// XXX: TODO - once we have more structure around which network is the default this should be smarter
	// than returning first network we have

	if len(networks) == 0 {
		return "", nil, fmt.Errorf("networks not configured")
	}
	if len(networks) > 1 {
		return "", nil, fmt.Errorf("multiple networks not supported")
	}

	// safely extract the only entry without depending on a hard coded notion of what chainID should be
	for chainID, network := range networks {
		return chainID, network, nil
	}
	return "", nil, fmt.Errorf("no networks configured")

}

func (x *QuorumNetwork) Clone(y *QuorumNetwork) {
	x.NodeAddress = y.NodeAddress
	x.TesseraAddress = y.TesseraAddress
	x.TesseraKey = y.TesseraKey
	x.IsPrivate = y.IsPrivate
}
