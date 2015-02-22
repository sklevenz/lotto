package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ts *httptest.Server

func startTestServer() {
	initLogger()
	loadConfig()
	ts = httptest.NewServer(createHandler())
}

func TestStaticFile(t *testing.T) {
	startTestServer()
	defer ts.Close()

	res, err := http.Get(ts.URL + "/static/test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", body)
}

func TestIndex(t *testing.T) {
	startTestServer()
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("%s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", body)
}
