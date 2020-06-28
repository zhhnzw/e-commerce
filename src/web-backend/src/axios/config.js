export var serviceDomain = 'http://127.0.0.1:8000';

let env = process.env.NODE_ENV;
if (env==='development'){

}else if (env === 'production') {
    serviceDomain = 'https://www.1024cx.top';
}