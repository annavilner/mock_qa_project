import { test, expect } from '../fixtures/apiFixtures';

test('GET /api/ validate response', async ({ request, basicApiURL }) => {
  const response = await request.get(await basicApiURL(), {
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

  // Validate field values & types
  expect(responseBody.account_id).toBe('ACC001');
  expect(typeof responseBody.balance).toBe('number');
});
