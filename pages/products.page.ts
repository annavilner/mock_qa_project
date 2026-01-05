import { Page, Locator, expect } from '@playwright/test';

export class ProductsPage {
  readonly page: Page;
  readonly cartBadge: Locator;

  constructor(page: Page) {
    this.page = page;
    this.cartBadge = page.locator('.shopping_cart_badge');
  }

  async addProductToCart(productName: string) {
    await this.page
      .locator('.inventory_item')
      .filter({ hasText: productName })
      .getByRole('button', { name: 'Add to cart' })
      .click();
  }

  async goToCart() {
    await this.page.locator('.shopping_cart_link').click();
  }

  async expectCartCount(count: string) {
    await expect(this.cartBadge).toHaveText(count);
  }
}
