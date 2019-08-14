package state

// KongState is an in-memory database representation
// of Kong's configuration.
type KongState struct {
	Services     *ServicesCollection
	Routes       *RoutesCollection
	Upstreams    *UpstreamsCollection
	Targets      *TargetsCollection
	Certificates *CertificatesCollection
	Plugins      *PluginsCollection
	Consumers    *ConsumersCollection

	KeyAuths   *KeyAuthsCollection
	HMACAuths  *HMACAuthsCollection
	JWTAuths   *JWTAuthsCollection
	BasicAuths *BasicAuthsCollection
}

// NewKongState creates a new in-memory KongState.
func NewKongState() (*KongState, error) {
	services, err := NewServicesCollection()
	if err != nil {
		return nil, err
	}
	routes, err := NewRoutesCollection()
	if err != nil {
		return nil, err
	}
	upstreams, err := NewUpstreamsCollection()
	if err != nil {
		return nil, err
	}
	targets, err := NewTargetsCollection()
	if err != nil {
		return nil, err
	}
	certificates, err := NewCertificatesCollection()
	if err != nil {
		return nil, err
	}
	plugins, err := NewPluginsCollection()
	if err != nil {
		return nil, err
	}
	consumers, err := NewConsumersCollection()
	if err != nil {
		return nil, err
	}

	keyAuths, err := NewKeyAuthsCollection()
	if err != nil {
		return nil, err
	}
	hmacAuths, err := NewHMACAuthsCollection()
	if err != nil {
		return nil, err
	}
	jwtAuths, err := NewJWTAuthsCollection()
	if err != nil {
		return nil, err
	}
	basicAuths, err := NewBasicAuthsCollection()
	if err != nil {
		return nil, err
	}

	return &KongState{
		Services:     services,
		Routes:       routes,
		Upstreams:    upstreams,
		Targets:      targets,
		Certificates: certificates,
		Plugins:      plugins,
		Consumers:    consumers,

		KeyAuths:   keyAuths,
		HMACAuths:  hmacAuths,
		JWTAuths:   jwtAuths,
		BasicAuths: basicAuths,
	}, nil
}
