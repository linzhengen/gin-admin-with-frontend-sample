import React, { PureComponent } from 'react';
import { Form, Input, Modal, message } from 'antd';
import { md5Hash } from '../../utils/utils';
import { updatePwd } from '@/services/login';

@Form.create()
class UpdatePasswordDialog extends PureComponent {
  state = {
    submitting: false,
  };

  onOKClick = () => {
    const { form } = this.props;

    form.validateFieldsAndScroll((err, values) => {
      if (err) {
        return;
      }
      if (values.new_password !== values.confirm_new_password) {
        message.warning('２回入力した新しいパスワードは一致していません');
        return;
      }

      this.setState({ submitting: true });
      const formData = {
        old_password: md5Hash(values.old_password),
        new_password: md5Hash(values.new_password),
      };
      updatePwd(formData).then(res => {
        if (res.status === 'OK') {
          message.success('パスワード更新に成功しました');
          this.handleCancel();
        }
        this.setState({ submitting: false });
      });
    });
  };

  handleCancel = () => {
    const { onCancel } = this.props;
    if (onCancel) {
      onCancel();
    }
  };

  dispatch = action => {
    const { dispatch } = this.props;
    dispatch(action);
  };

  render() {
    const {
      visible,
      form: { getFieldDecorator },
    } = this.props;

    const { submitting } = this.state;

    const formItemLayout = {
      labelCol: {
        span: 6,
      },
      wrapperCol: {
        span: 16,
      },
    };

    return (
      <Modal
        title="パスワード変更"
        width={450}
        visible={visible}
        maskClosable={false}
        confirmLoading={submitting}
        destroyOnClose
        onOk={this.onOKClick}
        onCancel={this.handleCancel}
        style={{ top: 20 }}
        bodyStyle={{ maxHeight: 'calc( 100vh - 158px )', overflowY: 'auto' }}
      >
        <Form>
          <Form.Item {...formItemLayout} label="現在のパスワード">
            {getFieldDecorator('old_password', {
              rules: [
                {
                  required: true,
                  message: '現在のパスワードを入力してください',
                },
              ],
            })(<Input type="password" placeholder="現在のパスワード" />)}
          </Form.Item>
          <Form.Item {...formItemLayout} label="新しいパスワード">
            {getFieldDecorator('new_password', {
              rules: [
                {
                  required: true,
                  message: '新しいパスワードを入力してください',
                },
              ],
            })(<Input type="password" placeholder="新しいパスワード" />)}
          </Form.Item>
          <Form.Item {...formItemLayout} label="現在のパスワードの確認">
            {getFieldDecorator('confirm_new_password', {
              rules: [
                {
                  required: true,
                  message: '現在のパスワードを入力してください',
                },
              ],
            })(<Input type="password" placeholder="現在のパスワード" />)}
          </Form.Item>
        </Form>
      </Modal>
    );
  }
}

export default UpdatePasswordDialog;
