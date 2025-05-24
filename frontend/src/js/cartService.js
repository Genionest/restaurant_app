import { calculateTotalPrice } from './apiService.js';

export class CartService {
    constructor() {
        this.cart = [];
    }

    async addItem(item) {
        const existingItem = this.cart.find(i => i.id === item.id);
        if (existingItem) {
            existingItem.quantity++;
        } else {
            this.cart.push({ ...item, quantity: 1 });
        }
    }

    removeItem(itemId) {
        const itemIndex = this.cart.findIndex(i => i.id === itemId);
        if (itemIndex >= 0) {
            this.cart.splice(itemIndex, 1);
        }
    }

    getCart() {
        return this.cart;
    }

    async getTotal() {
        try {
            const response = await calculateTotalPrice(this.cart);
            return response;
        } catch (error) {
            console.error('获取总价失败:', error);
            return 0;
        }
    }

    clearCart() {
        this.cart = [];
    }
}

export function renderCartItems(cartItems, onRemoveItem) {
    const cartItemsContainer = document.getElementById('cart-items');
    cartItemsContainer.innerHTML = '';
    
    if (cartItems.length === 0) {
        cartItemsContainer.innerHTML = '<p class="empty-cart">购物车为空</p>';
        return;
    }

    cartItems.forEach(item => {
        const cartItem = document.createElement('div');
        cartItem.className = 'cart-item';
        cartItem.innerHTML = `
            <div class="cart-item-info">
                <div class="cart-item-name">${item.name}</div>
                <div class="cart-item-price">¥${item.price.toFixed(2)}</div>
            </div>
            <div class="cart-item-quantity-controls">
                <button class="quantity-decrease" data-id="${item.id}">-</button>
                <span class="quantity-value">${item.quantity}</span>
                <button class="quantity-increase" data-id="${item.id}">+</button>
            </div>
            <button class="remove-item" data-id="${item.id}">×</button>
        `;
        cartItem.querySelector('.remove-item').addEventListener('click', () => onRemoveItem(item.id));
        cartItemsContainer.appendChild(cartItem);
    });
}
