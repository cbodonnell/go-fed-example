package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func fetchIRI(iri *url.URL) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", iri.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/activity+json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
