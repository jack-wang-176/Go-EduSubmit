# === 第一阶段：构建阶段 (Builder) ===
FROM node:20-alpine AS builder

WORKDIR /app

# 1. 复制依赖文件
COPY package*.json ./

# 2. 设置淘宝镜像源 (关键！否则国内下载依赖可能会失败或极慢)
RUN npm config set registry https://registry.npmmirror.com

# 3. 安装依赖
RUN npm install

# 4. 复制源代码并打包
COPY . .
RUN npm run build

# === 第二阶段：生产阶段 (Production) ===
FROM nginx:stable-alpine

# 1. 复制 Nginx 配置文件
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 2. 从第一阶段把 dist 文件夹里的东西“偷”过来
COPY --from=builder /app/dist /usr/share/nginx/html

# 3. 暴露端口
EXPOSE 80

# 4. 启动
CMD ["nginx", "-g", "daemon off;"]