package accessman

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

// Mock Fake accessman server for testing
type Mock struct {
	Port   string
	Router *mux.Router
}

// NewMockServer creates new mock server
func NewMockServer(port string) Mock {
	mock := Mock{
		Port: port,
	}

	mock.Create()

	return mock
}

// Run run fake accessman server
func (mock *Mock) Run() {
	http.ListenAndServe(":"+mock.Port, mock.Router)
}

// SourceValidation ...
func (mock *Mock) SourceValidation(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	w.WriteHeader(http.StatusNoContent)
}

// Create new fake accessman server
func (mock *Mock) Create() {
	router := mux.NewRouter()

	rURL := fmt.Sprintf(
		SourceValidationURL,
		"{a}",
		"{b}",
		"{c}",
		"{d}",
	)

	router.HandleFunc(rURL, mock.SourceValidation)

	mock.Router = router
}

// GetAddress get the IP address of this mock server
func (mock *Mock) GetAddress() string {
	var ip net.IP

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				break
			case *net.IPAddr:
				ip = v.IP
				break
			}
		}
	}
	return ip.String()
}

// GetServerURL mock server's URL
func (mock *Mock) GetServerURL() string {
	return mock.GetAddress() + ":" + mock.Port
}
