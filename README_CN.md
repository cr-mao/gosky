##   gosky  框架

 快速、简单、对phper友好

####  安装、使用
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
go run main.go sql struct --table=`table_name` --port=3308

# 编译成linux 二进制
CGO_ENABLED=0 GOOS=linux  go build -o gosky main.go
# 二进制运行 
gosky serve --env=production 
```


####  docker启动项目要的demo 数据库和表结构
也可以自行导入 [demo.sql](demo.sql)
```shell 
docker build -t mysql-for-go:v1 -f mysqlDockerfile .
docker run -itd --name mysqlforgodemo -e MYSQL_ROOT_PASSWORD=123456 -p 3308:3306 mysql-for-go:v1
# mysql -h 127.0.0.1 -u root -P 3308 -p 
```


#### 测试

apitests 目录下写http 接口测试，已做好初始化工作
```shell 
 go test -v ./...
```


#### framework struct

```text
├── apitests     http接口测试
├── app
│   ├── http  (http serve)
│   │   ├── controllers  
│   │   │     └── v1       
│   │   │        └── user_controller.go      
│   │   ├── middlewares  
│   │   │   └── header.go  
│   │   │   ├── logger.go   
│   │   │   └── recovery.go 
│   │   ├── routers 
│   │   │   ├── route_api.go 
│   │   │   └── router.go
│   │   └── serve.go
│   └── requests  （requests valid)
│   └── services   
│   │   └── user
│   │   │   ├── user_service.go 
│   └── models 
│   │   └── user_model
│   │   │   ├── user_model.go 
├── bootstrap (框架启动时初始化数据库，日志，redis 等）
├── build   (构建相关）
├── cmd    （命令行处理）
├── config  (配置函数，项目起来会注册到全局的viper对象中)
├── docs    
├── infra   
│   ├── app      
│   ├── conf     (config)
│   ├── console 
│   ├── db       (mysql)
│   ├── errcode  (response code)
│   ├── helpers  
│   ├── job     （job daemon or crontab）
│   ├── logger   
│   ├── redis    
│   └── response
├── local.config.yaml  
├── testing.config.yaml 
├── production.config.yaml 
├── storage            (log file directory)
├── main.go            

```

#### 接口开发文档

[docs](docs/index.md)