package client

import (
	"fmt"
	"time"
	"net/http"
	"errors"

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