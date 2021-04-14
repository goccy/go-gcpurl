# go-gcpurl

![Go](https://github.com/goccy/go-gcpurl/workflows/Go/badge.svg)
[![GoDoc](https://godoc.org/github.com/goccy/go-gcpurl?status.svg)](https://pkg.go.dev/github.com/goccy/go-gcpurl?tab=doc)

Parse the URL to get the GCP projectID in Go

# Synopsis

```go
package main

import (
	"fmt"

	"github.com/goccy/go-gcpurl"
)

func main() {
	url, err := gcpurl.ParseURL(`https://us-central1-awesome-project-id.cloudfunctions.net`)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.Host)          // us-central1-awesome-project-id.cloudfunctions.net
	fmt.Println(url.Region)        // us-central1
	fmt.Println(url.ProjectID)     // awesome-project-id
	fmt.Println(url.ServiceDomain) // cloudfunctions.net
}
```

# License

MIT