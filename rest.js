// B"H

const axios = require('axios');
const https = require('https');
const updateYaml = require(__dirname + '/updateYaml.js');

async function getToken(baseUrl, cookie, env) {
    let tokenPath = '/console/token';
    let refererPath = '/console/welcome';

    let headers = {
        'DNT': '1',
        'Accept-Encoding': 'gzip, deflate, br',
        'Accept-Language': 'en-US,en;q=0.8',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36',
        'Accept':
            'application/json',
        'Referer': baseUrl + refererPath,
        'Connection': 'keep-alive',
        'Cookie': cookie,
    };


    const instance = axios.create({
        baseURL: baseUrl,
        headers: headers,
        validateStatus: function (status) {
            return status >= 200 && status < 300; // default
        },
        httpsAgent: new https.Agent({
            rejectUnauthorized: false
        })
    });


    try {
        const response = await instance.get(tokenPath);
        console.log(response.data.id_token);
        updateYaml.updateKubeConfig(env, response.data.id_token);
    } catch (error) {
        console.error(error);
    }

}

// getToken();

exports.getToken = getToken;
