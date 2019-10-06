import React, { PureComponent } from 'react';
import { connect } from 'dva';
import moment from 'moment';
import PageHeaderLayout from '../../layouts/PageHeaderLayout';
import styles from './Home.less';

@connect(state => ({
  global: state.global,
}))
class Home extends PureComponent {
  state = {
    currentTime: moment().format('HH:mm:ss'),
  };

  componentDidMount() {
    this.interval = setInterval(() => {
      this.setState({ currentTime: moment().format('HH:mm:ss') });
    }, 1000);
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  getHeaderContent = () => {
    const {
      global: { user },
    } = this.props;

    const { role_names: roleNames } = user;

    const text = [];
    if (roleNames && roleNames.length > 0) {
      text.push(
        <span key="role" style={{ marginRight: 20 }}>{`所属ロール：${roleNames.join('/')}`}</span>
      );
    }

    if (text.length > 0) {
      return text;
    }
    return null;
  };

  render() {
    const {
      global: { user },
    } = this.props;

    const { currentTime } = this.state;

    const breadcrumbList = [{ title: 'トップ' }];

    return (
      <PageHeaderLayout
        title={`Hello，${user.real_name}，よい毎日を！`}
        breadcrumbList={breadcrumbList}
        content={this.getHeaderContent()}
        action={<span>現在時間：{currentTime}</span>}
      >
        <div className={styles.index} />
      </PageHeaderLayout>
    );
  }
}

export default Home;
