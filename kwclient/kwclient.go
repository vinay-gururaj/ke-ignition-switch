package kwclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:10000"
const KWHubHost string = "https://hub.csgjourney.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	SessionId  string
	CsrfToken  string
}

type KwProject struct {
	Name        string
	Description string
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func (c *Client) createKWProject(kwProject *KwProject) {
	// kwclient.CreateProject(kwProject.ProjectName, kwProject.ProjectDescription, kwProject.SessionId, kwProject.CsrfToken, KWHubHost)
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
