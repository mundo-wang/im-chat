package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/dao/model"
	"strconv"
	"text/template"
)

type FrontApi struct {
}

func GetFrontApi() *FrontApi {
	return &FrontApi{}
}

func (api *FrontApi) GetIndex(c *gin.Context) (interface{}, error) {
	files, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		wlog.Error("call template.ParseFiles failed").Err(err).Log()
		return nil, err
	}
	err = files.Execute(c.Writer, "index")
	if err != nil {
		wlog.Error("call files.Execute failed").Err(err).Log()
		return nil, err
	}
	return nil, nil
}

func (api *FrontApi) ToRegister(c *gin.Context) (interface{}, error) {
	files, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		wlog.Error("call template.ParseFiles failed").Err(err).Log()
		return nil, err
	}
	err = files.Execute(c.Writer, "register")
	if err != nil {
		wlog.Error("call files.Execute failed").Err(err).Log()
		return nil, err
	}
	return nil, nil
}

func (api *FrontApi) ToChat(c *gin.Context) (interface{}, error) {
	ind, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html",
	)
	if err != nil {
		wlog.Error("call template.ParseFiles failed").Err(err).Log()
		return nil, err
	}
	userIdStr := c.Query("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		wlog.Error("call strconv.Atoi failed").Err(err).Field("userId", userIdStr).Log()
		return nil, err
	}
	token := c.Query("token")
	user := model.Users{}
	user.ID = userId
	user.Identity = token
	err = ind.Execute(c.Writer, user)
	if err != nil {
		wlog.Error("call ind.Execute failed").Err(err).Field("user", user).Log()
		return nil, err
	}
	return nil, nil
}
