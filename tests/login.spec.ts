import { expect, test } from '@playwright/test';
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

  test('validate problem user', async ({ page }) => {
    const login = new LoginPage(page);
    await login.goto();
    await login.login('problem_user', 'secret_sauce');
    await page.waitForURL('**/inventory.html');
  });

  test('validate performance user', async ({ page }) => {
    const login = new LoginPage(page);
    await login.goto();
    await login.login('performance_glitch_user', 'secret_sauce');
    await page.waitForURL('**/inventory.html');
  });

  test ('validate visual user login' , async ({ page }) => {
    const login = new LoginPage(page);
    await login.goto();
    await login.login('visual_user', 'secret_sauce');
    
    await page.waitForURL('**/inventory.html');
    await expect(page.locator('.inventory_item_img')).toHaveCount(12);
  }); 
   });
  
  



