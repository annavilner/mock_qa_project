import { APIRequestContext } from '@playwright/test';
import {
  TransferRequest,
  TransferResponse,
  ErrorResponse,
  TransactionDetailsResponse
} from './types/api.types';

export class BankingAPI {
  constructor(private request: APIRequestContext, private baseUrl: string) {}

  async transferFunds(
    payload: TransferRequest,
    mockCode?: number
  ): Promise<TransferResponse | ErrorResponse> {
    const headers: Record<string, string> = {};

    if (mockCode) {
      headers['x-mock-response-code'] = String(mockCode);
    }

    const response = await this.request.post(`${this.baseUrl}/transaction`, {
      data: payload,
      headers
    });

    return await response.json();
  }

 async getBalance(accountId: string) {
  const res = await this.request.get(`${this.baseUrl}/${accountId}/balance`);
  const data = await res.json();

  return {
    account_id: data.account_id ?? data.accountId ?? data.id,
    balance: data.balance ?? data.amount
  };
}


  async getTransactionDetails(
    transactionId: string
  ): Promise<TransactionDetailsResponse> {
    const response = await this.request.get(
      `${this.baseUrl}/transactions/${transactionId}`
    );
    return await response.json();
  }
}
