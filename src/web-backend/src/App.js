import React, {Component} from 'react';
import './App.css';
import {Layout, Breadcrumb, Card, Menu} from "antd";
import Routers from './routers'
import {Link} from 'react-router-dom'
import DocumentTitle from 'react-document-title'
import {instanceOf} from "prop-types";
import {Cookies, withCookies} from "react-cookie";
import {cookieMaxAge, serviceDomain} from "./axios/config";
import {get} from "./axios/tools";
import LoginForm from "./components/forms/LoginForm"


const {Content, Sider, Header} = Layout;

class App extends Component{
  static propTypes = {
    cookies: instanceOf(Cookies).isRequired
  };

  constructor(props) {
    super(props);
    const {cookies} = props;
    console.log(props);
    this.state = {
      userData:{},
      secondPagePath: '概览',
      title: '',
      isSuper:false,
      userName: '',
      isLogin: cookies.get('isLogin') || false,
      popoverVisible: false
    }
  }

  showPopover = () => {
    this.setState({
      popoverVisible: true
    });
  };

  handleVisibleChange = visible =>{
    this.setState({popoverVisible: visible});
  };

  updatePagePath(secondPagePath) {
    this.setState({secondPagePath: secondPagePath})
  }

  updateUserInfo() {
    const {cookies} = this.props;
    get({
      url: serviceDomain+'/v1/sys/user?pageSize=1&pageIndex=1&userName='+cookies.get('userName'),
      callback: (d) =>{
        this.setState({userData: d.data.data.data[0]})
      }
    })
  }

  updateLoginState(isLogin, userName) {
    const {cookies} = this.props;
    if (isLogin===true){
      cookies.set('isLogin', true, {path:'/', maxAge:cookieMaxAge});
      cookies.set('userName', true, {path:'/', maxAge:cookieMaxAge});
      this.updateUserInfo();
    }else {
      cookies.remove('isLogin', {page:'/'});
      cookies.remove('userName', {page:'/'});
    }
  }

  componentDidMount() {
    const {cookies} = this.props;
    this.setState({isLogin: cookies.get('isLogin'), userName: cookies.get('userName')});
    if (cookies.get('isLogin')==='true'){
      this.updateUserInfo();
    }
  }

  render() {
    const {cookies} = this.props;
    if (cookies.get('isLogin')==='true'){
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
    }else {
      return (
          <DocumentTitle>
            <Layout>
              <LoginForm updateLoginStatus={this.updateLoginState.bind(this)}/>
            </Layout>
          </DocumentTitle>
      )
    }
  }
}

export default withCookies(App);
