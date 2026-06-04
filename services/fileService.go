package service

import (
	"context"
	"errors"
	"strings"

	"github.com/outsstill/gin-admin/core"
	"github.com/outsstill/gin-admin/global"
	"github.com/outsstill/gin-admin/model"
	fileModel "github.com/outsstill/gin-admin/model/file"
	"github.com/outsstill/gin-admin/pkg/file"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/pkg/paginator"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type FileService struct {
	storage file.IStorage
	app     *core.App
}

func NewFileService(app *core.App, drive ...string) *FileService {

	fileDrive := global.Config.Storage.Driver

	if len(drive) > 0 {
		fileDrive = drive[0]
	}

	fileConfig := file.Config{
		Driver: fileDrive,
		LocalConfig: file.LocalConfig{
			BasePath:      global.Config.Storage.Local.Path,
			PublicBaseURL: global.Config.Storage.Local.Domain,
		},
		OssConfig: file.OssConfig{
			Region:     global.Config.Storage.Oss.Region,
			BucketName: global.Config.Storage.Oss.Bucket,
			Key:        global.Config.Storage.Oss.KeyId,
			Secret:     global.Config.Storage.Oss.KeySecret,
		},
	}
	fileStorage := file.NewStorage(fileConfig)
	return &FileService{
		storage: fileStorage,
		app:     app,
	}
}

func (service *FileService) UploadFile(c *gin.Context) (*fileModel.File, error) {
	// 从 form-data 获取文件
	fileObj, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer fileObj.Close()

	// 验证大小
	if header.Size > cast.ToInt64(global.Config.Storage.SizeLimit) {
		return nil, errors.New("超过最大文件大小")
	}

	// 验证后缀
	extLimit := cast.ToStringSlice(global.Config.Storage.Ext)
	if helpers.FindElement(extLimit, strings.ToLower(helpers.GetFileExt(header.Filename))) < 0 {
		return nil, errors.New("文件格式不允许 只允许[ " + strings.Join(extLimit, " ") + " ]")
	}

	input := file.PutObjectInput{
		Key:         header.Filename,
		Size:        header.Size,
		ContentType: header.Header.Get("Content-Type"),
		Reader:      fileObj,
		File:        header,
		Meta:        map[string]string{},
	}

	// 上传
	obj, putErr := service.storage.Put(c, input)
	if putErr != nil {
		return nil, putErr
	}

	// 存入数据库
	fileStore := &fileModel.File{}
	fileStore.Bucket = obj.Bucket
	fileStore.Name = obj.Name
	fileStore.OriginName = obj.OriginName
	fileStore.Path = obj.Path
	fileStore.Key = obj.Key
	fileStore.Size = obj.Size
	fileStore.Ext = obj.Ext
	fileStore.Storage = obj.Storage
	fileStore.ETag = obj.ETag
	fileStore.ContentType = obj.ContentType
	fileStore.LastModified = obj.LastModified
	fileStore.Url = obj.URL
	fileStore.UserId = 99
	fileStore.GroupId = cast.ToInt(c.DefaultPostForm("group_id", "99"))
	fileStore.Type = cast.ToInt(c.DefaultPostForm("type", "1"))

	// 组装url
	fileStore.FullUrl = fileStore.GetFileFullUrl()
	return fileStore, nil
}

func (service *FileService) DeleteFile(id string) error {
	fileObj := service.Get(id)
	if fileObj.ID <= 0 {
		return errors.New("没有找到该条记录")
	}
	// 上传
	err := service.storage.Delete(context.Background(), fileObj.Path)
	if err != nil {
		return err
	}

	service.Delete(&fileObj)

	return nil
}

func (service *FileService) Create(model *fileModel.File) {
	service.app.DB.Create(model)
}

func (service *FileService) Save(model *fileModel.File) (rowsAffected int64) {
	result := service.app.DB.Save(model)
	return result.RowsAffected
}

func (service *FileService) Delete(model *fileModel.File) (rowsAffected int64) {
	result := service.app.DB.Delete(model)
	return result.RowsAffected
}

func (service *FileService) Get(idstr string) (model fileModel.File) {
	service.app.DB.Where("id", idstr).Preload("Roles").Preload("AvatarFile").First(&model)
	return
}

func (service *FileService) All() (models []fileModel.File) {
	service.app.DB.Find(&models)
	return
}

// Paginate 分页内容
func (service *FileService) Paginate(c *gin.Context, perPage int) (users []fileModel.File, paging paginator.Paging) {
	db := service.app.DB.Model(fileModel.File{})
	if c.Query("storage") != "" {
		db = db.Where("storage = ?", c.Query("storage"))
	}

	paging = paginator.Paginate(
		c,
		db,
		&users,
		global.Config.VADMINURL(model.TableName(&fileModel.File{})),
		perPage,
	)
	return
}
