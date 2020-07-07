import React from 'react'
import {
    Chart,
    Interval,
    Tooltip,
    Axis,
    Coordinate,
    Interaction,
    Annotation,
    Guide,
    Legend
} from 'bizcharts';
import {DataSet} from "@antv/data-set";
import {get} from "../axios/tools";
import {serviceDomain} from "../axios/config";
import {Card} from "antd";


class GoodsChart extends React.Component{
    render() {
        const { DataView } = DataSet;
        const { Html } = Guide;
        const dv = new DataView();
        dv.source(this.props.data).transform({
            type: "percent",
            field: "count",
            dimension: "goodsType",
            as: "percent"
        });
        const cols = {
            percent: {
                formatter: val =>{
                    val = (val*100).toFixed(2)+"%";
                    return val;
                }
            }
        };
        // let guideHtml = '<div style="color:#8c8c8c;font-size: 1.16em;text-align: center;width: 10em;">总计<br/><span style="color: #262626;font-size: 1.8em">' + this.props.total + '</span>个</div>';
        return (
            <div>
                <h1>商品类型分布</h1>
                <Chart
                    height={400}
                    data={dv}
                    scale={cols}
                    autoFit
                >
                    <Legend visible={false} />
                    <Coordinate type="theta" innerRadius={0.4}/>
                    <Axis visible={false}/>
                    <Interval
                        position="percent"
                        adjust="stack"
                        color="goodsType"
                        style={{
                            lineWidth: 1,
                            stroke: '#fff',
                        }}
                        label={['count', {
                            content: (data) => {
                                return `${data.goodsType}: ${(data.percent * 100).toFixed(2)}%`;
                            },
                        }]}
                    />
                    {/*<Guide>*/}
                    {/*    <Html*/}
                    {/*        position={["50%", "50%"]}*/}
                    {/*        html={guideHtml}*/}
                    {/*        alignX="middle"*/}
                    {/*        alignY="middle"*/}
                    {/*    />*/}
                    {/*</Guide>*/}
                    <Annotation.Text
                        position={['50%', '50%']}
                        content={this.props.total}
                        style={{
                            lineHeight: '240px',
                            fontSize: '30',
                            fill: '#262626',
                            textAlign: 'center',
                        }}
                    />
                    <Interaction type='element-single-selected' />
                    {/*<Legend*/}
                    {/*    position="right"*/}
                    {/*    offsetY={-window.innerHeight/2+230}*/}
                    {/*    offsetX={-20}*/}
                    {/*/>*/}
                    {/*<Tooltip*/}
                    {/*    showTitle={false}*/}
                    {/*    itemTpl="<li><span style=&quot;background-color:{color};&quot; class=&quot;g2-tooltip-market&quot;></span>{name}: {value}</li>"*/}
                    {/*/>*/}
                    {/*<Guide>*/}
                    {/*    <Html*/}
                    {/*        position={["50%", "50%"]}*/}
                    {/*        html={guideHtml}*/}
                    {/*        alignX="middle"*/}
                    {/*        alignY="middle"*/}
                    {/*    />*/}
                    {/*</Guide>*/}
                    {/*<Geom*/}
                    {/*    type="intervalStack"*/}
                    {/*    postion="count"*/}
                    {/*    color="item"*/}
                    {/*    style={{*/}
                    {/*        lineWidth: 1,*/}
                    {/*        stroke: "#fff"*/}
                    {/*    }}*/}
                    {/*>*/}
                    {/*    <Label*/}
                    {/*        content="percent"*/}
                    {/*        formatter={(val, item) => {*/}
                    {/*            return item.point.item + ": " + val*/}
                    {/*        }}*/}
                    {/*    />*/}
                    {/*</Geom>*/}
                </Chart>
            </div>
        );
    }
}

class OrderChart extends React.Component{
    render() {
        const { DataView } = DataSet;
        const dv = new DataView();
        dv.source(this.props.data).transform({
            type: "percent",
            field: "count",
            dimension: "goodsType",
            as: "percent"
        });
        const cols = {
            percent: {
                formatter: val =>{
                    val = (val*100).toFixed(2)+"%";
                    return val;
                }
            }
        };
        return (
            <div>
                <h1>订单类型分布</h1>
                <Chart
                    height={400}
                    data={dv}
                    scale={cols}
                    autoFit
                >
                    <Legend visible={false} />
                    <Coordinate type="theta" innerRadius={0.4}/>
                    <Axis visible={false}/>
                    <Interval
                        position="percent"
                        adjust="stack"
                        color="goodsType"
                        style={{
                            lineWidth: 1,
                            stroke: '#fff',
                        }}
                        label={['count', {
                            content: (data) => {
                                return `${data.goodsType}: ${(data.percent * 100).toFixed(2)}%`;
                            },
                        }]}
                    />
                    <Annotation.Text
                        position={['50%', '50%']}
                        content={this.props.total}
                        style={{
                            lineHeight: '240px',
                            fontSize: '30',
                            fill: '#262626',
                            textAlign: 'center',
                        }}
                    />
                    <Interaction type='element-single-selected' />
                </Chart>
            </div>
        );
    }
}

class Home extends React.Component{
    state = {
        userTotal: 0,
        goodsData: [],
        goodsTotal: 0,
        orderData: [],
        orderTotal: 0,
    };

    componentDidMount() {
        this.props.changePagePath('概览');
        this.updateData()
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
                let count = 0;
                for (let i in d.data.data.data) {
                    let c = d.data.data.data[i].count;
                    if (c!==undefined){
                        count += c;
                    }else {
                        d.data.data.data[i].count = 0;
                    }
                }
                this.setState({goodsData: d.data.data.data, goodsTotal:count})
            }
        });
        get({
            url: serviceDomain + '/v1/statistic/order',
            callback: (d) => {
                let count = 0;
                for (let i in d.data.data.data) {
                    let c = d.data.data.data[i].count;
                    if (c!==undefined){
                        count += c;
                    }else {
                        d.data.data.data[i].count = 0;
                    }
                }
                this.setState({orderData: d.data.data.data, orderTotal:count})
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
                        <label style={{minWidth: 100, display:'inline-block', marginTop:'10px'}}>用户总数</label>
                        <label style={{minWidth: 100, display:'inline-block', marginLeft:'250px'}}>商品总数</label>
                        <label style={{minWidth: 100, display:'inline-block', marginLeft:'250px'}}>订单总数</label>
                        <br/>
                        <div style={{width:'100px',display:'inline-block'}}>
                            <label style={{fontSize:'20px'}}>{this.state.userTotal}</label><label>个</label>
                        </div>
                        <div style={{width:'100px',display:'inline-block', margin:'10px 0 0 250px'}}>
                            <label style={{fontSize:'20px'}}>{this.state.goodsTotal}</label><label>个</label>
                        </div>
                        <div style={{width:'100px',display:'inline-block', margin:'10px 0 0 250px'}}>
                            <label style={{fontSize:'20px',display:'inline-block'}}>{this.state.orderTotal}</label>
                            <label style={{fontSize:'18px',display:'inline-block'}}>个</label>
                        </div>
                    </Card>
                </div>
                <div style={{fontSize:'24px', marginLeft:'20px', marginTop:'20px', height:'500px', width:window.innerWidth-272}}>
                    <div style={{width:(window.innerWidth-278)/2, height:'500px', display:'inline-block'}}>
                        <Card style={{width:(window.innerWidth-278)/2, height:'500px'}}>
                            <GoodsChart data={this.state.goodsData} total={this.state.goodsTotal}/>
                        </Card>
                    </div>
                    <div style={{width:100, height:'500px', display:'inline-block', marginLeft:'8px'}}>
                        <Card style={{width:(window.innerWidth-278)/2, height:'500px'}}>
                            <OrderChart data={this.state.orderData} total={this.state.orderTotal}/>
                        </Card>
                    </div>
                </div>
            </div>
        )
    }
}

export default Home;