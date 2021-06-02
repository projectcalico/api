# projectcalico/api

This is canonical source for API definitions of Projectcalico.

## How to use

One way is to import the clientset directly and use it. See [examples/list-gnp/main.go](examples/list-gnp/main.go) for some example code.

## Adding new APIs
1. Create a new .go file contains the spec struct to `pkg/apis/<apigroup>/<version>/calico`

1. Add the new types to `pkg/apis/<apigroup>/types.go`

1. Add the new types to `pkg/apis/<apigroup>/<version>/types.go`

1. Add the new types to `pkg/apis/<apigroup>/register.go`

1. Update generated code, including clients, informers, etc.

   ```
   make build
   ```
