# 用户管理相关接口



![image-20230203193452628](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203193452628.png)



## 查询用户信息

### 接口描述

查询用户信息

### 请求方法

GET /user/self

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

### 输出参数

| 参数名称 | 类型   | 描述         |
| -------- | ------ | ------------ |
| code     | String | 响应码       |
| data     | obj    | 返回数据信息 |
| msg      | string | 信息         |

```json
{
    "code": 200,
    "data": {
        "id": "string", // id
        "username": "string", // 用户名
        "phone": "string", // 电话
        "email": "string", // 邮箱
        "nickName": "string" // 昵称
    },
    "msg": "string"
}
```

### 请求示例

**输入示例**



**输出示例**

```json
{
    "code": 200,
    "data": {
        "id": "113056becae0487da784fbb5cf085c46",
        "username": "sys",
        "phone": "13011011010",
        "email": "turbochang@foxmail.com",
        "nickName": "vincez"
    },
    "msg": "查询成功"
}
```





## 更新用户信息

### 接口描述

更新用户信息

### 请求方法

PUT /user/self

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述   |
| :------- | ---- | :----- | :----- |
| email    | 否   | String | 邮箱   |
| id       | 是   | String | 用户ID |
| nickName | 否   | string | 昵称   |
| phone    | 否   | String | 电话   |

json传输：

```json
{
  "email": "string",
  "id": "string",
  "nickName": "string",
  "phone": "string"
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
    "data": {},
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
{
  "email": "test@126.com",
  "id": "3d1656becae0487da784fbb5cf085c46",
  "nickName": "TEST昵称",
  "phone": "13000000000"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "更新用户信息成功"
}
```



## 更改密码

### 接口描述

更改密码

### 请求方法

POST /user/self

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称   | 必选 | 类型   | 描述       |
| :--------- | ---- | :----- | :--------- |
| changeTime | 是   | String | 改密码时间 |
| id         | 是   | String | 用户ID     |
| password   | 是   | string | 密码       |

json传输：

```json
{
  "changeTime": "string",
  "id": "string",
  "password": "string"
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
    "data": {},
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
{
  "changeTime": "2023-01-31T13:58:18+08:00",
  "id": "3d1656becae0487da784fbb5cf085c46",
  "password": "123456"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "更改密码成功！"
}
```