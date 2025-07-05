# 餐厅点餐系统

## 项目概述
一个完整的餐厅点餐系统，包含前端点餐界面和后端管理系统。

## 功能说明
### 前端功能
- 菜单分类展示
- 菜品搜索
- 购物车管理（添加/删除/调整数量）
- 价格计算

### 后端功能
- 菜品管理（增删改查）
- 订单处理
- 价格计算API

## 技术栈
### 前端
- HTML5/CSS3
- JavaScript (ES6+)
- Fetch API
- Vue.js (Vite)

### 后端
- Go (Gin框架)
- MySQL/PostgreSQL

## 项目结构
```
restaurant_app/
├── frontend/        # 前端代码
│   ├── app.js       # 主应用逻辑
│   ├── apiService.js # API服务
│   └── style.css    # 样式表
├── controller/      # 控制器
│   ├── dish_controller.go # 菜品控制器
│   └── router.go    # 路由配置
└── config/          # 配置管理
```

## API文档
### 菜品相关
- `GET /api/get_dishes` - 获取所有菜品
- `GET /api/get_dish/:id` - 获取单个菜品
- `POST /api/get_total_price` - 计算总价

### 管理接口
- `POST /admin/add_dish` - 添加菜品
- `POST /admin/update_dish` - 更新菜品
- `POST /admin/delete_dish` - 删除菜品

## 部署指南
1. 后端部署：
```bash
go build -o restaurant_app
./restaurant_app
```

2. 前端部署：
```bash
cd frontend
# http-server --cors -p 5173
npm run dev
```

## 开发说明
- 前端开发：修改frontend目录下文件
- 后端开发：修改controller和config目录下文件
- 样式修改：编辑frontend/style.css


## Docker部署
docker compose up --build