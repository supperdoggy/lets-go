package test

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

func postRequest(){
    formData := url.Values{
        "name": {"Max"},
    }

    resp, err := http.PostForm("localhost:8090/hello", formData)
    if err != nil{
        fmt.Println(err)
    }
    result := make(map[string]interface{})
    json.NewDecoder(resp.Body).Decode(&result)

    fmt.Println(result["form"])
}

func main() {
	postRequest()
}


