# 使用Node官方镜像作为基础镜像
FROM node:22.15-alpine

# 设置工作目录
WORKDIR /app

# 复制package.json和package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install

# 复制所有文件
COPY . .

# 暴露Vite开发服务器端口
EXPOSE 5173

# 启动开发服务器
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
# npm run dev -- --host 0.0.0.0
# npm run dev 执行package.json中的dev脚本
#   这里dev脚本是: "vite"
# -- 分隔符， 表示后面的参数传递给dev脚本
# --host 0.0.0.0  设置vite服务器监听0.0.0.0，也就是所有网络
