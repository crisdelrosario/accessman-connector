package accessman

import (
	"testing"
)

const (
	Port = "8843"
	API  = "am"
)

func TestMain(m *testing.M) {

	mock := LoadAccessmanMockServer()
	testAccessVerification(mock)
}

func LoadAccessmanMockServer() Mock {
	// Create a mock server
	mock := NewMockServer(Port)
	mock.Run()

	return mock
}

func testAccessVerification(mock Mock) {
	am := NewAccessRequest(API, mock.GetServerURL())

	clientIP := "0.0.0.0"
	path := "/client"
	method := "Get"
	username := "OPK00000000000001"
	password := "password"

	am.VerifyAccess(clientIP, path, method, username, password)
}
