# ACCESSMAN LIB #

### Create A Mock Server For Testing ###


```
import (
    "fmt"
    "github.com/crisdelrosario/accessman"
)

Port := "8443"
mock := accessman.NewMockServer(Port)
mock.Run()

fmt.Println("Mock Server URL: " + mock.GetServerURL())
```

### Access Verification ###

```
import (
  "fmt"
  "github.com/crisdelrosario/accessman"
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
```
