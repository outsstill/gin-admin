package service

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outsstill/gin-admin/model"
	fileModel "github.com/outsstill/gin-admin/model/file"
	"github.com/outsstill/gin-admin/pkg/auth"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/pkg/paginator"
	gokit "github.com/outsstill/go-kit"
	"github.com/outsstill/go-kit/storage"
	"github.com/spf13/cast"
)

type FileService struct {
	storage storage.IStorage
}

func NewFileService(drive ...string) *FileService {
	storageConfig := gokit.Config().Storage.ToStorage()

	if len(drive) > 0 {
		storageConfig.Driver = drive[0]
	}

	storageDriver, _ := storage.New(storageConfig)

	return &FileService{
		storage: storageDriver.Driver(),
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
	if header.Size > cast.ToInt64(gokit.Config().Storage.SizeLimit) {
		return nil, errors.New("超过最大文件大小")
	}

	// 验证后缀
	extLimit := cast.ToStringSlice(gokit.Config().Storage.Ext)
	if helpers.FindElement(extLimit, strings.ToLower(helpers.GetFileExt(header.Filename))) < 0 {
		return nil, errors.New("文件格式不允许 只允许[ " + strings.Join(extLimit, " ") + " ]")
	}

	input := &storage.UploadRequest{
		Filename:    header.Filename,
		Path:        header.Filename,
		Size:        header.Size,
		ContentType: header.Header.Get("Content-Type"),
		Reader:      fileObj,
		Meta:        map[string]string{},
	}

	obj, err := service.storage.Put(c, input)

	if err != nil {
		return nil, err
	}

	// 存入数据库
	fileStore := &fileModel.File{}
	fileStore.Bucket = obj.Bucket
	fileStore.Name = obj.StoredName
	fileStore.OriginName = obj.OriginName
	fileStore.Path = obj.Path
	fileStore.Key = obj.Key
	fileStore.Size = obj.Size
	fileStore.Ext = obj.Ext
	fileStore.Storage = obj.Driver
	fileStore.ETag = obj.ETag
	fileStore.ContentType = obj.ContentType
	fileStore.LastModified = obj.LastModified
	fileStore.UserId = cast.ToUint64(auth.CurrentAdminUID(c))
	fileStore.GroupId = cast.ToInt(c.DefaultPostForm("group_id", "99"))
	fileStore.Type = cast.ToInt(c.DefaultPostForm("type", "1"))

	// 组装url
	fileStore.FullUrl = obj.URL
	return fileStore, nil
}

func (service *FileService) DeleteFile(id string) error {
	fileObj := service.Get(id)
	if fileObj.ID <= 0 {
		return errors.New("没有找到该条记录")
	}
	// 上传
	err := service.storage.Delete(context.Background(), fileObj.Key)
	if err != nil {
		return err
	}

	service.Delete(fileObj)

	return nil
}

func (service *FileService) Create(model *fileModel.File) {
	gokit.DB().Create(model)
}

func (service *FileService) Save(model *fileModel.File) (rowsAffected int64) {
	result := gokit.DB().Save(model)
	return result.RowsAffected
}

func (service *FileService) Delete(model *fileModel.File) (rowsAffected int64) {
	result := gokit.DB().Delete(model)
	return result.RowsAffected
}

func (service *FileService) Get(idstr string) (model *fileModel.File) {
	gokit.DB().Where("id", idstr).First(&model)
	return
}

func (service *FileService) GetByKey(key string) (model *fileModel.File) {
	gokit.DB().Where("key = ?", key).First(&model)
	return
}

func (service *FileService) All() (models []fileModel.File) {
	gokit.DB().Find(&models)
	return
}

// Paginate 分页内容
func (service *FileService) Paginate(c *gin.Context, perPage int) (users []fileModel.File, paging paginator.Paging) {
	db := gokit.DB().Model(fileModel.File{})
	if c.Query("storage") != "" {
		db = db.Where("storage = ?", c.Query("storage"))
	}

	paging = paginator.Paginate(
		c,
		db,
		&users,
		helpers.VADMINURL(model.TableName(&fileModel.File{})),
		perPage,
	)
	return
}
