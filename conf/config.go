package conf

import (
	"github.com/tvb-sz/serve-swagger-ui/define"
	"github.com/tvb-sz/serve-swagger-ui/utils/cfg"
)

// 项目config配置定义

// 暴露给全局使用的配置变量
var (
	Config config
	Cmd    cmdConfig
)

// config 项目配置上层结构
type config struct {
	Server       server  `json:"server"`  // server config
	Log          log     `json:"log"`     // log config
	Google       google  `json:"google"`  // google oauth config
	Swagger      swagger `json:"swagger"` // swagger json file config
	ConfigFile   string  `json:"-"`       // record config file path
	EnableGoogle bool    `json:"-"`       // record is set google client_id & client_secret
}

// cmdConfig command line args
type cmdConfig struct {
	ConfigFile         string // set config file path
	Host               string // set web server host ip
	Port               int    // set web server ports
	SwaggerPath        string // set swagger file path
	GoogleClientID     string // set google oauth app-key
	GoogleClientSecret string // set google oauth app-secret
	LogLevel           string // set logger level
	LogPath            string // set logger path
}

// parseAfterLoad 配置项加载完成后的统一处理流程逻辑
func (c config) parseAfterLoad() {

}

// region 初始化

// Init 初始化
func Init() {
	// ① set framework version
	Config.Server.Version = define.Version

	// read perhaps config
	if Cmd.ConfigFile != "" {
		var cfgLoader cfg.IFace
		cfgLoader = cfg.Viper{}
		_ = cfgLoader.Parse(Cmd.ConfigFile, "toml", &Config)
	}

	// ② command line args first
	if Cmd.Host != "" {
		Config.Server.Host = Cmd.Host
	}
	if Cmd.Port != 0 {
		Config.Server.Port = Cmd.Port
	}
	if Cmd.SwaggerPath != "" {
		Config.Swagger.Path = Cmd.SwaggerPath
	}
	if Cmd.GoogleClientID != "" {
		Config.Google.ClientID = Cmd.GoogleClientID
	}
	if Cmd.GoogleClientSecret != "" {
		Config.Google.ClientSecret = Cmd.GoogleClientSecret
	}
	if Cmd.GoogleClientSecret != "" {
		Config.Google.ClientSecret = Cmd.GoogleClientSecret
	}
	if Cmd.LogPath != "" {
		Config.Log.Path = Cmd.LogPath
	}
	if Cmd.LogLevel != "" {
		Config.Log.Level = Cmd.LogLevel
	}

	// ③ set default value
	if Config.Server.Host == "" {
		Config.Server.Host = define.DefaultHost
	}
	if Config.Server.Port == 0 {
		Config.Server.Port = define.DefaultPort
	}
	if Config.Log.Path == "" {
		Config.Log.Path = define.DefaultLogPath
	}
	if Config.Log.Level == "" {
		Config.Log.Level = define.DefaultLogLevel
	}

	// 配置加载并解析映射成功后统一处理逻辑：譬如Url统一处理后缀斜杠
	Config.parseAfterLoad()
}

// endregion

// region 热重载

// Reload 热重载
// 当配置变更时又无需重新启动进程触发，监听调用
func Reload() {

}

// endregion
