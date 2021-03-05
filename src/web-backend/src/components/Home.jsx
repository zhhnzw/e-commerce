import React from 'react'
import {get} from "../axios/tools";
import {serviceDomain} from "../axios/config";
import {Card} from "antd";
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
        orderData: {"title": {"text": ""}, "series": []},
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
        let options = {
            chart: {
                plotBackgroundColor: null,
                plotBorderWidth: null,
                plotShadow: false,
                type: 'pie'
            },
            title: {
                text: ''
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

        get({
            url: serviceDomain + '/v1/statistic/goods',
            callback: (d) => {
                let targetOptions = {...options};
                targetOptions.title.text = '商品类型分布';
                let data = [];
                d.data.data.data.forEach(one => {
                    data.push({'name': one.goodsType, 'y': one.count!==undefined?one.count:0})
                });
                targetOptions.series.push({
                    name: 'goodsType',
                    colorByPoint: true,
                    data: data});
                console.log(targetOptions);
                this.setState({goodsData: targetOptions});
            }
        });
        get({
            url: serviceDomain + '/v1/statistic/order',
            callback: (d) => {
                let targetOptions = {...options};
                targetOptions.title.text = '订单类型分布';
                let data = [];
                d.data.data.data.forEach(one => {
                    data.push({'name': one.goodsType, 'y': one.count!==undefined?one.count:0})
                });
                targetOptions.series.push({
                    name: 'goodsType',
                    colorByPoint: true,
                    data: data});
                console.log(targetOptions);
                this.setState({orderData: targetOptions});
            }
        })
    }

    render() {
        return (
            <div style={{height:'600px'}}>
                <h1 style={{fontSize:'24px', marginTop:'10px', marginLeft:'20px', width:window.innerWidth-180}}>概览</h1>
                <div style={{marginLeft:'20px', marginTop:'20px', height:150, width:window.innerWidth-272}}>
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
                <div style={{fontSize:'24px', marginLeft:'20px', marginTop:'20px', height:'500px', width:window.innerWidth-272}}>
                    <div style={{width:(window.innerWidth-278)/2, height:'500px', display:'inline-block'}}>
                        <Card style={{width:(window.innerWidth-278)/2, height:'500px'}}>
                            <PieChart options={this.state.goodsData} total={this.state.goodsTotal}/>
                        </Card>
                    </div>
                    <div style={{width:100, height:'500px', display:'inline-block', marginLeft:'8px'}}>
                        <Card style={{width:(window.innerWidth-278)/2, height:'500px'}}>
                            <PieChart options={this.state.orderData} total={this.state.orderTotal}/>
                        </Card>
                    </div>
                </div>
            </div>
        )
    }
}

export default Home;