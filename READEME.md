### 目录
- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据
  
### 本地开发
记得新增目录的使用,必须要在go.mod文件中增加replace代替路径,不然运行时会找不到文件


启动服务:

```js
go build main.go

//然后执行
./main


```
启动成功后，输出了pid号；可以在另外一个终端执行 kill -1 pid号 ，这样可以不必要ctrl+c每次终止再启动的运行方式
关闭:
sudo lsof -i :8000
kill -9 pid


### swag 安装二进制文件命令
```js
//https://github.com/swaggo/swag/blob/master/README_zh-CN.md#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B
go install github.com/swaggo/swag/cmd/swag@v1.6.5



```