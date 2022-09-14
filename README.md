##   my  go framework  base skeleton


####  how to use 
```shell
# 请下载 go 1.18.5 版本
export GO111MODULE=on
go mod download
go run main.go --help 
# 启动服务 示例，默认使用local环境
go run main.go serve --env=local 

# make model  
go run main.go make model `table_name`
# sql2struct  还需要手动改改即可
go run main.go sql struct --table=`table_name`


# 编译成linux 二进制
CGO_ENABLED=0 GOOS=linux  go build -o gosky main.go
# 二进制运行 
gosky serve --env=production 
```


####  docker初始化表结构
```shell 
docker build -t mysql-for-go:v1 -f mysqlDockerfile .
docker run -itd --name mysqlforgodemo -e MYSQL_ROOT_PASSWORD=123456 -p 3308:3306 mysql-for-go:v1
# mysql -h 127.0.0.1 -u root -P 3308 -p 
```


#### test 
http 接口测试 统一写在 apitests 下 
```shell 
 go test -v ./...
```


#### 文档

[docs](docs/index.md)