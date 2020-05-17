module github.com/king-tu/mallweb

go 1.14

require (
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190808125512-07798873deee
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/zap v0.0.1
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo/redis v0.0.0-20200429221454-e14091dffc1b
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.6.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nicksnyder/go-i18n v1.10.1
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.14.1
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	golang.org/x/net v0.0.0-20200505041828-1ed23360d12c
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/protobuf v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
