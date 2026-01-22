import { test } from '@playwright/test';
import { LoginPage } from '../pages/login.page';
import { ProductsPage } from '../pages/products.page';

import { expect, Page } from '@playwright/test';
import { Sidebar } from '../pages/cart.page';

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

test('validate cart count after adding and removing products', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.addProductToCart('Sauce Labs Bike Light');
  await products.expectCartCount('2');  
  await products.removeProductFromCart('Sauce Labs Bike Light');
  await products.expectCartCount('1');  
  await products.goToCart();
  await cart.expectProductInCart('Sauce Labs Backpack');
});  

test.describe('Cart page navigation',  { tag :'@cart'},() => {
  test('navigate to cart page from products page', async ({ page }) => {
    const login = new LoginPage(page);
    const products = new ProductsPage(page);
    const cart = new CartPage(page);  
    await login.goto();
    await login.login('standard_user', 'secret_sauce');

    await products.goToCart();
    await expect(page).toHaveURL(/.*cart.html/);    
  }  );

  test('navigate back to products page from cart page', async ({ page }) => {
    const login = new LoginPage(page);
    const products = new ProductsPage(page);
    const cart = new CartPage(page);  
    await login.goto();
    await login.login('standard_user', 'secret_sauce');

    await products.goToCart();
    await expect(page).toHaveURL(/.*cart.html/);  
    await cart.page.getByRole('button', { name: 'Continue Shopping' }).click();
    await products.expectOnProductsPage();  
  });

  test('navigate with side menu from cart to log out', async ({ page }) => {
  const login = new LoginPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

   await page.locator('[data-test="shopping-cart-link"]').click();
  await page.getByRole('button', { name: 'Open Menu' }).click();
  await page.locator('[data-test="logout-sidebar-link"]').click();
});
    test('navigate with side menu from cart to Reset App', async ({ page }) => {
  const login = new LoginPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

   await page.locator('[data-test="shopping-cart-link"]').click();
  await page.getByRole('button', { name: 'Open Menu' }).click();
  await page.locator('[data-test="reset-sidebar-link"]').click();

  expect(await page.locator('.shopping_cart_badge').count()).toBe(0);
});

test('navigate with side menu from cart to about side bar', async ({ page }) => {
  const login = new LoginPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

   await page.locator('[data-test="shopping-cart-link"]').click();
  await page.getByRole('button', { name: 'Open Menu' }).click();
  await page.locator('[data-test="about-sidebar-link"]').click();

  expect(await page).toHaveURL(/saucelabs.com/);
  expect(await page.title()).toContain('Testing');

});

});

