package main

import (
	"core-api/bootstrap"
	"core-api/web/middleware"
	"core-api/web/router"
	"flag"
	"log"
	"os"
	"strconv"
)

//在 Mac、Linux、Windows 下Go交叉编译
//https://blog.csdn.net/x356982611/article/details/80054314
/**Mac 下编译 Linux 和 Windows 64位可执行程序
* CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
* CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
nohup ./main -p 9001 ->info.log &
netstat -nap|grep 9001
kill -9 21215
**/

// 启动地址及端口
const DEFAULTPORT  = 8080

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Core-Api", "Saya")
	app.Bootstrap()
	app.Configure(middleware.Configure, router.Configure)
	return app
}

func main() {
	port := flag.Int("p", DEFAULTPORT, "Set The Http Port")
	flag.Parse()
	pwd,_ := os.Getwd()
	log.Printf("Listen On Port:%d pwd:%s\n", *port, pwd)

	app := newApp()
	app.Listen(":" + strconv.Itoa(*port))
}
