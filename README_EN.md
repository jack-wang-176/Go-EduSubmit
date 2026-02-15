
<div align="center">

# 📚 Go-EduSubmit

> **Go-EduSubmit** is a high-performance homework submission and grading system built with Go (Maple + GORM), Vue 3, and TypeScript. It provides clear and standardized RESTful APIs.
> It utilizes Docker for containerized deployment and offers convenient commands for direct download and execution.

**English Version** | [中文版本](./README.md)

</div>
<br>

---
##  Quick Start

###  1. Run directly via GitHub Packages
```bash
docker run -d -p 8080:8080 ghcr.io/jack-wang-176/maple-backend:v1.1

```

###  2. Deploy via Release

1. **Prerequisite** Ensure **Docker Desktop** is installed and running on your computer.
2. **Download Config** Download the `docker-compose.yml` file from the **Assets** section below to any directory.
3. **One-Click Start** Open a terminal (Terminal / CMD) in the directory where the file is located and execute the following command:
```bash
docker-compose up -d
```

##  Project Introduction

This project aims to implement a complete website with Docker deployment by combining **Go**, **Vue 3**, **Docker**, and other technologies, realizing the core business functions of a homework management system.

##  Features

###  Student Portal

* **Homework Submission**: Supports text content and file link submission.
* **Submission History (`/mysub`)**: View personal submission records.
* **View Excellent Homework**: Access highly-rated assignments.
* **API Features**: Returns nested, simplified `homework` information (title, department) while filtering out irrelevant fields.

###  Teacher/Admin Portal

* **Homework Overview (`/submission/homework/:id`)**: View all student submissions for a specific assignment.
* **API Features**: Returns nested `student` information (nickname, department) to facilitate identity verification for teachers.
* **Online Grading (`/submission/:id/review`)**: Supports scoring, commenting, and marking as excellent.
* **Technical Highlights**: Implements optimistic locking using the CAS (Compare-And-Swap) concept to ensure data consistency.
* **Excellent Homework Management (`/excellent`)**: One-click mark/unmark excellent assignments. Public display interface with department filtering support.

##  Tech Stack

* **Language**: Go (Golang) 1.25 | TypeScript
* **Web Framework**: [Maple](https://github.com/jack-wang-176/Maple)
* **Frontend Framework**: Vue 3
* **Docker**: maple-backend
* **ORM**: GORM v2
* **Database**: MySQL 8.0
* **Toolkit**: Custom error handling (pkg/errors)
##  Project Structure

Follows a standard Go project layered architecture with single responsibility and easy maintenance.

```text
.
├── backend/                # Backend Project (Go + Maple + GORM)
│   ├── cmd/                # Application Entry Point (main.go)
│   ├── dao/                # Data Access Layer (DAO)
│   ├── handler/            # Request Handling Layer (Controller/Handler)
│   ├── middleware/         # Middleware (Auth, CORS, etc.)
│   ├── model/              # Database Models (Structs)
│   ├── pkg/                # Common Utilities (Utils, Error handling)
│   ├── router/             # Router Configuration
│   ├── service/            # Business Logic Layer
│   ├── Dockerfile          # Backend Container Build File
│   └── go.mod              # Go Dependency Management
├── frontend/               # Frontend Project Directory
│   └── homework-frontend/  # Vue 3 Project Source Code
│       ├── public/         # Static Assets
│       ├── src/            # Frontend Source Code
│       │   ├── api/        # API Request Encapsulation
│       │   ├── assets/     # Static Assets (Images, CSS)
│       │   ├── components/ # Common Components
│       │   ├── layout/     # Layout Components
│       │   ├── router/     # Frontend Router Configuration
│       │   ├── utils/      # Utility Functions
│       │   ├── views/      # Page Views
│       │   ├── App.vue     # Root Component
│       │   ├── main.ts     # Entry File
│       │   └── style.css   # Global Styles
│       ├── Dockerfile      # Frontend Container Build File (Node build + Nginx)
│       ├── index.html      # HTML Entry Point
│       ├── nginx.conf      # Nginx Configuration
│       ├── package.json    # npm Dependency Configuration
│       └── vite.config.ts  # Vite Build Configuration
├── mysql-data/             # MySQL Data Volume Mount Directory (Auto-generated)
├── Assessment_API.md       # API Parameter Submission and Return Requirements
├── docker-compose.yml      # Docker Compose File (Core Entry)
├── LICENSE                 # Open Source License File
├── README.md               # Project Documentation (Chinese)
├── README_EN.md            # Project Documentation (English)
└── 接口测试.openapi.json    # API Definition File (OpenAPI 3.0)

```
##  Architecture & Design

###  Data Model

* 1. The project designs `user`, `submission`, and `homework` fields to abstract specific business entities. The corresponding data tables are created upon database initialization. To resolve the relationships between these three entities, I adopted the method of **embedded structs** to eliminate the need for creating intermediate tables. Since `gorm.Model` is used, all data tables utilize **soft deletion**.


* 2. In the specific implementation of structs, the `department` and `role` fields are mapped to integers (anchored to numbers) and built independently. Therefore, we need a `map` data structure to handle this correspondence. This requires converting the corresponding `string` to the corresponding number during frontend parameter transmission, while the business logic operates and stores data using these numeric wrappers.


* 3. In addition to business abstraction, I also designed return structs corresponding to the business logic and methods to convert them. This is detailed in the `Unified Response` section.



### Layered Architecture

* 1. The business level is divided into three layers: `dao`, `service`, and `handler`. The `dao` layer handles database operations; the `service` layer performs further data processing, validation, and assembly; the `handler` layer is for specific route mounting and handles data transfer within `context` and `json`.


* 2. `pkg` provides global utility functions, `middleware` handles authentication, `router` contains specific route designs, `model` stores business structs, and `cmd` is the program entry point.


* 3. We use internal structs and global variables (based on these structs) to manage the corresponding functions of each layer. This ensures that when making logical choices, we first select the variable and then the function, creating a strict separation between layers.



###  Error Handling

* 1. Our error handling is divided into two categories: **Business Errors** and **Concrete Errors**. Business errors have custom error messages and codes defined by us, while Concrete errors are encapsulated and propagated as-is. We designed a specific error structure specification where errors are classified based on the first digit of the status code:


* Type 1: `handler` level errors (mainly related to data transfer).
* Type 2: `user` errors.
* Type 3: `homework` errors.
* Type 4: `submission` errors.


* 2. Data from the `dao` layer is propagated as-is. The `service` layer judges whether an error belongs to our defined business errors. The `handler` layer propagates errors directly, except for Type 1 errors (e.g., parameter binding errors).



###  Unified Response

* 1. In the `handler`, I designed a specific `sendResponse` struct/function to throw corresponding business responses, using variadic parameters to meet specific reply requirements.


* 2. Corresponding response structs are set within the structure. Due to project requirements, the `submission` layer is difficult to unify, so specific response structs are used for the other two. In this layer, temporary structs are wrapped within specific routes to meet corresponding needs.


* 3. In the standard return function, we check if it is a business logic error. If it is, it returns directly; if not, it logs the error before returning.



###  Data Transmission & Context

* 1. Information is transmitted in two ways: via **Request Headers** and via **Request Parameters**.


* 2. Regarding the first method: We obtain the `accessToken` in the `login` route. Since specific content is declared when creating the `token`, in other routes, we set up a global middleware to parse the `token` and call the `set` method to store it. Thus, the corresponding user information is stored in an encrypted header (context), which serves the purpose of user information storage and encryption verification.


* 3. Regarding the second method: We can pass data through dynamic parameters in the `url` or via `params`.



### Token

* 1. We have discussed the role of `token` in data transmission; now let's discuss its entire lifecycle.


* 2. Token creation happens at `login`. The `accessToken` is carried in the header. The `refreshToken` is used for refreshing (providing a refresh interface that does not require verification) and is stored in the database. When setting a new `refreshToken`, the previous `token` is marked as **expired**.

##  API Overview

### [Detailed API Requirements](https://www.google.com/search?q=Assessment_API.md)

### Example: Get My Submissions

* **Endpoint**: `GET /submission/mysub`
* **Response**:

```json
{
  "code": 0,
  "data": {
    "list": [
      {
        "id": 100,
        "homework": { "id": 1, "title": "Go Concurrency", "department_label": "Backend" },
        "score": 95,
        "submitted_at": "2026-02-07 18:00:00"
      }
    ],
    "total": 1
  }
}

```

### 3. Get Excellent Gallery

* **Endpoint**: `GET /submission/excellent`
* **Query**: `page=1&page_size=10`
* **Response**: Includes complete nested information for both `homework` (assignment) and `student` (author).

##  API Documentation

This project is designed based on the OpenAPI 3.0 specification. To facilitate debugging for developers, we provide a complete interface definition file.

### Method 1: Local Import (Recommended)

The repository includes the exported API specification file, which supports direct import into **Postman**, **Apifox**, or **Swagger UI**.

*  **Interface Definition File**: [接口测试.openapi.json](https://www.google.com/search?q=%E6%8E%A5%E5%8F%A3%E6%B5%8B%E8%AF%95.openapi.json)
  *(Click the link to view source directly, or right-click and "Save As" to download)*

**How to use:**

1. Download the `openapi.json` file.
2. Open Postman / Apifox.
3. Select `Import` -> Drag and drop the file to generate a complete interface debugging environment.

### Method 2: Online Preview

---

##  Contribution

Issues and Pull Requests are welcome to improve this project!

##  License

[MIT License](https://www.google.com/search?q=LICENSE)