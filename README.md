<div align=center>
<img src="https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203225420916.png" width=400" height="200"/>
</div>
<div align=center>
<img src="https://img.shields.io/static/v1?logo=Gitee&label=build&message=passing&color=lightBlue">
<img src="https://img.shields.io/badge/golang-1.19-blue"/>
<img src="https://img.shields.io/badge/gin-1.8.2-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.24.3-ff69b4"/>
<img src="https://img.shields.io/static/v1?logo=Apache&label=license&message=Apache 2.0&color=orange">
</div>




# 基于Gin + Gorm - LMS图书借阅后端系统

基于Gin + Gorm 后端的图书借阅系统。Restful风格。

使用到了`gin`，`gorm`，`swagger`，`zap`，`viper`。

集成`jwt鉴权`，`Redis缓存`，`MySql`

可根据系统版本号，自动初始化SQL数据。

该项目是一个基础项目，可以自行扩展更多功能。

## 系统架构





### 系统架构图

![系统架构图](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/%E7%B3%BB%E7%BB%9F%E6%9E%B6%E6%9E%84.png)



### 目录结构

    ├── api             (api层)
    │   └── v1          (v1版本接口)
    ├── config          (配置包)
    ├── constant		(常量)
    ├── core            (核心文件)
    ├── docs            (swagger文档目录)
    ├── global          (全局对象)                    
    ├── initialize      (初始化)                        
    ├── middleware      (中间件层)                       
    ├── model           (模型层)                    
    │   ├── request     (入参结构体)                     
    │   └── response    (出参结构体)                     
    ├── routes          (路由层)                    
    ├── service         (service层)                    
    ├── source          (source层)  
    │   └── sql_init	(sql初始化脚本)
    │       └── v1.0 	(v1.0版本的初始化脚本)
    ├── unit_test    	(单元测试)
    └── utils           (工具包)                    
    ├── config-default.yaml
    ├── config-dev.yaml
    ├── config-docker.yaml
    ├── Dockerfile
    ├── go.mod
    ├── LICENSE
    ├── main.go			(程序启动入口)
    ├── Makefile
    └── README.md



### 主要功能

- 用户管理：用户信息的增、删、改、查。账户的角色信息的修改。登录、退出功能。
- 角色管理：角色信息的增、删、改、查。
- 授权管理：基于`JWT`实现授权（权限）管理。
- 图书管理：图书信息的增、删、改、查。
- 借阅管理：图书借阅信息的增、删、改、查。
- 分页封装：图书、借阅、用户信息查询通过分页查询。
- 条件搜索：根据条件进行数据的搜索。



## API介绍

LMS API 接口文档

- [系统相关接口](./docs/api/baseApi.md)
- [用户管理相关接口](./docs/api/userApi.md)
- [角色相关接口](./docs/api/roleApi.md)
- [图书管理相关接口](./docs/api/bookApi.md)
- [借阅管理相关接口](./docs/api/lendingRecApi.md)



### API 概览

#### 系统相关接口

> 详情：[系统相关接口](./docs/api/baseApi.md)

| 请求方式 | 接口      | 功能     |
| -------- | --------- | -------- |
| POST     | /login    | 登录     |
| POST     | /logout   | 退出登录 |
| POST     | /register | 账户注册 |



#### 用户管理相关接口

> 详情：[用户管理相关接口](./docs/api/userApi.md)

| 请求方式 | 接口       | 功能         |
| -------- | ---------- | ------------ |
| GET      | /user/self | 查询用户信息 |
| PUT      | /user/self | 更新用户信息 |
| POST     | /user/self | 更改密码     |



#### 角色相关接口

> 详情：[角色相关接口](./docs/api/roleApi.md)

| 请求方式 | 接口       | 功能             |
| -------- | ---------- | ---------------- |
| GET      | /role/list | 获取角色信息列表 |
| PUT      | /role/ur   | 更改用户角色     |
| POST     | /role      | 添加角色信息     |
| PUT      | /role      | 修改角色信息     |
| DELETE   | /role      | 删除角色信息     |



#### 图书管理相关接口

> 详情：[图书管理相关接口](./docs/api/bookApi.md)

| 请求方式 | 接口                    | 功能                 |
| -------- | ----------------------- | -------------------- |
| PUT      | /book                   | 更新图书             |
| POST     | /book                   | 添加图书             |
| DELETE   | /book                   | 删除图书             |
| GET      | /book/list/publisher    | 获取所有出版社名称   |
| GET      | /book/list/search/:page | 根据条件查询图书列表 |
| GET      | /book/list/:page        | 分页查询图书列表     |



#### 借阅管理相关接口

> 详情：[借阅管理相关接口](./docs/api/lendingRecApi.md)

| 请求方式 | 接口                  | 功能             |
| -------- | --------------------- | ---------------- |
| PUT      | /lr                   | 更新借阅信息     |
| POST     | /lr                   | 添加借阅信息     |
| DELETE   | /lr                   | 删除借阅信息     |
| PUT      | /lr/hs                | 还书             |
| GET      | /lr/list/search/:page | 根据条件查询     |
| GET      | /lr/list/:page        | 获取借阅信息列表 |



## 配置介绍



### 启动参数

> 启动参数读取优先级：传参 > 启动参数 > 环境变量 > Env配置
>
> 传参：`main.go` 中 `core.Viper(传参指定配置文件)`
>
> 环境变量：配置`LMS_CONFIG` 指定配置文件
>
> ENV配置：在` constant/configConstant.go` 中`ENV`常量设置

| 参数  | 值                 |
| ----- | ------------------ |
| -conf | 指定yaml配置文件。 |



在系统启动时：

- 若已编译成二进制可执行文件指定配置文件：

  ```bash
  ./lms_server -conf config-docker.yaml
  ```

- GOLAND编译器指定配置文件：

  ![image-20230203173709146](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203173709146.png)

  

### CONFIG YAML配置

配置文件：`config-[type].yaml`

#### SYSTEM

| KEY                | 描述                                                         |
| ------------------ | ------------------------------------------------------------ |
| name               | 表示本系统的名称，可自定义                                   |
| version            | 鉴别本系统的版本。影响sql的初始化。根据该版本号在程序启动时，会查看数据库是否存在该版本号记录，如不存在则会执行`./source/sql_init`目录下对应版本号的sql初始化脚本。 |
| db-type            | 用于指定使用的数据库。本系统仅支持mysql。可`databaseInit.go`文件自行扩展支持其他数据库。 |
| router-prefix      | api访问的前缀。                                              |
| transfer-pw-decode | http传输的登录密码是否需要解密。默认：false（即传入系统的密码是明文） |

> **IPS:**
>
> - **transfer-pw-decode**
>
>   若为`true`，可传以下json进行登录
>
>   ```json
>   {
>       "username":"admin",
>       "password":"MTI1zNDUp2oCiA=",
>       "loginTime":"2023-01-31T13:58:18+08:00"
>   }
>   ```
>
>   http传入系统的加密解密方式逻辑可参考：`utils/hashUtils.go:PasswordDecode()`
>
>   **解密逻辑：** 后台更具loginTime时间进行计算。计算规则为：小时%10取余，分钟%100取余，秒%10取余，获取的数进行从小到大排序，将其看成索引，将password对应索引位置的字符元素删除之后形成的字符串就是正确的base64码，然后进行base解码。
>
>   **tips： ** 本加密解密逻辑非常简单，不适合正式场景的使用。需要自行制定修改前后端传输密码的加密解密逻辑。



#### GIN

| KEY  | 描述                |
| ---- | ------------------- |
| PORT | 指定web服务启动端口 |



#### JWT

| KEY          | 描述                                                         |
| ------------ | ------------------------------------------------------------ |
| signing-key  | 签名密钥                                                     |
| expires-time | jwt多久过期。可设置天数`d`，小时`h`。如：一天则`1d`，一小时则`1h` |
| buffer-time  | jwt缓冲时间可设置天数`d`，小时`h`。如：一天则`1d`，一小时则`1h` |
| issuer       | 签发者                                                       |



#### ZAP

| KEY            | 描述                                                         |
| -------------- | ------------------------------------------------------------ |
| level          | 日志等级。`debug` `info` `warn` `error` `dpanic` `panic` `fatal` 默认`debug` |
| gorm-level     | gorm日志等级。`info` `warn` `error` `silent` 默认 `info`     |
| format         | 日志输出格式。可选 `console` 和`json`，默认 `console`        |
| path           | 日志存放路径。                                               |
| prefix         | 日志打印前缀。                                               |
| log-in-console | 日志是否打印到console                                        |

> **TIPS:**
>
> zap引入`file-rotatelogs`按天切割日志，需**注意**：按时间切割效率与按文件大小切割相比低且不能保证日志数据不被破坏



#### REDIS

| KEY      | 描述         |
| -------- | ------------ |
| db       | 指定存入的库 |
| addr     | `ip:port`    |
| password | 密码         |



#### MYSQL

| KEY            | 描述       |
| -------------- | ---------- |
| ip             | ip地址     |
| port           | 端口号     |
| db-name        | 数据库名   |
| username       | 用户名     |
| password       | 密码       |
| charset        | 字符       |
| max-idle-conns | 最大空闲数 |
| max-open-conns | 最大连接数 |
| log-mode       | 日志模式   |
| log-zap        | -          |



## 快速使用

### GOLAND编译器运行

将代码克隆下来之后，按照配置介绍相关信息自行进行配置。

启动后访问`http://127.0.0.1:8081/api/v1/swagger/index.html`即可。



### DOCKER部署运行

本步骤使用环境：CENTOS 8.0

> **TIPS**：项目克隆下来之后建议先修改`config-docker.yaml`文件的配置信息，若不修改镜像制作好之后启动由于`config-docker.yaml`配置错误会报错启动不起来，后续还是要在数据卷中去修改`config-docker.yaml`。



- **将项目克隆到自己制定的文件夹下，执行**

  ```bash
  docker build -f ./Dockerfile -t lms:v1.0.0 .
  ```

  完成之后，会提示成功的提示

- **查看docker镜像列表**

  ```bash
  docker images
  ```

  此时就有了lms的镜像：

  ![image-20230203181425924](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203181425924.png)

- **创建数据卷（可选）**

  ```bash
  # 创建名为lms_volume的数据卷
  docker volume create lms_volume
  
  # 查看所有数据卷列表
  docker volume ls
  ```

  ![image-20230203181916344](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203181916344.png)

  > **TIPS** ：如果没有创建数据卷，那么会在docker目录下自动生成绑定数据卷。一般自动生成的数据卷挂载路径为：`/var/lib/docker/volumes/` 该目录下

- **启动lms镜像**

  ```bash
  docker run -d -p 8888:8888 -v lms_volume:/go/src/github.com/lms --name lms_server [image id]
  ```

  > **TIPS**：若没有创建数据卷，则启动命令为：
  >
  > ```bash
  > docker run -d -p 8888:8888 --name lms_server [image id]
  > ```

- **查看镜像启动状态**

  ```bash
  docker ps -a
  ```

  若`STATUS`显示`EXIT`退出或者没有成功启动。可以到对应的数据卷进行查看日志或修改`config-docker.yaml`文件配置。

  ![image-20230203182918657](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203182918657.png)

  再重启`lms_server`容器

  ```bash
  # 重启命令
  docker restart [容器id]
  ```

- 容器启动成功访问：`http://ip:8088/api/v1/swagger/index.html`



### 启动成功展示页

![image-20230203183029821](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203183029821.png)



### 其他问题

**？** 如果启动之后访问不了，建议查看防火墙是否放开8888端口。



## 特别鸣谢

---

学习资料：

- [gin-vue-admin # 不错的模版框架](https://github.com/flipped-aurora/gin-vue-admin)

---



感谢 [JetBrains](https://www.jetbrains.com/?from=ferry) 为本开源项目提供免费的 [IntelliJ GoLand](https://www.jetbrains.com/go/?from=ferry) 授权

[![img](https://camo.githubusercontent.com/8cfd0f5e55bcb43e6e1db012f95fc5b5e4c6d553bd4114eea530c99c1ad7d595/68747470733a2f2f7777772e666465766f70732e636f6d2f77702d636f6e74656e742f75706c6f6164732f323032302f30392f313539393231333835372d6a6574627261696e732d76617269616e742d342e706e67)](https://www.jetbrains.com/?from=ferry)



Copyright (c) 2023 Vincez