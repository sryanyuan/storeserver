package main

import (
	"io/ioutil"

	"encoding/json"

	"net/http"

	"github.com/gin-gonic/gin"
)

func adminSetHandler(c *gin.Context) {
	rsp := HTTPResp{}
	data, err := ioutil.ReadAll(c.Request.Body)
	if nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	var pair StoreKVPair
	if err = json.Unmarshal(data, &pair); nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	if err = storeSet(pair.Key, pair.Value); nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	c.JSON(http.StatusOK, &rsp)
}

func adminDeleteHandler(c *gin.Context) {
	rsp := HTTPResp{}
	data, err := ioutil.ReadAll(c.Request.Body)
	if nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	var pair StoreKVPair
	if err = json.Unmarshal(data, &pair); nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	if err = storeDelete(pair.Key); nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	c.JSON(http.StatusOK, &rsp)
}

func getHandler(c *gin.Context) {
	var err error
	rsp := HTTPResp{}

	key := c.Request.FormValue("key")
	if 0 == len(key) {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = "invalid key"
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	var value string
	if value, err = storeGet(key); nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	rsp.Message = value
	c.JSON(http.StatusOK, &rsp)
}

func getAllHandler(c *gin.Context) {
	var err error
	var values []StoreKVPair
	rsp := HTTPResp{}

	if values, err = storeGetAll(); nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	data, err := json.Marshal(values)
	if nil != err {
		rsp.Code = HTTPRspCodeInternalError
		rsp.Message = err.Error()
		c.JSON(HTTPStatusError, &rsp)
		return
	}

	rsp.Message = string(data)
	c.JSON(http.StatusOK, &rsp)
}
