import React, { PureComponent } from 'react';
import { connect } from 'dva';
import { Form, Input, Modal, Radio } from 'antd';
import { md5Hash } from '../../utils/utils';
import RoleSelect from './RoleSelect';

@connect(state => ({
  user: state.user,
}))
@Form.create()
class UserCard extends PureComponent {
  onOKClick = () => {
    const { form, onSubmit } = this.props;

    form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        const formData = { ...values };
        formData.status = parseInt(formData.status, 10);
        if (formData.password && formData.password !== '') {
          formData.password = md5Hash(formData.password);
        }
        onSubmit(formData);
      }
    });
  };

  dispatch = action => {
    const { dispatch } = this.props;
    dispatch(action);
  };

  render() {
    const {
      onCancel,
      user: { formType, formTitle, formVisible, formData, submitting },
      form: { getFieldDecorator },
    } = this.props;

    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 6 },
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 },
      },
    };

    return (
      <Modal
        title={formTitle}
        width={600}
        visible={formVisible}
        maskClosable={false}
        confirmLoading={submitting}
        destroyOnClose
        onOk={this.onOKClick}
        onCancel={onCancel}
        style={{ top: 20 }}
        bodyStyle={{ maxHeight: 'calc( 100vh - 158px )', overflowY: 'auto' }}
      >
        <Form>
          <Form.Item {...formItemLayout} label="ユーザ名">
            {getFieldDecorator('user_name', {
              initialValue: formData.user_name,
              rules: [
                {
                  required: true,
                  message: 'ユーザ名を入力してください',
                },
              ],
            })(<Input placeholder="ユーザ名入力" />)}
          </Form.Item>
          <Form.Item {...formItemLayout} label="ログインパスワード">
            {getFieldDecorator('password', {
              initialValue: formData.password,
              rules: [
                {
                  required: formType === 'A',
                  message: 'ログインパスワードを入力してください',
                },
              ],
            })(
              <Input
                type="password"
                placeholder={
                  formType === 'A'
                    ? 'ログインパスワードを入力してください'
                    : '未入力の場合は更新しない'
                }
              />
            )}
          </Form.Item>
          <Form.Item {...formItemLayout} label="本名">
            {getFieldDecorator('real_name', {
              initialValue: formData.real_name,
              rules: [
                {
                  required: true,
                  message: '本名を入力してください',
                },
              ],
            })(<Input placeholder="本名入力" />)}
          </Form.Item>
          <Form.Item {...formItemLayout} label="所属ロール">
            {getFieldDecorator('roles', {
              initialValue: formData.roles,
              rules: [
                {
                  required: true,
                  message: 'ロールを選択してください',
                },
              ],
            })(<RoleSelect />)}
          </Form.Item>
          <Form.Item {...formItemLayout} label="ユーザステータス">
            {getFieldDecorator('status', {
              initialValue: formData.status ? formData.status.toString() : '1',
            })(
              <Radio.Group>
                <Radio value="1">有効</Radio>
                <Radio value="2">無効</Radio>
              </Radio.Group>
            )}
          </Form.Item>
          <Form.Item {...formItemLayout} label="メールアドレス">
            {getFieldDecorator('email', {
              initialValue: formData.email,
            })(<Input placeholder="メールアドレス" />)}
          </Form.Item>
          <Form.Item {...formItemLayout} label="電話番号">
            {getFieldDecorator('phone', {
              initialValue: formData.phone,
            })(<Input placeholder="電話番号を入力してください" />)}
          </Form.Item>
        </Form>
      </Modal>
    );
  }
}

export default UserCard;
