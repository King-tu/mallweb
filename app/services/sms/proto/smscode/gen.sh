
# 根据proto文件生成go源码
protoc --micro_out=. --go_out=. *.proto

# 生成API文档的swagger.json文件
protoc -I /usr/local/include -I . -I ./google/api --swagger_out=logtostderr=true:. *.proto

#if [ $? -eq 0 ]; then
#    docker run --rm -d -p 8888:8080 -e SWAGGER_JSON=/foo/hello.swagger.json -v $(pwd):/foo swaggerapi/swagger-ui
#fi
#
#if [ $? -eq 0 ]; then
#    echo "succeed"
#else
#    echo "failed"
#fi

