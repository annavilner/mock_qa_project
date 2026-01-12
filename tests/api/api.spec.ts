import { test, expect } from '@playwright/test';

test('GET /api/users/2 - validate response', async ({ request }) => {
  const response = await request.get('https://f6b2a073-3ae1-447f-9f7e-a9eb76502794.mock.pstmn.io/accounts/{account_id}/balance', {
    headers: {
      'User-Agent': 'Playwright API Test',
      'Accept': 'application/json'
    }
  });

  // Validate status code
  expect(response.status()).toBe(200);

  const responseBody = await response.json();

  // Validate response fields
  expect(responseBody.account_id).toBeDefined();
  
  expect(responseBody.account.first_name).toBeDefined();
});

