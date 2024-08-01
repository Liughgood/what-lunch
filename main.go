package main

import (
        "encoding/json"
        "fmt"
        "html/template"
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

        // 定义 HTML 模板
        tmpl := template.Must(template.New("lunch").Parse(`
        <!DOCTYPE html>
        <html>
        <head>
                <title>随机午饭</title>
        </head>
        <body>
                <h1>随机午饭</h1>
                <p>名称: {{.Name}}</p>
                <p>链接: <a href="{{.URL}}">大众点评</a></p>
        </body>
        </html>
`))

        // 定义 HTTP 处理函数
        http.HandleFunc("/lunch", func(w http.ResponseWriter, r *http.Request) {
                randomLunch := lunchs[rand.Intn(len(lunchs))]

                // 使用 HTML 模板渲染响应
                err := tmpl.Execute(w, randomLunch)
                if err != nil {
                        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                        return
                }
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