import './css/base.css';
import './css/layout.css';
import './css/components.css';
import './css/toast.css';
import { RestaurantApp } from './js/app.js';

document.addEventListener('DOMContentLoaded', () => {
    const app = new RestaurantApp();
    app.init();
});
