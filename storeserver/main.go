package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-validator/validator"
	"github.com/ngaut/log"
)

func main() {
	// read config
	config, err := loadConfigFromFile("app.json")
	if nil != err {
		log.Info(err)
		return
	}
	// validate config
	if err = validator.Validate(config); nil != err {
		log.Error(err)
		return
	}
	// get auth user
	if nil == config.AuthUser ||
		len(config.AuthUser) == 0 {
		log.Error("Invalid config.AuthUser")
		return
	}
	// init db
	if err = storeInit(config.StoreSource); nil != err {
		log.Error(err)
		return
	}
	defer storeUninit()
	// init auth user
	accounts := make(map[string]string)
	for _, v := range config.AuthUser {
		userInfo := strings.Split(v, ":")
		if 2 != len(userInfo) ||
			0 == len(userInfo[0]) ||
			0 == len(userInfo[1]) {
			log.Error("Invalid auth user : ", v)
			return
		}

		accounts[userInfo[0]] = userInfo[1]
	}

	r := gin.Default()
	// add auth middleware
	authMw := r.Group("/admin", gin.BasicAuth(gin.Accounts(accounts)))
	authMw.POST("/set", adminSetHandler)
	authMw.DELETE("/delete", adminDeleteHandler)
	// public handler
	r.GET("/get", getHandler)
	r.GET("/getall", getAllHandler)

	r.Run(config.HTTPAddress)
}
