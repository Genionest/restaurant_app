import { serviceConfig } from './config.js';

export async function fetchMenuData() {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/get_dishes`);
        if (!response.ok) throw new Error('获取菜单失败');
        const data = await response.json();
        
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
        console.error('获取菜单数据出错:', error);
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
                Count: item.quantity || 1,
                DishID: item.id
            })))
        });
        
        if (!response.ok) throw new Error('计算总价失败');
        const data = await response.json();
        return data.total_price;
    } catch (error) {
        console.error('计算总价出错:', error);
        return 0;
    }
}
