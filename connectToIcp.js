// B"H

const puppeteer = require('puppeteer-core');
const REST = require(__dirname + '/rest.js');


let scrape = async (BASE_URL, env) => {

  console.log(env);
  const CREDS = {
    username: process.env.USERNAME,
    password: process.env.PASSWORD
  }
  let chromiumPath = process.env.CHROMIUM_PATH;
  let PORT = port;
  BASE_URL = `https://${BASE_URL}:${PORT}`;

  const browser = await puppeteer.launch({
    executablePath: chromiumPath,
    ignoreHTTPSErrors: true
  });

  const page = await browser.newPage();
  await page.goto(BASE_URL);

  await page.waitForNavigation({
    waitUntil: 'domcontentloaded'
  });

  await page.goto(BASE_URL, { waitUntil: 'networkidle0' }); // wait until page load 
  await page.type('#username', CREDS.username);
  await page.type('#password', CREDS.password); // click and wait for navigation 
  await Promise.all([
    page.click('button[name="loginButton"]'),
    page.waitForNavigation({ waitUntil: 'networkidle0' }),
  ]);


  const cookies = await page.cookies()
  let cookie = ""
  cookies.forEach(function (element) {
    cookie += " " + element.name + "=" + element.value + ";";
  });

  await REST.getToken(BASE_URL, cookie, env);
  // console.log(BASE_URL);

  await page.screenshot({ path: 'screen.png', fullPage: true });
  browser.close();
};

exports.scrape = scrape;
