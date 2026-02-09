# === 第一阶段：构建阶段 (Builder) ===
FROM golang:alpine AS builder

# 1. 设置环境变量 (开启 Module，设置国内代理加速)
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 2. 在容器内创建工作目录
WORKDIR /app

# 3. 先复制依赖文件 (利用缓存机制)
COPY go.mod go.sum ./
# 4. 下载依赖
RUN go mod download

# 5. 复制所有源码
COPY . .

# 6. 编译代码！(关键修改：指向 cmd 目录)
# 你的 main.go 在 cmd 文件夹下，所以要指定 ./cmd
RUN go build -o main ./cmd

# === 第二阶段：运行阶段 (Runner) ===
# 使用超小的 Alpine 系统
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从第一阶段把编译好的 main 文件“偷”过来
COPY --from=builder /app/main .

# 如果你有 config.yaml，以后记得也要 COPY 进去，现在先不用
# COPY --from=builder /app/config.yaml .

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./main"]