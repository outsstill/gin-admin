package controllers

import (
	"errors"
	"slices"
	"strings"
	"time"

	fileModel "github.com/outsstill/gin-admin/model/file"
	"github.com/outsstill/gin-admin/pkg/auth"
	"github.com/outsstill/gin-admin/pkg/helpers"
	"github.com/outsstill/gin-admin/pkg/response"
	"github.com/outsstill/gin-admin/requests"
	"github.com/outsstill/gin-admin/services"

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

	data, pager := service.NewFileService(uc.App).Paginate(c, uc.GetPerPage(c))

	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (uc *AdminFileController) All(c *gin.Context) {

	menus := service.NewFileService(uc.App).All()

	response.Data(c, menus)
}

func (uc *AdminFileController) Get(c *gin.Context) {

	user := service.NewFileService(uc.App).Get(c.Param("id"))

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
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminFileStore); !ok {
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
	service.NewFileService(uc.App).Create(u)

	response.Data(c, u)
}

func (uc *AdminFileController) Upload(c *gin.Context) {

	uploadStorage := c.PostForm("uploadStorage")

	if err := uc.CheckUploadStorage(uploadStorage); err != nil {
		response.Fail(c, err.Error())
		return
	}

	obj, err := service.NewFileService(uc.App, uploadStorage).UploadFile(c)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	if obj != nil {
		service.NewFileService(uc.App).Create(obj)
	}

	response.Data(c, obj)
}

func (uc *AdminFileController) Update(c *gin.Context) {
	model := service.NewFileService(uc.App).Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}

	// 验证
	request := requests.AdminFileUpdateRequest{}
	request.ID = model.ID
	if ok := requests.ValidateFunc(c, uc.App, &request, requests.VerityAdminFileUpdate); !ok {
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
	service.NewFileService(uc.App).Save(model)

	response.Data(c, model)
}

func (uc *AdminFileController) Delete(c *gin.Context) {

	model := service.NewFileService(uc.App).Get(c.Param("id"))
	if model.ID <= 0 {
		response.Fail(c, "没有找到")
		return
	}
	err := service.NewFileService(uc.App, model.Storage).DeleteFile(cast.ToString(model.ID))

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c)
}
