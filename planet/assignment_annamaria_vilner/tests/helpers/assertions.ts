import type { Locator } from '@playwright/test';
import { expect } from '@playwright/test';

/**
 * Waits for the locator to contain the expected error text.
 * - locator should point to the element that will contain the error (e.g. page.locator('#fromAccountError')).
 */
export async function expectFieldError(
  locator: Locator,
  expectedText: string,
  options?: { timeout?: number }
): Promise<void> {
  const timeout = options?.timeout ?? 2000;

  // Ensure the element exists in the DOM
  await expect(locator).toHaveCount(1, { timeout });

  // Wait until the error is visible (or at least attached) and has the expected text
  // Use toHaveText which will retry until the text matches or timeout expires
  await expect(locator).toBeVisible({ timeout });
  await expect(locator).toHaveText(expectedText, { timeout });
}

/**
 * Asserts that the locator has no visible/meaningful error text.
 * This tolerates the element being hidden or present with empty/whitespace text.
 */
export async function expectNoFieldError(locator: Locator, options?: { timeout?: number }): Promise<void> {
  const timeout = options?.timeout ?? 1000;

  // If the element does not exist, that's fine (no error)
  const count = await locator.count();
  if (count === 0) return;

  // Wait briefly for any pending updates, then check textContent trimmed is empty
  // We avoid forcing visibility; an invisible element with empty text is acceptable.
  await locator.waitFor({ state: 'attached', timeout }).catch(() => {
    /* ignore: attached wait may time out if already attached */
  });

  const text = (await locator.textContent()) ?? '';
  expect(text.trim(), `expected no error text in locator`).toBe('');
}