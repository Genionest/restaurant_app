// 服务配置
export const serviceConfig = {
    frontend: {
        host: 'localhost',
        port: 5173
    },
    backend: {
        apiBaseUrl: 'http://127.0.0.1:8080' // 默认后端地址
    }
};

// 更新配置函数
export function updateServiceConfig(newConfig) {
    Object.assign(serviceConfig, newConfig);
}

// 获取前端服务地址
export function getFrontendUrl() {
    return `http://${serviceConfig.frontend.host}:${serviceConfig.frontend.port}`;
}
