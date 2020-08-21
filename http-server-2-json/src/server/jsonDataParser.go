package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getJsonData(req *http.Request) (url.Values, error) {
	// result
	mymap := make(map[string]string)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return url.Values{}, err
	}
	err = json.Unmarshal(body, &mymap)
	if err != nil {
		fmt.Println(err.Error())
		return url.Values{}, err
	}
	return jsonToUrlValues(&mymap), nil
}

func jsonToUrlValues(mymap *map[string]string) url.Values {
	values := make(url.Values)
	for k, v := range *mymap {
		values.Add(k, v)
	}
	return values
}
