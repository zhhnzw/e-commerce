export var serviceDomain = 'http://127.0.0.1:8000';
export const cookieMaxAge = 12*60*60;  //cookie的有效时长，单位s

let env = process.env.NODE_ENV;
if (env==='development'){

}else if (env === 'production') {
    serviceDomain = 'https://www.1024cx.top/api';
}