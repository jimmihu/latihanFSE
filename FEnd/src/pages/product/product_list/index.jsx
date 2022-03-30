import { PlusOutlined,EditOutlined,DeleteOutlined,EyeOutlined,LockOutlined,FieldStringOutlined,UserOutlined,MailOutlined } from '@ant-design/icons';
import { Button, message, Input, Drawer } from 'antd';
import React, { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormText, ProFormTextArea } from '@ant-design/pro-form';
import ProDescriptions from '@ant-design/pro-descriptions';
import UpdateForm from './components/UpdateForm';
import { GetProductList,GetProductDetail,CreateProduct,DeleteProduct } from '@/services/ant-design-pro/api';
import styles from './index.less';
/**
 * @en-US Add node
 * @zh-CN 添加节点
 * @param fields
 */


const handleCreate = async (value) => {
  const load = message.loading('Creating Product..');
  try{
    const res = await CreateProduct(value);
    if(res.status === 'ok'){
      load();
      message.success('Product Created Successfully');
    }
    return true
  }catch(error){
    load();
    message.error('Failed to Create Product');
    return true
  }
};
/**
 * @en-US Update node
 * @zh-CN 更新节点
 *
 * @param fields
 */

const handleUpdate = async (fields) => {
};
/**
 *  Delete node
 * @zh-CN 删除节点
 *
 * @param selectedRows
 */

const handleRemove = async (id) => {
  const load = message.loading('Deleting Product..');
  try{
    const res = await DeleteProduct(id);
    if(res.status === 'ok'){
      load();
      message.success('Product Deleted Successfully');
    }
    return true
  }catch(error){
    load();
    message.error('Failed to Delete Product');
    return true
  }
};

const ProductList = () => {
  const handleDetail = async (id) => {
    try{
      const res = await GetProductDetail(id);
      if(res.status === "ok"){
        setCurrentRow(res.data);
        setShowDetail(true);
      }
    }catch(error){
      message.error('User Detail Not Found');
    }
  }
  /**
   * @en-US Pop-up window of new window
   * @zh-CN 新建窗口的弹窗
   *  */
  const [createModalVisible, handleModalVisible] = useState(false);
  /**
   * @en-US The pop-up window of the distribution update window
   * @zh-CN 分布更新窗口的弹窗
   * */

  const [updateModalVisible, handleUpdateModalVisible] = useState(false);
  const [showDetail, setShowDetail] = useState(false);
  const actionRef = useRef();
  const [currentRow, setCurrentRow] = useState();
  const [selectedRowsState, setSelectedRows] = useState([]);
  /**
   * @en-US International configuration
   * @zh-CN 国际化配置
   * */

  const intl = useIntl();
  const columns = [
    {
      title: "ID",
      dataIndex: 'id',
      tip: 'The ID is the unique key',
      valueType: 'text',
    },
    {
      title: 'Name',
      dataIndex: 'name',
      valueType: 'text',
    },
    {
      title: 'Description',
      dataIndex: 'description',
      valueType: 'textarea',
    },
    {
      title: 'Active',
      dataIndex: 'active',
      valueType: 'text',
      render: (_, rowData) => {
        return <>{rowData.active ? 'Active' : 'Inactive'}</>;
      },
    },
    {
      title: <FormattedMessage id="pages.searchTable.titleOption" defaultMessage="Operating" />,
      dataIndex: 'option',
      valueType: 'option',
      render: (_, rowData) => {
        return (<>
            <Button
              type="dashed"
              key="dashed"
              title="Detail"
              onClick={() => {
                handleDetail(rowData.id)
                actionRef.current.reload();
              }}
            >
              <EyeOutlined />
            </Button>
            <Button
              type="primary"
              key="primary"
              title="Edit"
              onClick={() => {
                handleUpdateModalVisible(true)
              }}
            >
              <EditOutlined />
            </Button>
            &nbsp;
            <Button
              type="danger"
              key="danger"
              onClick={() => {
                handleRemove(rowData.id)
                actionRef.current.reload();
              }}
            >
              <DeleteOutlined />
            </Button>
          </>
        );
      },
    },
  ];

  const columns2 = [
    {title: 'ID',dataIndex: 'id',valueType: 'text',},
    {title: 'Name',dataIndex: 'name',valueType: 'text',},
    {title: 'Description',dataIndex: 'description',valueType: 'textarea',},
    {title: 'Active',dataIndex: 'active',valueType: 'text',
      render: (_, rowData) => {
        return <>{rowData.active ? 'Active' : 'Inactive'}</>;
      },
    },
    {title: 'Maker',dataIndex: 'Maker',valueType: 'text',
      render: (_, rowData) => {
        return <>{rowData.Maker.name}</>;
      },
    },
    {title: 'Checker',dataIndex: 'Checker',valueType: 'text',
      render: (_, rowData) => {
        return <>{rowData.Checker.name}</>;
      },
    },
    {title: 'Signer',dataIndex: 'Signer',valueType: 'text',
      render: (_, rowData) => {
        return <>{rowData.Signer.name}</>;
      },
    },
  ]
  return (
    <PageContainer>
      <ProTable
        headerTitle="Product List Table"
        actionRef={actionRef}
        rowKey="key"
        search={false}
        toolBarRender={() => [
          <Button
            type="primary"
            key="primary"
            onClick={() => {
              handleModalVisible(true);
            }}
          >
            <PlusOutlined /> Create Product
          </Button>,
        ]}
        request={GetProductList}
        columns={columns}
      />
      <ModalForm
        title="Create Product"
        width="400px"
        visible={createModalVisible}
        onVisibleChange={handleModalVisible}
        onFinish={async (value) => {
          const success = await handleCreate(value);

          if (success) {
            handleModalVisible(false);

            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
      >
        <ProFormText
          rules={[
            {
              required: true,
              message: 'This field required',
            },
          ]}
          fieldProps={{
            size: 'large',
            prefix: <FieldStringOutlined className={styles.prefixIcon} />,
          }}
          placeholder="Name"
          width="md"
          name="name"
        />
        <ProFormText
          rules={[
            {
              required: true,
              message: 'This field required',
            },
          ]}
          fieldProps={{
            size: 'large',
            prefix: <FieldStringOutlined className={styles.prefixIcon} />,
          }}
          placeholder="Description"
          width="md"
          name="description"
        />

      </ModalForm>
      <UpdateForm
        onVisibleChange={handleModalVisible}
        onSubmit={async (value) => {
          const success = await handleUpdate(value);

          if (success) {
            handleUpdateModalVisible(false);
            setCurrentRow(undefined);

            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
        onCancel={() => {
          handleUpdateModalVisible(false);

          if (!showDetail) {
            setCurrentRow(undefined);
          }
        }}
        updateModalVisible={updateModalVisible}
        values={currentRow || {}}
      />

      <Drawer
        width={600}
        visible={showDetail}
        onClose={() => {
          setCurrentRow(undefined);
          setShowDetail(false);
        }}
        closable={false}
      >
        {currentRow?.name && (  
          <ProDescriptions
            column={2}
            title={currentRow?.name}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns2}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default ProductList;
