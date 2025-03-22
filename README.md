# FashOJ
---
```
FashOJ
├── .github/
│   └── workflows/
│       └── greetings.yml
├── .gitignore
├── backend/
│   ├── config/
│   │   ├── config.go
│   │   ├── config.yaml
│   │   └── db.go
│   ├── controllers/
│   │   ├── auth.go
│   │   ├── problem.go
│   │   └── user.go
│   ├── global/
│   │   └── global.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middlewares/
│   │   └── auth.go
│   ├── models/
│   │   ├── problem.go
│   │   └── user.go
│   ├── router/
│   │   └── router.go
│   └── utils/
│       └── utils.go
├── config.yaml
├── docs/
│   ├── API.md
│   ├── DEPLOYMENT.md
│   └── SECURITY.md
├── go.mod
├── judger/
│   ├── compiler/
│   │   ├── c.go
│   │   ├── java.go
│   │   └── python.go
│   └── sandbox/
│       └── seccomp/
│           ├── c.json
│           ├── java.json
│           └── python.json
├── LICENSE
├── README.md
└── scripts/
    ├── benchmark.sh
    ├── deploy.sh
    └── setup_judger.sh
```