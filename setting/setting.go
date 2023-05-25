package setting

//
//import (
//	"github.com/go-ini/ini"
//	"log"
//	"time"
//)
//
//type App struct {
//	JwtSecret string
//	PageSize  int
//}
//
//var AppSetting = &App{}
//
//type Server struct {
//	RunMode          string
//	HttpPort         int
//	ReadTimeout      time.Duration
//	WriteTimeout     time.Duration
//	HighestAuthority string
//	HighestKey       string
//	GpuUrl           string
//}
//
//var ServerSetting = &Server{}
//
//type RPC struct {
//	Addr string
//}
//
//var RPCSetting = &RPC{}
//
//type Database struct {
//	Type     string
//	User     string
//	Password string
//	Host     string
//	Name     string
//	MaxIdle  int
//	MaxConn  int
//}
//
//var DatabaseSetting = &Database{}
//
//type Redis struct {
//	Host        string
//	Password    string
//	MaxIdle     int
//	MaxActive   int
//	IdleTimeout time.Duration
//}
//
//var RedisSetting = &Redis{}
//
//type WeiXin struct {
//	CertPath      string // 证书路径
//	AppId         string //商家后台管理小程序
//	Secret        string //商家后台管理小程序
//	BuyAppId      string //买家小程序
//	BuySecret     string //买家小程序
//	MchId         string
//	ApiKey        string
//	ServiceAppId  string //公众号appid
//	ServiceSecret string //公众号service
//	WebAppId      string //网站扫码登录
//	WebSecret     string //网站扫码登录
//}
//
//var WeiXinSetting = &WeiXin{}
//
//type Baidu struct {
//	Username string
//	Password string
//	Token    string
//	SiteId   string
//}
//
//var BaiduSetting = &Baidu{}
//
//type Harbor struct {
//	Url      string
//	Username string
//	Password string
//	Path     string
//}
//
//var HarborSetting = &Harbor{}
//
//type FastDfs struct {
//	Url          string
//	GoogleSecret string
//}
//
//var FastDfsSetting = &FastDfs{}
//
//type DefaultInfo struct {
//	Name      string
//	AvatarUrl string
//}
//
//var DefaultInfoSetting = &DefaultInfo{}
//
//var Cfg *ini.File
//
//// Setup initialize the configuration instance
//func Setup() {
//	var err error
//	Cfg, err = ini.Load("config/env.ini")
//	if err != nil {
//		log.Fatalf("setting.Setup, fail to parse 'config/env.ini': %v", err)
//	}
//
//	//mapTo("app", AppSetting)
//	mapTo("server", ServerSetting)
//	mapTo("rpc", RPCSetting)
//	mapTo("database", DatabaseSetting)
//	mapTo("redis", RedisSetting)
//	mapTo("weixin", WeiXinSetting)
//	mapTo("harbor", HarborSetting)
//	mapTo("fastdfs", FastDfsSetting)
//	mapTo("baidutongji", BaiduSetting)
//	mapTo("defaultinfo", DefaultInfoSetting)
//
//	HarborSetting.Path = "dk"
//	if Cfg.Section("global").Key("Test").MustInt(0) == 1 {
//		HarborSetting.Path = "test"
//	}
//	//
//	//AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
//	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
//	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
//	//RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
//}
//
//// mapTo map section
//func mapTo(section string, v interface{}) {
//	err := Cfg.Section(section).MapTo(v)
//	if err != nil {
//		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
//	}
//}
