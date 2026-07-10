package file

import (
	"strings"
	"time"

	"github.com/outsstill/gin-admin/model"
	"github.com/outsstill/gin-admin/pkg/helpers"
	gokit "github.com/outsstill/go-kit"
	"gorm.io/gorm"
)

type File struct {
	model.BaseModel
	OriginName   string     `gorm:"column:origin_name;type:varchar(255);index" json:"origin_name"`
	Name         string     `gorm:"column:name;type:varchar(255)" json:"name"`
	Key          string     `gorm:"column:key;type:varchar(255)" json:"key"`
	GroupId      int        `gorm:"column:group_id;index" json:"group_id"`
	Size         int64      `gorm:"column:size" json:"size"`
	Storage      string     `gorm:"column:storage;type:varchar(255);index" json:"storage"`
	Path         string     `gorm:"column:path;type:text" json:"path"`
	Type         int        `gorm:"column:type;type:tinyint" json:"type"`
	Ext          string     `gorm:"column:ext;type:varchar(255)" json:"ext"`
	UserId       uint64     `gorm:"column:user_id" json:"user_id"`
	Url          string     `gorm:"column:url;type:text" json:"url"`
	ContentType  string     `gorm:"column:content_type;type:varchar(255)" json:"content_type"`
	ETag         string     `gorm:"column:e_tag;type:varchar(255)" json:"e_tag"`
	Bucket       string     `gorm:"column:bucket;type:varchar(255)" json:"bucket"`
	LastModified time.Time  `gorm:"column:last_modified" json:"last_modified"`
	FullUrl      string     `gorm:"-" json:"full_url"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	model.CommonTimestampsField
}

func (model *File) TableName() string {
	return "files"
}

// 查询后
func (model *File) AfterFind(tx *gorm.DB) (err error) {
	model.FullUrl = model.GetFileFullUrl()
	return
}

// 获取文件完整访问路径
func (model *File) GetFileFullUrl() string {
	url := model.Url
	if helpers.Empty(url) {
		storageDrive := model.Storage
		url = ""
		if storageDrive == "local" {
			path := strings.ReplaceAll(model.Path, gokit.Config().Storage.Local.BasePath, gokit.Config().Storage.Local.StaticPrefix)
			url = gokit.Config().Storage.Local.BaseURL
			url = url + "/" + path
		} else if storageDrive == "oss" {
			url = gokit.Config().Storage.Oss.Domain
			url = url + "/" + model.Path
		}
	}
	return url
}

func (model *File) GetFileFullPath() string {
	path := model.Path
	if model.Storage == "local" {
		path = gokit.Config().Storage.Local.BasePath + "/" + path
	} else {
		path = model.Path
	}
	return path
}
