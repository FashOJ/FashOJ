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
generator: "@tarslib/widdershins v4.0.29"

---

# FashOJ

Base URLs:

# Authentication

# normal

## GET 获取题目信息

GET /api/problem/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 否 |none|


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

## POST 创建与修改题目

POST /api/problem

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|


> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取所有提交信息

GET /api/submit

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|p|query|integer| 否 |page|
|Authorization|header|string| 否 |none|


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
|Authorization|header|string| 否 |none|

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

## DELETE 删除题目

DELETE /api/problem/{problem_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|problem_id|path|string| 是 |none|
|Authorization|header|string| 否 |none|


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

# user

## POST 用户登录

POST /api/auth/login

> Body 请求参数

```json
{}
```

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

## GET 用户注册

GET /api/auth/register

> Body 请求参数

```json
{}
```

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

## GET 获取用户信息

GET /api/auth/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 否 |none|

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

# contest

## GET 获取比赛信息

GET /api/contest/{contest_id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|contest_id|path|string| 是 |none|
|Authorization|header|string| 否 |none|

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
|Authorization|header|string| 否 |none|

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

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|contest_id|path|string| 是 |none|
|Authorization|header|string| 否 |none|


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
|Authorization|header|string| 否 |none|

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

