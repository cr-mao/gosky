## 如何使用api ([返回](index.md))


1. 所有请求使用`post`请求，body为json格式，减少因为请求方式错误，导致线上异常 (**可自行修改**)

2. 公共请求头
自行去扩展 

| 字段           | 类型     | 描述                | required             |
|--------------|--------|-------------------|----------------------|
| guid         | string | 用户唯一标志            | yes                  |




3. 签名算法
这里不做展示
demo 中已经去掉


5. 接口版本号

任何接口都会带版本号,默认从v1开始，遇到要升级版本的时候v2,v3,依次升级

示例：
```text
api.xxxx.com/v1/config
```

5.响应

error_code=0表示成功， >0表示失败

**success示例**

http status code：200
```json
{
    "error_code":0,
    "msg":"",
    "data":{
        "user_info":{
          "user_name": "crmao",
        }
    }
}
```

**fail示例**

http status code: 400

```json
{
    "error_code":10010002,
    "msg":"Illegal params",
    "data":null
}
```


| 错误标识                | 错误码      | HTTP状态码 | 描述                                                        |
| ----------------------- |----------| ---------- |-----------------------------------------------------------|
| ErrNo                   | 0        | 200        | OK                                                        |
| ErrInternalServer       | 10010001 | 500        | Internal server error （服务器内部错误）                           |
| ErrParams               | 10010002 | 400        | Illegal params  (请求参数不合法)                                 |
| ErrAuthenticationHeader | 10010003 | 401        | Authentication header Illegal  (要登录的接口，guid对应的用户找不到)      |
| ErrNotFound             | 10010005 | 404        | Route not found     (请求路由找不到）                             |
| ErrPermission           | 10010006 | 403        | Permission denied (没有权限,一些接口可能没请求权限)                      |
| ErrTooFast              | 10010007 | 429        | Too Many Requests （用户在给定的时间内发送了太多请求）                      |
| ErrTimeout              | 10010008 | 504        | Server response timeout   （go服务这边不会返回，一般是nginx、网关超时 才返回504） |
| ErrMysqlServer          | 10010101 | 500        | Mysql server error      （mysql 服务错误)                      |
| ErrMysqlSQL             | 10010102 | 500        | Illegal SQL               (sql 代码错误）                      |
| ErrRedisServer          | 10010201 | 500        | Redis server error        （redis 服务错误）                    |

6. 数据安全 
- 接口请求数据，响应数据加密 
如aes， des 算法等 ，可以自行上网查找，与客户端对齐即可
