# projectcalico/api

This repository is a mirror of the canonical home for Calico API definitions.

It is recommended to use this repository for projects which use the Calico API, so as to avoid importing the full Calico source code.

To make changes to the Calico API, you must first submit a PR against [github.com/projectcalico/calico/api](https://github.com/projectcalico/calico).
Once merged, changes will be mirrored to this repository for consumption.

## How to use

One way is to import the clientset directly and use it. See [examples/list-gnp/main.go](examples/list-gnp/main.go) for some example code.

## Adding new APIs
1. Create a .go file which contains the new type to `pkg/apis/<apigroup>/<version>`

1. Add the new type to `pkg/apis/<apigroup>/<version>/register.go`

1. Update generated code, including clients, informers, etc.

   ```
   make build
   ```
