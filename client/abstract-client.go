package client

import (
	"fmt"
	"time"
	"net/url"
	"net/http"
	"errors"
	"encoding/json"

	"github.com/ivpusic/golog"
)

const (
	Pending string = "PENDING"
	Completed      = "COMPLETED"
	Failed         = "FAILED"
)

var (
	defaultMaxNumberOfRetries int = 5
    defaultTimeoutInSeconds int = 25
    defaultNumberOfResults int = 5
)

type AbstractClientOptions struct {
	Endpoint string
	Port int
	UseSSL bool
	LogLevel golog.Level // optional
	MaxNumberOfRetries int // optional
}

type abstractClient struct {
    LogLevel golog.Level
    MaxNumberOfRetries int
    nodeBaseUrl string
    Logger golog.Logger
    // parser
}

func NewAbstractClient(options AbstractClientOptions) (abstractClient, error) {
	ac := new(abstractClient)

	if options.LogLevel == nil {
		ac.LogLevel = golog.ERROR
	} else {
		ac.LogLevel = options.LogLevel
	}

	if options.MaxNumberOfRetries > 0 {
		ac.MaxNumberOfRetries = options.MaxNumberOfRetries
	} else {
		ac.MaxNumberOfRetries = defaultMaxNumberOfRetries // assign default number of tries
	}

	if options.Endpoint == "" || options.Port == 0 {
		return abstractClient{}, errors.New("Endpoint and port are required parameters")
	}

	var sslHeader string
	if options.UseSSL {
		sslHeader = "https://"
	} else {
		sslHeader = "http://"
	}

	ac.nodeBaseUrl = fmt.Sprintf("%s%s:%d", sslHeader, options.Endpoint, options.Port)

	if ac.sendNodeInfoRequest() != nil {
		return abstractClient{}, errors.New("Endpoint not available")
	}

	return *ac, nil

}

//
// Get node information (version, is auto upgrade enabled, is telemetry enabled)
//

func (ac *abstractClient) nodeInfo() (*http.Response, error) {
	ac.Logger.Debug("Sending node info request")

	client := http.Client{
    	Timeout: time.Duration(defaultTimeoutInSeconds * 1000) * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("%s/info", ac.nodeBaseUrl))
	if err != nil {
		return nil, err
	}
	return resp, nil	
}

func (ac *abstractClient) sendNodeInfoRequest() error {
	ac.Logger.Debug("Sending node info request")

	client := http.Client{
    	Timeout: time.Duration(defaultTimeoutInSeconds * 1000) * time.Second,
	}

	_, err := client.Get(fmt.Sprintf("%s/info", ac.nodeBaseUrl))
	if err != nil {
		return err
	}
	return nil
}

func (ac *abstractClient) publishRequest(options PublishRequestOptions) (*http.Response, error) {
	ac.Logger.Debug("Sending node info request")

    form := url.Values{}

    if (options.Filepath != "") {
		form.Set("file", "foo") // TODO
    } else {
    	form.Set("data", options.Data)
    }

    jsonKeywords, err := json.Marshal(options.Keywords)
    if err != nil {
    	return nil, errors.New("Could not convert the keywords to JSON")
    }
    form.Set("keywords", string(jsonKeywords))

    if options.UAL != "" {
    	form.Set("ual", options.UAL)
    }

    client := &http.Client{}

    formUrl := fmt.Sprintf("%s/%s", ac.nodeBaseUrl, options.Method)

    resp, err := client.PostForm(formUrl, form)
    if err != nil {
        return nil, errors.New("Could not send publish request form")
    }

    return resp, nil
}

func (ac *abstractClient) Resolve(options ResolveRequestOptions) (*http.Response, error) {
	if len(options.ids) == 0 {
		return nil, errors.New("Please provide resolve options in order to resolve")
	}

	resp, err := ac.resolveRequest(options)
	if err != nil {
		return nil, err
	}

	// TODO getResult

}

func (ac *abstractClient) resolveRequest(options ResolveRequestOptions) (*http.Response, error) {
	ac.Logger.Debug("Sending resolve request")

	idsForm := url.Values{}

	for i := range options.IDS {
        idsForm.Add("ids", options.IDS[i])
    }

    client := &http.Client{}

    formUrl := fmt.Sprintf("%s/resolve?%s", ac.nodeBaseUrl, idsForm.Encode())

    req, err := http.NewRequest(http.MethodGet, formUrl, nil)
    if err != nil {
    	return nil, errors.New("Wrong ids or url provided")
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Could not send resolve request form")
	}

    return resp, nil

}


func (ac *abstractClient) Search(options SearchRequestOptions) {

}

func (ac *abstractClient) searchRequest(options SearchRequestOptions) {
	ac.Logger.Debug("Sending search request")

	searchForm := url.Values{}

	var prefix bool
	if 

}


