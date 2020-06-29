import React, {Component} from 'react';
import {Form, Icon, Input, Button, message} from "antd";
import {post} from "../../axios/tools"
import {serviceDomain} from "../../axios/config";
import md5 from "md5";

class NormalLoginForm extends Component {
    handleSubmit = (e) => {
        e.preventDefault();
        this.props.form.validateFields((err, values) => {
            if (!err){
                console.log("receive values of form: ", values);
                let url = serviceDomain + "/v1/login";
                let data = {
                    "userName": values.userName,
                    "password": md5(values.password)
                };
                post({
                    url: url,
                    data: data,
                    callback: (d)=>{
                        if (d.data.code==='0'){
                            this.props.updateLoginStatus(true, data.userName);
                        }else {
                            message.warn(d.data.msg);
                        }
                    }
                })
            }
        })
    };

    render() {
        const {getFieldDecorator} = this.props.form;
        return (
            <div style={{width:window.innerWidth, height:window.innerHeight}}>
                <div style={{width:'400px', height:'400px', position:'absolute', left:0,top:0,bottom:0,right:0,margin:'auto',backgroundColor:'white'}}>
                    <label style={{marginTop:'50px', float:'left',position:'relative',marginLeft:'37%',fontSize:'30px'}}>管理后台</label>
                    <Form onSubmit={this.handleSubmit} style={{width:'300px', float:'left', position:'relative',textAlign:'center',margin:'auto',top:'20px',marginLeft:'13%'}} className="login-form">
                        <Form.Item>
                            {getFieldDecorator('userName', {
                                rules: [{required:true, message:'请输入用户名!'}],
                            })(
                                <Input prefix={<Icon type="user" style={{color:'rgba(0,0,0,.25)'}}/>} placeholder="用户名"/>
                            )}
                        </Form.Item>
                        <Form.Item>
                            {getFieldDecorator('password', {
                                rules: [{required:true, message:'请输入密码!'}],
                            })(
                                <Input prefix={<Icon type="lock" style={{color:'rgba(0,0,0,.25)'}}/>} type="password" placeholder="密码"/>
                            )}
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit" className="login-form-button" style={{width:'100%'}}>
                                登录
                            </Button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        )
    }
}

const LoginForm = Form.create()(NormalLoginForm);
export default LoginForm;