// B"H


const yaml = require('js-yaml');
const fs = require('fs');
const KUBE_CONFIG_PATH = '/home/USERNAME/.kube/config';
const BASE_USERNAME = 'USERNAME-';


function updateKubeConfig(username, token){
    username = BASE_USERNAME + username.toUpperCase();
    try {
        const config = yaml.safeLoad(fs.readFileSync(KUBE_CONFIG_PATH, 'utf8'));
        // const indentedJson = JSON.stringify(config, null, 4);
        config.users.forEach(element => {
            if(element.name === username){
                element.user.token = token ;
            }
        });
        
        fs.writeFileSync(KUBE_CONFIG_PATH , yaml.safeDump(config));
        // console.log(yaml.safeDump(config));
    } catch (e) {
        console.log(e);
    }
}


exports.updateKubeConfig = updateKubeConfig;
// updateKubeConfig('SHMUELRA-MGM', 'TEST THE UPDATE');
