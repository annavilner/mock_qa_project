import { test, expect, chromium } from '@playwright/test';

test('parallel excecution', async ({ page }) => {
  const browser = await chromium.launch();
  const context = await browser.newContext();
  
  //arrafy for promisses
   const promises = [
    context.newPage(), 
    context.newPage(),
    context.newPage()
  ];
    
    // wait for all pages to be created
    const pages = await Promise.all(promises);
    
    //autmateion for all pages
    await browser.close();
});

 