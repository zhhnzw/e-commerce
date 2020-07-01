import React from 'react';
import {Button, Input, message, Table} from "antd";
import {get} from '../../axios/tools'
import {serviceDomain} from "../../axios/config";


const {Search} = Input;

class TableForGoods extends React.Component{
    pageIndex = 1;
    pageSize = 10;
    count = 0;
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
        this.pageIndex += 1;
        this.updateData();
    };

    updateData() {
        get({
            url: serviceDomain+'/v1/goods?primaryType=clothes&secondaryType=shirt&pageSize='+this.pageSize+'&pageIndex='+this.pageIndex,
            callback: (d) => {
                let data = [];
                for (let i in d.data.data.data) {
                    let goodsUuid = d.data.data.data[i].goodsUuid;
                    let primaryType = d.data.data.data[i].primaryType;
                    let title = d.data.data.data[i].title;
                    let item = {
                        'key': i+(this.pageIndex-1)*this.pageSize, 'title': title, 'primaryType': primaryType, 'goodsUuid': goodsUuid
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

    handleSearch() {

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
            title: '商品类型',
            dataIndex: 'primaryType',
            key: 'primaryType',
            filters: [{
                text: '衣服',
                value: 'clothes',
            }, {
                text: '袜子',
                value: 'shoes',
            }],
            onFilter: (value, record) => {
                if (record.type !== undefined) {
                    return record.type.indexOf(value) === 0
                } else {
                    return false
                }
            }
        }, {
            title: '操作',
            key: 'action',
            render: (text, record) => (
                <span>
                    <Button type='link' onClick={this.handleEditGoods.bind(this,record)}>编辑</Button>
                </span>
            ),
        }];
        return (
            <div style={{margin:'20px 12px 0 12px'}}>
                <Button type='primary' onClick={this.handleAddGoods} style={{marginBottom:'20px'}}>添加商品</Button>
                <Search
                    placeholder="请输入商品ID"
                    onSearch={value => this.handleSearch(value)}
                    style={{width: 200, marginLeft:'20px'}}
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