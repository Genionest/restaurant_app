import { serviceConfig } from './config.js';
import { fetchMenuData, submitOrder, getHotDishes } from './apiService.js';
import { renderCategories, renderMenuItems } from './menuView.js';
import { CartService, renderCartItems } from './cartService.js';
import { ErrorHandler } from './errorHandler.js';

class RestaurantApp {
    constructor() {
        this.menuData = {};
        this.hotDishes = null; // 缓存热销菜品
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
        try {
            let items = [];
            if (category === '全部') {
                items = Object.values(this.menuData).flat();
            } else if (category === '热销') {
                if (!this.hotDishes) {
                    this.hotDishes = await getHotDishes();
                }
                items = this.hotDishes;
                // alert(items)
            } else {
                items = this.menuData[category] || [];
            }
            renderMenuItems(items);
        } catch (error) {
            console.error('加载菜单项失败:', error);
            renderMenuItems([]);
        }
    }

    setupEventListeners() {
        // 分类点击事件
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
            } else if (e.target.classList.contains('checkout-btn')) {
                this.submitOrder();
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

    async submitOrder() {
        const cart = this.cartService.getCart();
        if (cart.length === 0) {
            ErrorHandler.showError('购物车为空，无法提交订单');
            return;
        }

        // 添加确认提示
        const isConfirmed = confirm('确定要提交订单吗？');
        if (!isConfirmed) {
            return;
        }

        try {
            const result = await submitOrder(cart);
            this.cartService.clearCart();
            this.updateCart();
            if (result && result.msg) {
                ErrorHandler.showSuccess(result.msg);
            } else {
                ErrorHandler.showSuccess('订单提交成功');
            }
        } catch (error) {
            console.error('提交订单失败:', error);
            ErrorHandler.showError(error.message || '提交订单失败');
        }
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

export { RestaurantApp };
