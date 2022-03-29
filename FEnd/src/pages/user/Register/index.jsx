import {
  AlipayCircleOutlined,
  LockOutlined,
  MobileOutlined,
  TaobaoCircleOutlined,
  MailOutlined,
  FieldStringOutlined,
  UserOutlined,
  WeiboCircleOutlined,
} from '@ant-design/icons';
import { Alert, message, Tabs } from 'antd';
import React, { useState } from 'react';
import { ProFormCaptcha, ProFormCheckbox, ProFormText, LoginForm } from '@ant-design/pro-form';
import { useIntl, history, FormattedMessage, SelectLang, useModel } from 'umi';
import Footer from '@/components/Footer';
import { login, LoginUser,CreateUser } from '@/services/ant-design-pro/api';
import { getFakeCaptcha } from '@/services/ant-design-pro/login';
import styles from './index.less';
import { PathRegisterContext } from 'rc-menu/lib/context/PathContext';

const LoginMessage = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Register = () => {
  const [type, setType] = useState('register');
  const intl = useIntl();

  const handleSubmit = async (values) => {
      const { personal_number, name, email, password } = values;
      try {
        // 登录
        const msg = await CreateUser({ personal_number, name, email, password });

        if (msg.status === 'ok') {
          const defaultLoginSuccessMessage = intl.formatMessage({
            id: 'pages.login.success',
            defaultMessage: 'Registered Successfully！',
          });
          message.success(defaultLoginSuccessMessage);
          history.push('/user/login');
          
          /** 此方法会跳转到 redirect 参数所在的位置 */
        }
      } catch (error) {
        const defaultLoginFailureMessage = intl.formatMessage({
          id: 'pages.login.failure',
          defaultMessage: 'Register Failed！',
        });
        message.error(defaultLoginFailureMessage);
      }
  };

  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <LoginForm
          logo={<img alt="logo" src="/logo.svg" />}
          title="Ant Design/LatihanFSE"
          submitter={{
            render: (props, doms) => {
              console.log(props);
              return (
                <div style={{ width: '100%' }}>
                  <button
                    type="button "
                    key="submit "
                    style={{
                      width: '100%',
                      backgroundColor: '#1890ff',
                      border: 'none',
                      padding: 8,
                      color: 'white',
                    }}
                  >
                    Register
                  </button>
                </div>
              );
            },
          }}
          subTitle={intl.formatMessage({
            id: 'pages.layouts.userLayout.title',
          })}
          // actions={[
          //   <FormattedMessage
          //     key="loginWith"
          //     id="pages.login.loginWith"
          //     defaultMessage="其他登录方式"
          //   />,
          //   <AlipayCircleOutlined key="AlipayCircleOutlined" className={styles.icon} />,
          //   <TaobaoCircleOutlined key="TaobaoCircleOutlined" className={styles.icon} />,
          //   <WeiboCircleOutlined key="WeiboCircleOutlined" className={styles.icon} />,
          // ]}
          onFinish={async (values) => {
            await handleSubmit(values);
          }}
        >
          <Tabs activeKey={type} onChange={setType}>
            <Tabs.TabPane
              key="register"
              tab="Register"
            />
          </Tabs>

          {status === 'error' && loginType === 'login' && (
            <LoginMessage
              content={intl.formatMessage({
                id: 'pages.login.accountLogin.errorMessage',
                defaultMessage: '账户或密码错误(admin/ant.design)',
              })}
            />
          )}
          {type === 'login' && (
            <>
              <ProFormText
                name="personal_number"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder="Personal Number"
                rules={[
                  {
                    required: true,
                    message: "Personal Number Required!",
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                placeholder="Password"
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.password.required"
                        defaultMessage="请输入密码！"
                      />
                    ),
                  },
                ]}
              />
            </>
          )}

          {status === 'error' && loginType === 'register' && <LoginMessage content="验证码错误" />}
          {type === 'register' && (
            <>
            <ProFormText
              name="personal_number"
              fieldProps={{
                size: 'large',
                prefix: <UserOutlined className={styles.prefixIcon} />,
              }}
              placeholder="Personal Number"
              rules={[
                {
                  required: true,
                  message: "Personal Number Required!",
                },
              ]}
            />
            <ProFormText
              name="name"
              fieldProps={{
                size: 'large',
                prefix: <FieldStringOutlined className={styles.prefixIcon} />,
              }}
              placeholder="Name"
              rules={[
                {
                  required: true,
                  message: "Name Required!",
                },
              ]}
            />
            <ProFormText
              name="email"
              fieldProps={{
                size: 'large',
                prefix: <MailOutlined className={styles.prefixIcon} />,
              }}
              placeholder="Email"
              rules={[
                {
                  required: true,
                  message: "Email Required!",
                },
              ]}
            />
            <ProFormText.Password
              name="password"
              fieldProps={{
                size: 'large',
                prefix: <LockOutlined className={styles.prefixIcon} />,
              }}
              placeholder="Password"
              rules={[
                {
                  required: true,
                  message: (
                    <FormattedMessage
                      id="pages.login.password.required"
                      defaultMessage="请输入密码！"
                    />
                  ),
                },
              ]}
            />
          </>
          )}
          <div
            style={{
              marginBottom: 24,
              paddingBottom: 24
            }}
          >
            <a
              style={{
                float: 'right',
              }}
              onClick={() => history.push('/user/login')}
            >
              <FormattedMessage
                id="pages.login.login"
                defaultMessage="Already Registered? Login now!"
              />
            </a>
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Register;
