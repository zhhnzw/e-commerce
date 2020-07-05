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
        let goodsTypeId = 0;
        if (filter!==undefined && filter.goodsType!==undefined && filter.goodsType.length>0){
            goodsTypeId = filter.goodsType[0];
            this.count = this.countMap[filter.goodsType[0][1]];  // 如果筛选了商品类型，则商品页码也要跟随变动
        }
        let url = serviceDomain + '/v1/goods?goodsTypeId=' + goodsTypeId + '&pageSize=' + this.pageSize + '&pageIndex=' + this.pageIndex;
        this.setState({loading:true});
        get({
            url: url,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let secondaryType = d.data.data.data[i].secondaryType;
                    let title = d.data.data.data[i].title;
                    let price = d.data.data.data[i].price/100+"元";
                    let stock = d.data.data.data[i].stock;
                    let publishDate = d.data.data.data[i].publishDate;
                    publishDate = publishDate.replace('T', " ").replace('+08:00', "");
                    let item = {
                        'key': i, 'title': title, 'goodsType': secondaryType, 'goodsUuid': goodsUuid, 'price':price, 'publishDate': publishDate, 'stock':stock
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
                this.setState({pagination:pagination, loading:false});
            }
        })
    }

    handleSearch(value) {
        this.setState({loading:true});
        get({
            url: serviceDomain+'/v1/goods?pageSize=10&pageIndex=1&goodsUuid='+value,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let secondaryType = d.data.data.data[i].secondaryType;
                    let title = d.data.data.data[i].title;
                    let price = d.data.data.data[i].price/100+'元';
                    let stock = d.data.data.data[i].stock;
                    let publishDate = d.data.data.data[i].publishDate;
                    publishDate = publishDate.replace('T', " ").replace('+08:00', "");
                    let item = {
                        'key': i, 'title': title, 'goodsType': secondaryType, 'goodsUuid': goodsUuid, 'price':price, 'publishDate': publishDate, 'stock':stock
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
            title: '商品ID',
            dataIndex: 'goodsUuid',
            key: 'goodsUuid',
            // sorter: (a, b) => a.goodsUuid - b.goodsUuid,
            // render: (text, record) => (
            //     <div>{record.goodsUuid}</div>
            // ),
        }, {
            title: '标题',
            dataIndex: 'title',
            key: 'title'
        }, {
            title: '价格',
            dataIndex: 'price',
            key: 'price'
        }, {
            title: '库存',
            dataIndex: 'stock',
            key: 'stock'
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
                value: 1,
            }, {
                text: '夹克',
                value: 2,
            }, {
                text: '休闲裤',
                value: 3,
            }, {
                text: '运动裤',
                value: 4,
            }, {
                text: '篮球鞋',
                value: 5,
            }, {
                text: '休闲鞋',
                value: 6,
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
                    loading={this.state.loading}
                />
            </div>
        )
    }
}

export default TableForGoods;