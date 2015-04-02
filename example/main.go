package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wolfeidau/mpu"
	"gopkg.in/alecthomas/kingpin.v1"
)

const uri = "http://localhost:9090/upload"

var (
	app    = kingpin.New("uploader", "A command-line file upload example.")
	client = app.Command("client", "uploader client.")
	server = app.Command("server", "uploader server.")
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case client.FullCommand():
		startClient()
	case server.FullCommand():
		startServer()
	}

}

func startClient() {

	extraParams := map[string]string{
		"author":   "Mark Wolfe",
		"hostname": "wolfesmachine.local",
	}

	uploader := mpu.Uploader(mpu.DefaultConfig()) // gzip encoding by default.

	req, err := uploader.NewFileRequest(uri, extraParams, "fileUpload", "/tmp/output.log")

	if err != nil {
		log.Fatalf("building req failed: %s", err)
	}

	start := time.Now()

	resp, err := uploader.Do(req)

	if err != nil {
		log.Fatalf("post failed: %s", err)
	}

	defer resp.Body.Close()

	log.Printf("success status=%d timetaken=%s", resp.StatusCode, time.Now().Sub(start))

}

func startServer() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("reading body")
		written, err := io.Copy(ioutil.Discard, r.Body)

		if err != nil {
			log.Printf("failed to read body: %s", err)
			w.WriteHeader(500)
			return
		}

		defer r.Body.Close()

		log.Printf("read len=%d", written)

		w.WriteHeader(200)
	})

	http.ListenAndServe(":9090", nil)
}
