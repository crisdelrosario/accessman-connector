# ACCESSMAN LIB #

### Create A Mock Server For Testing ###


```
#!go

import (
    "fmt"
    "bitbucket.org/cdelrosario03/accessman"
)

Port := "8443"
mock := accessman.NewMockServer(Port)
mock.Run()

fmt.Println("Mock Server URL: " + mock.GetServerURL())
```

### Access Verification ###

```
#!go
import (
  "fmt"
  "bitbucket.org/cdelrosario03/accessman"
)

const (
  API = "am"
  AccessmanURL = "http://0.0.0.0:8843"
)

am := accessman.NewAccessRequest(API, AccessmanURL)

clientIP := "0.0.0.0"
path := "/client"
method := "Get"
username := "OPK00000000000001"
password := "password"

if err := am.VerifyAccess(clientIP, path, method, username, password); err !=  nil {
  fmt.Error(err.Error())
}
```
