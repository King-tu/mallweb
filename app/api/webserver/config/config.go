package config

//mysql 服务
const (
	//MySQL_ADDR  = "192.168.5.130"
	MySQL_ADDR   = "127.0.0.1"
	MySQL_PORT   = "3306"
	MySQL_USER   = "root"
	MySQL_PASSWD = "123456"
	MySQL_DB     = "IHomeDB"

	//最大连接数
	MySQL_MaxOpenConns = 20
	//最大空闲数
	MySQL_MaxIdleConns = 10
	//最大生存时间
	MySQL_ConnMaxLifetime = 100
)

//redis 服务
const (
	Redis_ADDR = "192.168.5.130"
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

//consul 服务(注册中心)
const (
	Consul_ADDR = "192.168.5.130"
	Consul_PORT = "8500"
)

//获取区域信息 服务
const (
	//服务名
	AreaSrv_Name = "go.micro.srv.areaSrv"
	AreaSrv_ADDR = "192.168.5.130"
	AreaSrv_PORT = "51234"
)

//获取图片验证码 服务
const (
	ImageCodeSrv_Name = "go.micro.srv.imageCodeSrv"
	ImageCodeSrv_ADDR = "192.168.5.130"
	ImageCodeSrv_PORT = "51235"
)

//短信验证码 服务
const (
	SmsCodeSrv_Name = "go.micro.srv.smsCodeSrv"
	SmsCodeSrv_ADDR = "192.168.5.130"
	SmsCodeSrv_PORT = "51236"
)

//发送短信服务
const (
	RegisterSrv_Name = "go.micro.srv.registerSrv"
	RegisterSrv_ADDR = "192.168.5.130"
	RegisterSrv_PORT = "51237"
)

//登录 服务
const (
	LoginSrv_Name = "go.micro.srv.loginSrv"
	LoginSrv_ADDR = "192.168.5.130"
	LoginSrv_PORT = "51238"
)

//获取用户信息
const (
	UserInfoSrv_Name = "go.micro.srv.userInfoSrv"
	UserInfoSrv_ADDR = "192.168.5.130"
	UserInfoSrv_PORT = "51239"
)

//nginx ip Port
const (
	Nginx_IP = "192.168.5.130"
	Nginx_PORT = "8888"

	Nginx_ADDR = Nginx_IP +":"+ Nginx_PORT
)

//头像
const (
	AvatarSrv_Name = "go.micro.srv.avatarSrv"
	AvatarSrv_IP = "192.168.5.130"
	AvatarSrv_PORT = "51240"

	AvatarSrv_ADDR = AvatarSrv_IP +":"+ AvatarSrv_PORT
)

//更新用户名
const (
	UpdateUserNameSrv_Name = "go.micro.srv.updateUserNameSrv"
	UpdateUserNameSrv_IP = "192.168.5.130"
	UpdateUserNameSrv_PORT = "51241"

	UpdateUserNameSrv_ADDR = UpdateUserNameSrv_IP +":"+ UpdateUserNameSrv_PORT
)

//发送用户实名认证信息服务
const (
	UpdateUserAuthSrv_Name = "go.micro.srv.updateUserAuthSrv"
	UpdateUserAuthSrv_IP = "192.168.5.130"
	UpdateUserAuthSrv_PORT = "51242"

	UpdateUserAuthSrv_ADDR = UpdateUserAuthSrv_IP +":"+ UpdateUserAuthSrv_PORT
)


//发送（发布）房源信息服务
const (
	PostHousesSrv_Name = "go.micro.srv.postHousesSrv"
	PostHousesSrv_IP = "192.168.5.130"
	PostHousesSrv_PORT = "51243"

	PostHousesSrv_ADDR = PostHousesSrv_IP +":"+ PostHousesSrv_PORT
)

//获取用户已发布房源信息服务
const (
	GetUserHousesSrv_Name = "go.micro.srv.getUserHousesSrv"
	GetUserHousesSrv_IP = "192.168.5.130"
	GetUserHousesSrv_PORT = "51244"

	GetUserHousesSrv_ADDR = PostHousesSrv_IP +":"+ PostHousesSrv_PORT
)

//发送（上传）房屋图片服务
const (
	PostHousesImageSrv_Name = "go.micro.srv.postHousesImageSrv"
	PostHousesImageSrv_IP = "192.168.5.130"
	PostHousesImageSrv_PORT = "51245"

	PostHousesImageSrv_ADDR = PostHousesImageSrv_IP +":"+ PostHousesImageSrv_PORT
)

//获取房屋详细信息服务
const (
	GetHouseInfoSrv_Name = "go.micro.srv.getHouseInfoSrv"
	GetHouseInfoSrv_IP = "192.168.5.130"
	GetHouseInfoSrv_PORT = "51246"

	GetHouseInfoSrv_ADDR = GetHouseInfoSrv_IP +":"+ GetHouseInfoSrv_PORT
)

//获取（搜索）房源服务
const (
	SearchHouseSrv_Name = "go.micro.srv.searchHouseSrv"
	SearchHouseSrv_IP = "192.168.5.130"
	SearchHouseSrv_PORT = "51247"

	SearchHouseSrv_ADDR = SearchHouseSrv_IP +":"+ SearchHouseSrv_PORT
)

//发送（发布）订单服务
const (
	PostOrdersSrv_Name = "go.micro.srv.postOrdersSrv"
	PostOrdersSrv_IP = "192.168.5.130"
	PostOrdersSrv_PORT = "51248"

	PostOrdersSrv_ADDR = PostOrdersSrv_IP +":"+ PostOrdersSrv_PORT
)