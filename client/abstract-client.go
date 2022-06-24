package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"dkg-client-go/parser"

	"github.com/ivpusic/golog"
)

const (
	Pending   string = "PENDING"
	Completed        = "COMPLETED"
	Failed           = "FAILED"
)

var (
	defaultMaxNumberOfRetries int = 5
	defaultTimeoutInSeconds   int = 25
	defaultNumberOfResults    int = 5
)

type AbstractClientOptions struct {
	Endpoint           string
	Port               int
	UseSSL             bool
	LogLevel           golog.Level // optional
	MaxNumberOfRetries int         // optional
}

type abstractClient struct {
	LogLevel           golog.Level
	MaxNumberOfRetries int
	nodeBaseUrl        string
	Logger             golog.Logger
	Parser			   int
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

	ac.Parser = parser.NewSparqlSyntaxCheck()

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
		Timeout: time.Duration(defaultTimeoutInSeconds*1000) * time.Second,
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
		Timeout: time.Duration(defaultTimeoutInSeconds*1000) * time.Second,
	}

	_, err := client.Get(fmt.Sprintf("%s/info", ac.nodeBaseUrl))
	if err != nil {
		return err
	}
	return nil
}

type PublishRequestOptions struct {
	Method   string
	Data     string
	Filepath string
	Keywords []string
	UAL      string
}

func (ac *abstractClient) publishRequest(options PublishRequestOptions) (*http.Response, error) {
	ac.Logger.Debug("Sending node info request")

	form := url.Values{}

	if options.Filepath != "" {
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

type ResolveRequestOptions struct {
	IDS []string
}

func (ac *abstractClient) Resolve(options ResolveRequestOptions) (*http.Response, error) {
	if len(options.IDS) == 0 {
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

type SearchRequestOptions struct {
	Query      string
	Limit      int
	ResultType string
	Prefix     string
}

func (ac *abstractClient) Search(options SearchRequestOptions) (*http.Response, error) {
	if options.Query == "" || options.ResultType == "" {
		return nil, errors.New("Please provide search options in order to search")
	}
	resp, err := ac.searchRequest(options)
	if err != nil {
		return nil, err
	}
	// TODO getResult
}

func (ac *abstractClient) searchRequest(options SearchRequestOptions) (*http.Response, error) {
	ac.Logger.Debug("Sending search request")

	searchForm := url.Values{}

	searchForm.Add("query", options.Query)

	if options.ResultType == "entities" {
		searchForm.Add("limit", fmt.Sprintf("%d", options.Limit))
		searchForm.Add("prefix", options.Prefix)
	}

	client := &http.Client{}

	formUrl := fmt.Sprintf("%s/search?%s", ac.nodeBaseUrl, searchForm.Encode())

	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong query or url provided")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Could not send search form")
	}

	return resp, nil
}

func (ac *abstractClient) getSearchRequest(options SearchResultOptions) (*http.Response, error) {
	if options.HandlerId == 0 {
		return nil, errors.New("Unable to get results, need handler id")
	}

	client := &http.Client{}

	formUrl := fmt.Sprintf("%s/%s:search/result/%s", ac.nodeBaseUrl, options.ResultType, options.HandlerId)

	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong ids or url provided")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	timeoutFlag := make(chan bool, 1)
	go func() {
		time.Sleep(options.Timeout * time.Second)

		timeoutFlag <- true
	}()

	for {
		sleepTimeout := make(chan bool, 1)
		go func() {
			time.Sleep(5 * time.Second)
	
			sleepTimeout <- true
		}()
		<- sleepTimeout

		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.New("Could not send search result form")
		}

		//TODO if data.status == FAILED

		if <-timeoutFlag && failed && options.NumbersOfResults > currentNumberOfResults {
			break
	}

}

func (ac *abstractClient) Query(options QueryOptions) (*http.Response, error) {
	if options.Query == "" {
		return nil, errors.New("Please provide options in order to query")
	}

	resp, err := ac.queryRequest(options)
	if err != nil {
		return nil, err
	}

	// TODO GET RESULT
}

func (ac *abstractClient) queryRequest(options QueryOptions) (*http.Response, error) {
	ac.Logger.Debug("Sending query request")

	err := ac.Parser.Check(options.Query)
	if err != nil {
		return nil, err
	}

	queryForm := url.Values{}
	queryForm.Add("query", options.Query)

	queryUrl := url.Values{}
	queryUrl.Add("type", options.Type)

	formUrl := fmt.Sprintf("%s/query?%s", ac.nodeBaseUrl, queryUrl.Encode())

	client := &http.Client{}

	resp, err := client.PostForm(formUrl, queryForm)
	if err != nil {
		return nil, errors.New("Could not send publish request form")
	}

	return resp, nil

}

func (ac *abstractClient) Validate(options ValidateOptions) (*http.Response, error) {}

func (ac *abstractClient) getProofsRequest(option ValidateOptions) (*http.Response, error) {
	ac.Logger.Debug("Sending get proofs request")

	jsonNquads, err := json.Marshal(options.Nquads)
	if err != nil {
		return nil, errors.New("Could not convert the nquads to JSON")
	}

	validateForm := url.Values{}
	validateForm.Add("nquads", jsonNquads)

	formUrl := fmt.Sprintf("%s/proofs:get", ac.nodeBaseUrl)

	resp, err := client.PostForm(formUrl, validateForm)
	if err != nil {
		return nil, errors.New("Could not send proofs form")
	}
}

func (ac *abstractClient) performValidation(assertions)  {
	// TODO response.data
}