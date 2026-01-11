import { Page, Locator, expect } from '@playwright/test';

export class CartPage {
  readonly page: Page;
  readonly checkoutButton: Locator;

  constructor(page: Page) {
    this.page = page;
    this.checkoutButton = page.getByRole('button', { name: 'Checkout' });
  }

  async expectProductInCart(productName: string) {
    await expect(
      this.page.locator('.inventory_item_name', { hasText: productName })
    ).toBeVisible();
  }
async removeProductFromCart(productName: string) {
  // Navigate to cart first
  await this.page.click('.shopping_cart_link');

  const removeButton = this.page.locator('.cart_item')
    .filter({ hasText: productName })
    .getByRole('button', { name: 'Remove' });

  await expect(removeButton).toBeVisible({ timeout: 5000 }); // wait for button
  await removeButton.click();
}

  async checkout() {
    await this.checkoutButton.click();
  }
}
