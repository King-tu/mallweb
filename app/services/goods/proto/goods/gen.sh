
# 根据proto文件生成go源码
protoc --micro_out=. --go_out=. *.proto

# 生成API文档的swagger.json文件
protoc -I /usr/local/include -I . -I ./google/api --swagger_out=logtostderr=true:. *.proto

# cp ./*.proto ../../../../api/apidocs/

docker stop swaggerui
docker run --name swaggerui --rm -d -p 9090:8080 -e SWAGGER_JSON=/foo/*.swagger.json -v $(pwd):/foo swaggerapi/swagger-ui

#if [ $? -eq 0 ]; then
#    echo "succeed"
#else
#    echo "failed"
#fi

