package client

import "github.com/ivpusic/golog"

type dkgClient struct {
	Client nativeAbstractClient
	Assets assetsClient
}

type DkgClientOptions struct {
	Endpoint           string
	Port               int
	UseSSL             bool
	LogLevel           golog.Level
	MaxNumberOfRetries int
}

func NewDkgClient(options DkgClientOptions) (dkgClient, error) {
	clientOpt := AbstractClientOptions{options.Endpoint, options.Port, options.UseSSL, options.LogLevel, options.MaxNumberOfRetries}
	c, err := newNativeAbstractClient(clientOpt)
	if err != nil {
		return dkgClient{}, err
	}

	a, err := newAssetsClient(&c)
	if err != nil {
		return dkgClient{}, err
	}

	return dkgClient{c, a}, nil
}

// golog.ERROR
