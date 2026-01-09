import { TransferRequest } from '../api/types/api.types';

export const validTransfer: TransferRequest = {
  from_account: 'ACC001',
  to_account: 'ACC002',
  amount: 50,
  currency: 'USD'
};

export const insufficientFundsTransfer: TransferRequest = {
  from_account: 'ACC001',
  to_account: 'ACC002',
  amount: 99999,
  currency: 'USD' 
};

