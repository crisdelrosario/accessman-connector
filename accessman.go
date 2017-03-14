package accessman

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// Accessman custom Accessman API lib
type Accessman struct {
	API  string
	URL  string
	Auth Auth
}

const (
	// SourceValidationURL Accessman source validation URL
	SourceValidationURL = "/source/%s/validation/%s/%s/%s"
)

// ErrRequestForAccessDenied Error message if Accessman denied the access to API
var ErrRequestForAccessDenied = errors.New("Access to API is denied by accessman")

// NewAccessRequest creates a new accessman instance
func NewAccessRequest(API string, host string) Accessman {
	return Accessman{
		API: API,
		URL: host,
	}
}

// VerifyAccess Checks if client is allowed to access this API
func (am *Accessman) VerifyAccess(clientIP string, path string, method string, username string, password string) error {
	return am.Validate(clientIP, path, method, username, password)
}

// Validate source validatation
func (am *Accessman) Validate(clientIP string, path string, method string, username string, password string) error {
	sourceValidationURL := am.CreateSourceValidationURL(clientIP, am.API, path, method)

	connection := Connection{}
	response, err := connection.Get(am.URL+sourceValidationURL, am.Auth)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return ErrRequestForAccessDenied
	}

	return nil
}

// CreateSourceValidationURL Create source validation url
func (am *Accessman) CreateSourceValidationURL(clientIP string, api string, path string, method string) string {
	return fmt.Sprintf(SourceValidationURL, clientIP, api, Hash{path}.MD5(), strings.ToLower(method))
}

// SetAuth sets the username and password that we will be using for basic authentication
func (am *Accessman) SetAuth(user string, pass string) {
	am.Auth.User = user
	am.Auth.Pass = pass
}

// AddNew add and account and generate OPK
func (am *Accessman) AddNew(email string) {
}

// GetClientIP get the remote ip of the client
func (am *Accessman) GetClientIP(r *http.Request) string {
	ipAddress := r.RemoteAddr

	if ip := r.Header.Get("X-Forwarded-For"); "" != ip {
		ipAddress = ip

		// X-Forwarded-For might contain multiple IPs. Get the first one.
		if strings.Contains(ipAddress, ",") {
			ips := strings.Split(ipAddress, ",")
			ipAddress = strings.Trim(ips[0], " ")
		}
	}

	var ip net.IP
	var err error

	if -1 != strings.Index(ipAddress, ":") {
		if ipAddress, _, err = net.SplitHostPort(ipAddress); err != nil {
		}
	}

	if err := ip.UnmarshalText([]byte(ipAddress)); err != nil {
	}

	return ipAddress
}
