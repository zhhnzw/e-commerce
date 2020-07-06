import React from 'react';
import {Button, Input, message, Table} from "antd";
import {get} from '../../axios/tools'
import {serviceDomain} from "../../axios/config";


const {Search} = Input;

class TableForOrder extends React.Component{
    pageIndex = 1;
    pageSize = 10;
    count = 0;
    countMap = {};
    goodsMap = {
        1:'shirt',
        2:'jacket',
        3:'casual_pants',
        4:'sports_pants',
        5:'basketball_shoes',
        6:'casual_shoes'
    };
    state = {
        data : [],
        pagination: {showQuickJumper:true},
        selectedRowKeys: [],
        loading: false,
    };

    handleAddGoods = () => {
        this.props.history.push({pathname:'/goods',state:{pageType:'create'}})
    };

    handleEditGoods = () => {
        this.props.history.push({pathname:'/goods',state:{pageType:'edit'}})
    };

    handleTableChange = (pagination, filter, sorter) => {
        const pager = {...this.state.pagination};
        pager.current = pagination.current;
        this.setState({
            pagination: pager,
        });
        console.log(filter, sorter);
        this.pageIndex = pager.current;
        this.updateData(filter);
    };

    updateData(filter) {
        let goodsTypeId = 0;
        if (filter!==undefined && filter.goodsType!==undefined && filter.goodsType.length>0){
            goodsTypeId = filter.goodsType[0];
            this.count = this.countMap[filter.goodsType[0][1]];  // 如果筛选了商品类型，则商品页码也要跟随变动
        }
        let url = serviceDomain + '/v1/order?goodsTypeId=' + goodsTypeId + '&pageSize=' + this.pageSize + '&pageIndex=' + this.pageIndex;
        this.setState({loading:true});
        get({
            url: url,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let orderId = d.data.data.data[i].orderId;
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let goodsType = this.goodsMap[d.data.data.data[i].goodsTypeId];
                    // let title = d.data.data.data[i].title;
                    let price = d.data.data.data[i].price/100+"元";
                    let orderStatus = d.data.data.data[i].orderStatus;
                    let userName = d.data.data.data[i].userName;
                    let item = {
                        'key': i, 'orderId': orderId, 'goodsType': goodsType, 'goodsUuid': goodsUuid, 'price':price, 'orderStatus':orderStatus, 'userName':userName
                    };
                    data.push(item)
                }
                const pagination = {...this.state.pagination};
                pagination.total = this.count;
                this.setState({data:data, pagination:pagination, loading:false});
            }
        })
    }

    getCount() {
        this.setState({loading:true});
        get({
            url: serviceDomain+'/v1/statistic/order',
            callback: (d) => {
                console.log(d);
                let count = 0;
                for (let i in d.data.data.data) {
                    this.countMap[d.data.data.data[i].goodsType] = d.data.data.data[i].count;
                    let c = d.data.data.data[i].count;
                    count += c;
                }
                this.count = count;
                const pagination = {...this.state.pagination};
                pagination.total = count;
                this.setState({pagination:pagination, loading:false});
            }
        })
    }

    handleSearch(value) {
        this.setState({loading:true});
        get({
            url: serviceDomain+'/v1/order?pageSize=10&pageIndex=1&orderId='+value,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let orderId = d.data.data.data[i].orderId;
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let goodsType = this.goodsMap[d.data.data.data[i].goodsTypeId];
                    // let title = d.data.data.data[i].title;
                    let price = d.data.data.data[i].price/100+"元";
                    let orderStatus = d.data.data.data[i].orderStatus;
                    let userName = d.data.data.data[i].userName;
                    let item = {
                        'key': i, 'orderId': orderId, 'goodsType': goodsType, 'goodsUuid': goodsUuid, 'price':price, 'orderStatus':orderStatus, 'userName':userName
                    };
                    data.push(item)
                }
                this.count = 1;
                const pagination = {...this.state.pagination};
                pagination.total = this.count;
                this.setState({data:data, pagination:pagination, loading:false});
            }
        })
    }

    componentDidMount() {
        this.getCount();
        this.updateData();
    }

    onSelectChange = selectedRowKeys => {
        this.setState({selectedRowKeys});
    };

    render() {
        const {loading, selectedRowKeys} = this.state;
        const rowSelection = {
            selectedRowKeys,
            onChange: this.onSelectChange,
        };
        const hasSelected = selectedRowKeys.length > 0;
        const columns = [{
            title: '订单ID',
            dataIndex: 'orderId',
            key: 'orderId',
        }, {
            title: '用户',
            dataIndex: 'userName',
            key: 'userName'
        }, {
            title: '商品ID',
            dataIndex: 'goodsUuid',
            key: 'goodsUuid',
        }, {
            title: '价格',
            dataIndex: 'price',
            key: 'price'
        }, {
            title: '订单状态',
            dataIndex: 'orderStatus',
            key: 'orderStatus'
        }, {
            title: '商品类型',
            dataIndex: 'goodsType',
            key: 'goodsType',
            // filterMultiple: false,
            // filters: [{
            //     text: '衬衫',
            //     value: 1,
            // }, {
            //     text: '夹克',
            //     value: 2,
            // }, {
            //     text: '休闲裤',
            //     value: 3,
            // }, {
            //     text: '运动裤',
            //     value: 4,
            // }, {
            //     text: '篮球鞋',
            //     value: 5,
            // }, {
            //     text: '休闲鞋',
            //     value: 6,
            // }]
        }];
        return (
            <div style={{margin:'20px 12px 0 12px'}}>
                <Search
                    placeholder="请输入订单id"
                    onSearch={value => this.handleSearch(value)}
                    style={{width: 400, marginBottom:'8px'}}
                    enterButton
                />
                <Table
                    rowSelection={rowSelection}
                    columns={columns}
                    dataSource={this.state.data}
                    pagination={this.state.pagination}
                    onChange={this.handleTableChange}
                    loading={this.state.loading}
                />
            </div>
        )
    }
}

export default TableForOrder;