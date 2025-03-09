# 目录结构
OJ
├── backend
│   ├── Dockerfile
│   ├── requirements.txt
│   └── src
│       ├── app.py
│       ├── config
│       │   └── settings.py
│       ├── controllers
│       │   ├── admin.py
│       │   ├── auth.py
│       │   ├── problem.py
│       │   └── submit.py
│       ├── models
│       │   ├── problem.py
│       │   ├── submit.py
│       │   └── user.py
│       ├── routes
│       ├── services
│       └── utils
├── docker-compose.yml
├── docs
│   ├── API.md
│   ├── DEPLOYMENT.md
│   └── SECURITY.md
├── frontend
│   ├── Dockerfile
│   ├── public
│   └── src
├── judger
│   ├── Dockerfile
│   ├── src
│   │   ├── compiler
│   │   │   ├── c.py
│   │   │   ├── java.py
│   │   │   └── python.py
│   │   ├── executor.py
│   │   ├── judge.py
│   │   └── sandbox
│   │       ├── docker
│   │       └── seccomp
│   │           ├── c.json
│   │           ├── java.json
│   │           └── python.json
│   └── tests
│       ├── integration
│       ├── malicious_codes
│       └── unit
├── README.md
└── scripts
    ├── benchmark.sh
    ├── deploy.sh
    └── setup_judger.sh
