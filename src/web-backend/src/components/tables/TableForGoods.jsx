import React from 'react';
import {Button, Input, message, Table} from "antd";
import {get} from '../../axios/tools'
import {serviceDomain} from "../../axios/config";


const {Search} = Input;

class TableForGoods extends React.Component{
    limit = 10;
    offset = 0;
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
        this.offset = (pagination.current-1)*10;
        this.updateData();
    };

    updateData() {
        get({
            url: serviceDomain+'/v1/goods?primaryType=clothes&secondaryType=shirt&pageSize=20&pageIndex=1',
            callback: (d) => {
                console.log(d);
                let data = [];
                for (let i in d.data.data.data) {
                    let goodsUuid = d.data.data.data[i].GoodsUuid;
                    let primaryType = d.data.data.data[i].PrimaryType;
                    let title = d.data.data.data[i].Title;
                    let item = {
                        'key': i, 'title': title, 'primaryType': primaryType, 'goodsUuid': goodsUuid
                    };
                    data.push(item)
                }
                const pagination = {...this.state.pagination};
                pagination.total = 100; //TODO: 临时写死
                this.setState({data:data, pagination:pagination})
            }
        })
    }

    componentDidMount() {
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