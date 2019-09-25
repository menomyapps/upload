// B"H

colors = {
    reset: '\033[0m',
    //text color
    black: '\033[30m',
    red: '\033[31m',
    green: '\033[32m',
    yellow: '\033[33m',
    blue: '\033[34m',
    magenta: '\033[35m',
    cyan: '\033[36m',
    white: '\033[37m',
    //background color
    blackBg: '\033[40m',
    redBg: '\033[41m',
    greenBg: '\033[42m',
    yellowBg: '\033[43m',
    blueBg: '\033[44m',
    magentaBg: '\033[45m',
    cyanBg: '\033[46m',
    whiteBg: '\033[47m'
}


function getTimeStamp(){
    let timeStamp = `${new Date().toISOString()} ::`;
    return timeStamp;
}


function resetColor(){
  return colors.reset;
}


function getColorAndPrefix(type){
  let colorAndPrefix;
  switch (type) {
    case 'INFO':
      colorAndPrefix = colors.cyan + ' Info : ';
      break;
    case 'ERROR':
      colorAndPrefix = colors.red  + ' Error : ';
      break;
    default:
      colorAndPrefix = resetColor();
      break;
  }
    
  return colorAndPrefix;
}


function logger(msg, type) {
    // Method overload workaround;
    try {
      type = type.toUpperCase();
    } catch (error){
      type = null; 
    }

    // Constract message.
    let message = getTimeStamp() + getColorAndPrefix(type) + msg + resetColor();

    switch (type) {
        case 'INFO' :
          console.log(message);
          resetColor();
          break;
        case 'ERROR' :
          console.error(message);
          resetColor();
          break;
        default:
          console.log(msg);
    }
}

module.exports = logger;
