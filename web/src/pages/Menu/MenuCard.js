import React, { PureComponent } from 'react';
import { connect } from 'dva';
import {
  Form,
  Input,
  Card,
  Radio,
  Modal,
  Icon,
  TreeSelect,
  Tooltip,
  InputNumber,
  Row,
  Col,
} from 'antd';

import MenuAction from './MenuAction';
import MenuResource from './MenuResource';

@connect(({ menu }) => ({
  menu,
}))
@Form.create()
class MenuCard extends PureComponent {
  onOKClick = () => {
    const { form, onSubmit } = this.props;
    form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        const formData = { ...values };
        formData.hidden = parseInt(formData.hidden, 10);
        formData.sequence = parseInt(formData.sequence, 10);
        onSubmit(formData);
      }
    });
  };

  dispatch = action => {
    const { dispatch } = this.props;
    dispatch(action);
  };

  toTreeSelect = data => {
    if (!data) {
      return [];
    }
    const newData = [];
    for (let i = 0; i < data.length; i += 1) {
      const item = { ...data[i], title: data[i].name, value: data[i].record_id };
      if (item.children && item.children.length > 0) {
        item.children = this.toTreeSelect(item.children);
      }
      newData.push(item);
    }
    return newData;
  };

  render() {
    const {
      menu: { formVisible, formTitle, formData, submitting, treeData },
      form: { getFieldDecorator },
      onCancel,
    } = this.props;

    const formItemLayout = {
      labelCol: {
        span: 6,
      },
      wrapperCol: {
        span: 18,
      },
    };

    return (
      <Modal
        title={formTitle}
        width={900}
        visible={formVisible}
        maskClosable={false}
        confirmLoading={submitting}
        destroyOnClose
        onOk={this.onOKClick}
        onCancel={onCancel}
        style={{ top: 20 }}
        bodyStyle={{ maxHeight: 'calc( 100vh - 158px )', overflowY: 'auto' }}
      >
        <Card bordered={false}>
          <Form>
            <Row>
              <Col span={12}>
                <Form.Item {...formItemLayout} label="メニュー名">
                  {getFieldDecorator('name', {
                    initialValue: formData.name,
                    rules: [
                      {
                        required: true,
                        message: 'メニュー名を入力してください',
                      },
                    ],
                  })(<Input placeholder="入力してください" />)}
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item {...formItemLayout} label="親メニュー">
                  {getFieldDecorator('parent_id', {
                    initialValue: formData.parent_id,
                  })(
                    <TreeSelect
                      showSearch
                      treeNodeFilterProp="title"
                      style={{ width: '100%' }}
                      dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                      treeData={this.toTreeSelect(treeData)}
                      placeholder="選択してください"
                    />
                  )}
                </Form.Item>
              </Col>
            </Row>
            <Row>
              <Col span={12}>
                <Form.Item {...formItemLayout} label="アイコン">
                  <Row>
                    <Col span={20}>
                      {getFieldDecorator('icon', {
                        initialValue: formData.icon,
                        rules: [
                          {
                            required: true,
                            message: 'アイコンを入力してください',
                          },
                        ],
                      })(<Input placeholder="入力してください" />)}
                    </Col>
                    <Col span={4} style={{ textAlign: 'center' }}>
                      <Tooltip title="公式アイコンのみ使用可">
                        <Icon type="question-circle" />
                      </Tooltip>
                    </Col>
                  </Row>
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item {...formItemLayout} label="ルーター">
                  {getFieldDecorator('router', {
                    initialValue: formData.router,
                  })(<Input placeholder="入力してください" />)}
                </Form.Item>
              </Col>
            </Row>
            <Row>
              <Col span={12}>
                <Form.Item {...formItemLayout} label="表示順">
                  {getFieldDecorator('sequence', {
                    initialValue: formData.sequence ? formData.sequence.toString() : '1000000',
                    rules: [
                      {
                        required: true,
                        message: '表示順を入力してください',
                      },
                    ],
                  })(<InputNumber min={1} style={{ width: '80%' }} />)}
                </Form.Item>
              </Col>
              <Col span={12}>
                <Form.Item {...formItemLayout} label="ステータス">
                  {getFieldDecorator('hidden', {
                    initialValue: formData.hidden ? formData.hidden.toString() : '0',
                  })(
                    <Radio.Group>
                      <Radio value="0">表示</Radio>
                      <Radio value="1">非表示</Radio>
                    </Radio.Group>
                  )}
                </Form.Item>
              </Col>
            </Row>
            <Row>
              <Col span={24}>
                <Card title="アクション管理" bordered={false}>
                  {getFieldDecorator('actions', {
                    initialValue: formData.actions,
                  })(<MenuAction />)}
                </Card>
              </Col>
            </Row>
            <Row>
              <Col span={24}>
                <Card title="リソース管理" bordered={false}>
                  {getFieldDecorator('resources', {
                    initialValue: formData.resources,
                  })(<MenuResource />)}
                </Card>
              </Col>
            </Row>
          </Form>
        </Card>
      </Modal>
    );
  }
}

export default MenuCard;
