# 系统相关接口

 ![image-20230203191543191](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203191543191.png)



## 用户登录

### 接口描述

用户登录

### 请求方法

POST /login

### 输入参数

**Body 参数**

| 参数名称 | 类型   | 描述         |
| -------- | ------ | ------------ |
| code     | String | 响应码       |
| data     | obj    | 返回数据信息 |
| msg      | string | 信息         |

json传输：

```json
{
    "username":"admin", // 用户名
    "password":"123456", // 密码
    "loginTime":"2023-01-31T13:58:18+08:00" // 登录时间
}
```

### 输出参数

| 参数名称 | 类型   | 描述         |
| -------- | ------ | ------------ |
| code     | String | 响应码       |
| data     | obj    | 返回数据信息 |
| msg      | string | 信息         |

```json
{
    "code": 0,
    "data": {
        "id": "string",
        "username": "string",
        "roleEn": "string",
        "authorization": "string"
    },
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
{
    "username":"admin",
    "password":"123456",
    "loginTime":"2023-01-31T13:58:18+08:00"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {
        "id": "3d1656becae0487da784fbb5cf085c46",
        "username": "admin",
        "roleEn": "gly",
        "authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjNkMTY1NmJlY2FlMDQ4N2RhNzg0ZmJiNWNmMDg1YzQ2IiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi566h55CG55So5oi3IiwiaXNzIjoiTE1TIiwiZXhwIjoxNjc1MzQ4Mjk5LCJuYmYiOjE2NzUyNjE4OTksImlhdCI6MTY3NTI2MTg5OX0.PIX6kx-CFa1K94n0BV0OLtk74AbUc5MfGIci0JCwjiE"
    },
    "msg": "查询成功"
}
```



## 用户登出

### 接口描述

用户登出

### 请求方法

POST /logout

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

### 输出参数

```json
{
  "code": 0,
  "data": {},
  "msg": "string"
}
```



## 用户注册

### 接口描述

用户注册

### 请求方法

POST /register

### 输入参数

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述     |
| :------- | ---- | :----- | :------- |
| email    | 否   | String | 邮箱     |
| nickName | 否   | String | 昵称     |
| password | 是   | string | 密码     |
| phone    | 否   | String | 电话     |
| roleEn   | 否   | String | 角色标识 |
| roleId   | 否   | String | 角色ID   |
| username | 是   | String | 登录名   |

json传输：

```json
{
  "email": "string",
  "nickName": "string",
  "password": "string",
  "phone": "string",
  "roleEn": "string",
  "roleId": "string",
  "username": "string"
}
```

### 输出参数

| 参数名称 | 类型   | 描述         |
| -------- | ------ | ------------ |
| code     | String | 响应码       |
| data     | obj    | 返回数据信息 |
| msg      | string | 信息         |

```json
{
    "code": 0,
    "data": {
        "id": "string",
        "username": "string",
        "roleEn": "string",
        "authorization": "string"
    },
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
{
  "email": "turbochang@foxmail.com",
  "nickName": "vince",
  "password": "123456",
  "phone": "13011011010",
  "roleEn": "gly",
  "roleId": "3cea30e811823c11dc224ace2cfd11cf",
  "username": "sys"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "注册成功"
}
```





