package main

import (
	"fmt"
	"ncfwxen/common/router"
	"ncfwxen/common/utils"
	"ncfwxen/config"
	"ncfwxen/controller/admin"
	"ncfwxen/controller/web"
	"net/http"
)

func main() {
	router := httprouter.New()

	//web route
	router.GET("/", web.Index)
	router.GET("/article", web.Index)
	router.GET("/article/:id", web.Article)
	router.GET("/admin/delcate", admin.DelCategory)
	router.POST("/admin/delarticle", admin.DelArticle)

	//admin route
	router.GET("/admin", admin.Index)
	router.GET("/admin/article", admin.Article)
	router.GET("/admin/login", admin.Login)
	router.POST("/admin/dologin", admin.DoLogin)
	utils.Nlog.Println("Start Ncfwxen...")
	sysConf := *config.SysConfig()
	addr := fmt.Sprintf("%s:%s", sysConf.Domain, sysConf.HttpPort)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	utils.Nlog.Fatal(http.ListenAndServe(addr, router))
}
