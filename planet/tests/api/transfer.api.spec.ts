import { test, expect } from '@playwright/test';
import { BankingAPI } from '../../src/api/banking.api';
import { API_BASE_URL } from '../../src/utils/env';
import { validTransfer } from '../../src/data/transfer.testdata'
 //parallele excecution, however fullyparallel mode is enabled
test.describe.configure({mode: "parallel"})

test('API: Fetch account balance validation', async ({ request }) => {
  const api = new BankingAPI(request, API_BASE_URL);

  const response = await api.transferFunds(validTransfer, 200);

  if (response.success) {
  expect(typeof response.new_balance).toBe("number");
} else {
  throw new Error(`Expected success but got error: ${response.error}`);
}


});

test('API: Successful fund transfer', async ({ request }) => {
  const api = new BankingAPI(request, API_BASE_URL);

  const response = await api.transferFunds(validTransfer, 200);

  if (response.success) { expect(response.transaction_id).toBe('TXN123456'); } else { throw new Error(`Expected success but got error: ${response.error}`);

} });
import { insufficientFundsTransfer } from '../../src/data/transfer.testdata';

test('API: Insufficient funds error', async ({ request }) => {
  const api = new BankingAPI(request, API_BASE_URL);

  const response = await api.transferFunds(insufficientFundsTransfer, 400);


  expect(response.success).toBe(false);

  if ("error" in response) { expect(response.error).toBe("Insufficient funds"); } else { throw new Error("Expected an error response but got success"); }
  
});

