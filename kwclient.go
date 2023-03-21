package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"xdk/kwclient"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:8001"
const KWHubHost string = "https://hub.csgjourney.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	SessionId  string
	CsrfToken  string
}

type xpProject struct {
	Name        string
	Description string
	SessionId   string
	CsrfToken   string
	Host        string
}

type sqsConnection struct {
	Name       string
	Key        string
	Secret     string
	ReadQueue  string
	WriteQueue string
	ProjectId  string
	SessionId  string
	CsrfToken  string
	Host       string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/createkwProject", createKWProject)
	http.HandleFunc("/createSQSConnection", createSQSConnection)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func createKWProject(w http.ResponseWriter, r *http.Request) {
	var newXPProject xpProject
	json.NewDecoder(r.Body).Decode(&newXPProject)
	newXPProject.Host = KWHubHost
	projectId := kwclient.CreateProject(newXPProject.Name, newXPProject.Description, newXPProject.SessionId, newXPProject.CsrfToken, newXPProject.Host)
	w.Write([]byte(projectId))
}

func createSQSConnection(w http.ResponseWriter, r *http.Request) {
	var newSQSConnection sqsConnection
	json.NewDecoder(r.Body).Decode(&newSQSConnection)
	newSQSConnection.Host = KWHubHost
	kwclient.CreateConnection(newSQSConnection.Name, newSQSConnection.Key, newSQSConnection.Secret, newSQSConnection.ReadQueue, newSQSConnection.WriteQueue, newSQSConnection.ProjectId, newSQSConnection.SessionId, newSQSConnection.CsrfToken, newSQSConnection.Host)
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

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

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
