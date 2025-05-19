export function renderCategories(menuData, onCategoryClick) {
    const categoriesContainer = document.querySelector('.categories ul');
    categoriesContainer.innerHTML = '';
    
    // 添加"全部"分类
    const allItem = document.createElement('li');
    allItem.textContent = '全部';
    allItem.classList.add('active');
    allItem.addEventListener('click', function() {
        document.querySelectorAll('.categories li').forEach(li => li.classList.remove('active'));
        this.classList.add('active');
        onCategoryClick('全部');
    });
    categoriesContainer.appendChild(allItem);
    
    // 添加"热销"分类
    const hotItem = document.createElement('li');
    hotItem.textContent = '热销';
    hotItem.addEventListener('click', function() {
        document.querySelectorAll('.categories li').forEach(li => li.classList.remove('active'));
        this.classList.add('active');
        onCategoryClick('热销');
    });
    categoriesContainer.appendChild(hotItem);
    
    // 添加其他分类
    Object.keys(menuData).forEach(category => {
        const item = document.createElement('li');
        item.textContent = category;
        item.addEventListener('click', function() {
            document.querySelectorAll('.categories li').forEach(li => li.classList.remove('active'));
            this.classList.add('active');
            onCategoryClick(category);
        });
        categoriesContainer.appendChild(item);
    });
}

export function renderMenuItems(items) {
    const menuItemsContainer = document.querySelector('.items-grid');
    menuItemsContainer.innerHTML = '';
    
    items.forEach(item => {
        const itemElement = document.createElement('div');
        itemElement.className = 'menu-item';
        itemElement.innerHTML = `
            <div class="item-image"></div>
            <div class="item-info">
                <h4>${item.name}</h4>
                <p>¥${item.price.toFixed(2)}</p>
                <button class="add-to-cart" data-id="${item.id}">+</button>
            </div>
        `;
        menuItemsContainer.appendChild(itemElement);
    });
}
