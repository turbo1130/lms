## 角色相关接口



![image-20230203194441697](https://gitee.com/turbo30/study_pic/raw/master/pictureForLms/image-20230203194441697.png)



## 获取角色信息列表

### 接口描述

获取角色信息列表

### 请求方法

GET /role/list

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
    "data": [
        {
            "id": "string", // 角色ID
            "createdAt": "string", // 创建时间
            "updatedAt": "string", // 更新时间
            "deletedAt": 0, // 是否删除
            "roleEn": "string", // 角色标识
            "roleName": "string" // 角色名称
        }
    ],
    "msg": "string"
}
```

### 请求示例

**输入示例**



**输出示例**

```json
{
    "code": 200,
    "data": [
        {
            "id": "3cea30e811823c11dc224ace2cfd11cf",
            "createdAt": "2023-02-01T17:36:58+08:00",
            "updatedAt": "2023-02-01T17:36:58+08:00",
            "deletedAt": null,
            "roleEn": "gly",
            "roleName": "管理员"
        },
        {
            "id": "688e78155f42205fbabdd96f812cda3f",
            "createdAt": "2023-02-01T17:36:58+08:00",
            "updatedAt": "2023-02-01T17:36:58+08:00",
            "deletedAt": null,
            "roleEn": "ptzy",
            "roleName": "普通职员"
        }
    ],
    "msg": "查询成功"
}
```



## 更改用户角色权限

### 接口描述

更改用户角色权限

### 请求方法

POST /role/ur

### 输入参数

**Header参数**

| 参数名称      | 必选 | 类型   | 描述            |
| ------------- | ---- | ------ | --------------- |
| Authorization | 是   | String | "Bearer " + JWT |

**Body 参数**

| 参数名称 | 必选 | 类型   | 描述     |
| :------- | ---- | :----- | :------- |
| id       | 是   | String | 用户ID   |
| roleEn   | 是   | String | 角色标识 |
| roleId   | 是   | string | 角色ID   |

json传输：

```json
{
  "id": "string",
  "roleEn": "string",
  "roleId": "string"
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
    "code": 200,
    "data": {},
    "msg": "string"
}
```

### 请求示例

**输入示例**

```json
{
  "id": "2ceec9a8253c41019426aaffeff444be",
  "roleEn": "ptzy",
  "roleId": "688e78155f42205fbabdd96f812cda3f"
}
```

**输出示例**

```json
{
    "code": 200,
    "data": {},
    "msg": "更新用户角色权限成功"
}
```