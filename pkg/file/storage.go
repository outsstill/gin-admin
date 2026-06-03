package file

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type PutObjectInput struct {
	Key         string
	Size        int64
	ContentType string
	Reader      io.Reader
	File        *multipart.FileHeader
	// 供直传/分片扩展使用
	Meta map[string]string
}

type ObjectInfo struct {
	Bucket       string
	Key          string
	OriginName   string
	Name         string
	Size         int64
	Storage      string
	ContentType  string
	Ext          string
	Path         string
	ETag         string
	URL          string // 公开或带签名的可访问 URL（可选）
	LastModified time.Time
}

type PresignParams struct {
	Key         string
	Method      string // PUT/GET
	Expiry      time.Duration
	ContentType string
}

type IStorage interface {
	Put(ctx context.Context, in PutObjectInput) (ObjectInfo, error)
	//Get(ctx context.Context, key string) (io.ReadCloser, ObjectInfo, error)
	Delete(ctx context.Context, full string) error
	BackendName() string
}

type Config struct {
	Driver string
	LocalConfig
	OssConfig
}

func NewStorage(cfg Config) IStorage {
	switch cfg.Driver {
	case "local":
		return NewLocalStorage(cfg.LocalConfig)
	case "oss":
		return NewOssStorage(cfg.OssConfig)
	case "other":
		return NewOtherStorage()
	default:
		panic("unsupported storage type")
	}
}

func safeJoin(base, p string) (string, error) {
	clean := filepath.Clean("/" + p)
	// 防目录穿越
	if strings.Contains(clean, "..") {
		return "", fmt.Errorf("invalid path")
	}
	return filepath.Join(base, clean), nil
}
