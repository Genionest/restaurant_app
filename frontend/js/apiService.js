import { serviceConfig } from './config.js';
import { ErrorHandler } from './errorHandler.js';

async function handleResponse(response) {
    const data = await response.json().catch(() => ({}));
    if (!response.ok) {
        const errorMessage = data.error || data.message || '请求失败';
        throw new Error(errorMessage);
    }
    return data;
}

export async function fetchMenuData() {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/get_dishes`);
        const data = await handleResponse(response);
        
        return data.reduce((acc, item) => {
            if (!acc[item.Category]) acc[item.Category] = [];
            acc[item.Category].push({
                id: item.ID,
                name: item.Name,
                price: item.Price,
                img: item.Img
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
                Count: item.quantity,
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

export async function getHotDishes() {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/get_hot_dishes`);
        const data = await handleResponse(response);
        
        return data.map(item => ({
            id: item.ID,
            name: item.Name, 
            price: item.Price,
            category: item.Category,
            img: item.Img
        }));
    } catch (error) {
        ErrorHandler.showError(error.message);
        return [];
    }
}

export async function submitOrder(cartItems) {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/submit_order`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(cartItems.map(item => ({
                Count: item.quantity,
                DishID: item.id
            })))
        });
        
        const data = await handleResponse(response);
        ErrorHandler.showSuccess(data.msg || '订单提交成功');
        return data;
    } catch (error) {
        ErrorHandler.showError(error.message);
        throw error;
    }
}
