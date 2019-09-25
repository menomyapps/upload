// B"H
const CONNECT = require(__dirname + '/connectToIcp.js');
require('dotenv').config();

console.info('Starting.');


// Get all envs hostnmaes.
let envs = {
    int: process.env.INT_HOSTNAME,
    qa: process.env.QA_HOSTNAME,
    prod: process.env.PROD_HOSTNAME,
}


/**
 * Start a chain of functions in difference files.
 * Flow:
 * scrape(connectToIcp.js) > getToken(rest.js) > updateKubeConfig(updateYaml.js).
 * Signatures:
 * async scrape(BASE_URL, env);
 * async getToken(baseUrl, cookie, env);
 * updateKubeConfig(username, token);
 **/ 
async function update(){
    let keysIndex = 0;
    for (var env in envs) {
        let keys = Object.keys(envs);
        console.log("Start getting " + envs[env] + " token.");
        await CONNECT.scrape(envs[env], keys[keysIndex]);
        console.log("env " + envs[env] + " token is updated.");
        keysIndex++;
    }
}

update();

