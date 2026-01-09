import { test } from '@playwright/test';
import { LoginPage } from '../pages/login.page';
import { ProductsPage } from '../pages/products.page';
import { CartPage } from '../pages/cart.page';
import { CheckoutPage } from '../pages/checkout.page';

test('complete checkout flow', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);
  const checkout = new CheckoutPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.goToCart();
  await cart.checkout();

  await checkout.fillInformation('Anna', 'Doe', '5623');
  await checkout.finishCheckout();
  await checkout.expectSuccess();
});


test('complete checkout flow with multiple products', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);
  const checkout = new CheckoutPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.addProductToCart('Sauce Labs Bike Light');
  await products.goToCart();
  await cart.checkout();

  await checkout.fillInformation('John', 'Smith', '12345');
  await checkout.finishCheckout();
  await checkout.expectSuccess();
});

test('incomplete checkout flow with empty names', async ({ page }) => {
  const login = new LoginPage(page);
  const products = new ProductsPage(page);
  const cart = new CartPage(page);
  const checkout = new CheckoutPage(page);

  await login.goto();
  await login.login('standard_user', 'secret_sauce');

  await products.addProductToCart('Sauce Labs Backpack');
  await products.goToCart();
  await cart.checkout();

  // Empty first and last name
  await checkout.fillInformation('', '', '12345');
  await checkout.expectError('Error: First Name is required');
});
