export class ErrorHandler {
    static showError(message) {
        this.showToast(message, 'error-toast');
    }

    static showSuccess(message) {
        this.showToast(message, 'success-toast');
    }

    static showToast(message, className) {
        console.log(`Showing toast: ${message}`); // 调试日志
    
        // 移除已存在的提示
        const existingToast = document.querySelector(`.${className}`);
        if (existingToast) {
            existingToast.remove();
        }

        // 创建提示元素
        const toast = document.createElement('div');
        toast.className = `${className} toast`;
        // 处理多行信息
        const formattedMessage = message.replace(/\n/g, '<br>');
        toast.innerHTML = `
            <div class="toast-content">
                <div class="toast-message">${formattedMessage}</div>
                <button class="toast-close">×</button>
            </div>
        `;

        // 添加到页面
        document.body.appendChild(toast);
        console.log('Toast element created:', toast); // 调试日志

        // 自动消失
        setTimeout(() => {
            toast.classList.add('fade-out');
            setTimeout(() => toast.remove(), 300);
        }, 5000);

        // 点击关闭
        toast.querySelector('.toast-close').addEventListener('click', () => {
            toast.classList.add('fade-out');
            setTimeout(() => toast.remove(), 300);
        });
    }
}
