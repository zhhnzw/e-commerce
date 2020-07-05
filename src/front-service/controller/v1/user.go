package v1

import (
	"front-service/models"
	"front-service/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

const (
	StatusOk       string = "操作成功!"
	GetFailed      string = "操作失败!"
	CreateFailed   string = "创建失败!"
	CreateRepeated string = "重复创建!"
	DeleteFailed   string = "删除失败!"
	UpdateFailed   string = "更新失败!"
)

type UserForm struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	NickName  string `json:"nickName"`
	Password  string `json:"password"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	IsValid   bool   `json:"isValid"`
	PageSize  int    `gorm:"-"`
	PageIndex int    `gorm:"-"`
}

func Login(c *gin.Context) {
	var form UserForm
	err := c.ShouldBind(&form)
	utils.CheckErr(err, "")
	model := models.User{
		UserName: form.UserName,
		Password: form.Password,
	}
	resp := utils.Resp{Data: make(map[string]string), Code: "1"}
	var model1 models.User
	db := models.DB.Table("sys_user").Select(models.SysUserQueryFields).Where("user_name=?", model.UserName).First(&model1)
	if db.Error != nil && !strings.Contains(db.Error.Error(), "record not found") {
		resp.Message = "服务端故障, 查询用户信息失败!"
		utils.Logf(db.Error, "")
		c.JSON(http.StatusInternalServerError, resp)
		return
	} else if db.RowsAffected == 0 {
		resp.Message = "没有这个账号!"
		c.JSON(http.StatusBadRequest, resp)
		return
	} else if model1.IsValid == false {
		resp.Message = "该账号已被禁用, 请与管理员联系!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	utils.CheckErr(db.Error, "")
	err = bcrypt.CompareHashAndPassword([]byte(model1.Password), []byte(model.Password))
	if err != nil {
		resp.Message = "密码错误!"
	} else {
		resp.Message = StatusOk
		data := make(map[string]interface{})
		data["avatar"] = model1.Avatar
		data["email"] = model1.Email
		data["mobile"] = model1.Mobile
		data["nickName"] = model1.NickName
		data["userName"] = model1.UserName
		resp.Data = data
		resp.Code = "0"
	}
	session := sessions.Default(c)
	session.Set("userName", model.UserName)
	err = session.Save()
	if err != nil {
		utils.Logf(err, "")
		c.JSON(http.StatusInternalServerError, utils.Resp{Message: "服务端故障, session操作失败", Code: "1"})
		return
	}
	if err := utils.RedisClient.SAdd(utils.RedisKeyLoginUsers, model.UserName).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Resp{Message: "服务端故障, redis增加登录用户失败", Code: "1"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("userName")
	if err := utils.RedisClient.SRem(utils.RedisKeyLoginUsers, userName).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Resp{Message: "服务端故障, redis剔除登录用户失败", Code: "1"})
		return
	}
	session.Clear()
	c.JSON(http.StatusOK, utils.Resp{Code: "0"})
}

type alterPwdStr struct {
	OldPwd string `json:"oldPwd"`
	NewPwd string `json:"newPwd"`
}

func AlterPwd(c *gin.Context) {
	var model alterPwdStr
	err := c.ShouldBind(&model)
	if err != nil {
		utils.Logf(err, "")
		c.JSON(http.StatusBadRequest, utils.Resp{Data: nil, Message: "参数不正确", Code: "1"})
		return
	}
	session := sessions.Default(c)
	userName := session.Get("userName")
	resp := utils.Resp{Code: "1"}
	var model1 models.User
	db := models.DB.Table("sys_user").Select(models.SysUserQueryFields).Where("user_name=?", userName).First(&model1)
	if db.Error != nil {
		resp.Message = "没有这个账号!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	utils.CheckErr(db.Error, "")
	err = bcrypt.CompareHashAndPassword([]byte(model1.Password), []byte(model.OldPwd))
	if err != nil {
		resp.Message = "原密码错误!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	encodePassword, err := bcrypt.GenerateFromPassword([]byte(model.NewPwd), bcrypt.DefaultCost)
	utils.CheckErr(err, "")
	userModel := models.User{Password: string(encodePassword), UserName: userName.(string), IsValid: model1.IsValid}
	if _, err := userModel.UpdateUser(); err != nil {
		utils.Logf(err, "")
		resp.Message = UpdateFailed
		c.JSON(http.StatusBadRequest, resp)
		return
	} else {
		resp.Message = StatusOk
		resp.Code = "0"
		c.JSON(http.StatusOK, resp)
		return
	}
}

func CreateUser(c *gin.Context) {
	var model UserForm
	err := c.ShouldBind(&model)
	if err != nil {
		utils.Logf(err, "")
		c.JSON(http.StatusBadRequest, utils.Resp{Data: nil, Message: "参数不正确", Code: "1"})
		return
	}
	if len(model.UserName) < 1 {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: "userName 必传", Code: "1"})
		return
	}
	userModel := models.User{
		UserName: model.UserName,
		NickName: model.NickName,
		Password: model.Password,
		Mobile:   model.Mobile,
		Email:    model.Email,
		Avatar:   model.Avatar,
		IsValid:  true,
	}
	resp := utils.Resp{Data: make(map[string]string), Code: "1"}
	encodePassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	utils.CheckErr(err, "")
	userModel.Password = string(encodePassword)
	if succeed, e := userModel.CreateUser(); succeed {
		resp.Code = "0"
		c.JSON(http.StatusOK, resp)
	} else if strings.Contains(e.Error(), "Duplicate entry") {
		resp.Message = CreateRepeated
		c.JSON(http.StatusBadRequest, resp)
	} else {
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
}

func GetUsers(c *gin.Context) {
	var model models.User
	err := c.ShouldBind(&model)
	if err != nil {
		utils.Logf(err, "")
		c.JSON(http.StatusBadRequest, utils.Resp{Data: nil, Message: "参数不正确", Code: "1"})
		return
	}
	resp := utils.Resp{Data: make(map[string]string), Message: "", Code: "1"}
	if results, rows, err := model.GetUsers(); err != nil {
		utils.Logf(err, "")
		resp.Message = GetFailed
		resp.Data = results
	} else {
		resp.Message = StatusOk
		resp.Data = map[string]interface{}{"data": results, "count": rows}
		resp.Code = "0"
	}
	c.JSON(http.StatusOK, resp)
}
