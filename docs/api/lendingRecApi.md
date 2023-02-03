# 借阅管理相关接口



![image-20230203204433457](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203204433457.png)



## 更新借阅记录

### 接口描述

更新借阅记录

### 请求方法

PUT /lr

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述           |
| :------- | ---- | :----- | :------------- |
| id       | 是   | String | 借阅ID         |
| jydjr    | 否   | string | 借阅登记人     |
| jyr      | 否   | String | 借阅人         |
| jysj     | 否   | String | 借阅时间       |

json传输：

```json
{
  "id": "string",
  "jydjr": "string",
  "jyr": "string",
  "jysj": "string",
  "xh": "string"
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
  "id":"70fddd791f8b43b6a2f1933a00064c9d",
  "jydjr": "U_TEST",
  "jyr": "U_借阅人",
  "jysj": "2023-01-31T13:58:18+08:00",
  "xh": "U_2017108152"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "更新借阅成功！"
}
```



## 添加借阅记录

### 接口描述

添加借阅记录

### 请求方法

POST /lr

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述           |
| :------- | ---- | :----- | :------------- |
| bookId   | 是   | String | 图书ID         |
| jydjr    | 否   | string | 借阅登记人     |
| jyr      | 否   | String | 借阅人         |
| jysj     | 否   | String | 借阅时间       |
| xh       | 否   | String | 学号or身份证号 |

json传输：

```json
{
  "bookId": "string",
  "jydjr": "string",
  "jyr": "string",
  "jysj": "string",
  "xh": "string"
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
  "bookId": "22a32d9240a26915369e9daef7046e79",
  "jydjr": "TEST",
  "jyr": "借阅人",
  "jysj": "2023-02-01T17:03:18+08:00",
  "xh": "2017108152"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "添加借阅成功！"
}
```



## 删除借阅记录

### 接口描述

删除借阅记录

### 请求方法

DELETE /lr

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述   |
| :------- | ---- | :----- | :----- |
| id       | 是   | String | 借阅ID |

json传输：

```json
{
	"id":"string"
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
	"id":"70fddd791f8b43b6a2f1933a00064c9d"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "删除借阅记录成功！"
}
```



## 还书

### 接口描述

还书

### 请求方法

PUT /lr/hs

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述       |
| :------- | ---- | :----- | :--------- |
| bookId   | 是   | String | 图书ID     |
| hsdjr    | 是   | String | 还书登记人 |
| hssj     | 是   | String | 还书时间   |
| id       | 是   | String | 借阅ID     |

json传输：

```json
{
  "bookId": "string",
  "hsdjr": "string",
  "hssj": "string",
  "id": "string"
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
  "id":"70fddd791f8b43b6a2f1933a00064c9d",
  "hsdjr": "Vince",
  "hssj": "2023-02-01T17:42:00+08:00",
  "bookId":"22a32d9240a26915369e9daef7046e79"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "还书成功！"
}
```



## 根据条件查询借阅记录

### 接口描述

根据条件查询借阅记录

### 请求方法

GET /lr/list/search/:page

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**path 参数**

| 参数名称 | 必选 | 类型   | 描述 |
| -------- | ---- | ------ | ---- |
| page     | 是   | String | 页码 |

**Params 参数**

| 参数名称 | 必选 | 类型   | 描述         |
| :------- | ---- | :----- | :----------- |
| xh       | 否   | String | 学号or身份证 |
| jyr      | 否   | String | 借阅人       |
| jysj     | 否   | String | 借阅时间     |
| bookName | 否   | String | 书名         |
| isbn     | 否   | String | ISBN号       |

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
           "list": [
            {
                "id": "string", // 借阅id
                "createdAt": "string", //创建时间
                "updatedAt": "string", // 更新时间
                "deletedAt": "string", // 是否删除
                "bookId": "string", // 图书ID
                "jyzhId": "string", // 借阅账号Id（借阅登记人ID）
                "hszhId": "string", // 还书账号Id（还书登记人ID）
                "xh": "string", // 学号or身份证号
                "jyr": "string", // 借阅人
                "jysj": "string", // 借阅时间
                "hssj": "string", // 还书时间
                "jydjr": "string", // 借阅登记人
                "hsdjr": "string", // 还书登记人
                "isH": true // 是否已还书
            }
        ],
        "total": 0, // 总记录数
        "page": 0, // 当前页码
        "pageSize": 0 // 总页码数
    },
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
http://127.0.0.1:8081/api/v1/lr/list/search/1?xh=U_2017108152&jyr=U_Vince&jysj=2023-02-01 17:03:18&bookName=活着&isbn=9787530221532
```

**输出示例**

```json
{
    "code": 200,
    "data": {
        "list": [
            {
                "id": "70fddd791f8b43b6a2f1933a00064c9d",
                "createdAt": "2023-02-01T17:37:20Z",
                "updatedAt": "2023-02-01T09:47:40Z",
                "deletedAt": null,
                "bookId": "22a32d9240a26915369e9daef7046e79",
                "jyzhId": "3d1656becae0487da784fbb5cf085c46",
                "hszhId": "3d1656becae0487da784fbb5cf085c46",
                "xh": "U_2017108152",
                "jyr": "U_Vince",
                "jysj": "2023-02-01T17:03:18Z",
                "hssj": "2023-02-01T09:42:00Z",
                "jydjr": "U_TEST",
                "hsdjr": "admin",
                "isH": true
            }
        ],
        "total": 1,
        "page": 1,
        "pageSize": 1
    },
    "msg": "查询成功"
}
```



## 获取借阅信息列表

### 接口描述

获取借阅信息列表

### 请求方法

GET /lr/list/:page

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**path 参数**

| 参数名称 | 必选 | 类型   | 描述 |
| -------- | ---- | ------ | ---- |
| page     | 是   | String | 页码 |

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
           "list": [
            {
                "id": "string", // 借阅id
                "createdAt": "string", //创建时间
                "updatedAt": "string", // 更新时间
                "deletedAt": "string", // 是否删除
                "bookId": "string", // 图书ID
                "jyzhId": "string", // 借阅账号Id（借阅登记人ID）
                "hszhId": "string", // 还书账号Id（还书登记人ID）
                "xh": "string", // 学号or身份证号
                "jyr": "string", // 借阅人
                "jysj": "string", // 借阅时间
                "hssj": "string", // 还书时间
                "jydjr": "string", // 借阅登记人
                "hsdjr": "string", // 还书登记人
                "isH": true // 是否已还书
            }
        ],
        "total": 0, // 总记录数
        "page": 0, // 当前页码
        "pageSize": 0 // 总页码数
    },
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
curl -X 'GET' \
  'http://127.0.0.1:8888/api/v1/lr/list/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjNkMTY1NmJlY2FlMDQ4N2RhNzg0ZmJiNWNmMDg1YzQ2IiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi566h55CG55So5oi3IiwiaXNzIjoiTE1TIiwiZXhwIjoxNjc1NTEzOTc4LCJuYmYiOjE2NzU0Mjc1NzgsImlhdCI6MTY3NTQyNzU3OH0.E4RssZ3U6vpVW9ldS3bWUuNINJe24W12knvBtgNHCCI'
```

**输出示例**

```json
{
  "code": 200,
  "data": {
    "page": 1,
    "pageSize": 1,
    "total": 1,
    "list": [
      {
        "id": "70fddd791f8b43b6a2f1933a00064c9d",
        "createdAt": "2023-02-01T17:37:20+08:00",
        "updatedAt": "2023-02-01T09:47:40+08:00",
        "deletedAt": null,
        "bookId": "22a32d9240a26915369e9daef7046e79",
        "jyzhId": "3d1656becae0487da784fbb5cf085c46",
        "hszhId": "3d1656becae0487da784fbb5cf085c46",
        "xh": "U_2017108152",
        "jyr": "U_Vince",
        "jydjr": "U_TEST",
        "hsdjr": "admin",
        "jysj": "2023-02-01T17:03:18+08:00",
        "hssj": "2023-02-01T09:42:00+08:00",
        "isH": true
      }
    ]
  },
  "msg": "查询成功"
}
```