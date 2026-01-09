import { Page, Locator, expect } from '@playwright/test';

export class CheckoutPage {
  readonly page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  async fillInformation(first: string, last: string, zip: string) {
    await this.page.getByPlaceholder('First Name').fill(first);
    await this.page.getByPlaceholder('Last Name').fill(last);
    await this.page.getByPlaceholder('Zip/Postal Code').fill(zip);
    await this.page.getByRole('button', { name: 'Continue' }).click();
  }

  async finishCheckout() {
    await this.page.getByRole('button', { name: 'Finish' }).click();
  }

  async expectSuccess() {
    await expect(
      this.page.getByText('Thank you for your order!')
    ).toBeVisible();
  }
async expectError(message: string) {
  await expect(this.page.locator('[data-test="error"]')).toHaveText(message);
}

} 