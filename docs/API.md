---
title: FashOJ
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# FashOJ

Base URLs:

# Authentication

# normal

## GET 获取题目信息

GET /api/problem/{problem_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|problem_id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "problem_id": "string",
  "author": "string",
  "limit": {
    "time_limit": "string",
    "memory_limit": "string"
  },
  "difficulty": 9,
  "try": 0,
  "ac": 0,
  "title": "string",
  "content": "string",
  "example": [
    {
      "input": "string",
      "output": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» problem_id|string|true|none||none|
|» author|string|true|none||none|
|» limit|object|true|none||none|
|»» time_limit|string|true|none||none|
|»» memory_limit|string|true|none||none|
|» difficulty|integer|true|none||none|
|» try|integer|true|none||none|
|» ac|integer|true|none||none|
|» title|string|true|none||none|
|» content|string|true|none||none|
|» example|[object]|true|none||none|
|»» input|string|true|none||none|
|»» output|string|true|none||none|

## DELETE 删除题目

DELETE /api/problem/{problem_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|problem_id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": "string",
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|string|true|none||none|
|» message|string|true|none||none|

## PATCH 修改题目

PATCH /api/problem/{problem_id}

修改题目

> Body 请求参数

```json
{
  "type": "object",
  "properties": {
    "problem_id": {
      "type": "string"
    },
    "limit": {
      "type": "object",
      "properties": {
        "time_limit": {
          "type": "number"
        },
        "memory_limit": {
          "type": "number"
        }
      },
      "required": [
        "time_limit",
        "memory_limit"
      ]
    },
    "difficulty": {
      "type": "integer",
      "minimum": 0,
      "maximum": 9
    },
    "title": {
      "type": "string"
    },
    "content": {
      "type": "string"
    },
    "example": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "input": {
            "type": "string"
          },
          "output": {
            "type": "string"
          }
        },
        "required": [
          "input",
          "output"
        ]
      }
    }
  },
  "required": [
    "problem_id",
    "limit",
    "difficulty",
    "title",
    "content",
    "example"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|problem_id|path|string| 是 |none|
|body|body|object| 否 |none|
|» problem_id|body|string| 否 |none|
|» limit|body|object| 否 |none|
|»» time_limit|body|string| 是 |none|
|»» memory_limit|body|string| 是 |none|
|» difficulty|body|integer| 否 |none|
|» title|body|string| 否 |none|
|» content|body|string| 否 |none|
|» example|body|[object]| 否 |none|
|»» input|body|string| 是 |none|
|»» output|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 创建题目

POST /api/problem

> Body 请求参数

```json
{
  "problem_id": "1",
  "limit": {
    "time_limit": 1750160029783,
    "memory_limit": 61
  },
  "difficulty": 9,
  "title": "减去因此敲挖干脆",
  "content": "irure ad culpa Excepteur",
  "example": [
    {
      "input": "cillum pariatur sed deserunt consequat",
      "output": "in ipsum veniam"
    },
    {
      "input": "veniam cupidatat",
      "output": "amet enim"
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» problem_id|body|string| 是 |none|
|» limit|body|object| 是 |none|
|»» time_limit|body|number| 是 |none|
|»» memory_limit|body|number| 是 |none|
|» difficulty|body|integer| 是 |none|
|» title|body|string| 是 |none|
|» content|body|string| 是 |none|
|» example|body|[object]| 是 |none|
|»» input|body|string| 是 |none|
|»» output|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 提交代码

POST /api/submit

> Body 请求参数

```json
{
  "problem_id": "24",
  "code": "4",
  "lang": "laborum ipsum eu"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» problem_id|body|string| 是 |none|
|» code|body|string| 是 |none|
|» lang|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "submit_id": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» submit_id|string|true|none||none|

## GET 获取所有提交信息

GET /api/submit

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|p|query|integer| 否 |page|

> 返回示例

> 200 Response

```json
{
  "problems": [
    {
      "problem_id": "string",
      "title": "string",
      "author": "string",
      "limit": {
        "time_limit": "string",
        "memory_limit": "string"
      },
      "difficulty": 9,
      "try": 0,
      "ac": 0
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» problems|[object]|true|none||none|
|»» problem_id|string|false|none||none|
|»» title|string|false|none||none|
|»» author|string|false|none||none|
|»» limit|object|false|none||none|
|»»» time_limit|string|true|none||none|
|»»» memory_limit|string|true|none||none|
|»» difficulty|integer|false|none||none|
|»» try|integer|false|none||none|
|»» ac|integer|false|none||none|

## GET 获取指定的提交信息

GET /api/submit/{submit_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|submit_id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "result": "string",
  "message": "string",
  "testcases": [
    {
      "id": 0,
      "status": "string",
      "time": "string",
      "memory": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» result|string|true|none||none|
|» message|string|true|none||none|
|» testcases|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» status|string|true|none||none|
|»» time|string|true|none||none|
|»» memory|string|true|none||none|

## POST 上传题目测试数据

POST /api/problem/{problem_id}/upload

> Body 请求参数

```yaml
file: file://C:\Users\zine_\Data\新建文件夹\2.zip

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|problem_id|path|string| 是 |none|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# auth

## POST 用户登录

POST /api/auth/login

> Body 请求参数

```json
{
  "username": "zine",
  "password": "123456"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI3MzM2NTYsInVzZXJuYW1lIjoiemluMTIzZSJ9.X9MHeM379XmGC1yax4pYSREqjk3Pgz84o4o_LyrzhtU"
}
```

> 401 Response

```json
{
  "error": "wrong password or username"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|string|true|none||none|
|» message|string|true|none||none|
|» token|string|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» error|string|true|none||none|

## POST 用户注册

POST /api/auth/register

> Body 请求参数

```json
{
  "username": "zine",
  "password": "123456",
  "email": "zine@qq.com"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|
|» email|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": "string",
  "message": "string",
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|string|true|none||none|
|» message|string|true|none||none|
|» token|string|true|none||none|

# contest

## GET 获取比赛信息

GET /api/contest/{contest_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|contest_id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "title": "string",
  "start_time": 0,
  "end_time": 0,
  "type": "string",
  "content": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» title|string|true|none||none|
|» start_time|integer|true|none||none|
|» end_time|integer|true|none||none|
|» type|string|true|none||none|
|» content|string|true|none||none|

## GET 获取比赛列表

GET /api/contest

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|p|query|integer| 否 |none|

> 返回示例

> 200 Response

```json
{
  "contests": [
    {
      "title": "string",
      "start_time": 0,
      "end_time": 0
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» contests|[object]|true|none||none|
|»» title|string|true|none||none|
|»» start_time|integer|true|none||none|
|»» end_time|integer|true|none||none|

## GET 获取题目列表

GET /api/contest/{contest_id}/problem

> Body 请求参数

```json
{
  "problems": [
    {
      "problem_id": "string",
      "title": "string"
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|contest_id|path|string| 是 |none|
|body|body|object| 否 |none|
|» problems|body|[object]| 是 |none|
|»» problem_id|body|string| 是 |none|
|»» title|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "problems": [
    {
      "problem_id": "string",
      "title": "string",
      "author": "string",
      "limit": {
        "time_limit": "string",
        "memory_limit": "string"
      },
      "difficulty": 9,
      "try": 0,
      "ac": 0
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» problems|[object]|true|none||none|
|»» problem_id|string|false|none||none|
|»» title|string|false|none||none|
|»» author|string|false|none||none|
|»» limit|object|false|none||none|
|»»» time_limit|string|true|none||none|
|»»» memory_limit|string|true|none||none|
|»» difficulty|integer|false|none||none|
|»» try|integer|false|none||none|
|»» ac|integer|false|none||none|

## GET 获取题目信息

GET /api/contest/{contest_id}/{problem_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|contest_id|path|string| 是 |none|
|problem_id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "problem_id": "string",
  "author": "string",
  "limit": {
    "time_limit": "string",
    "memory_limit": "string"
  },
  "difficulty": 9,
  "try": 0,
  "ac": 0,
  "title": "string",
  "content": "string",
  "example": [
    {
      "input": "string",
      "output": "string"
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» problem_id|string|true|none||none|
|» author|string|true|none||none|
|» limit|object|true|none||none|
|»» time_limit|string|true|none||none|
|»» memory_limit|string|true|none||none|
|» difficulty|integer|true|none||none|
|» try|integer|true|none||none|
|» ac|integer|true|none||none|
|» title|string|true|none||none|
|» content|string|true|none||none|
|» example|[object]|true|none||none|
|»» input|string|true|none||none|
|»» output|string|true|none||none|

# user

## GET 获取用户信息

GET /api/user/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "user_name": "string",
  "password": "string",
  "avator": "string",
  "email": "string",
  "content": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» user_name|string|true|none||none|
|» password|string|true|none||none|
|» avator|string|true|none||none|
|» email|string|true|none||none|
|» content|string|true|none||none|

## POST 修改用户权限

POST /api/user/changepermission

> Body 请求参数

```json
{
  "username": "zine",
  "permission": -1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» right|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# announcement

## POST 创建公告

POST /api/announcement

> Body 请求参数

```json
{
  "title": "喝唉一齐同样地好些哇稍微拍",
  "abstract": "in ut ut labore eu",
  "text": "做王空验个。从主口。资做说各完完议段法非。"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» abstract|body|string| 是 |none|
|» text|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取公告列表

GET /api/announcement

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取公告详细信息

GET /api/announcement/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 删除公告

DELETE /api/announcement/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PATCH 修改公告

PATCH /api/announcement/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_problem">problem</h2>

<a id="schemaproblem"></a>
<a id="schema_problem"></a>
<a id="tocSproblem"></a>
<a id="tocsproblem"></a>

```json
{
  "problem_id": "string",
  "author": "string",
  "limit": {
    "time_limit": "string",
    "memory_limit": "string"
  },
  "difficulty": 9,
  "try": 0,
  "ac": 0,
  "title": "string",
  "content": "string",
  "example": [
    {
      "input": "string",
      "output": "string"
    }
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|problem_id|string|false|none||none|
|author|string|false|none||none|
|limit|object|false|none||none|
|» time_limit|string|true|none||none|
|» memory_limit|string|true|none||none|
|difficulty|integer|false|none||none|
|try|integer|false|none||none|
|ac|integer|false|none||none|
|title|string|false|none||none|
|content|string|false|none||none|
|example|[object]|false|none||none|
|» input|string|true|none||none|
|» output|string|true|none||none|

