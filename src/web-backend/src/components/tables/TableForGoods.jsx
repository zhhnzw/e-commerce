import React from 'react';
import {Button, Input, message, Table} from "antd";
import {get} from '../../axios/tools'
import {serviceDomain} from "../../axios/config";


const {Search} = Input;

class TableForGoods extends React.Component{
    pageIndex = 1;
    pageSize = 10;
    count = 0;
    countMap = {};
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
        let primaryType = '';
        let secondaryType = '';
        if (filter!==undefined && filter.goodsType!==undefined && filter.goodsType.length>0){
            primaryType = filter.goodsType[0][0];
            secondaryType = filter.goodsType[0][1];
            this.count = this.countMap[filter.goodsType[0][1]];  // 如果筛选了商品类型，则商品页码也要跟随变动
        }
        let url = serviceDomain + '/v1/goods?primaryType=' + primaryType + '&secondaryType=' + secondaryType + '&pageSize=' + this.pageSize + '&pageIndex=' + this.pageIndex
        get({
            url: url,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let secondaryType = d.data.data.data[i].secondaryType;
                    let title = d.data.data.data[i].title;
                    let price = d.data.data.data[i].price+"元";
                    let publishDate = d.data.data.data[i].publishDate;
                    publishDate = publishDate.replace('T', " ").replace('+08:00', "");
                    let item = {
                        'key': i, 'title': title, 'goodsType': secondaryType, 'goodsUuid': goodsUuid, 'price':price, 'publishDate': publishDate
                    };
                    data.push(item)
                }
                const pagination = {...this.state.pagination};
                pagination.total = this.count;
                this.setState({data:data, pagination:pagination});
            }
        })
    }

    getCount() {
        get({
            url: serviceDomain+'/v1/statistic/goods',
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
                this.setState({pagination:pagination});
            }
        })
    }

    handleSearch(value) {
        console.log(value);
        get({
            url: serviceDomain+'/v1/goods?pageSize=10&pageIndex=1&goodsUuid='+value,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let secondaryType = d.data.data.data[i].secondaryType;
                    let title = d.data.data.data[i].title;
                    let price = d.data.data.data[i].price+'元';
                    let publishDate = d.data.data.data[i].publishDate;
                    publishDate = publishDate.replace('T', " ").replace('+08:00', "");
                    let item = {
                        'key': i, 'title': title, 'goodsType': secondaryType, 'goodsUuid': goodsUuid, 'price':price, 'publishDate': publishDate
                    };
                    data.push(item)
                }
                this.count = 1;
                const pagination = {...this.state.pagination};
                pagination.total = this.count;
                this.setState({data:data, pagination:pagination});
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
            title: '商品ID',
            key: 'goodsUuid',
            sorter: (a, b) => a.goodsUuid - b.goodsUuid,
            render: (text, record) => (
                <div>{record.goodsUuid}</div>
            ),
        }, {
            title: '标题',
            dataIndex: 'title',
            key: 'title'
        }, {
            title: '价格',
            dataIndex: 'price',
            key: 'price'
        }, {
            title: '发布日期',
            dataIndex: 'publishDate',
            key: 'publishDate'
        }, {
            title: '商品类型',
            dataIndex: 'goodsType',
            key: 'goodsType',
            filterMultiple: false,
            filters: [{
                text: '衬衫',
                value: ['clothes', 'shirt'],
            }, {
                text: '夹克',
                value: ['clothes', 'jacket'],
            }, {
                text: '休闲裤',
                value: ['pants', 'casual_pants'],
            }, {
                text: '运动裤',
                value: ['pants', 'sports_pants'],
            }, {
                text: '篮球鞋',
                value: ['shoes', 'basketball_shoes'],
            }, {
                text: '休闲鞋',
                value: ['shoes', 'casual_shoes'],
            }],
            // onFilter: (value, record) => {
            //     console.log(value, record);
            //     if (record.primaryType === value[0]) {
            //         return true
            //     } else {
            //         return false
            //     }
            // }
        }, {
            title: '操作',
            key: 'action',
            render: (text, record) => (
                <span>
                    <Button type='link' disabled={true} onClick={this.handleEditGoods.bind(this,record)}>编辑</Button>
                    <Button type='link' disabled={true} onClick={this.handleEditGoods.bind(this,record)}>下架</Button>
                </span>
            ),
        }];
        return (
            <div style={{margin:'20px 12px 0 12px'}}>
                <Button type='primary' onClick={this.handleAddGoods} style={{marginBottom:'20px'}}>添加商品</Button>
                <Search
                    placeholder="请输入商品ID"
                    onSearch={value => this.handleSearch(value)}
                    style={{width: 400, marginLeft:'20px'}}
                    enterButton
                />
                <Table
                    rowSelection={rowSelection}
                    columns={columns}
                    dataSource={this.state.data}
                    pagination={this.state.pagination}
                    onChange={this.handleTableChange}
                />
            </div>
        )
    }
}

export default TableForGoods;