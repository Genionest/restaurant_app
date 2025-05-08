// 引入配置
import { serviceConfig } from './config.js';

// 全局数据
let menuData = {};
let cart = [];
let total = 0;

// DOM元素
const menuItemsContainer = document.querySelector('.items-grid');
const cartCount = document.querySelector('.cart-count');
const cartTotal = document.querySelector('.cart-total');
const searchInput = document.querySelector('.search-bar input');
const searchBtn = document.querySelector('.search-bar button');
const cartSummary = document.getElementById('cart-summary');
const cartDetails = document.getElementById('cart-details');
const cartItemsContainer = document.getElementById('cart-items');

// 从API获取菜单数据
async function fetchMenuData() {
    try {
        const response = await fetch(`${serviceConfig.backend.apiBaseUrl}/api/get_dishes`);
        if (!response.ok) throw new Error('获取菜单失败');
        const data = await response.json();
        
        // 按分类组织数据
        menuData = data.reduce((acc, item) => {
            if (!acc[item.Category]) acc[item.Category] = [];
            acc[item.Category].push({
                id: item.ID,
                name: item.Name,
                price: item.Price
            });
            return acc;
        }, {});
        
        return menuData;
    } catch (error) {
        console.error('获取菜单数据出错:', error);
        return {};
    }
}

// 渲染分类菜单
function renderCategories() {
    const categoriesContainer = document.querySelector('.categories ul');
    categoriesContainer.innerHTML = '';
    
    // 添加"全部"分类
    const allItem = document.createElement('li');
    allItem.textContent = '全部';
    allItem.classList.add('active');
    allItem.addEventListener('click', () => loadMenuItems('全部'));
    categoriesContainer.appendChild(allItem);
    
    // 添加其他分类
    Object.keys(menuData).forEach(category => {
        const item = document.createElement('li');
        item.textContent = category;
        item.addEventListener('click', () => loadMenuItems(category));
        categoriesContainer.appendChild(item);
    });
}

// 加载菜单项
async function loadMenuItems(category) {
    menuItemsContainer.innerHTML = '';
    const items = category === '全部' ? Object.values(menuData).flat() : menuData[category];
    
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

// 初始化
async function init() {
    await fetchMenuData();
    renderCategories();
    await loadMenuItems('全部');
    setupEventListeners();
}

// 设置事件监听
function setupEventListeners() {
    // 分类点击
    document.querySelector('.categories').addEventListener('click', (e) => {
        if (e.target.tagName === 'LI') {
            document.querySelectorAll('.categories li').forEach(i => i.classList.remove('active'));
            e.target.classList.add('active');
        }
    });

    // 添加到购物车
    menuItemsContainer.addEventListener('click', (e) => {
        if (e.target.classList.contains('add-to-cart')) {
            const itemId = parseInt(e.target.dataset.id);
            addToCart(itemId);
        }
    });

    // 搜索
    searchBtn.addEventListener('click', searchItems);
    searchInput.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') searchItems();
    });

    // 购物车展开/收起
    cartSummary.addEventListener('click', toggleCartDetails);

    // 删除商品
    cartItemsContainer.addEventListener('click', (e) => {
        if (e.target.classList.contains('remove-item')) {
            const itemId = parseInt(e.target.dataset.id);
            removeFromCart(itemId);
        }
    });
}

// 添加到购物车
function addToCart(itemId) {
    const allItems = Object.values(menuData).flat();
    const item = allItems.find(i => i.id === itemId);
    
    if (item) {
        const existingItem = cart.find(i => i.id === itemId);
        if (existingItem) {
            existingItem.quantity++;
        } else {
            cart.push({ ...item, quantity: 1 });
        }
        total += item.price;
        updateCart();
    }
}

// 从购物车移除商品
function removeFromCart(itemId) {
    const itemIndex = cart.findIndex(i => i.id === itemId);
    if (itemIndex >= 0) {
        total -= cart[itemIndex].price * cart[itemIndex].quantity;
        cart.splice(itemIndex, 1);
        updateCart();
    }
}

// 切换购物车详情
function toggleCartDetails() {
    cartDetails.classList.toggle('expanded');
}

// 更新购物车
function updateCart() {
    cartCount.textContent = cart.reduce((sum, item) => sum + item.quantity, 0);
    cartTotal.textContent = `¥${total.toFixed(2)}`;
    renderCartItems();
}

// 渲染购物车商品
function renderCartItems() {
    cartItemsContainer.innerHTML = '';
    
    if (cart.length === 0) {
        cartItemsContainer.innerHTML = '<p class="empty-cart">购物车为空</p>';
        return;
    }

    cart.forEach(item => {
        const cartItem = document.createElement('div');
        cartItem.className = 'cart-item';
        cartItem.innerHTML = `
            <div class="cart-item-info">
                <div class="cart-item-name">${item.name}</div>
                <div class="cart-item-price">¥${item.price.toFixed(2)}</div>
            </div>
            <div class="cart-item-quantity">x${item.quantity}</div>
            <button class="remove-item" data-id="${item.id}">×</button>
        `;
        cartItemsContainer.appendChild(cartItem);
    });
}

// 搜索功能
async function searchItems() {
    const keyword = searchInput.value.trim().toLowerCase();
    if (!keyword) return;
    
    const allItems = Object.values(menuData).flat();
    const filteredItems = allItems.filter(item => 
        item.name.toLowerCase().includes(keyword)
    );
    
    menuItemsContainer.innerHTML = '';
    filteredItems.forEach(item => {
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

// 启动应用
document.addEventListener('DOMContentLoaded', init);
