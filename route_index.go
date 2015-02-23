package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {

	x := requestRandom()

	templates := template.Must(template.ParseFiles("template/index.html"))
	err := templates.Execute(writer, x)
	if err != nil {
		log.Fatal(err)
	}

}

func requestRandom() []interface{} {
	url := "https://api.random.org/json-rpc/1/invoke"

	var jsonStr = []byte(`{"jsonrpc":"2.0","method":"generateIntegers","params":{"apiKey":"30dc7b26-988c-49cf-9732-0b24c23d203e","n":6,"min":1,"max":49,"replacement":false},"id":3076}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var msg interface{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&msg)
	log.Println(msg)

	m := msg.(map[string]interface{})
	mm := m["result"].(map[string]interface{})
	mmm := mm["random"].(map[string]interface{})
	d := mmm["data"].([]interface{})

	log.Println(m)
	log.Println(mm)
	log.Println(mmm)
	log.Println(d)

	return d
}
