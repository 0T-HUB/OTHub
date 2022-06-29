package client

type dkgClient struct {
	client nativeAbstractClient
	assets *assetsClient
}

func newDkgClient() (dkgClient, error) {
	c, err := newNativeAbstractClient()
	if err != nil {
		return dkgClient{}, err
	}

	a, err := newAssetsClient(opt, &c)
	if err != nil {
		return dkgClient{}, err
	}

	return dkgClient{c, a}, nil
}
