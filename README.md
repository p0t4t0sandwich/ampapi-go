# ampapi-go

## This implementation is under active development, please feel free to contribute or create an issue if you've found anything that needs fixing

This API allows you to communicate with AMP installations from within Go.

Documentation for available API calls can be found by appending /API to the URL of any existing AMP installation.

Please Note: This program is directly based on the [ampapi-node](https://github.com/CubeCoders/ampapi-node) implementation and is almost verbatim in most aspects.

## Installation

```
Need to setup as go package
```

## Example

```go
package main

import (
    "ampapi/ampapi"
    "fmt"
)

func main() {
    API := ampapi.NewAMPAPI("http://localhost:8080/")

    res := API.Login("username", "password", "", false)

    fmt.Println(res)

    res2 := API.Call("ADSModule", "GetInstances", nil)

    fmt.Println(res2)
}
```

## Additional Notes

This implementation is basic at most, I've had issues implementing the init() functionality of the ampapi-node and ampapi-python libraries. I don't know enough Go to prevent the `module` and `method` parameters from pulling the latest value of the variable (in Python I had just force-set the default function parameters which decoupled the string object).

The alternative to this will be to manually auto-generate the API methods based on the latest APISpec.

## ToDo

- [ ] Get ampapi-go set up as a Go package

- [ ] Implement `ampapi.*AMPAPI.Init()`

- [ ] Implement async methods (be it using lambdas or auto-generated functions)
