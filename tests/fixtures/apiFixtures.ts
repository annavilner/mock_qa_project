import { test as base } from '@playwright/test';

type ApiFixtures = {
  basicApiURL: (accountId: number | string) => string;
};

export const test = base.extend<ApiFixtures>({
  basicApiURL: async ({}, use) => {
    await use((accountId) =>
      `https://f6b2a073-3ae1-447f-9f7e-a9eb76502794.mock.pstmn.io/accounts/${accountId}/balance`
    );
  },
});

export { expect } from '@playwright/test';
