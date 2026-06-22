package setting

import (
	"errors"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Setting struct {
	App     AppConfig     `mapstructure:"app" yaml:"app"`
	Storage StorageConfig `mapstructure:"storage" yaml:"storage"`
	JWT     JWTConfig     `mapstructure:"jwt" yaml:"jwt"`
	Captcha CaptchaConfig `mapstructure:"captcha" yaml:"captcha"`
	Paging  PagingConfig  `mapstructure:"paging" yaml:"paging"`
	Limit   LimitConfig   `mapstructure:"limit" yaml:"limit"`
}

var cfg *Setting

func Load(path string) (*Setting, error) {

	if len(path) == 0 {
		return nil, errors.New("config file path is empty")
	}

	s := &Setting{}

	v := viper.New()

	if len(path) > 0 {
		v.SetConfigType("yaml") // 类型
		v.AddConfigPath(".")    // 当前目录
		v.SetConfigFile(path)

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	if err := v.Unmarshal(s); err != nil {
		return nil, err
	}

	cfg = s
	return s, nil
}

func Get() *Setting {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg
}

func App() AppConfig {
	return Get().App
}

func Storage() StorageConfig {
	return Get().Storage
}

func JWT() JWTConfig {
	return Get().JWT
}

func Captcha() CaptchaConfig {
	return Get().Captcha
}

func Paging() PagingConfig {
	return Get().Paging
}

func Limit() LimitConfig {
	return Get().Limit
}

func IsLocal() bool {
	return App().Env == "local"
}

func IsProduction() bool {
	return App().Env == "production"
}

func IsTesting() bool {
	return App().Env == "testing"
}

func IsDebug() bool {
	return App().Debug == true
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(App().Timezone)
	return time.Now().In(chinaTimezone)
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return App().Url + "/" + path
}

// VADMINURL 拼接带 admin 标示 URL
func VADMINURL(path string) string {
	return URL("/admin/" + path)
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}

func GetFileStoragePath() string {
	formatted := time.Now().Format("20060102")

	return App().Name + "/" + formatted
}

// 获取文件存储名称(包含完整路径)
func GetFileStorageFullPath(fileName string, isOriginName bool) (string, string) {
	originFileName := fileName
	if !isOriginName {
		fileOriExt := filepath.Ext(fileName) // 获取文件扩展名 这里包含了 .
		//randomNumber := app.GetRandomNumber(16)
		randomNumber := uuid.New().String()
		// fileNameNoExt := fileName[:len(fileName)-len(fileOriExt)] // 文件名称 不含 .和后缀
		originFileName = cast.ToString(randomNumber) + fileOriExt
	}

	objectName := GetFileStoragePath() + "/" + originFileName

	return objectName, originFileName
}

type AppConfig struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Key      string `mapstructure:"key" yaml:"key"`
	Url      string `mapstructure:"url" yaml:"url"`
	HttpPort string `mapstructure:"http_port" yaml:"http_port"`
	FileUrl  string `mapstructure:"file_url" yaml:"file_url"`
	Env      string `mapstructure:"env" yaml:"env"`
	Version  string `mapstructure:"version" yaml:"version"`
	Debug    bool   `mapstructure:"debug" yaml:"debug"`
	Timezone string `mapstructure:"timezone" yaml:"timezone"`
}

type StorageConfig struct {
	Driver    string              `mapstructure:"driver" yaml:"driver"`
	SizeLimit int64               `mapstructure:"size_limit" yaml:"size_limit"`
	Ext       []string            `mapstructure:"ext" yaml:"ext"`
	Local     *LocalStorageConfig `mapstructure:"local" yaml:"local"`
	Oss       *OssStorageConfig   `mapstructure:"oss" yaml:"oss"`
}

type LocalStorageConfig struct {
	Path         string `mapstructure:"path" yaml:"path"`
	Domain       string `mapstructure:"domain" yaml:"domain"`
	StaticPrefix string `mapstructure:"static" yaml:"static"`
}

type OssStorageConfig struct {
	KeyId     string `mapstructure:"key_id" yaml:"key_id"`
	KeySecret string `mapstructure:"key_secret" yaml:"key_secret"`
	Region    string `mapstructure:"region" yaml:"region"`
	Bucket    string `mapstructure:"bucket" yaml:"bucket"`
	Domain    string `mapstructure:"domain" yaml:"domain"`
}

type PagingConfig struct {
	PerPage         int    `mapstructure:"perpage" yaml:"perpage"`
	UrlQueryOrder   string `mapstructure:"url_query_order" yaml:"url_query_order"`
	UrlQuerySort    string `mapstructure:"url_query_sort" yaml:"url_query_sort"`
	UrlQueryPage    string `mapstructure:"url_query_page" yaml:"url_query_page"`
	UrlQueryPerPage string `mapstructure:"url_query_per_page" yaml:"url_query_per_page"`
}

type CaptchaConfig struct {
	Height          int     `mapstructure:"height" yaml:"height"`
	Width           int     `mapstructure:"width" yaml:"width"`
	Length          int     `mapstructure:"length" yaml:"length"`
	Maxskew         float64 `mapstructure:"maxskew" yaml:"maxskew"`
	Dotcount        int     `mapstructure:"dotcount" yaml:"dotcount"`
	ExpireTime      int     `mapstructure:"expire_time" yaml:"expire_time"`
	DebugExpireTime int     `mapstructure:"debug_expire_time" yaml:"debug_expire_time"`
	TestingKey      string  `mapstructure:"testing_key" yaml:"testing_key"`
}

type JWTConfig struct {
	ExpireTime     int `mapstructure:"expire_time" yaml:"expire_time"`           // 过期时间，单位是分钟，一般不超过两个小时
	MaxReFreshTime int `mapstructure:"max_refresh_time" yaml:"max_refresh_time"` // 允许刷新时间，单位分钟，从 Token 的签名时间算起
}

type LimitConfig struct {
	Rate        string `mapstructure:"rate" yaml:"rate"` // 过期时间，单位是分钟，一般不超过两个小时
	TestRate    string `mapstructure:"test_rate" yaml:"test_rate"`
	LoginRate   string `mapstructure:"login_rate" yaml:"login_rate"`
	CaptchaRate string `mapstructure:"captcha_rate" yaml:"captcha_rate"`
	StoreRate   string `mapstructure:"store_rate" yaml:"store_rate"`
	UpdateRate  string `mapstructure:"update_rate" yaml:"update_rate"`
	DeleteRate  string `mapstructure:"delete_rate" yaml:"delete_rate"`
	QueryRate   string `mapstructure:"query_rate" yaml:"query_rate"`
}
