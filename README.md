# mpu [![GoDoc](https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat)](http://godoc.org/github.com/wolfeidau/mpu)

This library provides a simple server and client for uploading files using multipart forms and supports gzip encoding.

# overview


```
go get github.com/wolfeidau/mpu
```

# Client Example

```go

uri := "https://localhost:9090/uploads"

extraParams := map[string]string{
	"author":   "Mark Wolfe",
	"hostname": "wolfesmachine.local",
}

uploader := mpu.Uploader(mpu.DefaultConfig()) // gzip coming soon!

req, err := uploader.NewFileRequest(uri, extraParams, "fileUpload", "/tmp/output.log")

// is this a local issue that I probably want to quit based on.
if err != nil {
	log.Fatalf("building req failed: %s", err)
}

start := time.Now()

resp, err := uploader.Do(req)

// is this a network issue which is "hopefully" transient.
if err != nil {
	log.Fatalf("post failed: %s", err)
}

defer resp.Body.Close()

log.Printf("Success status=%d timetaken=%s", resp.StatusCode, time.Now().Sub(start))

```


# License

This code is Copyright (c) 2014 Mark Wolfe and licenced under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.