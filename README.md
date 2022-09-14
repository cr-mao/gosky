##   gosky  framework

fast, easy , friendly for phper. 
[中文文档](README_CN.md)

####  install and use 
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


####  docker init  demo table  
```shell 
docker build -t mysql-for-go:v1 -f mysqlDockerfile .
docker run -itd --name mysqlforgodemo -e MYSQL_ROOT_PASSWORD=123456 -p 3308:3306 mysql-for-go:v1
# mysql -h 127.0.0.1 -u root -P 3308 -p 
```


#### test 
```shell 
 go test -v ./...
```


#### framework struct 


```text
├── apitests     (api test)
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
│   └── models 
│   │   └── user_model
│   │   │   ├── user_model.go 
│   └── requests  （requests valid)
│   └── services   
│   │   └── user
│   │   │   ├── user_service.go 
├── bootstrap 
├── build  
├── cmd    
├── config  ( config function ,initial,will register to global viper object)
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

#### develop doc 

[docs](docs/index.md)