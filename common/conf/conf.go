package conf

// mysql 服务
const (
	MySQL_IP     = "192.168.3.233"
	MySQL_PORT   = "3306"
	MySQL_ADDR   = MySQL_IP + ":" + MySQL_PORT
	MySQL_USER   = "root"
	MySQL_PASSWD = "root"
	MySQL_DBNAME = "mallweb"

	//最大连接数
	MySQL_MaxOpenConns = 20
	//最大空闲数
	MySQL_MaxIdleConns = 10
	//最大生存时间
	MySQL_ConnMaxLifetime = 10
)

// redis 服务
const (
	Redis_ADDR = "192.168.3.233"
	Redis_PORT = "6379"

	//最大连接数
	Redis_MaxActive = 20
	//最大空闲数
	Redis_MaxIdle = 10
	//空闲超时时间(超过该时间后将关闭连接,单位 s)
	Redis_IdleTimeout = 20
	//需要操作的数据库
	Redis_DB = 1
)

// etcd 服务(注册中心)
const (
	Consul_ADDR = "192.168.3.233"
	Consul_PORT = "2379"
)
