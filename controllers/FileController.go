package controllers

import (
	"errors"
	"slices"
	"strings"
	"time"

	fileModel "github.com/outsstill/gin-admin/model/file"
	"github.com/outsstill/gin-admin/pkg/auth"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/requests"
	"github.com/outsstill/gin-admin/services"
	"github.com/outsstill/go-kit/response"
	"github.com/outsstill/go-kit/storage"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type AdminFileController struct {
	*BaseAPIController
}

func NewAdminFileController(base *BaseAPIController) *AdminFileController {
	return &AdminFileController{
		BaseAPIController: base,
	}
}

func (uc *AdminFileController) Index(c *gin.Context) {

	data, pager := uc.App.GetFileService().Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminFileController) All(c *gin.Context) {

	menus := uc.App.GetFileService().All()

	response.Data(c, menus)
}

func (uc *AdminFileController) Get(c *gin.Context) {

	user := uc.App.GetFileService().Get(c.Param("id"))

	response.Data(c, user)
}

func (uc *AdminFileController) CheckStore(storage string, path string, url string) error {

	storage = strings.ToLower(storage)

	if !slices.Contains([]string{"oss", "local", "other"}, storage) {
		return errors.New("错误的存储引擎")
	}

	if slices.Contains([]string{"oss", "local"}, storage) && helpers.Empty(path) {
		return errors.New("必须填入 path")
	}

	if slices.Contains([]string{"other"}, storage) && helpers.Empty(url) {
		return errors.New("选择外链存储,必须填入 url")
	}

	return nil
}

func (uc *AdminFileController) CheckUploadStorage(storage string) error {
	storage = strings.ToLower(storage)

	if !slices.Contains([]string{"oss", "local"}, storage) {
		return errors.New("错误的存储引擎")
	}
	return nil
}

func (uc *AdminFileController) Store(c *gin.Context) {
	// 验证
	request := requests.AdminFileStoreRequest{}
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminFileStore); !ok {
		return
	}

	if err := uc.CheckStore(request.Storage, request.Path, request.Url); err != nil {
		response.Fail(c, err.Error())
		return
	}

	u := &fileModel.File{
		Name:         request.Name,
		OriginName:   request.OriginName,
		GroupId:      request.GroupId,
		Size:         request.Size,
		Ext:          request.Ext,
		Type:         request.Type,
		Storage:      request.Storage,
		Url:          request.Url,
		Path:         request.Path,
		LastModified: time.Now(),
		UserId:       cast.ToUint64(auth.CurrentAdminUID(c)),
	}
	uc.App.GetFileService().Create(u)

	response.Data(c, u)
}

func (uc *AdminFileController) Upload(c *gin.Context) {

	uploadStorage := c.PostForm("uploadStorage")

	if err := uc.CheckUploadStorage(uploadStorage); err != nil {
		response.Fail(c, err.Error())
		return
	}

	obj, err := service.NewFileService(uploadStorage).UploadFile(c)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	fileStore := &fileModel.File{}
	if obj != nil {
		// 存入数据库
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
		fileStore.UploadType = cast.ToInt(storage.UPLOAD_TYPE_CROSS_SERVER)
		uc.App.GetFileService().Create(fileStore)
	}

	response.Data(c, fileStore)
}

func (uc *AdminFileController) Update(c *gin.Context) {
	model := uc.App.GetFileService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminFileUpdateRequest{}
	request.ID = model.ID
	if ok := requests.ValidateFunc(c, &request, requests.VerityAdminFileUpdate); !ok {
		return
	}

	if err := uc.CheckStore(request.Storage, request.Path, request.Url); err != nil {
		response.Fail(c, err.Error())
		return
	}

	model.Bucket = request.Bucket
	model.Name = request.Name
	model.OriginName = request.OriginName
	model.Path = request.Path
	model.Key = request.Key
	model.Size = request.Size
	model.Ext = request.Ext
	model.Storage = request.Storage
	model.ETag = request.ETag
	model.ContentType = request.ContentType
	model.LastModified = time.Now()
	model.Url = request.Url
	model.GroupId = request.GroupId
	model.Type = request.Type
	model.UserId = cast.ToUint64(auth.CurrentAdminUID(c))

	uc.App.GetFileService().Save(model)

	response.Data(c, model)
}

func (uc *AdminFileController) Delete(c *gin.Context) {

	model := uc.App.GetFileService().Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}
	err := service.NewFileService(model.Storage).DeleteFile(cast.ToString(model.ID))

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c)
}
