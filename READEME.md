### 目录
- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据
  
### 本地开发
记得新增目录的使用,必须要在go.mod文件中增加replace代替路径,不然运行时会找不到文件