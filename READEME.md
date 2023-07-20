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
kill -9 pid(上面命令执行后会有)


### swag 生成命令
```js

swag init

```

docker数据库挂载
```js
docker pull mysql
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin123 -v /Users/luogengzhong/Documents/study/docker-mysql:/var/lib/mysql -d mysql
```
docker 数据库登录 密码rootroot

 远程服务器执行的命令是:
 ```js
 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin123 -v /root/mydata/docker-mysql:/var/lib/mysql -d mysql
 ```
```js
 登录数据库后,记得执行创建数据库和表命令


docker exec -it mysql mysql -uroot -p
```

docker 项目镜像
```js
//先创建Dokerfile文件
//根目录下执行构建镜像 创建后查看docker images
docker build -t gin-blog-docker .
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .
docker build -t gin-blog-docker-scratch .
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker-scratch
```



查看影像
docker ps -as

删除影像
docker rm 影像name

启动/停止 影像
docker start/stop 影像NAMES


查看文件夹路径
pwd

模拟定时任务删除
```js
go run cron.go
```