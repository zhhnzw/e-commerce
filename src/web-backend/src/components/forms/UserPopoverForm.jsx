import React from 'react';
import {Button, Modal, Form, Input, message} from "antd";
import {serviceDomain} from "../../axios/config";
import md5 from 'md5'
import {post} from "../../axios/tools";
import Cookies from 'universal-cookie'

const cookies = new Cookies();

const CollectionCreateForm = Form.create({name:'form_in_modal'})(
    class extends React.Component {
        render() {
            const {visible, onCancel, onCreate, form} = this.props;
            const {getFieldDecorator} = form;
            return (
                <Modal
                    visible={visible}
                    title="修改密码"
                    okText="确认"
                    onCancel={onCancel}
                    onOk={onCreate}
                >
                    <Form layout="vertical">
                        <Form.Item label="原密码">
                            {getFieldDecorator('oldPwd', {
                                rules: [{required: true, message: '请输入旧密码!'}],
                            })(<Input type="password"/>)}
                        </Form.Item>
                        <Form.Item label="新密码">
                            {getFieldDecorator('newPwd', {
                                rules: [{required: true, message: '请输入新密码!'}],
                            })(<Input type="password"/>)}
                        </Form.Item>
                        <Form.Item label="确认密码">
                            {getFieldDecorator('confirmPwd', {
                                rules: [{required: true, message: '请输入新密码!'}],
                            })(<Input type="password"/>)}
                        </Form.Item>
                    </Form>
                </Modal>
            )
        }
    }
);

class UserPopoverForm extends React.Component{
    state = {
        visible: false,
    };

    showModal = () => {
        this.props.parentVisible(false);
        this.setState({visible:true})
    };

    handleCancel = () => {
        this.setState({visible:false})
    };

    handleCreate = () => {
        const {form} = this.formRef.props;
        form.validateFields((err, values) => {
            if (!err){
                console.log('Received values of form:', values);
                if (values.newPwd !== values.confirmPwd){
                    message.warn("确认密码输入不相同!");
                    return
                }
                let url = serviceDomain + '/v1/alterPwd';
                let data = {
                    "oldPwd": md5(values.oldPwd),
                    "newPwd": md5(values.newPwd),
                };
                post({
                    url:url,
                    data:data,
                    callback:(d)=>{
                        if(d.data.code==='0'){
                            message.info(d.data.msg);
                            this.setState({visible:false});
                        }else {
                            message.warn(d.data.msg);
                        }
                    }
                });
            }
        })
    };

    saveFormRef = formRef => {
        this.formRef = formRef;
    };

    logout = () => {
        let url = serviceDomain + '/v1/logout';
        post({
            url:url,
            callback:(d)=>{
                if(d.data.code==='0'){
                    this.setState({visible:false});
                    cookies.remove('isLogin', {path:'/'});
                    cookies.remove('userName', {path:'/'});
                    window.location.href = "/"
                }else {
                    message.warn(d.data.msg);
                }
            }
        });
    };

    render() {
        return (
            <div>
                <div style={{display:'flex', justifyContent:'center', alignItems:'center', marginTop:'12px'}}>
                    <Button onClick={this.showModal}>
                        修改密码
                    </Button>
                    <Button onClick={this.logout}>退出登录</Button>
                </div>
                <CollectionCreateForm
                    wrappedComponentRef={this.saveFormRef}
                    visible={this.state.visible}
                    onCancel={this.handleCancel}
                    onCreate={this.handleCreate}
                />
            </div>
        );
    }

}

export default UserPopoverForm;