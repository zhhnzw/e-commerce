import React from 'react'
import {get} from "../axios/tools";
import {serviceDomain} from "../axios/config";
import {Card, Row, Col} from "antd";
import HighchartsReact from 'highcharts-react-official'
var Highcharts = require('highcharts');


class PieChart extends React.Component{
    render() {
        return <div>
            <HighchartsReact
                highcharts={Highcharts}
                constructorType={'chart'}
                options={this.props.options}
            />
        </div>
    }
}

class Home extends React.Component{
    state = {
        userTotal: 0,
        goodsData: {"title": {"text": ""}, "series": []},
        goodsTotal: 0,
        orderData: {"title": {"text": ""}, "series": []},
        orderTotal: 0
    };

    componentDidMount() {
        this.props.changePagePath('概览');
        this.updateData()
        // G6图使用示例
        // this.graph = new G6.Graph(
        //     container: 'graph',
        // )
        // 然后在render函数中创建
        // <div id=graph>
    }

    updateData() {
        get({
            url: serviceDomain + '/v1/statistic/user',
            callback: (d) => {
                this.setState({userTotal:d.data.data.count})
            }
        });
        get({
            url: serviceDomain + '/v1/statistic/goods',
            callback: (d) => {
                let options = {
                    chart: {
                        plotBackgroundColor: null,
                        plotBorderWidth: null,
                        plotShadow: false,
                        type: 'pie'
                    },
                    title: {
                        text: '商品类型分布'
                    },
                    tooltip: {
                        pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
                    },
                    plotOptions: {
                        pie: {
                            allowPointSelect: true,
                            cursor: 'pointer',
                            dataLabels: {
                                enabled: true,
                                format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                                style: {
                                    color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                                }
                            }
                        }
                    },
                    series: []
                };
                let total = 0;
                let data = [];
                d.data.data.data.forEach(one => {
                    let count = one.count!==undefined?one.count:0;
                    total += count;
                    data.push({'name': one.goodsType, 'y': count})
                });
                options.series.push({
                    name: 'goodsType',
                    colorByPoint: true,
                    data: data});
                this.setState({goodsData: options, goodsTotal:total});
            }
        });
        get({
            url: serviceDomain + '/v1/statistic/order',
            callback: (d) => {
                let options = {
                    chart: {
                        plotBackgroundColor: null,
                        plotBorderWidth: null,
                        plotShadow: false,
                        type: 'pie'
                    },
                    title: {
                        text: '订单类型分布'
                    },
                    tooltip: {
                        pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
                    },
                    plotOptions: {
                        pie: {
                            allowPointSelect: true,
                            cursor: 'pointer',
                            dataLabels: {
                                enabled: true,
                                format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                                style: {
                                    color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                                }
                            }
                        }
                    },
                    series: []
                };
                let total = 0;
                let data = [];
                d.data.data.data.forEach(one => {
                    let count = one.count!==undefined?one.count:0;
                    total += count;
                    data.push({'name': one.goodsType, 'y': count})
                });
                options.series.push({
                    name: 'orderType',
                    colorByPoint: true,
                    data: data});
                this.setState({orderData: options, orderTotal: total});
            }
        })
    }

    render() {
        return (
            <div>
                <h1 style={{fontSize:'24px', marginTop:'10px', marginLeft:'20px', marginRight:'20px'}}>概览</h1>
                <div style={{marginLeft:'20px', marginTop:'20px', height:150, marginRight:'20px'}}>
                    <Card>
                        <label style={{minWidth: window.innerWidth-100, display:'inline-block', fontSize:'16px'}}>大盘数据</label>
                        <label style={{minWidth: 120, display:'inline-block', marginTop:'10px'}}>用户总数</label>
                        <label style={{minWidth: 120, display:'inline-block', marginLeft:'250px'}}>商品总数</label>
                        <label style={{minWidth: 120, display:'inline-block', marginLeft:'250px'}}>订单总数</label>
                        <br/>
                        <div style={{width:'120px',display:'inline-block'}}>
                            <label style={{fontSize:'20px'}}>{this.state.userTotal}</label><label>个</label>
                        </div>
                        <div style={{width:'120px',display:'inline-block', margin:'10px 0 0 250px'}}>
                            <label style={{fontSize:'20px'}}>{this.state.goodsTotal}</label><label>个</label>
                        </div>
                        <div style={{width:'120px',display:'inline-block', margin:'10px 0 0 250px'}}>
                            <label style={{fontSize:'20px',display:'inline-block'}}>{this.state.orderTotal}</label>
                            <label style={{fontSize:'18px',display:'inline-block'}}>个</label>
                        </div>
                    </Card>
                </div>
                <Row style={{marginLeft: '20px' , marginTop:'20px', minHeight: '480px', marginRight:'20px'}}>
                    <Col span={12}>
                        <Card style={{height:'460px', marginRight: '4px'}}>
                            <PieChart options={this.state.goodsData}/>
                        </Card>
                    </Col>
                    <Col span={12}>
                        <Card style={{height:'460px', marginLeft: '4px'}}>
                            <PieChart options={this.state.orderData}/>
                        </Card>
                    </Col>
                </Row>
            </div>
        )
    }
}

export default Home;