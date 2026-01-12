import { test, expect } from '@playwright/test';

test('GET /api/users/2 - validate response', async ({ request }) => {
  const response = await request.get('https://reqres.in/api/users/2', {
    headers: {
      'User-Agent': 'Playwright API Test',
      'Accept': 'application/json'
    }
  });

  // Validate status code
  expect(response.status()).toBe(200);

  const responseBody = await response.json();

  // Validate response fields
  expect(responseBody.data.id).toBeDefined();
  expect(responseBody.data.email).toBeDefined();
  expect(responseBody.data.first_name).toBeDefined();
});

