package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Lunch struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 从 lunchs.json 文件中读取食物列表
	lunchs, err := loadLunchsFromFile("lunchs.json")
	if err != nil {
		fmt.Println("Error loading lunchs:", err)
		return
	}

	// 定义 HTTP 处理函数
	http.HandleFunc("/lunch", func(w http.ResponseWriter, r *http.Request) {
		// 获取随机食物
		randomLunch := lunchs[rand.Intn(len(lunchs))]

		// 将午饭信息输出为 JSON 格式
		fmt.Fprintf(w, `{"name": "%s", "url": "%s"}`, randomLunch.Name, randomLunch.URL)
	})

	// 在端口 8080 上启动 HTTP 服务器
	fmt.Println("Starting server on port 9999...")
	http.ListenAndServe(":9999", nil)
}

// loadLunchsFromFile 从指定的 JSON 文件中读取食物列表
func loadLunchsFromFile(filename string) ([]Lunch, error) {
	// 读取文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解码 JSON 数据
	var lunchs []Lunch
	err = json.Unmarshal(data, &lunchs)
	if err != nil {
		return nil, err
	}

	return lunchs, nil
}
