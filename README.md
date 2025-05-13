# FashOJ - 编程竞赛平台

![Static Badge](https://img.shields.io/badge/vue-3-blue)
![Static Badge](https://img.shields.io/badge/go-1.24.1-blue)

## 项目简介

FashOJ 是一个高性能的在线编程竞赛平台，提供题目评测、比赛管理和用户排名功能。

## 功能特性

-   用户认证系统
-   题目管理与提交
-   实时评测与结果返回
-   比赛创建与管理
-   用户排名系统

## 技术栈

-   后端: Go (Gin 框架)
-   判题机: Go 独立模块
-   前端: Vue.js
-   数据库: MySQL

## 快速开始

```bash
# 克隆项目
git clone https://github.com/your-repo/FashOJ.git
cd FashOJ

# 启动服务
./scripts/deploy.sh
```

## 支持语言

-   C++
-   Python
-   Java

## 文档

-   [API 接口文档](docs/API.md)
-   [部署指南](docs/DEPLOYMENT.md)
-   [安全规范](docs/SECURITY.md)

<!-- ## 贡献指南

欢迎提交 Pull Request 或 Issue 报告问题。贡献前请阅读：

1. 遵循现有代码风格
2. 确保通过所有测试
3. 更新相关文档 -->

## 贡献指南
非常欢迎大家为FashOJ贡献代码！如果你想参与项目开发，可以参考以下步骤：
1.  Fork本仓库到你自己的GitHub账号。
2.  克隆你Fork后的仓库到本地：
```bash
git clone https://github.com/[yourname]/FashOJ.git
```
3.  创建一个新的分支：
```bash
git checkout -b [分支名称]
```
4.  在新分支上进行代码修改和功能开发。请确保你的代码遵循项目现有的代码风格和规范。
5.  完成修改后，提交你的代码并推送到远程仓库：
```bash
git add .
git commit -m "message"
git push origin [分支名称]
```
6.  回到GitHub，在你的仓库页面发起Pull Request，详细描述你的改动内容和目的。我们会尽快对你的PR进行审核和反馈。

在贡献代码之前，建议先阅读项目的 [行为准则](https://ys.mihoyo.com/main/) ，确保你的贡献符合项目的整体理念和社区规范。

## 许可证

[MIT License](LICENSE)
