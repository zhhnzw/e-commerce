package models

import (
	"backend-service/utils"
	"errors"
	"log"
	"strconv"
	"time"
)

const SysUserTableName string = "sys_user"

type SysUser struct {
	Id          int            `json:"id"`
	Province    string         `json:"province"`
	UserName    string         `json:"userName"`
	NickName    string         `json:"nickName"`
	Password    string         `json:"-"`
	Mobile      string         `json:"mobile"`
	Email       string         `json:"email"`
	IsValid     bool           `json:"isValid"`
	IsSuper     bool           `json:"isSuper"`
	Avatar      string         `json:"avatar"`
	CreatedTime time.Time      `json:"-" form:"-" gorm:"-"`
	UpdatedTime utils.JSONTime `json:"updateTime" form:"-" gorm:"-"`
	PageSize    int            `gorm:"-" json:"-" form:"-"`
	PageIndex   int            `gorm:"-" json:"-" form:"-"`
	FilterValue string         `gorm:"-" json:"-" form:"-"`
}

var SysUserQueryFields []string

func init() {
	SysUserQueryFields = []string{"id", "user_name", "nick_name", "password", "mobile", "email", "is_valid", "is_super", "avatar", "updated_time"}
}

func (model *SysUser) CreateUser() (bool, error) {
	DB.NewRecord(model)
	d := DB.Table(SysUserTableName).Create(model)
	if DB.NewRecord(model) {
		log.Printf("mysql %s 插入失败, model:%+v", SysUserTableName, model)
		return false, d.Error
	}
	return true, nil
}

func (model *SysUser) GetUsers() ([]SysUser, int, error) {
	results := make([]SysUser, 0, model.PageSize)
	db := DB.Table(SysUserTableName).Select(SysUserQueryFields)
	if len(model.FilterValue) > 0 {
		if _, err := strconv.Atoi(model.FilterValue); err != nil {
			db = db.Where("real_name LIKE ?", "%"+model.FilterValue+"%")
		} else {
			db = db.Where("mobile LIKE ?", "%"+model.FilterValue+"%")
		}
	}
	if len(model.Province) > 0 {
		db = db.Where("province=?", model.Province)
	}
	var rows int
	db.Count(&rows)
	db = db.Order("updated_time DESC").
		Limit(model.PageSize).Offset((model.PageIndex - 1) * model.PageSize).Find(&results)
	return results, rows, db.Error
}

func (model *SysUser) GetUser() (*SysUser, error) {
	db := DB.Table(SysUserTableName).Select(SysUserQueryFields).Where("id=?", model.Id).First(model)
	return model, db.Error
}

func (model *SysUser) UpdateUser() (int64, error) {
	if model.Id < 1 && len(model.UserName) < 1 {
		return 0, errors.New("id 和 user_name 必传其一")
	}
	db := DB.Table(SysUserTableName)
	if model.Id > 0 {
		db = db.Where("id=?", model.Id)
	} else if len(model.UserName) > 0 {
		db = db.Where("user_name=?", model.UserName)
	}
	db.Update("is_valid", model.IsValid)
	db = db.Updates(*model)
	return db.RowsAffected, db.Error
}

func (model *SysUser) DeleteUser() error {
	if model.Id < 1 {
		return errors.New("必须按id号删除记录")
	}
	var userModel SysUser
	db := DB.Select("id, user_name").Table(SysUserTableName).Where("id=?", model.Id).Find(&userModel)
	db = DB.Table(SysUserTableName).Where("id=?", model.Id).Delete(&userModel)
	if db.Error != nil {
		return db.Error
	} else {
		return nil
	}
}
