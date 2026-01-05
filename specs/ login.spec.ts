import { test } from '@playwright/test';
import { LoginPage } from '../pages/login.page';


test.describe('Login', () => {
  test('successful login', async ({ page }) => {
    const login = new LoginPage(page);
    await login.goto();
    await login.login('standard_user', 'secret_sauce');
    await page.waitForURL('**/inventory.html');
  });

  test('invalid credentials show error', async ({ page }) => {
    const login = new LoginPage(page);
    await login.goto();
    await login.login('wrong', 'wrong');
    await login.expectError(
      'Epic sadface: Username and password do not match any user in this service'
    );
  });
});
