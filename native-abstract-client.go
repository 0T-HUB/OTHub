// EXPLANATION
// Go doesn't support struct extending/inheritance, so to deal with this,
// this file includes NativeClient and AbstractClient in one Single Client

package client

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"dkg-client-go/parser"

	"github.com/ivpusic/golog"
	"github.com/nebulouslabs/merkletree"
)

// Constants signifying the state of a request
const (
	Pending   string = "PENDING"
	Completed        = "COMPLETED"
	Failed           = "FAILED"
)

// Some default values used in the API
var (
	defaultMaxNumberOfRetries int = 5
	defaultTimeoutInSeconds   int = 25
	defaultNumberOfResults    int = 5
)

// Options for building an AbstractClient
type AbstractClientOptions struct {
	Endpoint           string
	Port               int
	UseSSL             bool
	LogLevel           golog.Level
	MaxNumberOfRetries int
}

// The struct that mixes an AbstractClient and a NativeClient
type nativeAbstractClient struct {
	LogLevel           golog.Level
	MaxNumberOfRetries int
	nodeBaseUrl        string
	Logger             golog.Logger
	Parser             parser.SparqlSyntaxCheck
}

// Function to return a new NativeAbstractClient
func newNativeAbstractClient(options AbstractClientOptions) (nativeAbstractClient, error) {
	ac := new(nativeAbstractClient)

	// Set the logger
	ac.LogLevel = options.LogLevel

	// Change the MaxNumberOfRetries variable
	if options.MaxNumberOfRetries > 0 {
		ac.MaxNumberOfRetries = options.MaxNumberOfRetries
	} else {
		ac.MaxNumberOfRetries = defaultMaxNumberOfRetries // assign default number of tries
	}

	// Check if some arguments are missing
	if options.Endpoint == "" || options.Port == 0 {
		return nativeAbstractClient{}, errors.New("Endpoint and port are required parameters")
	}

	// Set the nodeBaseUrl SSL's
	var sslHeader string
	if options.UseSSL {
		sslHeader = "https://"
	} else {
		sslHeader = "http://"
	}

	// Format the URL
	ac.nodeBaseUrl = fmt.Sprintf("%s%s:%d", sslHeader, options.Endpoint, options.Port)

	// Create a new SPARQL syntax checker
	ac.Parser = parser.NewSparqlSyntaxCheck()

	// Check if the endpoint is reachable
	if ac.sendNodeInfoRequest() != nil {
		return nativeAbstractClient{}, errors.New("Endpoint not available")
	}

	return *ac, nil

}

//
// Get node information (version, is auto upgrade enabled, is telemetry enabled)
//

func (ac *nativeAbstractClient) NodeInfo() (*http.Response, error) {
	ac.Logger.Debug("Sending node info request")

	// Create a new client
	client := http.Client{
		Timeout: time.Duration(defaultTimeoutInSeconds*1000) * time.Second,
	}

	// Send the request
	resp, err := client.Get(fmt.Sprintf("%s/info", ac.nodeBaseUrl))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ac *nativeAbstractClient) sendNodeInfoRequest() error {
	ac.Logger.Debug("Sending node info request")

	// Create a new client
	client := http.Client{
		Timeout: time.Duration(defaultTimeoutInSeconds*1000) * time.Second,
	}

	// Send the request
	_, err := client.Get(fmt.Sprintf("%s/info", ac.nodeBaseUrl))
	if err != nil {
		return err
	}
	return nil
}

// Options for building a PublishRequest
type PublishRequestOptions struct {
	Method   string
	Data     string
	Filepath string
	Keywords []string
	UAL      string
}

// Function that "publishes" a file to the network
func (ac *nativeAbstractClient) Publish(options PublishRequestOptions) ([]byte, error) {
	// Check arguments
	if options.Filepath == "" && options.Data == "" {
		return nil, errors.New("Please provide atleast Filepath or Data")
	}

	// Make a query request
	resp, err := ac.publishRequest(options)
	if err != nil {
		return nil, err
	}

	queryResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal(resp, &queryResponse); err != nil {
		return nil, errors.New("Could not unmarshal query request response")
	}

	// Get the handler id
	opt := GetResultOptions{queryResponse["handler_id"].(int), "publish"}

	// Get the actual result
	respJson, err := ac.getResult(opt)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func (ac *nativeAbstractClient) publishRequest(options PublishRequestOptions) ([]byte, error) {
	ac.Logger.Debug("Sending node info request")

	form := url.Values{}

	// Add the data to the form
	if options.Filepath != "" {
		// Read the file and add it to the form
		buf, err := ioutil.ReadFile(options.Filepath)
		if err != nil {
			return nil, errors.New("Could not open file")
		}

		form.Set("file", string(buf))
	} else {
		// Directly add the data to the form
		form.Set("data", options.Data)
	}

	// Convert the keywords to a JSON string
	jsonKeywords, err := json.Marshal(options.Keywords)
	if err != nil {
		return nil, errors.New("Could not convert the keywords to JSON")
	}
	form.Set("keywords", string(jsonKeywords))

	// Check if we have an UAL to add
	if options.UAL != "" {
		form.Set("ual", options.UAL)
	}

	// Create a new client
	client := &http.Client{}

	// Form the request
	formUrl := fmt.Sprintf("%s/%s", ac.nodeBaseUrl, options.Method)

	// Send the request
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

// Options for building a ResolveRequest
type ResolveRequestOptions struct {
	IDS []string
}

// Function that emits a resolve request
func (ac *nativeAbstractClient) Resolve(options ResolveRequestOptions) ([]byte, error) {
	// Check arguments
	if len(options.IDS) == 0 {
		return nil, errors.New("Please provide resolve options in order to resolve")
	}

	// Create a request
	resp, err := ac.resolveRequest(options)
	if err != nil {
		return nil, err
	}

	// Read the response
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read in resolve body")
	}
	resp.Body.Close()

	resolveResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal([]byte(b), &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	// Get the handler id
	opt := GetResultOptions{resolveResponse["handler_id"].(int), "resolve"}

	// Get the result with that id
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

	// Create a new client
	client := &http.Client{}

	// Form the request
	formUrl := fmt.Sprintf("%s/resolve?%s", ac.nodeBaseUrl, idsForm.Encode())

	// Create the request
	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong ids or url provided")
	}

	// Set the headers for this http request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Could not send resolve request form")
	}

	return resp, nil

}

// Options for building a SearchtRequest
type SearchRequestOptions struct {
	Query            string
	ResultType       string
	Prefix           bool
	Limit            int
	Issuers          []string
	SchemaTypes      string
	NumbersOfResults int
	Timeout          int
}

// Function that emits a search request
func (ac *nativeAbstractClient) Search(options SearchRequestOptions) ([]byte, error) { //DONE
	if options.Query == "" || options.ResultType == "" {
		return nil, errors.New("Please provide search options in order to search")
	}

	// Get search request
	resp, err := ac.searchRequest(options)
	if err != nil {
		return nil, err
	}

	// Variable that stores the json as a map
	searchRequestResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal(resp, &searchRequestResponse); err != nil {
		return nil, errors.New("Could not unmarshal search result response")
	}

	// Generate options
	searchResultOptions := SearchResultOptions{
		searchRequestResponse["handler_id"].(int),
		options.ResultType,
		options.Timeout,
		options.NumbersOfResults,
	}

	// Get search result
	resp, err = ac.getSearchResult(searchResultOptions)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ac *nativeAbstractClient) searchRequest(options SearchRequestOptions) ([]byte, error) { //DONE
	ac.Logger.Debug("Sending search request")

	searchForm := url.Values{}

	// Add certain values to the form to be sent
	searchForm.Add("query", options.Query)

	if options.ResultType == "entities" {
		searchForm.Add("limit", fmt.Sprintf("%d", options.Limit))
		searchForm.Add("prefix", fmt.Sprintf("%t", options.Prefix))
	}

	// Create the client
	client := &http.Client{}

	// Format the URL
	formUrl := fmt.Sprintf("%s/search?%s", ac.nodeBaseUrl, searchForm.Encode())

	// Make the request
	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong query or url provided")
	}

	// Set the headers and make the request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Could not send search form")
	}

	// Convert resp.Body to []byte
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

// Options used when a search request is done
type SearchResultOptions struct {
	HandlerId        int
	ResultType       string
	Timeout          int
	NumbersOfResults int
}

func (ac *nativeAbstractClient) getSearchResult(options SearchResultOptions) ([]byte, error) { //DONE
	if options.HandlerId == 0 {
		return nil, errors.New("Unable to get results, need handler id")
	}

	// Create the http client
	client := &http.Client{}

	// Format the URL
	formUrl := fmt.Sprintf("%s/%s:search/result/%d", ac.nodeBaseUrl, options.ResultType, options.HandlerId)

	// Create the request
	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong ids or url provided")
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Channel for setting the timeouts
	timeoutFlag := make(chan bool)

	// Variable to show when the request has failed
	failed := false

	// Variable showing the current number of results received
	currentNumberOfResults := 0

	// Variable with the final JSON response
	b := make([]byte, 0)

	// Goroutine that will make the function sleep for the desired Timeout
	go func() {
		time.Sleep(time.Duration(options.Timeout) * time.Second)

		timeoutFlag <- true
	}()

	// Loop until the timeout is done, the request fails and the number of results is reached
	for {
		// Goroutine that will make the response go every 5 seconds
		sleepTimeout := make(chan bool, 1)
		go func() {
			time.Sleep(5 * time.Second)

			sleepTimeout <- true
		}()
		<-sleepTimeout

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.New("Could not send search result form")
		}

		// Read the response into a variable
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("Could not read search result body")
		}
		resp.Body.Close()

		searchResponse := make(map[string]interface{})

		// Transform response to json struct
		if err := json.Unmarshal([]byte(b), &searchResponse); err != nil {
			return nil, errors.New("Could not unmarshal search result response")
		}

		// Check if the request has failed
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

// Function that will send a query to the node
func (ac *nativeAbstractClient) Query(options QueryOptions) ([]byte, error) {
	// Check arguments
	if options.Query == "" {
		return nil, errors.New("Please provide options in order to query")
	}

	// Make a query request
	resp, err := ac.queryRequest(options)
	if err != nil {
		return nil, err
	}

	queryResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal(resp, &queryResponse); err != nil {
		return nil, errors.New("Could not unmarshal query request response")
	}

	// Get the handler id
	opt := GetResultOptions{queryResponse["handler_id"].(int), "query"}

	// Get the actual result
	respJson, err := ac.getResult(opt)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

// Options used when you want to send a query
type QueryOptions struct {
	Query string
	Type  string
}

func (ac *nativeAbstractClient) queryRequest(options QueryOptions) ([]byte, error) { //DONE
	ac.Logger.Debug("Sending query request")

	// Check that the SPARQL query's syntax is correct
	err := ac.Parser.Check(options.Query)
	if err != nil {
		return nil, err
	}

	// Add values to the form to be sended
	queryForm := url.Values{}
	queryForm.Add("query", options.Query)

	queryUrl := url.Values{}
	queryUrl.Add("type", options.Type)

	// Format the URL
	formUrl := fmt.Sprintf("%s/query?%s", ac.nodeBaseUrl, queryUrl.Encode())

	// Create the client
	client := &http.Client{}

	// Make the request
	resp, err := client.PostForm(formUrl, queryForm)
	if err != nil {
		return nil, errors.New("Could not send publish request form")
	}

	// Read the JSON response
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read search result body")
	}
	resp.Body.Close()

	return b, nil

}

// Options used when you want to make a validate request
type ValidateOptions struct {
	Nquads []string
}

// Function to send a valdiate request to the node
func (ac *nativeAbstractClient) Validate(options ValidateOptions) ([]validatedTriple, error) {

	// Check the arugments
	if len(options.Nquads) == 0 {
		return nil, errors.New("Please provide assertions and nquads in order to get proofs")
	}

	// First do a request
	resp, err := ac.getProofsRequest(options)
	if err != nil {
		return nil, errors.New("Couldnt not send proofs request")
	}

	resolveResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal(resp, &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	// Get the handler id
	opt := GetResultOptions{resolveResponse["handler_id"].(int), "proofs:get"}

	// Get the actual result
	respJson, err := ac.getResult(opt)
	if err != nil {
		return nil, err
	}

	// Transform the new response to json struct
	if err := json.Unmarshal(respJson, &resolveResponse); err != nil {
		return nil, errors.New("Could not unmarshal resolve request response")
	}

	// Check if the status equals to completed
	if resolveResponse["status"] != Completed {
		return nil, errors.New("Unable to get proofs for given nquads")
	}

	// Perform the validation of the nquads
	respFinal, err := ac.performValidation(resp)
	if err != nil {
		return nil, err
	}

	return respFinal, nil

}

func (ac *nativeAbstractClient) getProofsRequest(options ValidateOptions) ([]byte, error) {
	ac.Logger.Debug("Sending get proofs request")

	// Convert an slice into a JSON array
	jsonNquads, err := json.Marshal(options.Nquads)
	if err != nil {
		return nil, errors.New("Could not convert the nquads to JSON")
	}

	// Add it to the form
	validateForm := url.Values{}
	validateForm.Add("nquads", string(jsonNquads))

	// Format the URL
	formUrl := fmt.Sprintf("%s/proofs:get", ac.nodeBaseUrl)

	// Create the URL
	client := &http.Client{}

	// Do the request
	resp, err := client.PostForm(formUrl, validateForm)
	if err != nil {
		return nil, errors.New("Could not send proofs form")
	}

	// Read the response into a variable
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read search result body")
	}
	resp.Body.Close()

	return b, nil
}

// Struct that represents an Assertion in JSON
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

// Struct that represent a validatedTriple
type validatedTriple struct {
	Triple string
	Valid  bool
}

func (ac *nativeAbstractClient) performValidation(assertions []byte) ([]validatedTriple, error) {
	// Create a list of validated triples
	validationResult := make([]validatedTriple, 0)

	// Variable holding the JSON as struct form
	var out AssertionType

	// Convert the JSON to the struct
	err := json.Unmarshal([]byte(assertions), &out)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through all the assertions
	for _, assertion := range out {
		// Grab the rooHash
		rootHash, err := ac.fetchRootHash(assertion.AssertionID)
		if err != nil {
			return nil, err
		}

		// Loop through all the Proofs
		for _, obj := range assertion.Proofs {
			// Create a validated triple
			v := validatedTriple{obj.Triple, false}

			// Check if assertion has proof, if no continue to next
			if len(obj.Proof) == 0 {
				ac.Logger.Debug(fmt.Sprintf("%s has no proof in assertion %s", obj.Triple, assertion.AssertionID))
				continue
			}

			// Generate an slice of the proofs
			proof := make([][]byte, 0)
			for i := range obj.Proof {
				if obj.Proof[i].Left != "" {
					proof = append(proof, []byte(obj.Proof[i].Left))
				} else {
					proof = append(proof, []byte(obj.Proof[i].Right))
				}
			}

			// Validate the proof cryptographically
			verified := ac.validateProof([]byte(rootHash), proof)
			v.Valid = verified
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

func (ac *nativeAbstractClient) validateProof(rootHash []byte, proof [][]byte) bool {
	// Go through every proof
	// If any can't be verified, it returns false, if all can be verified return true
	for x := range proof {
		if !merkletree.VerifyProof(sha256.New(), rootHash, proof, uint64(x), uint64(len(proof))) {
			return false
		}
	}
	return true
}

// WARNING
// This function may have been written incorrectly since the first original js code
func (ac *nativeAbstractClient) fetchRootHash(assertionId string) (string, error) {
	result, err := ac.Resolve(ResolveRequestOptions{[]string{assertionId}})
	if err != nil {
		return "", err
	}

	resolveResponse := make([]map[string]map[string]string, 0)

	// Transform response to json struct
	if err := json.Unmarshal(result, &resolveResponse); err != nil {
		return "", errors.New("Could not unmarshal resolve request response")
	}

	// Trying to be similiar to this js line:
	// return result.data[0][assertionId].rootHash;
	return resolveResponse[0][assertionId]["rootHash"], err

}

type GetResultOptions struct {
	HandlerId int
	Operation string
}

func (ac *nativeAbstractClient) getResult(options GetResultOptions) ([]byte, error) { //DONE

	// Channel that will receive when 500 ms have passed
	timeoutFlag := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(500) * time.Millisecond)

		timeoutFlag <- true
	}()

	<-timeoutFlag

	// Check options
	if options.HandlerId == 0 || options.Operation == "" {
		return nil, errors.New("Unable to get results, need handler id and operation")
	}

	// Variable that stores the json as a map
	resultResponse := make(map[string]interface{})

	// Assign a pending until we get an answer
	resultResponse["status"] = Pending

	// Number of "attempts" to get the result
	retries := 0

	// Url to get the result
	formUrl := fmt.Sprintf("%s/%s/result/%d", ac.nodeBaseUrl, options.Operation, options.HandlerId)

	// Create the http request
	req, err := http.NewRequest(http.MethodGet, formUrl, nil)
	if err != nil {
		return nil, errors.New("Wrong operation, url or id provided")
	}

	// Set the header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Variable storing the JSON response
	b := make([]byte, 0)

	timerFlag := make(chan bool, 1)

	for {
		// If we do more retries that the max number of retries, return an error
		if retries > ac.MaxNumberOfRetries {
			return nil, errors.New("Unable to get results. Max number of retries reached")
		}
		retries++

		// Run this function every 5 seconds
		go func() {
			time.Sleep(time.Duration(5) * time.Second)

			timerFlag <- true
		}()
		<-timerFlag

		// Create http client
		client := &http.Client{}

		// Actually "execute" the request
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.New("Could not send GetResult request")
		}

		// Read the response as a byte array
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("Could not read GetResult body")
		}
		resp.Body.Close()

		// Transform response to json struct
		if err := json.Unmarshal(b, &resultResponse); err != nil {
			return nil, errors.New("Could not unmarshal GetResult response")
		}

		ac.Logger.Debug(fmt.Sprintf("%s result status: %s", options.Operation, resultResponse["status"]))

		// If the result is not pending, break the loop
		if resultResponse["status"] != Pending {
			break
		}
	}

	// If failed, raise an error
	if resultResponse["status"] == Failed {
		return nil, errors.New(fmt.Sprintf("Get %s failed. Reason: %s", options.Operation, resultResponse["message"]))
	}

	// If no error found, return the JSON response
	return b, nil

}
