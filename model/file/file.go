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
	OriginName   string    `gorm:"column:origin_name" json:"origin_name"`
	Name         string    `gorm:"column:name" json:"name"`
	Key          string    `gorm:"column:key" json:"key"`
	GroupId      int       `gorm:"column:group_id;index" json:"group_id"`
	Size         int64     `gorm:"column:size" json:"size"`
	Storage      string    `gorm:"column:storage" json:"storage"`
	UploadType   int       `gorm:"column:upload_type" json:"upload_type"`
	Path         string    `gorm:"column:path" json:"-"`
	Type         int       `gorm:"column:type;default:1" json:"type"`
	Ext          string    `gorm:"column:ext" json:"ext"`
	UserId       uint64    `gorm:"column:user_id" json:"-"`
	Url          string    `gorm:"column:url" json:"url"`
	MimeType     string    `gorm:"column:mime_type" json:"mime_type"`
	ContentType  string    `gorm:"column:content_type" json:"content_type"`
	ETag         string    `gorm:"column:e_tag" json:"e_tag"`
	Bucket       string    `gorm:"column:bucket" json:"bucket"`
	LastModified time.Time `gorm:"column:last_modified" json:"last_modified"`
	FullUrl      string    `gorm:"-" json:"full_url"`
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
			url = url + "/" + model.Key
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
