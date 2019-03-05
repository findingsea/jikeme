package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://app.jike.ruguoapp.com/1.0/squarePosts/list"
	jsonStr := []byte(`{"topicId": "5618c159add4471100150637", "limit": 1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	contentData := dat["data"].([]interface{})
	content := contentData[0].(map[string]interface{})
	fmt.Println(content["content"].(string))
	fmt.Println("--", (content["topic"].(map[string]interface{})["content"].(string)))
}
