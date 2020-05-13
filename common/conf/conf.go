package conf

import "time"

var (
	// 服务版本号
	Version = "0.1.0"
	// 日志
	LogFileName = "/tmp/log/mallweb/mallweb.log"
)

//设置了30秒的TTL生存期，并设置了每15秒一次的重注册。
const (
	//TTL指定一次注册在注册中心的有效期，过期后便删除
	RegisterTTLSec = time.Second * 20
	//间隔时间注册则是定时向注册中心重新注册以保证服务仍在线
	RegisterIntervalSec = time.Second * 15
)

// etcd 服务(注册中心)
const (
	ETCD_ADDR = "192.168.3.233:2379"
)

// mysql 服务
const (
	MySQL_ADDR   = "192.168.3.233:3306"
	MySQL_USER   = "root"
	MySQL_PASSWD = "root"
	MySQL_DBNAME = "mallweb"

	//最大连接数
	MySQL_MaxOpenConns = 30
	//最大空闲数
	MySQL_MaxIdleConns = 10
	//最大生存时间
	MySQL_ConnMaxLifetime = 100
)

// redis 服务
const (
	Redis_ADDR = "192.168.3.233:6379"

	//最大连接数
	Redis_MaxActive = 20
	//最大空闲数
	Redis_MaxIdle = 20
	//空闲超时时间(超过该时间后将关闭连接,单位 s)
	Redis_IdleTimeout = 100
)
