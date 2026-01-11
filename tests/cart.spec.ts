import { test } from '@playwright/test';
import { LoginPage } from '../pages/login.page';
import { ProductsPage } from '../pages/products.page';

import { expect, Page } from '@playwright/test';

export class CartPage {
  constructor(public page: Page) {}

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
test('validate cart is empty after removing all products', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.expectCartCount('1');
await products.goToCart();

 const removeButton = cart.page.locator('.cart_item')
  .getByRole('button', { name: 'Remove' });

  await expect(removeButton).toBeVisible({ timeout: 5000 }); // wait for button
  await removeButton.click();

  const removedProduct = page.locator('text=Sauce Labs Backpack');

await expect(removedProduct).toHaveCount(0);

});
