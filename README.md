# FashOJ
---
```
FashOJ
├── .github
│   ├── workflows
│   │   └── greetings.yml
├── .gitignore
├── LICENSE
├── README.md
├── backend
│   ├── config
│   │   ├── config.go
│   │   ├── config.yaml
│   │   ├── db.go
│   ├── controllers
│   │   ├── auth.go
│   │   ├── problem.go
│   │   ├── user.go
│   ├── global
│   │   ├── global.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middlewares
│   │   ├── auth.go
│   ├── models
│   │   ├── problem.go
│   │   ├── user.go
│   ├── router
│   │   ├── router.go
│   ├── utils
│   │   ├── logger.go
│   │   └── utils.go
├── config.yaml
├── docs
│   ├── API.md
│   ├── DEPLOYMENT.md
│   ├── SECURITY.md
├── judger
│   ├── cmd
│   │   ├── main.go
│   ├── go.mod
│   ├── internal
│   │   ├── config
│   │   ├── judge
│   │   │   ├── comparer.go
│   │   │   ├── complier.go
│   │   │   ├── executor.go
│   │   │   ├── judger.go
│   │   ├── sandbox
│   │   ├── utils
│   ├── testTEMP
│   │   ├── execTemp
│   │   ├── outputCase
│   │   │   ├── output1.txt
│   │   ├── sourceCode
│   │   │   ├── test.cpp
│   │   └── testCase
│   │       └── input1.txt
└── scripts
    ├── benchmark.sh
    ├── deploy.sh
    └── setup_judger.sh
```