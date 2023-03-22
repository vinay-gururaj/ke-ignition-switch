package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"xdk/kwclient"

	"github.com/gorilla/mux"
)

// HostURL - Default Hashicups URL
const (
	HostURL     = "http://localhost:8001"
	KWHubHostV1 = "https://hub.csgjourney.com"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	SessionId  string
	CsrfToken  string
}

type xpProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type sqsConnection struct {
	Name       string `json:"name"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	ReadQueue  string `json:"readQueue"`
	WriteQueue string `json:"writeQueue"`
	ProjectId  string `json:"projectId"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/createkwProject", createKWProject).Methods("POST")
	router.HandleFunc("/createSQSConnection", createSQSConnection).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func createKWProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newXPProject xpProject
	json.NewDecoder(r.Body).Decode(&newXPProject)
	projectId := kwclient.CreateProject(newXPProject.Name, newXPProject.Description, r.Header.Get("SessionId"), r.Header.Get("CsrfToken"), KWHubHostV1)
	json.NewEncoder(w).Encode(projectId)
}

func createSQSConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newSQSConnection sqsConnection
	json.NewDecoder(r.Body).Decode(&newSQSConnection)
	kwclient.CreateConnection(newSQSConnection.Name, newSQSConnection.Key, newSQSConnection.Secret, newSQSConnection.ReadQueue, newSQSConnection.WriteQueue, newSQSConnection.ProjectId, r.Header.Get("SessionId"), r.Header.Get("CsrfToken"), KWHubHostV1)
}

func main() {
	handleRequests()
}

// NewClient -
func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	c.SessionId = "s:NHdrh-HDWmutvmy0osWOZNHTAQjIoKo3.pGG1jd59eJkA6F5mprb3hJUtou+oEHbEcbbSdhTwqT8"
	c.CsrfToken = "dZrhGvzj-IIkUmHDzFvBa83zczHM_ao5BhEg"

	return &c, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
