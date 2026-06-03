package config

import (
	"github.com/outsstill/gin-admin/model"
)

type Config struct {
	model.BaseModel
	ConfigKey   string `gorm:"column:config_key;index;type:varchar(255)" json:"config_key"`
	ConfigValue string `gorm:"column:config_value;type:varchar(255)" json:"config_value"`
	ConfigLabel string `gorm:"column:config_label;type:varchar(255);index" json:"config_label"`
	Type        int    `gorm:"column:type;type:tinyint" json:"type"`
	Options     string `gorm:"column:options;type:text" json:"options"`
	Describe    string `gorm:"column:describe;type:text" json:"describe"`
	IsCanFront  int    `gorm:"column:is_can_front;type:tinyint" json:"is_can_front"`
	IsRequired  uint   `gorm:"column:is_required;type:tinyint" json:"is_required"`
	Order       uint   `gorm:"column:order" json:"order"`
	GroupId     uint   `gorm:"column:group_id" json:"group_id"`
	State       uint   `gorm:"column:state;type:tinyint" json:"state"`
	ShowType    string `gorm:"column:show_type;type:varchar(255)" json:"show_type"`
	Placeholder string `gorm:"column:placeholder;type:varchar(255)" json:"placeholder"`
	model.CommonTimestampsField
}

func (model *Config) TableName() string {
	return "configs"
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
