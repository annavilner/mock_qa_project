// src/api/types/api.types.ts

export interface TransferRequest {
  from_account: string;
  to_account: string;
  amount: number;
  currency: 'USD' | 'EUR' | 'GBP';
}

export interface TransferResponse {
  transaction_id: string;
  success: true;
  new_balance: number;
  timestamp: string;
}

export interface ErrorResponse {
  success: false;
  error: string;
  current_balance?: number;
}

export interface BalanceResponse {
  account_id: string;
  balance: number;
  currency: string;
  last_updated: string;
}

export interface TransactionDetailsResponse {
  transaction_id: string;
  from_account: string;
  to_account: string;
  amount: number;
  currency: string;
  status: string;
  timestamp: string;
}
