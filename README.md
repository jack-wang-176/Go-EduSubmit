
<div align="center">

# 📚 Go-EduSubmit

> **Go-EduSubmit** 是一个基于Go (Maple + GORM)、Vue 3、TypeScript构建的高性能作业提交与批改系统。提供清晰、规范的 RESTFUL API。
> 采用 Docker 实现项目的容器化部署，并提供了便捷的指令来实现项目的直接下载运行

[English Version](./README_EN.md) | **中文版本**

</div>
<br>

---
##  快速开始 (Getting Started)

* 前往release版本里面下载`user-docker-compose.yml`文件
* 在对应文件下的目录里运行一下指令
```bash
   docker-compose up -d
```
* 在浏览器中访问`http://localhost`


##  项目介绍
本项目旨在通过结合`Go` `Vue 3` `Docker`等多种技术栈，实现完整的网站实现和docker部署，实现了作业管理系统的核心业务功能


##  业务功能 (Features)

###  学生端 (Student)

* **作业提交**：支持文本内容与文件链接提交。
* **提交历史 (`/mysub`)**： 查看自己的提交记录。
* **获取优秀作业**
* **接口特点**：返回结果中嵌套简化的 `homework` 信息（标题、部门），屏蔽无关字段。


###  教师/管理员端 (Admin)

* **作业概览 (`/submission/homework/:id`)**： 查看特定作业下的所有学生提交。
* **接口特点**：返回结果中嵌套 `student` 信息（昵称、部门），方便老师确认身份。

* **在线批改 (`/submission/:id/review`)**： 支持评分 (Score)、评语 (Comment) 及优秀标记。
* **技术亮点**：使用 CAS (Compare-And-Swap) 思想实现乐观锁，保障数据一致性。

* **优秀作业管理 (`/excellent`)**： 一键标记/取消优秀作业。 公共展示接口，支持按部门筛选。

##  技术栈 (Tech Stack)

* **语言**: Go (Golang) 1.25 | TypeScript
* **Web 框架**: [Maple](https://github.com/jack-wang-176/Maple) 
* **前端框架**:Vue 3
* **Docker**: maple-backend
* **ORM**: GORM v2
* **数据库**: MySQL 8.0
* **工具包**: 自定义错误处理 (pkg/errors)

##  项目结构 (Structure)

遵循标准的 Go 项目分层架构，职责单一，易于维护。

```text
.
├── backend/                # 后端项目 (Go + Maple + GORM)
│   ├── cmd/                # 程序入口 (main.go)
│   ├── dao/                # 数据访问层 (Database Access Object)
│   ├── handler/            # 请求处理层 (Controller/Handler)
│   ├── middleware/         # 中间件 (Auth, CORS, etc.)
│   ├── model/              # 数据库模型 (Structs)
│   ├── pkg/                # 公共工具包 (Utils, Error handling)
│   ├── router/             # 路由配置
│   ├── service/            # 业务逻辑层
│   ├── Dockerfile          # 后端容器构建文件
│   └── go.mod              # Go 依赖管理
├── frontend/               # 前端项目目录
│   └── homework-frontend/  # Vue3 项目源码
│       ├── public/         # 静态资源
│       ├── src/            # 前端源代码
│       │   ├── api/        # 接口请求封装
│       │   ├── assets/     # 静态资源 (Images, CSS)
│       │   ├── components/ # 公共组件
│       │   ├── layout/     # 页面布局组件
│       │   ├── router/     # 前端路由配置
│       │   ├── utils/      # 工具函数
│       │   ├── views/      # 页面视图 (Page Views)
│       │   ├── App.vue     # 根组件
│       │   ├── main.ts     # 入口文件
│       │   └── style.css   # 全局样式
│       ├── Dockerfile      # 前端容器构建文件 (Node build + Nginx)
│       ├── index.html      # HTML 入口
│       ├── nginx.conf      # Nginx 配置文件
│       ├── package.json    # npm 依赖配置
│       └── vite.config.ts  # Vite 构建配置
├── mysql-data/             # MySQL 数据卷挂载目录 (自动生成)
├── Assessment_API.md       # 接口参数提交和返回要求
├── docker-compose.yml      # Docker 编排文件 (核心入口)
├── LICENSE                 # 开源协议文件
├── README.md               # 项目说明文档 (中文)
├── README_EN.md            # 项目说明文档 (英文)
└── 接口测试.openapi.json    # API 接口定义文件 (OpenAPI 3.0)

```
## 项目架构设计 (Architecture & Design)
###  数据模型 (Data Model)
* 1.本项目设计了`user` `submission` `homework`字段来抽象具体业务，对应数据表在初始化数据库时创建，
为了解决三个实体间的对应关系，我在这里采用了内嵌其他结构体方法，这样来省略创建中间表，
在这里使用了`gorm.model`，因此所有的数据表都是采用的软删除的方式。
* 2.在结构体的具体实现上，`department` 和 `role` 字段和数字锚定，本身成为独立构建，
因此我们需要用`map`数据结构来对这种对应关系进行锚定，这就要求在前端进行传参的时候都要将对应`string`转换为对应数字，
而在我们的业务逻辑中，使用对应数字的包装进行操作和数据储存
* 3.除了业务抽象外，我还设计了与业务逻辑对应的返回结构体，并且设计对应方法对其进行转化，
这一点我们在`response`处进行详细说明。
###  分层架构 (Layered Architecture)
* 1.在业务层次上分为三层`dao` `service` `handler`，其中`dao`层是对数据库进行操作，`service`是对数据做进一步处理和验证组装，
`handler`是具体的路由挂载，用于处理`context`和`json`中的数据传递。
* 2.`pkg`提供了全局性的函数工具,`middleware`是鉴权的中间件，`router`是具体的路由挂载设计,
`model`层储存业务结构体,`cmd`是程序的入口。
* 3.具体每一层的对应函数我们使用包内部的结构体和全局性的变量(基于这个结构体)来进行管理，这样使得我们在进行逻辑选择时首先选择对应变量再选择对应函数，
这样一种方式让各个层次之间严密分明
###  错误处理 (Error Handling)
* 1.在这里我们的处理分为两类，一类是业务错误，一类是具体错误，业务错误由我们自己定义错误信息和错误码，
而具体错误我们封装了对应方法进行原样抛出，在这里我们设计了具体的错误结构规范，错误根据状态码的第一位进行分类处理，
一类是`handler`级别的错误，这类错误主要和数据传递联系在一块，二类是`user`，三类是`homework`,四类是`submission`
* 2.对于`dao`层的数据我们原样抛出，`service`层中要判断是否属于我们划定的业务错误，
`handler`层中除了一类错误(例如参数绑定出错)，其他直接抛出
###  统一响应 (Unified Response)
* 1.在`handler`中我设计了具体的`sendResponse`结构体来进行相应业务抛出，采用不定参数来满足特定的回复信息要求
* 2.在结构体里面我们设置了对应相应结构体，由于项目要求，在`submisson`层中难以统一，因此只有另两个结构体我们使用相应结构体，
在这个层里面在每一个具体路由里面包装临时结构体来满足相应需求
* 3.在标准的返回函数里面，我们要去判断是否是业务逻辑错误，如果是业务逻辑那么直接返回，如果不是那么日志记录并返回
###  数据传输与上下文 (Data Transmission)
* 1.在这里我们进行信息传递的方式有两种，一种是在通过请求头进行传递，另一种是通过请求参数进行传递，
* 2.首先来谈论第一种，我们在`login`路由里面获得`accessToken`，因为我们在进行创建`token`的时候就声明了具体的内容，
在其他路由中，我们都在这之前设置了全局性的中间键对`token`进行解析而调用`set`方法进行储存，
这样对应的用户信息就储存在一个加密头中，这个加密头本身提供了用户信息储存和加密验证的作用
* 3.再来讨论第二种，我们可以在`url`通过动态参数进行传递或者通过`params`进行参数传递。
### token
* 1.之前我们已经讨论了`token`在信息传递上的作用，接下来我们来讨论一下其整个生命周期
* 2.`token`的创建，通过`login`获取对应的`token`其中`accessToken`被我们带在相应头中，
`refreshToken`用作刷新，并提供不用验证的刷新接口，其中`refreshToken`在数据库中进行储存，
在新的`refreshToken`设置之前的`token`通过对应的字段标记为过期

  
##  API 接口概览 (API Overview)

### [具体接口要求](Assessment_API.md)
### 例子：获取我的提交 (My Submissions)

* **Endpoint**: `GET /submission/mysub`
* **Response**:
```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 100,
        "homework": { "id": 1, "title": "Go Concurrency", "department_label": "后端" },
        "score": 95,
        "submitted_at": "2026-02-07 18:00:00"
      }
    ],
    "total": 1
  }
}

```



### 3. 获取优秀作业广场 (Excellent Gallery)

* **Endpoint**: `GET /submission/excellent`
* **Query**: `page=1&page_size=10`
* **Response**: 同时包含 `homework` (题目) 和 `student` (作者) 的完整嵌套信息。

##  接口测试文档 (API Documentation)

本项目基于 OpenAPI 3.0 规范设计。为了方便开发者调试，我们提供了完整的接口定义文件。

### 方式 1：本地导入 (推荐)
仓库中已包含导出的 API 规范文件，支持直接导入 **Postman**、**Apifox** 或 **Swagger UI**。

*  **接口定义文件**：[接口测试.openapi.json](接口测试.openapi.json)
  *(点击链接可直接查看源码，或右键 "另存为" 下载)*

**如何使用：**
1.  下载 `openapi.json` 文件。
2.  打开 Postman / Apifox。
3.  选择 `Import` (导入) -> 拖入该文件即可生成完整的接口调试环境。

### 方式 2：在线预览

[![Apifox Docs](https://img.shields.io/badge/Apifox-在线文档-FF4400?style=flat&logo=apifox&logoColor=white)](https://s.apifox.cn/e4106f50-0404-4c41-81bc-f45308f92ccb)

---




##  贡献 (Contribution)

欢迎提交 Issue 或 Pull Request 来改进本项目！

##  许可证 (License)

[MIT License](LICENSE)
