export class ErrorHandler {
    static showError(message) {
        // 移除已存在的错误提示
        const existingError = document.querySelector('.error-toast');
        if (existingError) {
            existingError.remove();
        }

        // 创建错误提示元素
        const errorToast = document.createElement('div');
        errorToast.className = 'error-toast';
        // 处理多行错误信息
        const formattedMessage = message.replace(/\n/g, '<br>');
        errorToast.innerHTML = `
            <div class="error-content">
                <div class="error-message">${formattedMessage}</div>
                <button class="error-close">×</button>
            </div>
        `;

        // 添加到页面
        document.body.appendChild(errorToast);

        // 自动消失
        setTimeout(() => {
            errorToast.classList.add('fade-out');
            setTimeout(() => errorToast.remove(), 300);
        }, 5000);

        // 点击关闭
        errorToast.querySelector('.error-close').addEventListener('click', () => {
            errorToast.classList.add('fade-out');
            setTimeout(() => errorToast.remove(), 300);
        });
    }
}
