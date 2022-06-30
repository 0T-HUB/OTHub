package client

type assetsProxyPath struct {
	nodeBaseUrl string
}

func NewAssetsProxyPath(url string) assetsProxyPath {
	return assetsProxyPath{url}
}

type CreatePathOptions struct {
	Settings map[string]string
	Data     map[string]string
	ID       string
	Loaded   bool
}

func (aspp *assetsProxyPath) CreatePath(options CreatePathOptions) {
	if options.Data == nil {
		options.Data = options.Settings
		options.Settings = make(map[string]string)
	}
}
