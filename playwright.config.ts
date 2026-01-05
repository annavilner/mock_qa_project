import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  retries: 1,
  reporter: [['html', { open: 'never' }]],
  use: {
    browserName: 'chromium',
    baseURL: 'https://www.saucedemo.com',
    screenshot: 'only-on-failure',
    trace: 'on-first-retry',
  },
});
