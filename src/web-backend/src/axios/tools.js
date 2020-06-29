import axios from 'axios';
import {message} from "antd";
import Cookies from 'universal-cookie';

const cookies = new Cookies();

axios.defaults.withCredentials = true;

axios.interceptors.response.use(
    response => {
        // 如果不是200状态码，抛出错误
        if (response.status === 200){
            return Promise.resolve(response);
        } else {
            return Promise.reject(response);
        }
    },
    error => {
        if (error.response.status) {
            switch (error.response.status) {
                // 401 未登录
                case 401:
                    cookies.remove('isLogin', {path: '/'})
                    cookies.remove('account', {path: '/'})
                    window.location.href = '/';
                    break;
                default:
            }
            return Promise.reject(error.response)
        }
    }
);

export const get = ({url, msg='接口异常', config, callback = res => res.data}) =>
    axios.get(url, config).then(callback).catch(err=>{
        console.log(err);
        message.warn(err.data.msg);
    });

export const post = ({url, data, msg='接口异常', config, callback = res => res.data}) =>
    axios.post(url, data, config).then(callback).catch(err=>{
        console.log(err);
        message.warn(err.data.msg);
    });

export const put = ({url, data, msg='接口异常', config, callback = res => res.data}) =>
    axios.put(url, data, config).then(callback).catch(err=>{
        console.log(err);
        message.warn(err.data.msg);
    });

export const deleteRequest = ({url, msg='接口异常', config, callback = res => res.data}) =>
    axios.put(url, config).then(callback).catch(err=>{
        console.log(err);
        message.warn(err.data.msg);
    });