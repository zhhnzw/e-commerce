import React, {Component} from 'react';
import './App.css';
import {Layout, Breadcrumb, Card, Menu} from "antd";
import Routers from './routers'
import {Link} from 'react-router-dom'
import DocumentTitle from 'react-document-title'
import {instanceOf} from "prop-types";
import {Cookies} from "react-cookie";

const {Content, Sider, Header} = Layout;

class App extends Component{
  static propTypes = {
    cookies: instanceOf(Cookies).isRequired
  };

  constructor(props) {
    super(props);
    const {cookies} = props;
    this.state = {
      secondPagePath: '概览',
      title: ''
    }
  }

  updatePagePath(secondPagePath) {
    this.setState({secondPagePath: secondPagePath})
  }

  render() {
    return (
        <DocumentTitle>
          <Layout>
            <Header style={{height:'54px'}}>
              <div style={{textAlign:'center', height:'54px', lineHeight:'54px'}}>
                <label style={{frontSize:'16px', color:'white'}}>电商</label>
              </div>
            </Header>
            <Layout>
              <Sider style={{backgroundColor: '#f7f7f7', height:window.innerHeight-54}}>
                <Card bordered={false} title='菜单' bodyStyle={{paddingLeft:'0', paddingTop:'1px', paddingBottom:'0'}} headStyle={{backgroundColor: '#f7f7f7'}}>
                  <Menu
                      mode='inline'
                      defaultSelectedKeys={['1']}
                      defaultOpenKeys={['sub1']}
                      style={{width:'200px', borderRight: 0, backgroundColor:'#f7f7f7'}}
                  ><Menu.Item key={'/'}>
                    <Link to={"/"}><span>概览</span></Link>
                  </Menu.Item>
                    <Menu.Item key="/goods/list">
                      <Link to={"/goods/list"}><span>商品</span></Link>
                    </Menu.Item>
                  </Menu> {/* borderRight:0 隐藏菜单右侧灰色的边线 */}
                </Card>
              </Sider>
              <Layout>
                <Content style={{backgroundColor:'white'}}>
                  <Breadcrumb style={{margin:'20px 0 10px 20px', fontSize:'12px'}}>
                    <Breadcrumb.Item>电商</Breadcrumb.Item>
                    <Breadcrumb.Item>{this.state.secondPagePath}</Breadcrumb.Item>
                  </Breadcrumb>
                  <Routers changePagePath={this.updatePagePath.bind(this)}/>
                </Content>
              </Layout>
            </Layout>
          </Layout>
        </DocumentTitle>
    )
  }

}

export default App;
