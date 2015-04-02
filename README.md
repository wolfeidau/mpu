# mpu [![GoDoc](https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat)](http://godoc.org/github.com/wolfeidau/mpu)

This library provides a simple server and client for uploading files using multipart forms and supports gzip encoding.

# overview


```
go get github.com/wolfeidau/mpu
```

# Client Example

```

uri := "https://localhost:9090/uploads"

extraParams := map[string]int{
    "author":   "Mark Wolfe",
    "hostname": "wolfesmachine.local",
}

uploader := mpu.Uploader(http.Client{}, mpu.DefaultConfig()) // gzip encoding by default.

req, err := uploader.NewFileRequest(uri, extraParams, "fileUpload", "/tmp/output.log")

if err != nil {
	log.Fatalf("building req failed: %s", err)
}

req.ContentType("text/csv")

client := &http.Client{}

resp, err := client.Do(request)

if err != nil {
	log.Fatalf("post failed: %s", err)
}

log.Printf("success status=%d msg=%s", resp.StatusCode, resp.Body)
```


# License

This code is Copyright (c) 2014 Mark Wolfe and licenced under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.