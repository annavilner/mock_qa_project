import { test, expect } from '@playwright/test';

test('GET /api/ validate response', async ({ request }) => {
  const response = await request.get('https://f6b2a073-3ae1-447f-9f7e-a9eb76502794.mock.pstmn.io/accounts/{account_id}/balance', {
    headers: {
      'User-Agent': 'Playwright API Test',
      'Accept': 'application/json'
    }
  });

  // Validate status code
  expect(response.status()).toBe(200);
 expect(response.ok()).toBeTruthy();
  const responseBody = await response.json();

  // Validate response fields
  expect(responseBody).toHaveProperty('account_id');
  expect(responseBody).toHaveProperty('balance');

  // ✅ Field values & types
  expect(responseBody.account_id).toBe('ACC001');
  expect(typeof responseBody.balance).toBe('number');

});

