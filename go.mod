module github.com/king-tu/mallweb

go 1.14

require (
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190808125512-07798873deee
	github.com/coreos/etcd v3.3.21+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/zap v0.0.1
	github.com/gin-gonic/gin v1.6.2
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.4 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/lucas-clemente/quic-go v0.15.7 // indirect
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/go-plugins v1.5.1 // indirect
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.5.0
	github.com/micro/micro v1.18.0 // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nicksnyder/go-i18n v1.10.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	google.golang.org/genproto v0.0.0-20200519141106-08726f379972
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
