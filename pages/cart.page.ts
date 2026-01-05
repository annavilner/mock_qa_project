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

  async checkout() {
    await this.checkoutButton.click();
  }
}
