# 图书管理相关接口



![image-20230203201802906](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203201802906.png)



## 更新图书

### 接口描述

更新图书

### 请求方法

PUT /book

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称  | 必选 | 类型   | 描述                                          |
| :-------- | ---- | :----- | :-------------------------------------------- |
| id        | 是   | String | 图书ID                                        |
| author    | 否   | String | 作者                                          |
| isbn      | 否   | string | ISBN号                                        |
| jcs       | 否   | int    | 借出数                                        |
| isJcs     | 否   | bool   | 若jcs为 0 则isJcs必须传true，才能正确更新     |
| name      | 否   | String | 书名                                          |
| num       | 否   | String | 书号                                          |
| price     | 否   | String | 价格                                          |
| publisher | 否   | string | 出版社                                        |
| stock     | 否   | int    | 库存                                          |
| isStock   | 否   | bool   | 若stock为 0 则isStock必须传true，才能正确更新 |

json传输：

```json
{
  "id":"string",
  "author": "string",
  "isbn": "string",
  "jcs": 0,
  "isJcs": true,
  "name": "string",
  "num": "string",
  "price": "string",
  "publisher": "string",
  "stock": 0,
  "isStock": false
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
  "id":"a350dd8e084f407ca742e54996b02517",
  "author": "U_TEST_AUTHOR",
  "isbn": "U23455688",
  "jcs": 0,
  "isJcs": true,
  "name": "U_TEST",
  "num": "U_101-test",
  "price": "66.66",
  "publisher": "U_TEST出版社",
  "stock": 23,
  "isStock": false
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "更新图书成功！"
}
```



## 添加图书

### 接口描述

添加图书

### 请求方法

POST /book

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称  | 必选 | 类型   | 描述   |
| :-------- | ---- | :----- | :----- |
| author    | 是   | String | 作者   |
| isbn      | 是   | string | ISBN号 |
| jcs       | 是   | int    | 借出数 |
| name      | 是   | String | 书名   |
| num       | 是   | String | 书号   |
| price     | 是   | String | 价格   |
| publisher | 是   | String | 出版社 |
| stock     | 是   | int    | 库存量 |

json传输：

```json
{
  "author": "string",
  "isbn": "string",
  "jcs": 0,
  "name": "string",
  "num": "string",
  "price": "string",
  "publisher": "string",
  "stock": 0
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
  "author": "TEST_AUTHOR",
  "isbn": "23455688",
  "jcs": 0,
  "name": "TEST",
  "num": "101-test",
  "price": "66.66",
  "publisher": "TEST出版社",
  "stock": 12
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "添加图书成功！"
}
```



## 删除图书

### 接口描述

删除图书

### 请求方法

DELETE /book

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述   |
| :------- | ---- | :----- | :----- |
| id       | 是   | String | 图书ID |

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
  "id":"67c65d5d3fba4184b9cf02c710335913"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "删除图书成功！"
}
```



## 获取所有出版社名称

### 接口描述

获取所有出版社名称

### 请求方法

GET /book/list/publisher

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
    "code": 0,
    "data": [],
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
    "page": 0,
    "pageSize": 0,
    "total": 0,
    "list": [
      "北京十月文艺出版社",
      "北方文艺出版社",
      "中国经济出版社",
      "浙江文艺出版社",
      "重庆出版社",
      "海南出版社",
      "湖南文艺出版社",
      "中信出版社",
      "花城出版社",
      "上海三联书店",
      "西北大学出版社",
      "译林出版社"
    ]
  },
  "msg": "查询成功"
}
```



## 根据条件查询图书列表信息

### 接口描述

根据条件查询图书列表信息

### 请求方法

GET /book/list/search/:page

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

| 参数名称  | 必选 | 类型   | 描述   |
| :-------- | ---- | :----- | :----- |
| num       | 否   | String | 书号   |
| name      | 否   | String | 书名   |
| author    | 否   | String | 作者   |
| isbn      | 否   | String | ISBN号 |
| publisher | 否   | String | 出版社 |

### 输出参数

| 参数名称 | 类型   | 描述         |
| -------- | ------ | ------------ |
| code     | String | 响应码       |
| data     | obj    | 返回数据信息 |
| msg      | string | 信息         |

```json
{
    "code": 0,
    "data": [],
    "msg": "string"
}
```

### 请求示例

**输入示例**

```bash
curl -X 'GET' \
  'http://127.0.0.1:8888/api/v1/book/list/search/1?name=%E4%B8%89%E4%BD%93%EF%BC%88%E4%B8%80%EF%BC%89' \
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
        "id": "6668eec82131325def8134afe5031d59",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "101-55683",
        "name": "三体（一）",
        "author": "刘慈欣",
        "isbn": "9787536692930",
        "publisher": "重庆出版社",
        "price": "23.00",
        "jcs": 0,
        "stock": 43
      }
    ]
  },
  "msg": "查询成功"
}
```



## 分页查询图书列表

### 接口描述

获取图书列表

### 请求方法

GET /book/list/:page

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
        "page": 0, // 当前页
    	"pageSize": 0, // 总页数
    	"total": 0, // 总记录数
    	"list": [] // 返回的数据
    },
    "msg": "string"
}
```

### 请求示例

**输入示例**

```bash
curl -X 'GET' \
  'http://127.0.0.1:8888/api/v1/book/list/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjNkMTY1NmJlY2FlMDQ4N2RhNzg0ZmJiNWNmMDg1YzQ2IiwiVXNlcm5hbWUiOiJhZG1pbiIsIk5pY2tOYW1lIjoi566h55CG55So5oi3IiwiaXNzIjoiTE1TIiwiZXhwIjoxNjc1NTEzOTc4LCJuYmYiOjE2NzU0Mjc1NzgsImlhdCI6MTY3NTQyNzU3OH0.E4RssZ3U6vpVW9ldS3bWUuNINJe24W12knvBtgNHCCI'
```

**输出示例**

```json
{
  "code": 200,
  "data": {
    "page": 1,
    "pageSize": 2,
    "total": 12,
    "list": [
      {
        "id": "688e78155f42205fbabdd96f812cda3f",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "102-75395",
        "name": "回望：一个经济学家是如何长成的",
        "author": "张维迎",
        "isbn": "9787573008237",
        "publisher": "海南出版社",
        "price": "53.00",
        "jcs": 0,
        "stock": 8
      },
      {
        "id": "bdb3705ed79fa7ec4871e027bbcc70d9",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "102-24562",
        "name": "人间漂流",
        "author": "小杜",
        "isbn": "9787542678645",
        "publisher": "上海三联书店",
        "price": "59.00",
        "jcs": 0,
        "stock": 12
      },
      {
        "id": "26ae8d491ed42c67f4a1461260a58159",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "103-45612",
        "name": "雅舍谈吃",
        "author": "梁实秋",
        "isbn": "9787531756934",
        "publisher": "北方文艺出版社",
        "price": "17.40",
        "jcs": 0,
        "stock": 15
      },
      {
        "id": "c0e50e74b3040090ccce268926293ef0",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "104-12645",
        "name": "市场的逻辑",
        "author": "张维迎著，理想国",
        "isbn": "9787560443621",
        "publisher": "西北大学出版社",
        "price": "58.00",
        "jcs": 0,
        "stock": 16
      },
      {
        "id": "aaa0c048426382d1deaf8d5cb2f49438",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "102-24531",
        "name": "一地鸡毛",
        "author": "刘震云",
        "isbn": "9787536097254",
        "publisher": "花城出版社",
        "price": "30.00",
        "jcs": 0,
        "stock": 19
      },
      {
        "id": "3cea58e811823c11dc224ace2cfd83cf",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "101-55685",
        "name": "投资改变人生：那些滚雪球的人（第二辑）",
        "author": "雪球",
        "isbn": "9787513670678",
        "publisher": "中国经济出版社",
        "price": "46.70",
        "jcs": 0,
        "stock": 23
      },
      {
        "id": "73e055675966774d7791557b44d4fd0a",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "103-11300",
        "name": "迟来的告白（年历版）",
        "author": "陶立夏",
        "isbn": "29492449",
        "publisher": "湖南文艺出版社",
        "price": "128.00",
        "jcs": 0,
        "stock": 28
      },
      {
        "id": "9dc8a12725a8ada622a69b6be9da4aa1",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "104-15462",
        "name": "人类简史：从动物到上帝",
        "author": "[以色列] 尤瓦尔·赫拉利（Yuval Noah Harari）",
        "isbn": "9787508660752",
        "publisher": "中信出版社",
        "price": "68.00",
        "jcs": 0,
        "stock": 37
      },
      {
        "id": "ca47d246b3ac38ab8fb5ffca493bbf65",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "101-55684",
        "name": "我只知道人是什么",
        "author": "余华",
        "isbn": "9787544788182",
        "publisher": "译林出版社",
        "price": "29.00",
        "jcs": 0,
        "stock": 37
      },
      {
        "id": "6668eec82131325def8134afe5031d59",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "101-55683",
        "name": "三体（一）",
        "author": "刘慈欣",
        "isbn": "9787536692930",
        "publisher": "重庆出版社",
        "price": "23.00",
        "jcs": 0,
        "stock": 43
      },
      {
        "id": "5a18b9bce91070fab6ff85914f864120",
        "createdAt": "2023-02-01T17:36:58+08:00",
        "updatedAt": "2023-02-01T17:36:58+08:00",
        "deletedAt": null,
        "num": "102-24563",
        "name": "望江南",
        "author": "王旭烽",
        "isbn": "9787533963354",
        "publisher": "浙江文艺出版社",
        "price": "35.68",
        "jcs": 0,
        "stock": 45
      }
    ]
  },
  "msg": "查询成功"
}
```

