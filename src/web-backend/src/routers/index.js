import React, { Component } from 'react';
import { Route, Redirect, Switch } from 'react-router-dom';
import AllComponents from '../components';
import routesConfig from './config';
import { DocumentTitle } from 'react-document-title';
import queryString from 'query-string'


export default class Routers extends Component {
    render() {
        return (
            <Switch>
                {Object.keys(routesConfig).map((key) =>
                    routesConfig[key].map(r => {
                        const route = r => {
                            const Component = AllComponents[r.component];
                            return (
                                <Route
                                    key={r.route || r.key}
                                    exact
                                    path={r.route || r.key}
                                    render={props => {
                                        const reg = /\?\S*/g;
                                        // 匹配？及以后的字符串
                                        const queryParams = window.location.hash.match(reg);
                                        // 去除？的参数
                                        const {params} = props.match;
                                        Object.keys(params).forEach(key => {
                                            params[key] = params[key] && params[key].replace(reg, '');
                                        });
                                        props.match.params = {...params};
                                        const merge = {
                                            ...props,
                                            query: queryParams ? queryString.parse(queryParams[0]) : {},
                                        };
                                        return (
                                            <DocumentTitle title={r.title}>
                                                <Component {...merge} changePagePath={this.props.changePagePath}/>
                                            </DocumentTitle>
                                        );
                                    }}
                                />
                            );
                        };
                        return r.component ? route(r) : r.subs.map(r => route(r))
                    })
                )}

                <Route render={() => <Redirect to="/404" />} />
            </Switch>
        );
    }
}