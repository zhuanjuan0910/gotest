package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// 下面测试binding数据
	// 首先测试binding-JSON,
	// 注意Body中的数据必须是JSON格式
	resp, _ := http.Post("http://0.0.0.0:8000/bindJSON", "application/json", strings.NewReader("{\"user\":\"TAO\", \"password\": \"123\"}"))
	helpRead(resp)

	// 下面测试bind FORM数据
	resp, _ = http.Post("http://0.0.0.0:8000/bindForm", "application/x-www-form-urlencoded", strings.NewReader("user=TAO&password=123"))
	helpRead(resp)

	// 下面测试接收JSON和XML数据
	resp, _ = http.Get("http://0.0.0.0:8000/someJSON")
	helpRead(resp)
	resp, _ = http.Get("http://0.0.0.0:8000/moreJSON")
	helpRead(resp)
	resp, _ = http.Get("http://0.0.0.0:8000/someXML")
	helpRead(resp)
}

// 用于读取resp的body
func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}
