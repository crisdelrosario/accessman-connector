package accessman

import (
	"net/http"
	"strings"
)

// Connection connection structure
type Connection struct {
	Headers map[string]string
}

const (
	// ContentTypeJSON default content type
	ContentTypeJSON = "application/json"
)

// Send send the http request
func (connection *Connection) Send(method, url, payload string, auth Auth) (*http.Response, error) {
	request, _ := http.NewRequest(method, url, strings.NewReader(payload))

	if connection.Headers != nil {
		for key, value := range connection.Headers {
			request.Header.Set(key, value)
		}
	}

	request.Header.Set("Content-type", ContentTypeJSON)

	if auth.User != "" && auth.Pass != "" {
		request.SetBasicAuth(auth.User, auth.Pass)
	}

	return http.DefaultClient.Do(request)
}

// Get http get request
func (connection *Connection) Get(url string, auth Auth) (*http.Response, error) {
	return connection.Send(http.MethodGet, url, "", auth)
}
