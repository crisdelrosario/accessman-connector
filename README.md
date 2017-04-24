# accessman
--
    import "."

### Access Verification

    import (
      "fmt"
      "github.com/crisdelrosario/accessman-connector"
    )

    const (
      API = "am"
      AccessmanURL = "http://0.0.0.0:8843"
    )

    am := accessman.NewAccessRequest(API, AccessmanURL)

    clientIP := "0.0.0.0"
    path := "/test"
    method := "Get"
    username := "username"
    password := "password"

    if err := am.VerifyAccess(clientIP, path, method, username, password); err !=  nil {
      fmt.Error(err.Error())
    }

### Create A Mock Server For Testing

    import (
    		"fmt"
    		"github.com/crisdelrosario/accessman-connector"
    )

    Port := "8443"
    mock := accessman.NewMockServer(Port)
    mock.Run()

    fmt.Println("Mock Server URL: " + mock.GetServerURL())

## Usage

```go
const (
	// ContentTypeJSON default content type
	ContentTypeJSON = "application/json"
)
```

```go
const (
	// SourceValidationURL Accessman source validation URL
	SourceValidationURL = "/source/%s/validation/%s/%s/%s"
)
```

```go
var ErrRequestForAccessDenied = errors.New("Access to API is denied by accessman")
```
ErrRequestForAccessDenied Error message if Accessman denied the access to API

#### type Accessman

```go
type Accessman struct {
	API  string
	URL  string
	Auth Auth
}
```

Accessman custom Accessman API lib

#### func  NewAccessRequest

```go
func NewAccessRequest(API string, host string) Accessman
```
NewAccessRequest creates a new accessman instance

#### func (*Accessman) AddNew

```go
func (am *Accessman) AddNew(email string)
```
AddNew add and account and generate OPK

#### func (*Accessman) CreateSourceValidationURL

```go
func (am *Accessman) CreateSourceValidationURL(clientIP string, api string, path string, method string) string
```
CreateSourceValidationURL Create source validation url

#### func (*Accessman) GetClientIP

```go
func (am *Accessman) GetClientIP(r *http.Request) string
```
GetClientIP get the remote ip of the client

#### func (*Accessman) SetAuth

```go
func (am *Accessman) SetAuth(user string, pass string)
```
SetAuth sets the username and password that we will be using for basic
authentication

#### func (*Accessman) Validate

```go
func (am *Accessman) Validate(clientIP string, path string, method string, username string, password string) error
```
Validate source validatation

#### func (*Accessman) VerifyAccess

```go
func (am *Accessman) VerifyAccess(clientIP string, path string, method string, username string, password string) error
```
VerifyAccess Checks if client is allowed to access this API

#### type Auth

```go
type Auth struct {
	User string
	Pass string
}
```

Auth holds the username and password for basic authentication

#### type Connection

```go
type Connection struct {
	Headers map[string]string
}
```

Connection connection structure

#### func (*Connection) Get

```go
func (connection *Connection) Get(url string, auth Auth) (*http.Response, error)
```
Get http get request

#### func (*Connection) Send

```go
func (connection *Connection) Send(method, url, payload string, auth Auth) (*http.Response, error)
```
Send send the http request

#### type Hash

```go
type Hash struct {
}
```

Hash represents a Hashable string

#### func (Hash) MD5

```go
func (h Hash) MD5() string
```
MD5 calculates the MD5 hash of a string

#### func (Hash) SHA1

```go
func (h Hash) SHA1() string
```
SHA1 calculates the SHA1 hash of a string

#### type Mock

```go
type Mock struct {
	Port   string
	Router *mux.Router
}
```

Mock Fake accessman server for testing

#### func  NewMockServer

```go
func NewMockServer(port string) Mock
```
NewMockServer creates new mock server

#### func (*Mock) Create

```go
func (mock *Mock) Create()
```
Create new fake accessman server

#### func (*Mock) GetAddress

```go
func (mock *Mock) GetAddress() string
```
GetAddress get the IP address of this mock server

#### func (*Mock) GetServerURL

```go
func (mock *Mock) GetServerURL() string
```
GetServerURL mock server's URL

#### func (*Mock) Run

```go
func (mock *Mock) Run()
```
Run run fake accessman server

#### func (*Mock) SourceValidation

```go
func (mock *Mock) SourceValidation(w http.ResponseWriter, r *http.Request)
```
SourceValidation ...
