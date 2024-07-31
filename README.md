## 随机选择吃什么
数据保存在 lunchs.json

### 打包
静态编译，arm64 linux

    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o what-lunch-linux main.go

### 部署

    ./what-lunch-linux