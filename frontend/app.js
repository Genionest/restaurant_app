import { serviceConfig } from './config.js';
import { fetchMenuData } from './apiService.js';
import { renderCategories, renderMenuItems } from './menuView.js';
import { CartService, renderCartItems } from './cartService.js';

class RestaurantApp {
    constructor() {
        this.menuData = {};
        this.cartService = new CartService();
        this.initElements();
    }

    initElements() {
        this.elements = {
            menuItemsContainer: document.querySelector('.items-grid'),
            cartCount: document.querySelector('.cart-count'),
            cartTotal: document.querySelector('.cart-total'),
            searchInput: document.querySelector('.search-bar input'),
            searchBtn: document.querySelector('.search-bar button'),
            cartSummary: document.getElementById('cart-summary'),
            cartDetails: document.getElementById('cart-details')
        };
    }

    async init() {
        await this.loadData();
        this.setupEventListeners();
    }

    async loadData() {
        this.menuData = await fetchMenuData();
        renderCategories(this.menuData, (category) => this.loadMenuItems(category));
        await this.loadMenuItems('全部');
        this.updateCart();
    }

    async loadMenuItems(category) {
        const items = category === '全部' 
            ? Object.values(this.menuData).flat() 
            : this.menuData[category];
        renderMenuItems(items);
    }

    setupEventListeners() {
        // 分类点击事件已在renderCategories中设置
        
        // 添加到购物车
        this.elements.menuItemsContainer.addEventListener('click', (e) => {
            if (e.target.classList.contains('add-to-cart')) {
                const itemId = parseInt(e.target.dataset.id);
                this.addToCart(itemId);
            }
        });

        // 搜索
        this.elements.searchBtn.addEventListener('click', () => this.searchItems());
        this.elements.searchInput.addEventListener('keypress', (e) => {
            if (e.key === 'Enter') this.searchItems();
        });

        // 购物车展开/收起
        this.elements.cartSummary.addEventListener('click', this.toggleCartDetails.bind(this));

        // 购物车数量调整
        document.addEventListener('click', (e) => {
            if (e.target.classList.contains('quantity-increase')) {
                const itemId = parseInt(e.target.dataset.id);
                this.adjustQuantity(itemId, 1);
            } else if (e.target.classList.contains('quantity-decrease')) {
                const itemId = parseInt(e.target.dataset.id);
                this.adjustQuantity(itemId, -1);
            }
        });
    }

    async adjustQuantity(itemId, delta) {
        const item = this.cartService.getCart().find(i => i.id === itemId);
        if (item) {
            item.quantity += delta;
            if (item.quantity <= 0) {
                this.cartService.removeItem(itemId);
            }
            await this.updateCart();
        }
    }

    addToCart(itemId) {
        const allItems = Object.values(this.menuData).flat();
        const item = allItems.find(i => i.id === itemId);
        if (item) {
            this.cartService.addItem(item);
            this.updateCart();
        }
    }

    removeFromCart(itemId) {
        this.cartService.removeItem(itemId);
        this.updateCart();
    }

    async updateCart() {
        const cart = this.cartService.getCart();
        this.elements.cartCount.textContent = cart.reduce((sum, item) => sum + item.quantity, 0);
        
        try {
            const total = await this.cartService.getTotal();
            this.elements.cartTotal.textContent = `¥${total.toFixed(2)}`;
        } catch (error) {
            console.error('更新购物车总价失败:', error);
            this.elements.cartTotal.textContent = '¥0.00';
        }
        
        renderCartItems(cart, (itemId) => this.removeFromCart(itemId));
    }

    toggleCartDetails() {
        this.elements.cartDetails.classList.toggle('expanded');
    }

    async searchItems() {
        const keyword = this.elements.searchInput.value.trim().toLowerCase();
        if (!keyword) return;
        
        const allItems = Object.values(this.menuData).flat();
        const filteredItems = allItems.filter(item => 
            item.name.toLowerCase().includes(keyword)
        );
        
        renderMenuItems(filteredItems);
    }
}

// 启动应用
document.addEventListener('DOMContentLoaded', () => {
    const app = new RestaurantApp();
    app.init();
});
