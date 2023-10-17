package quorumnetwork

// This commentary was originally in a unittest for validating the described
// assumptions for the tenancy record quorumnetwork configuration in
// tenanciesv1. It was pretty ineficient to do that on every request so we rely
// (mostly) on verification at the time tenants are assigned to nodes.

/*
	Each chain will/may have a single 'free' shared node. tsp data goes there
	by default. on that same chain a single customer may pay for a private
	node. That private node customer can then chose to have their tsp data go
	to a private data store.

	Private node customers may allow/want to onboard partners and 'share' their
	private node with them. in this scenario all customers on that private node
	must be configured to use the same storage for the tsp asset data. (node
	aka tessera public key).

	It is not yet fully resolved how we will onboard

	Scenario: swiss bank and their cleaning company. Swiss bank pay for private
	node and there form *must* get their own tessera public key.  The swiss
	bank set things up so that events interacting with the cleaning company use
	simple_hash asset. This means that the asset and event data goes on the
	"tenant storage path".

	The swiss bank are subject to data residency rules. By default TSP asset
	data will be on the shared node (aka but not always freemium).  To support
	the data residency case we need to support private 'node' storage for the
	TSP path. As it is an extra cost, we want this to be optional. Hence we
	have both tessera key and topic in the network configuration. For a private
	node customer the tessara key will be unique. For all tenants 'co-resident'
	on that private node they will be the same. The private node customer gets
	to choose (for themselves and all co-resident tenants), whether to use the
	default shared storage OR the private store for TSP (simple hash) assets.

	The following cover the expected legitimate storage topic selection cases
	in the tenant record. They are specifically about the relationship between
	customers who have paid for a node and customers who have not. And the
	necessary support for customers being logically resident (same tessera pub)
	on a private node customer node.
*/

/*
	co-resident tenants correctly configured (same topic name)

	swiss bank + cleaning company are correctly configured
	co-resident tenants. both explicitly configured for the shared
	topic for tsp.

	swiss bank is a private node customer and has selected shared
	storage for tsp.  Cleaning company is a customer specifically on
	boarded to be resident on the same node as as the swiss bank (same
	tessera pub). Here the cleaning company's tenant record correctly
	uses the shared topic (matching the choice of the swiss bank they
	are working with)

	request: &tenancies.TopicFromPubKeyRequest{
		ChainID:       "1",
		TesseraPubKey: "swissbank-key",
	},

	tenants: []*tenanciesDocument{
		{
			// This is the swiss bank. a private node customer who has
			// explicitly selected the default shared topic
			Doc: &tenancies.TenantResponse{
				Identity: "swissbank",
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						Topic:      defaultSharedTopic,
						TesseraKey: "swissbank-key",
					},
				},
			},
		},
		{
			// This is the cleaning company. a private node customer
			// who has explicitly selected the default shared topic
			Doc: &tenancies.TenantResponse{
				Identity: "cleaning company",
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						Topic:      defaultSharedTopic,
						TesseraKey: "swissbank-key",
					},
					"2": {
						Topic:      "german-bank-topic",
						TesseraKey: "german-bank-key",
					},
				},
			},
		},
	},

	This request would succes and return the defaultSharedTopic topic
*/

/*
	co-resident tenants correctly configured for private store",

	swiss bank + cleaning company are correctly configured co-resident tenants.
	both explicitly configured for the swiss banks private tsp store

	swiss bank is a private node customer and has paid for private storage for
	their tsp evidence.  cleaning company is a customer specifically on boarded
	to be resident on the same node as as the swiss bank (same tessera pub).
	here the cleaning company's tenant record correctly uses the swiss banks
	private store topic

	request: &tenancies.TopicFromPubKeyRequest{
		ChainID:       "1",
		TesseraPubKey: "swissbank-key",
	},

	tenants: []*tenanciesDocument{
		{
			// A swiss bank who has paid for a private node and has explicitly
			// opted for private storage for there tsp data.
			Doc: &tenancies.TenantResponse{
				Identity: "swissbank",
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						Topic:      "swissbank-topic",
						TesseraKey: "swissbank-key",
					},
				},
			},
		},
		{
			// Cleaning company. a private node customer who has explicitly
			// selected the private topic selected by their node owning parter.
			Doc: &tenancies.TenantResponse{
				Identity: "cleaning company",
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						Topic:      "swissbank-topic",
						TesseraKey: "swissbank-key",
					},
					"2": {
						Topic:      "german-bank-topic",
						TesseraKey: "german-bank-key",
					},
				},
			},
		},
	},

	this request would succeed and return the swissbank-topic
},
*/

/*
	co-resident tenants inconsistently configured (miss matched topic name)

	swiss bank has chosen private data store (or switched to it) for their tsp
	assets. The co-resident cleaning company tenancy record has not been
	updated to be consistent with the node owners (swiss banks) choice

	request: &tenancies.TopicFromPubKeyRequest{
		ChainID:       "1",
		TesseraPubKey: "swissbank-key",
	},

	tenants: []*tenanciesDocument{
		{
			// swiss bank is a private node customer who has explicitly
			// selected private store for tsp assets.
			Doc: &tenancies.TenantResponse{
				Identity: "swissbank",
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						Topic:      "swissbank-topic",
						TesseraKey: "swissbank-key",
					},
				},
			},
		},
		{
			// Here we have failed (broken scripts/onboarding or customer
			// management) to update the co-resident cleaning company to follow
			// the tsp choice of the node owning swiss bank they are
			// co-resident with.
			Doc: &tenancies.TenantResponse{
				Identity: "cleaning company",
				QuorumNetworks: map[string]*quorumnetwork.QuorumNetwork{
					"1": {
						Topic:      defaultSharedTopic,
						TesseraKey: "swissbank-key",
					},
					"2": {
						Topic:      "german-bank-topic",
						TesseraKey: "german-bank-key",
					},
				},
			},
		},
	},

	This is a catastrophic mis configuration (akin to putting the wrong tessera
	public key in the database). Arrangements in eventsender guarantee that
	assets are created on the owners node and any events on that asset are
	correctly shared back to the owner. However, with the data base in this
	state, data shared to other nodes by the policy would possibly be sent to
	the wrong topic (the defaultSharedTopic in this example).
*/
