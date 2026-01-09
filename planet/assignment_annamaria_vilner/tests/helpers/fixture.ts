import fs from 'fs';
import path from 'path';
import { test as base, expect as baseExpect, type Locator } from '@playwright/test';

type FormHelpers = {
  fillField(name: string, value: string): Promise<void>;
  blurField(name: string): Promise<void>;
  submit(): Promise<void>;
  getError(name: string): Locator;
};

type MyFixtures = {
  /**
   * Load the given HTML file (path relative to repo root) into the test page.
   * Defaults to "fund-transfer.html".
   */
  loadHtml: (htmlPath?: string) => Promise<void>;

  /**
   * Helpers for interacting with the transfer form.
   * Fields are referenced by their id (e.g. "fromAccount", "amount").
   */
  form: FormHelpers;
};

export const test = base.extend<MyFixtures>({
  loadHtml: async ({ page }, use) => {
    const load = async (htmlFile = 'fund-transfer.html') => {
      const resolved = path.resolve(process.cwd(), htmlFile);
      const html = fs.readFileSync(resolved, 'utf8');
      // Set the page content to the static HTML
      await page.setContent(html, { waitUntil: 'domcontentloaded' });
    };
    await use(load);
  },

  form: async ({ page }, use) => {
    const helpers: FormHelpers = {
      fillField: async (name: string, value: string) => {
        await page.fill(`#${name}`, value);
      },
      blurField: async (name: string) => {
        await page.dispatchEvent(`#${name}`, 'blur');
      },
      submit: async () => {
        await page.click('#transferButton');
      },
      getError: (name: string) => page.locator(`#${name}Error`),
    };

    await use(helpers);
  },
});

// Re-export Playwright expect so tests can import a single module
export const expect = baseExpect;