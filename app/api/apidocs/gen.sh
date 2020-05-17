
# 根据proto文件生成go源码
#protoc --micro_out=. --go_out=. *.proto

# 生成API文档的swagger.json文件
protoc -I /usr/local/include -I . -I ./google/api --swagger_out=allow_merge=true,merge_file_name=apidocs:. --swagger_out=logtostderr=true:. *.proto

docker stop swaggerui

docker run --name swaggerui --rm -d -p 9090:8080 -e SWAGGER_JSON=/foo/apidocs.swagger.json -v $(pwd):/foo swaggerapi/swagger-ui


