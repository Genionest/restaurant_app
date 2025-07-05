export const renderLoginModal = () => {
    const modal = document.createElement('div');
    modal.id = 'login-modal';
    modal.className = 'modal';
    modal.innerHTML = `
        <div class="modal-content">
            <h2>用户登录</h2>
            <form id="login-form">
                <div class="form-group">
                    <input type="text" id="username" placeholder="用户名" required>
                </div>
                <div class="form-group">
                    <input type="password" id="password" placeholder="密码" required>
                </div>
                <button type="submit" class="submit-btn">登录</button>
            </form>
        </div>
    `;
    document.body.appendChild(modal);
    return modal;
};

export const setupLoginModal = (onSubmit) => {
    const modal = renderLoginModal();
    const form = modal.querySelector('#login-form');
    
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const username = form.querySelector('#username').value;
        const password = form.querySelector('#password').value;
        await onSubmit(username, password);
    });

    return {
        show: () => modal.style.display = 'block',
        hide: () => modal.style.display = 'none'
    };
};
