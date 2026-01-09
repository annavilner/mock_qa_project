import fs from 'fs';
import path from 'path';
import { test, expect } from '@playwright/test';

test.describe.configure({mode: "parallel"})

const htmlPath = path.resolve(process.cwd(), 'fund-transfer.html');
const html = fs.readFileSync(htmlPath, 'utf8');

test.describe('Fund Transfer UI', () => {
  test.beforeEach(async ({ page }) => {
    
    await page.setContent(html, { waitUntil: 'domcontentloaded' });
  });

  test('Complete fund transfer flow (happy path)', async ({ page }) => {
    // Fill form
    await page.fill('#fromAccount', 'ACC001');
    await page.fill('#toAccount', 'ACC002');
    await page.fill('#amount', '100');
    await page.selectOption('#currency', 'USD');

    // Submit
    await page.click('#transferButton');

    // Loading message should show while "processing"
    await expect(page.locator('#loadingMessage')).toBeVisible();

    // Wait for confirmation to appear (script uses setTimeout 1500ms)
    await page.waitForSelector('#confirmationMessage', { state: 'visible', timeout: 3000 });

    // Assertions on confirmation
    await expect(page.locator('#transactionId')).toHaveText(/TXN\d+/);
    await expect(page.locator('#transferredAmount')).toHaveText('USD 100.00');
    await expect(page.locator('#newBalance')).toHaveText('$1400.75');
    await expect(page.locator('#currentBalance')).toHaveText('$1400.75');

    // Button should be re-enabled and form reset
    await expect(page.locator('#transferButton')).toBeEnabled();
    await expect(page.locator('#fromAccount')).toHaveValue('');
    await expect(page.locator('#toAccount')).toHaveValue('');
    await expect(page.locator('#amount')).toHaveValue('');
  });

  test('Form validation for empty/invalid fields', async ({ page }) => {
    // Submit with all empty fields
    await page.click('#transferButton');
  
    // Expect errors for account IDs (length < 3)
    await expect(page.locator('#fromAccountError')).toHaveText('Account ID must be at least 3 characters');
    await expect(page.locator('#toAccountError')).toHaveText('Cannot transfer to the same account');

    // Fill accounts with 2-char values to trigger same validation again
    await page.fill('#fromAccount', 'AB');
    await page.fill('#toAccount', 'CD');
    // Invalid amount (0)
    await page.fill('#amount', '0');

    await page.click('#transferButton');

    // Amount must be greater than 0
    await expect(page.locator('#amountError')).toHaveText('Amount must be greater than 0');
  });

  test('Validation for same account transfer attempt', async ({ page }) => {
    await page.fill('#fromAccount', 'SAMEACC');
    await page.fill('#toAccount', 'SAMEACC');
    await page.fill('#amount', '50');

    await page.click('#transferButton');

    // Should show cannot transfer to same account
    await expect(page.locator('#toAccountError')).toHaveText('Cannot transfer to the same account');

    // Ensure no confirmation message was shown
    await expect(page.locator('#confirmationMessage')).toBeHidden();
  });

  test('Insufficient funds error message display', async ({ page }) => {
    await page.fill('#fromAccount', 'ACC100');
    await page.fill('#toAccount', 'ACC200');
    // Amount greater than mock balance in page (1500.75)
    await page.fill('#amount', '2000');

    await page.click('#transferButton');

    // Should display insufficient funds error
    await expect(page.locator('#amountError')).toHaveText('Insufficient funds');

    // Ensure no confirmation was displayed
    await expect(page.locator('#confirmationMessage')).toBeHidden();
  });
});