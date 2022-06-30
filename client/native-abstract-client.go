// EXPLANATION
// Go doesn't support struct extending/inheritance, so to deal with this,
// this file includes NativeClient and AbstractClient in one Single Client

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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
	LogLevel           golog.Level
	MaxNumberOfRetries int
}

type nativeAbstractClient struct {
	LogLevel           golog.Level
	MaxNumberOfRetries int
	nodeBaseUrl        string
	Logger             golog.Logger
	Parser             parser.SparqlSyntaxCheck
}

func newNativeAbstractClient(options AbstractClientOptions) (nativeAbstractClient, error) {
	ac := new(nativeAbstractClient)

	ac.LogLevel = options.LogLevel

	if options.MaxNumberOfRetries > 0 {
		ac.MaxNumberOfRetries = options.MaxNumberOfRetries
	} else {
		ac.MaxNumberOfRetries = defaultMaxNumberOfRetries // assign default number of tries
	}

	if options.Endpoint == "" || options.Port == 0 {
		return nativeAbstractClient{}, errors.New("Endpoint and port are required parameters")
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
		return nativeAbstractClient{}, errors.New("Endpoint not available")
	}

	return *ac, nil

}

//
// Get node information (version, is auto upgrade enabled, is telemetry enabled)
//

func (ac *nativeAbstractClient) nodeInfo() (*http.Response, error) {
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

func (ac *nativeAbstractClient) sendNodeInfoRequest() error {
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

func (ac *nativeAbstractClient) publishRequest(options PublishRequestOptions) ([]byte, error) {
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

	// convert resp.Body to []byte
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read response body")
	}
	resp.Body.Close()

	return b, nil
}

type ResolveRequestOptions struct {
	IDS []string
}

func (ac *nativeAbstractClient) Resolve(options ResolveRequestOptions) ([]byte, error) { //DONE
	if len(options.IDS) == 0 {
		return nil, errors.New("Please provide resolve options in order to resolve")
	}

	resp, err := ac.resolveRequest(options)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read in resolve body")
	}
	resp.Body.Close()

	resolveResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal([]byte(b), &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	opt := GetResultOptions{resolveResponse["handler_id"].(int), "resolve"}

	respJson, err := ac.getResult(opt)
	if err != nil {
		return nil, err
	}

	return respJson, nil

}

func (ac *nativeAbstractClient) resolveRequest(options ResolveRequestOptions) (*http.Response, error) {
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
	Query            string
	ResultType       string
	Prefix           string
	Limit            int
	Issuers          []string
	SchemaTypes      string
	NumbersOfResults int
	Timeout          int
}

func (ac *nativeAbstractClient) Search(options SearchRequestOptions) ([]byte, error) { //DONE
	if options.Query == "" || options.ResultType == "" {
		return nil, errors.New("Please provide search options in order to search")
	}

	// get search request
	resp, err := ac.searchRequest(options)
	if err != nil {
		return nil, err
	}

	// variable that stores the json as a map
	searchRequestResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal(resp, &searchRequestResponse); err != nil {
		return nil, errors.New("Could not unmarshal search result response")
	}

	// generate options
	searchResultOptions := SearchResultOptions{
		searchRequestResponse["handler_id"].(int),
		options.ResultType,
		options.Timeout,
		options.NumbersOfResults,
	}

	// get search result
	resp, err = ac.getSearchResult(searchResultOptions)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ac *nativeAbstractClient) searchRequest(options SearchRequestOptions) ([]byte, error) { //DONE
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

	// convert resp.Body to []byte
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read response body")
	}
	resp.Body.Close()

	return b, nil
}

// type searchResultResponse struct {
// 	Results []struct {
// 		ID string `json:"id"`
// 	} `json:"results"`
// }

type SearchResultOptions struct { //DONE
	HandlerId        int
	ResultType       string
	Timeout          int
	NumbersOfResults int
}

func (ac *nativeAbstractClient) getSearchResult(options SearchResultOptions) ([]byte, error) { //DONE
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
	failed := false
	currentNumberOfResults := 0
	b := make([]byte, 0)

	go func() {
		time.Sleep(time.Duration(options.Timeout) * time.Second)

		timeoutFlag <- true
	}()

	for {
		sleepTimeout := make(chan bool, 1)
		go func() {
			time.Sleep(5 * time.Second)

			sleepTimeout <- true
		}()
		<-sleepTimeout

		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.New("Could not send search result form")
		}

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("Could not read search result body")
		}
		resp.Body.Close()

		searchResponse := make(map[string]interface{})

		// transform response to json struct
		if err := json.Unmarshal([]byte(b), &searchResponse); err != nil {
			return nil, errors.New("Could not unmarshal search result response")
		}

		if searchResponse["status"] == "FAILED" {
			failed = true
		} else {
			currentNumberOfResults = len(searchResponse["itemListElement"].([]interface{})[0].(map[string]interface{}))
		}

		if <-timeoutFlag && failed && options.NumbersOfResults > currentNumberOfResults {
			break
		}
	}

	return b, nil

}

func (ac *nativeAbstractClient) Query(options QueryOptions) ([]byte, error) {
	if options.Query == "" {
		return nil, errors.New("Please provide options in order to query")
	}

	resp, err := ac.queryRequest(options)
	if err != nil {
		return nil, err
	}

	queryResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal(resp, &queryResponse); err != nil {
		return nil, errors.New("Could not unmarshal query request response")
	}

	opt := GetResultOptions{queryResponse["handler_id"].(int), "query"}

	respJson, err := ac.getResult(opt)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

type QueryOptions struct {
	Query string
	Type  string
}

func (ac *nativeAbstractClient) queryRequest(options QueryOptions) ([]byte, error) { //DONE
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read search result body")
	}
	resp.Body.Close()

	return b, nil

}

type ValidateOptions struct {
	Nquads []string
}

func (ac *nativeAbstractClient) Validate(options ValidateOptions) ([]byte, error) {
	if len(options.Nquads) == 0 {
		return nil, errors.New("Please provide assertions and nquads in order to get proofs")
	}

	resp, err := ac.getProofsRequest(options)
	if err != nil {
		return nil, errors.New("Couldnt not send proofs request")
	}

	resolveResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal(resp, &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	opt := GetResultOptions{resolveResponse["handler_id"].(int), "proofs:get"}

	respJson, err := ac.getResult(opt)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respJson, &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	if resolveResponse["status"] != Completed {
		return nil, errors.New("Unable to get proofs for given nquads")
	}

	// TODO Perform validation

	return resp, nil

}

func (ac *nativeAbstractClient) getProofsRequest(options ValidateOptions) ([]byte, error) {
	ac.Logger.Debug("Sending get proofs request")

	jsonNquads, err := json.Marshal(options.Nquads)
	if err != nil {
		return nil, errors.New("Could not convert the nquads to JSON")
	}

	validateForm := url.Values{}
	validateForm.Add("nquads", string(jsonNquads))

	formUrl := fmt.Sprintf("%s/proofs:get", ac.nodeBaseUrl)

	client := &http.Client{}

	resp, err := client.PostForm(formUrl, validateForm)
	if err != nil {
		return nil, errors.New("Could not send proofs form")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read search result body")
	}
	resp.Body.Close()

	return b, nil
}

type AssertionType []struct {
	AssertionID string `json:"assertionId"`
	Proofs      []struct {
		Triple     string `json:"triple"`
		TripleHash string `json:"tripleHash"`
		Proof      []struct {
			Right string `json:"right,omitempty"`
			Left  string `json:"left,omitempty"`
		} `json:"proof"`
	} `json:"proofs"`
}

type validatedTriple struct {
	Triple string
	Valid  bool
}

func (ac *nativeAbstractClient) performValidation(assertions []byte) ([]validatedTriple, error) {
	validationResult := make([]validatedTriple, 0)

	var out AssertionType
	err := json.Unmarshal([]byte(assertions), &out)
	if err != nil {
		log.Fatal(err)
	}

	for _, assertion := range out {
		rootHash, err := ac.fetchRootHash(assertion.AssertionID)
		if err != nil {
			return nil, err
		}
		for _, obj := range assertion.Proofs {
			v := validatedTriple{obj.Triple, false}

			if len(obj.Proof) == 0 {
				ac.Logger.Debug(fmt.Sprintf("%s has no proof in assertion %s", obj.Triple, assertion.AssertionID))
				continue
			}

			verified := ac.validateProof(obj, rootHash)

			v.Valid = true
			validationResult = append(validationResult, v)

			if verified {
				ac.Logger.Debug(fmt.Sprintf("Validation successful for data: %v", obj))
			} else {
				ac.Logger.Debug(fmt.Sprintf("Invalid data: %v", obj))
			}
		}
	}

	return validationResult, nil
}

// TODO because of wrong api?
func (ac *nativeAbstractClient) fetchRootHash(assertionId string) (string, error) {
	result, err := ac.Resolve(ResolveRequestOptions{[]string{assertionId}})
	if err != nil {
		return "", err
	}

	resolveResponse := make(map[string]interface{})

	// transform response to json struct
	if err := json.Unmarshal([]byte(b), &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	return result, err

}

type GetResultOptions struct {
	HandlerId int
	Operation string
}

func (ac *nativeAbstractClient) getResult(options GetResultOptions) ([]byte, error) { //DONE

	// channel that will receive when 500 ms have passed
	timeoutFlag := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(500) * time.Millisecond)

		timeoutFlag <- true
	}()

	<-timeoutFlag

	// check options
	if options.HandlerId == 0 || options.Operation == "" {
		return nil, errors.New("Unable to get results, need handler id and operation")
	}

	// variable that stores the json as a map
	resultResponse := make(map[string]interface{})

	// assign a pending until we get an answer
	resultResponse["status"] = Pending

	// number of "attempts" to get the result
	retries := 0

	// url to get the result
	formUrl := fmt.Sprintf("%s/%s/result/%d", ac.nodeBaseUrl, options.Operation, options.HandlerId)

	// create the http request
	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong operation, url or id provided")
	}

	// set the header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// variable storing the JSON response
	b := make([]byte, 0)

	timerFlag := make(chan bool, 1)

	for {
		// if we do more retries that the max number of retries, return an error
		if retries > ac.MaxNumberOfRetries {
			return nil, errors.New("Unable to get results. Max number of retries reached")
		}
		retries++

		// run this function every 5 seconds
		go func() {
			time.Sleep(time.Duration(5) * time.Second)

			timerFlag <- true
		}()
		<-timerFlag

		// create http client
		client := &http.Client{}

		// actually "execute" the request
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.New("Could not send GetResult request")
		}

		// read the response as a byte array
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("Could not read GetResult body")
		}
		resp.Body.Close()

		// transform response to json struct
		if err := json.Unmarshal(b, &resultResponse); err != nil {
			return nil, errors.New("Could not unmarshal GetResult response")
		}

		ac.Logger.Debug(fmt.Sprintf("%s result status: %s", options.Operation, resultResponse["status"]))

		// if the result is not pending, break the loop
		if resultResponse["status"] != Pending {
			break
		}
	}

	// if failed, raise an error
	if resultResponse["status"] == Failed {
		return nil, errors.New(fmt.Sprintf("Get %s failed. Reason: %s", options.Operation, resultResponse["message"]))
	}

	// if no error found, return the JSON response
	return b, nil

}
