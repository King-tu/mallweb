package config

type App struct {
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	ETCD      ETCD      `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	Elastic   Elastic   `mapstructure:"elastic" json:"elastic" yaml:"elastic"`
	Jaeger    Jaeger    `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	SmsConfig SmsConfig `mapstructure:"sms_config" json:"sms_config" yaml:"sms_config"`
}

type System struct {
	Version     string `mapstructure:"version" json:"version" yaml:"version"`
	ApiAddr     string `mapstructure:"api_addr" json:"api_addr" yaml:"api_addr"`
	ProjectName string `mapstructure:"project_name" json:"project_name" yaml:"project_name"`
}

type Mysql struct {
	Username        string `mapstructure:"username" json:"username" yaml:"username"`
	Password        string `mapstructure:"password" json:"password" yaml:"password"`
	Addr            string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Dbname          string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Config          string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxConnLifetime int    `mapstructure:"max_conn_lifetime" json:"max_conn_lifetime" yaml:"max_conn_lifetime"`
}

type Redis struct {
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	DB          int    `mapstructure:"db" json:"db" yaml:"db"`
	MaxActive   int    `mapstructure:"max_active" json:"max_active" yaml:"max_active"`
	MaxIdle     int    `mapstructure:"max_idle" json:"max_idle" yaml:"max_idle"`
	IdleTimeout int    `mapstructure:"idle_timeout" json:"idle_timeout" yaml:"idle_timeout"`
}

type ETCD struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	// 指定一次注册在注册中心的有效期，过期后便删除
	RegisterTTL int `mapstructure:"register_ttl" json:"register_ttl" yaml:"register_ttl"`
	// 定时向注册中心重新注册以保证服务仍在线
	RegisterInterval int `mapstructure:"register_interval" json:"register_interval" yaml:"register_interval"`
}

// es
type Elastic struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	UserName string `mapstructure:"user_name" json:"user_name" yaml:"user_name"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type Jaeger struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type Log struct {
	FileName string `mapstructure:"file_name" json:"file_name" yaml:"file_name"`
	Level    string `mapstructure:"level" json:"level" yaml:"level"`
	Format   string `mapstructure:"format" json:"format" yaml:"format"`
	// 日志归档配置项
	// 每个日志文件保存的最大尺寸 单位：M
	MaxSize int `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	// 文件最多保存多少天
	MaxBackups int `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	// 日志文件最多保存多少个备份
	MaxAge int `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	// 是否压缩
	Compress bool `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type JWT struct {
	Secrect     string `mapstructure:"secrect" json:"secrect" yaml:"secrect"`
	ExpiresTime int    `mapstructure:"expires_ime" json:"expires_ime" yaml:"expires_ime"`
}

type SmsConfig struct {
	RegionId        string `mapstructure:"region_id" json:"region_id" yaml:"region_id"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secrect" json:"access_key_secrect" yaml:"access_key_secrect"`
	SignName        string `mapstructure:"sign_name" json:"sign_name" yaml:"sign_name"`
	TemplateCode    string `mapstructure:"template_code" json:"template_code" yaml:"template_code"`
}
