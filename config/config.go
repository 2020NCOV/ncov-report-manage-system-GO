package config


import (
	"html/template"
)

type Database struct {
	Host         string  // 地址
	Port         string  // 端口
	User         string  // 用户名
	Pwd          string  // 密码
	Name         string  // 数据库名
	MaxIdleCon   int     // 最大闲置连接数
	MaxOpenCon   int     // 最大打开连接数
	Driver       string  // 驱动名
	File         string  // 文件名
}

// 数据库配置
// 为一个map，其中key为数据库连接的名字，value为对应的数据配置
// 注意：key为default的数据库是默认数据库，也是框架所用的数据库，而你可以
// 配置多个数据库，提供给你的业务表使用，实现对不同数据库的管理。
type DatabaseList map[string]Database

// 存储目录：存储头像等上传文件
type Store struct {
	Path   string
	Prefix string
}

// FileUploadEngine is a file upload engine.
type FileUploadEngine struct {
	Name   string
	Config map[string]interface{}
}

type Config struct {
	// 数据库配置
	Databases DatabaseList `json:"database"`

	// 登录域名
	Domain string `json:"domain"`

	// 网站语言
	Language string `json:"language"`

	// 全局的管理前缀
	UrlPrefix string `json:"prefix"`

	// 主题名
	Theme string `json:"theme"`

	// 上传文件存储的位置
	Store Store `json:"store"`

	// 网站的标题
	Title string `json:"title"`

	// 侧边栏上的Logo
	Logo template.HTML `json:"logo"`

	// 侧边栏上的Logo缩小版
	MiniLogo template.HTML `json:"mini_logo"`

	// 登录后跳转的路由
	IndexUrl string `json:"index"`

	// 是否开始debug模式
	Debug bool `json:"debug"`

	// Info日志路径
	InfoLogPath string `json:"info_log"`

	// Error log日志路径
	ErrorLogPath string `json:"error_log"`

	// Access log日志路径
	AccessLogPath string `json:"access_log"`

	// 是否开始数据库Sql操作日志
	SqlLog bool `json:"sql_log"`

	// 是否关闭access日志
	AccessLogOff bool `json:"access_log_off"`
	// 是否关闭info日志
	InfoLogOff   bool `json:"info_log_off"`
	// 是否关闭error日志
	ErrorLogOff  bool `json:"error_log_off"`

	// 网站颜色主题
	ColorScheme string `json:"color_scheme"`

	// Session的有效时间，单位为秒
	SessionLifeTime int `json:"session_life_time"`

	// Cdn链接，为全局js/css配置cdn链接
	AssetUrl string `json:"asset_url"`

	// 文件上传引擎
	FileUploadEngine FileUploadEngine `json:"file_upload_engine"`

	// 自定义头部js/css
	CustomHeadHtml template.HTML `json:"custom_head_html"`

	// 自定义尾部js/css
	CustomFootHtml template.HTML `json:"custom_foot_html"`

	// 登录页面标题
	LoginTitle string `json:"login_title"`

	// 登录页面logo
	LoginLogo template.HTML `json:"login_logo"`
}