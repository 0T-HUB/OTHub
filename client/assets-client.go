package client

import (
	"encoding/json"
	"errors"
)

type assetsClient struct {
	Client          *nativeAbstractClient
	AssetsProxyPath int //FOR NOW...
}

func newAssetsClient(nac *nativeAbstractClient, AssetsProxyPath int) (assetsClient, error) {
	return assetsClient{nac, 0}, nil
}

type CreateOptions struct {
	Filepath string
	Data     string
	Keywords []string
}

func (ac *assetsClient) Create(options CreateOptions) ([]byte, error) {
	if options.Filepath == "" || options.Data == "" {
		return nil, errors.New("Please provide publish options in order to publish")
	}

	opt := PublishRequestOptions{"provision", options.Data, options.Filepath, options.Keywords, ""}

	resp, err := ac.Client.publishRequest(opt)
	if err != nil {
		return nil, err
	}

	publishResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal(resp, &publishResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	respJson, err := ac.Client.getResult(GetResultOptions{publishResponse["handler_id"].(int), opt.Method})
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

type UpdateOptions struct {
	Filepath string
	Data     string
	Keywords []string
}

func (ac *assetsClient) Update(ual string, options CreateOptions) ([]byte, error) {
	if options.Filepath == "" || options.Data == "" {
		return nil, errors.New("Please provide publish options in order to publish")
	}

	opt := PublishRequestOptions{"update", options.Data, options.Filepath, options.Keywords, ual}

	resp, err := ac.Client.publishRequest(opt)
	if err != nil {
		return nil, err
	}

	publishResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal(resp, &publishResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	respJson, err := ac.Client.getResult(GetResultOptions{publishResponse["handler_id"].(int), opt.Method})
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
