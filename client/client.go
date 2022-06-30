package client

import "github.com/ivpusic/golog"

type dkgClient struct {
	client nativeAbstractClient
	assets assetsClient
}

type DkgClientOptions struct {
	Endpoint           string
	Port               int
	UseSSL             bool
	LogLevel           golog.Level
	MaxNumberOfRetries int
}

func newDkgClient(options DkgClientOptions) (dkgClient, error) {
	clientOpt := AbstractClientOptions{options.Endpoint, options.Port, options.UseSSL, options.LogLevel, options.MaxNumberOfRetries}
	c, err := newNativeAbstractClient(clientOpt)
	if err != nil {
		return dkgClient{}, err
	}

	a, err := newAssetsClient(&c, 2)
	if err != nil {
		return dkgClient{}, err
	}

	return dkgClient{c, a}, nil
}

// golog.ERROR
