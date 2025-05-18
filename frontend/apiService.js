import { serviceConfig } from './config.js';
import { ErrorHandler } from './errorHandler.js';

async function handleResponse(response) {
    const data = await response.json().catch(() => ({}));
    if (!response.ok) {
        const errorMessage = data.message || 
                            (data.error && typeof data.error === 'string' ? data.error : '请求失败');
        const errorDetails = data.details ? `\n详情: ${JSON.stringify(data.details)}` : '';
        throw new Error(`${errorMessage}${errorDetails}`);
    }
    return data;
}

export async function fetchMenuData() {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/get_dishes`);
        const data = await handleResponse(response);
        
        // 按分类组织数据
        return data.reduce((acc, item) => {
            if (!acc[item.Category]) acc[item.Category] = [];
            acc[item.Category].push({
                id: item.ID,
                name: item.Name,
                price: item.Price
            });
            return acc;
        }, {});
    } catch (error) {
        ErrorHandler.showError(error.message);
        return {};
    }
}

export async function calculateTotalPrice(cartItems) {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/get_total_price`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(cartItems.map(item => ({
                Price: item.price,
                DishID: item.id
            })))
        });
        
        const data = await handleResponse(response);
        return data.total_price;
    } catch (error) {
        ErrorHandler.showError(error.message);
        return 0;
    }
}
