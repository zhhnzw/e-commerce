import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import {ConfigProvider} from "antd";
import './style.less'
import {CookiesProvider} from 'react-cookie'
import moment from 'moment'
import {BrowserRouter as Router} from 'react-router-dom'
import zh_CN from 'antd/es/locale-provider/zh_CN'
moment.locale('zh-cn');

ReactDOM.render(
    <ConfigProvider locale={zh_CN}>
        <CookiesProvider>
            <Router>
                <App/>
            </Router>
        </CookiesProvider>
    </ConfigProvider>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
