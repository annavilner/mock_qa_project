import { test } from '@playwright/test';
import { LoginPage } from '../pages/login.page';
import { ProductsPage } from '../pages/products.page';

import { expect, Page } from '@playwright/test';

export class CartPage {
  constructor(private page: Page) {}

 async expectProductInCart(productName: string) {
  const product = this.page
    .locator('.cart_item')
    .filter({ hasText: productName });

  await expect(product).toBeVisible();
}

}

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

test('add multuiple to cart', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');


  await products.addProductToCart('Sauce Labs Backpack');
  await products.addProductToCart('Sauce Labs Fleece Jacket');
  await products.addProductToCart('Sauce Labs Onesie');
 
  await products.goToCart();
  await cart.expectProductInCart('Sauce Labs Backpack');
  await cart.expectProductInCart('Sauce Labs Fleece Jacket');
  await cart.expectProductInCart('Sauce Labs Onesie');
  await products.expectCartCount('3');
});

test ('remove product from cart', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.expectCartCount('1');

  await products.removeProductFromCart('Sauce Labs Backpack');
  await products.expectCartCount('0');


});
