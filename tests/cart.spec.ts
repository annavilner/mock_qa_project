import { test } from '@playwright/test';
import { LoginPage } from '../pages/login.page';
import { ProductsPage } from '../pages/products.page';
import { CartPage } from '../pages/cart.page';

test('add product to cart', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.expectCartCount('1');

  await products.goToCart();
  await cart.expectProductInCart('Sauce Labs Backpack');
});
