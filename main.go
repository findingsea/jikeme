package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	var topics = []string{
		"5618c159add4471100150637", // 浴室沉思
		"557ed045e4b0a573eb66b751", // 无用但有趣的冷知识
		"5a82a88df0eddb00179c1df7", // 今日烂梗
		"572c4e31d9595811007a0b6b", // 弱智金句病友会
		"56d177a27cb3331100465f72", // 今日金句
		"5aa21c7ae54af10017dc93f8", // 一个想法不一定对
		"5bf22b38ffa4f00017e1a8ff", // 有一点哲学在里面
		"5ab89ff5892a1a0011d1ba87", // 记一件小事
		"597ae4ac096cde0012cf6c06", // 科技圈大小事
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	topicdIndex := r.Intn(len(topics))
	url := "https://app.jike.ruguoapp.com/1.0/squarePosts/list"
	jsonStr := []byte(`{"topicId": "` + topics[topicdIndex] + `", "limit": 20}`)
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

	dataArr := dat["data"].([]interface{})
	dataIndex := r.Intn(len(dataArr))
	data := dataArr[dataIndex].(map[string]interface{})
	content := data["content"].(string)
	user := data["user"].(map[string]interface{})["screenName"].(string)
	topic := data["topic"].(map[string]interface{})["content"].(string)
	fmt.Println(content)
	fmt.Println("--", user+", 「"+topic+"」")
}
