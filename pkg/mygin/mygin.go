package mygin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/XOS/Probe/model"
	"github.com/XOS/Probe/service/dao"
)

const defaultUserAgent = "Mozilla/5.0 (compatible; NGBot/2.1; +https://server.nange.cn/)"

var adminPage = map[string]bool{
	"/server":       true,
	"/monitor":      true,
	"/setting":      true,
	"/notification": true,
	"/cron":         true,
}

func CommonEnvironment(c *gin.Context, data map[string]interface{}) gin.H {
	data["MatchedPath"] = c.MustGet("MatchedPath")
	data["Version"] = dao.Version
	data["Conf"] = dao.Conf
	// 是否是管理页面
	data["IsAdminPage"] = adminPage[data["MatchedPath"].(string)]
	// 站点标题
	if t, has := data["Title"]; !has {
		data["Title"] = dao.Conf.Site.Brand
	} else {
		data["Title"] = fmt.Sprintf("%s - %s", t, dao.Conf.Site.Brand)
	}
	u, ok := c.Get(model.CtxKeyAuthorizedUser)
	if ok {
		data["Admin"] = u
	}
	return data
}

func RecordPath(c *gin.Context) {
	c.Header("User-Agent", "Mozilla/5.0 (compatible; NGBot/2.1; +https://server.nange.cn/)")
	url := c.Request.URL.String()
	for _, p := range c.Params {
		url = strings.Replace(url, p.Value, ":"+p.Key, 1)
	}
	c.Set("MatchedPath", url)
}
