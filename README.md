# NewtonGolang - A Newton Wrapper for Golang

This library is a simple microlibrary for the Newton math API. The package exposes a struct (newton.Data) for wrapping the data sent to and from the service as well as a general method newton.Calc() to call the API.

Here is a sample code snippet that may use the package.

```go
import (
    "github.com/dmps/newton",
    "fmt"
    )

getResultFromNewton(operation,expression string){
    data := newton.Data{Operation:operation,Expression:expression}
    response := newton.Calc(data)
    fmt.Println(response.Result)
}
```

If you want to learn more, feel free to delve into the code or the tests.