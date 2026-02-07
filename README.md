

# 📚 Go-EduSubmit

> **Go-EduSubmit** 是一个基于 Go (Gin + GORM) 构建的高性能作业提交与批改系统后端。它专注于解决复杂的关联数据查询与并发批改冲突问题，提供清晰、规范的 RESTful API。

[English](https://www.google.com/search?q=%23-english-introduction) | [中文介绍](https://www.google.com/search?q=%23-%E9%A1%B9%E7%9B%AE%E4%BB%8B%E7%BB%8D)

---

## 📖 项目介绍

本项目旨在通过go语言构建职责分明的多层级系统

* **数据清洗 (DTO)**：在 Handler 层实现了灵活的数据转换（Data Transfer Object），拒绝向前端返回数据库的“脏数据”（如零值时间、空对象），只返回前端需要的嵌套结构（如 `homework` 与 `student` 信息）。
* **并发控制 (Optimistic Locking)**：在批改作业时引入**乐观锁**机制（基于 `version` 版本号），有效防止多名助教同时修改同一份作业导致的数据覆盖问题。
* **关联查询优化**：运用 GORM 的 `Preload` 机制，高效处理 User、Homework、Submission 三表关联。

## ✨ 核心功能 (Features)

### 👨‍🎓 学生端 (Student)

* **作业提交**：支持文本内容与文件链接提交。
* **提交历史 (`/mysub`)**：
* 查看自己的提交记录。
* **接口特点**：返回结果中嵌套简化的 `homework` 信息（标题、部门），屏蔽无关字段。



### 👩‍🏫 教师/管理员端 (Admin)

* **作业概览 (`/submission/homework/:id`)**：
* 查看特定作业下的所有学生提交。
* **接口特点**：返回结果中嵌套 `student` 信息（昵称、部门），方便老师确认身份。


* **在线批改 (`/submission/:id/review`)**：
* 支持评分 (Score)、评语 (Comment) 及优秀标记。
* **技术亮点**：使用 CAS (Compare-And-Swap) 思想实现乐观锁，保障数据一致性。


* **优秀作业管理 (`/excellent`)**：
* 一键标记/取消优秀作业。
* 公共展示接口，支持按部门筛选。



## 🛠 技术栈 (Tech Stack)

* **语言**: Go (Golang) 1.25
* **Web 框架**: Gin
* **ORM**: GORM v2
* **数据库**: MySQL 8.0
* **配置**: Viper (可选)
* **工具包**: 自定义错误处理 (pkg/errors)

## 📂 项目结构 (Structure)

遵循标准的 Go 项目分层架构，职责单一，易于维护。

```text
.
├── cmd/                # 启动入口 (Main application entry)
├── dao/                # 数据访问层 (GORM operations & SQL)
├── handler/            # 控制层 (Request binding, DTO conversion, Response)
├── model/              # 数据库模型 (Database structs & hooks)
├── pkg/                # 公共工具 (Error codes, Response wrapper)
├── router/             # 路由定义 (Gin router setup)
└── service/            # 业务逻辑层 (Business logic & Transactions)

```
## 项目设计(idea)
### 结构体(struct)
* 1.本项目设计了`user` `submission` `homework`字段来抽象具体业务，对应数据表在初始化数据库时创建，
为了解决三个实体间的对应关系，我在这里采用了内嵌其他结构体方法，这样来省略创建中间表，
在这里使用了`gorm.model`，因此所有的数据表都是采用的软删除的方式。
* 2.在结构体的具体实现上，`department` 和 `role` 字段和数字锚定，本身成为独立构建，
因此我们需要用`map`数据结构来对这种对应关系进行锚定，这就要求在前端进行传参的时候都要将对应`string`转换为对应数字，
而在我们的业务逻辑中，使用对应数字的包装进行操作和数据储存
* 3.除了业务抽象外，我还设计了与业务逻辑对应的返回结构体，并且设计对应方法对其进行转化，
这一点我们在`response`处进行详细说明。
### 分层思路(idea)
* 1.在业务层次上分为三层`dao` `service` `handler`，其中`dao`层是对数据库进行操作，`service`是对数据做进一步处理和验证组装，
`handler`是具体的路由挂载，用于处理`context`和`json`中的数据传递。
* 2.`pkg`提供了全局性的函数工具,`middleware`是鉴权的中间件，`router`是具体的路由挂载设计,
`model`层储存业务结构体,`cmd`是程序的入口。
* 3.具体每一层的对应函数我们使用包内部的结构体和全局性的变量(基于这个结构体)来进行管理，这样使得我们在进行逻辑选择时首先选择对应变量再选择对应函数，
这样一种方式让各个层次之间严密分明
### 错误上抛
* 1.在这里我们的处理分为两类，一类是业务错误，一类是具体错误，业务错误由我们自己定义错误信息和错误码，
而具体错误我们封装了对应方法进行原样抛出，在这里我们设计了具体的错误结构规范，错误根据状态码的第一位进行分类处理，
一类是`handler`级别的错误，这类错误主要和数据传递联系在一块，二类是`user`，三类是`homework`,四类是`submission`
* 2.对于`dao`层的数据我们原样抛出，`service`层中要判断是否属于我们划定的业务错误，
`handler`层中除了一类错误(例如参数绑定出错)，其他直接抛出
### 业务回应(response)
* 1.在`handler`中我设计了具体的`sendResponse`结构体来进行相应业务抛出，采用不定参数来满足特定的回复信息要求
* 2.在结构体里面我们设置了对应相应结构体，由于项目要求，在`submisson`层中难以统一，因此只有另两个结构体我们使用相应结构体，
在这个层里面在每一个具体路由里面包装临时结构体来满足相应需求
* 3.在标准的返回函数里面，我们要去判断是否是业务逻辑错误，如果是业务逻辑那么直接返回，如果不是那么日志记录并返回
### 信息传递(information transformation)
* 1.在这里我们进行信息传递的方式有两种，一种是在通过请求头进行传递，另一种是通过请求参数进行传递，
* 2.首先来谈论第一种，我们在`login`路由里面获得`accessToken`，因为我们在进行创建`token`的时候就声明了具体的内容，
在其他路由中，我们都在这之前设置了全局性的中间键对`token`进行解析而调用`set`方法进行储存，
这样对应的用户信息就储存在一个加密头中，这个加密头本身提供了用户信息储存和加密验证的作用
* 3.再来讨论第二种，我们可以在`url`通过动态参数进行传递或者通过`params`进行参数传递。
### token
* 1.之前我们已经讨论了`token`在信息传递上的作用，接下来我们来讨论一下其整个生命周期
* 2.`token`的创建，通过`login`获取对应的`token`其中`accessToken`被我们带在相应头中，
`refreshToken`用作刷新，并提供不用验证的刷新接口，其中`refreshToken`在数据库中进行储存，
在新的`refreshToken`设置之前的`token`通过对应的字段标记为国企



## 📝 API 接口概览 (API Overview)

### 1. 获取我的提交 (My Submissions)

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



### 2. 批改作业 (Review Submission)

* **Endpoint**: `PUT /submission/:id/review`
* **Body**:
```json
{
  "score": 90,
  "comment": "Good job!",
  "is_excellent": true
}

```


* **Logic**: 检查版本号，原子更新分数与状态。

### 3. 获取优秀作业广场 (Excellent Gallery)

* **Endpoint**: `GET /submission/excellent`
* **Query**: `page=1&page_size=10`
* **Response**: 同时包含 `homework` (题目) 和 `student` (作者) 的完整嵌套信息。

## 🚀 快速开始 (Getting Started)

1. **克隆仓库**
```bash
git clone https://github.com/your-username/Go-EduSubmit.git

```


2. **配置数据库**
   在 MySQL 中创建数据库，并修改配置文件中的 DSN。
3. **运行项目**
```bash
go mod tidy
go run cmd/main.go

```


*项目启动时会自动运行 `AutoMigrate` 迁移数据库表结构。*
4. **数据库补丁 (Optional)**
   如果遇到 `reviewed_at` 字段缺失报错，请执行：
```sql
ALTER TABLE submissions ADD COLUMN reviewed_at DATETIME NULL COMMENT '批改时间';

```



## 🤝 贡献 (Contribution)

欢迎提交 Issue 或 Pull Request 来改进本项目！

## 📄 许可证 (License)

[MIT License](https://www.google.com/search?q=LICENSE)
