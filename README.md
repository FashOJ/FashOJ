# FashOJ
---
```
FashOJ
├── backend                # 后端服务（Golang）
│   ├── cmd               # 入口文件
│   │   ├── main.go       # 主程序入口
│   │   └── judge.go      # 评测服务入口
│   ├── config            # 配置文件
│   │   ├── config.yaml   # 主配置文件
│   │   └── config.go     # 解析配置
│   ├── controllers       # 处理HTTP请求
│   │   ├── admin.go      
│   │   ├── auth.go       
│   │   ├── problem.go    
│   │   ├── submit.go     
│   │   └── user.go       
│   ├── models            # 数据库模型
│   │   ├── problem.go    
│   │   ├── submit.go     
│   │   └── user.go       
│   ├── repository        # 数据库操作
│   │   ├── problem_repo.go
│   │   ├── submit_repo.go
│   │   ├── user_repo.go
│   │   └── redis_repo.go
│   ├── routes            # 路由
│   │   ├── router.go     
│   │   ├── api.go        
│   │   └── admin.go      
│   ├── services          # 业务逻辑
│   │   ├── problem.go    
│   │   ├── submit.go     
│   │   ├── user.go       
│   │   ├── judge.go      # 评测逻辑
│   │   └── auth.go       
│   ├── utils             # 工具类
│   │   ├── logger.go     
│   │   ├── response.go   
│   │   ├── config.go     
│   │   └── redis.go      
│   ├── middleware        # 中间件
│   │   ├── auth.go       
│   │   ├── cors.go       
│   │   └── logging.go    
│   ├── tests             # 测试
│   │   ├── integration   
│   │   ├── unit          
│   │   ├── auth_test.go  
│   │   ├── problem_test.go
│   │   └── submit_test.go
│   ├── go.mod            # 依赖管理
│   └── go.sum            
├── frontend              # 前端（Vue）
│   ├── public           
│   ├── src              
│   │   ├── assets        
│   │   ├── components    
│   │   ├── views         
│   │   ├── router        
│   │   ├── store         
│   │   ├── utils         
│   │   └── App.vue       
│   ├── package.json     
│   ├── vite.config.js   
│   └── tsconfig.json    
├── judger                # 代码评测系统
│   ├── compiler          # 代码编译
│   │   ├── c.go         
│   │   ├── java.go      
│   │   └── python.go    
│   ├── executor.go       # 代码执行
│   ├── judge.go          # 评测逻辑
│   ├── sandbox           # 沙箱
│   │   ├── seccomp       
│   │   │   ├── c.json    
│   │   │   ├── java.json 
│   │   │   └── python.json
│   ├── tests            
│   └── README.md        
├── scripts               # 自动化脚本
│   ├── deploy.sh        
│   ├── setup_judger.sh  
│   └── benchmark.sh     
├── docs                  # 文档
│   ├── API.md           
│   ├── DEPLOYMENT.md    
│   └── SECURITY.md      
├── config.yaml           # 全局配置
├── README.md             # 项目介绍
├── LICENSE               # 开源协议
└── .gitignore            # Git忽略文件
```