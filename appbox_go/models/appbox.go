package models

import (
	"appbox_go_v/utils"
)

// App_info  Model
type App_info struct {
	ID          int    `json:"id"`
	App_name    string `json:"app_name"`
	App_type    string `json:"app_type"`
	Domain      string `json:"domain"`
	Stack       string `json:"stack"`
	Fontport    string `json:"fontport"`
	Backendport string `json:"backendport"`
	Delete_flg  string `gorm:"default:'0'"` //gorm 设置默认值
	Others      string `gorm:"default: ''"` //gorm 设置默认值
}

// TableName 解决gorm表明映射
// gorm框架表名自动加s问题
func (App_info) TableName() string {
	return "app_info"
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// Appbox 创建todo
// 添加

func CreateDt(dt *App_info) (err error) {
	err = utils.DB.Create(&dt).Error
	return
}

// 查看所有的待办事项

func GetAllDt() (dtList []*App_info, err error) {
	if err = utils.DB.Debug().Where("delete_flg=?", "0").Find(&dtList).Error; err != nil {
		return nil, err
	}
	return
}

// 查一条todo

func GetOneDt(id string) (dt *App_info, err error) {
	dt = new(App_info)
	if err = utils.DB.Debug().Where("id=?", id).Find(dt).Error; err != nil {
		return nil, err
	}
	return
}

// 修改某一个待办事项

func UpdateOneDt(dt *App_info) (err error) {
	err = utils.DB.Save(dt).Error
	return
}

// 删除某一个待办事项

func DeleteOneDt(id string) (dt *App_info, err error) {
	dt = new(App_info)
	if err = utils.DB.Debug().Where("id=?", id).Find(dt).Update("delete_flg", "1").Error; err != nil {
		return nil, err
	}

	return
}
