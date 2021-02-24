package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"test/common"
	"test/config"
	"test/types"
)

func DingDingTalk(c *gin.Context) {
	req := types.Content{}
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	req.Content = common.FilterContent(data)
	dingData := types.DingTalkReq{
		Msgtype: "text",
		Text:    req,
	}
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + config.AccessToken
	resp, data, err := common.HttpPost(url, dingData)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	log.Println(string(data))
	if resp.StatusCode != 200 {
		c.JSON(http.StatusOK, "resp.StatusCode != 200")
		return
	}
	c.JSON(http.StatusOK, "success")
	return
}
