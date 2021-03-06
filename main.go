package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/xiet16/go_gateway_gin_scaffold/golang_common/lib"
	"github.com/xiet16/go_gateway_gin_scaffold/http_proxy_router"
	"github.com/xiet16/go_gateway_gin_scaffold/router"
)

//构建终端endpoint dashboard后台管理 server代理服务器
//config  ./conf/prod 对应配置文件夹

var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *endpoint == "dashboard" {
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		router.HttpServerStop()
	} else {
		//加载配置
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		go func() {
			http_proxy_router.HttpServerRun()
			fmt.Println("start proxy server")
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		http_proxy_router.HttpServerStop()
	}
}
