import { serviceConfig } from './config.js';
import { ErrorHandler } from './errorHandler.js';

export const loginUser = async (username, password) => {
    try {
        const response = await fetch(`${serviceConfig.apiBaseUrl}/user/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username,
                password
            })
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || '登录失败');
        }

        return await response.json();
    } catch (error) {
        ErrorHandler.showError(error.message);
        throw error;
    }
};
