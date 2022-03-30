// @ts-ignore

/* eslint-disable */
import { request } from 'umi';
/** 获取当前的用户 GET /api/currentUser */

const BaseURL = 'http://localhost:8000';

export async function LoginUser(payload){
  return request(BaseURL +'/login',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  })
}

export async function CreateUser(payload){
  return request(BaseURL +'/users',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  })
}

export async function GetUserDetail(id) {
  return request(BaseURL+'/users/'+id, {
    method: 'GET',
    skipErrorHandler: true,
  });
}

export async function DeleteUser(id) {
  return request(BaseURL+'/users/'+id, {
    method: 'DELETE',
    skipErrorHandler: true,
  });
}

export async function GetUserList() {
  return request(BaseURL+'/users', {
    method: 'GET',
    skipErrorHandler: true,
  });
}

export async function CreateProduct(payload){
  return request(BaseURL +'/products',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  })
}

export async function GetProductList() {
  return request(BaseURL+'/products', {
    method: 'GET',
    skipErrorHandler: true,
  });
}

export async function GetProductDetail(id) {
  return request(BaseURL+'/products/'+id, {
    method: 'GET',
    skipErrorHandler: true,
  });
}

export async function DeleteProduct(id) {
  return request(BaseURL+'/products/'+id, {
    method: 'DELETE',
    skipErrorHandler: true,
  });
}
// export async function currentUser(options) {
//   return request('/api/currentUser', {
//     method: 'GET',
//     ...(options || {}),
//   });
// }
/** 退出登录接口 POST /api/login/outLogin */

// export async function outLogin(options) {
//   return request('/api/login/outLogin', {
//     method: 'POST',
//     ...(options || {}),
//   });
// }
// /** 登录接口 POST /api/login/account */

// export async function login(body, options) {
//   return request('/api/login/account', {
//     method: 'POST',
//     headers: {
//       'Content-Type': 'application/json',
//     },
//     data: body,
//     ...(options || {}),
//   });
// }
// /** 此处后端没有提供注释 GET /api/notices */

// export async function getNotices(options) {
//   return request('/api/notices', {
//     method: 'GET',
//     ...(options || {}),
//   });
// }
// /** 获取规则列表 GET /api/rule */

// export async function rule(params, options) {
//   return request('/api/rule', {
//     method: 'GET',
//     params: { ...params },
//     ...(options || {}),
//   });
// }
// /** 新建规则 PUT /api/rule */

// export async function updateRule(options) {
//   return request('/api/rule', {
//     method: 'PUT',
//     ...(options || {}),
//   });
// }
// /** 新建规则 POST /api/rule */

// export async function addRule(options) {
//   return request('/api/rule', {
//     method: 'POST',
//     ...(options || {}),
//   });
// }
// /** 删除规则 DELETE /api/rule */

// export async function removeRule(options) {
//   return request('/api/rule', {
//     method: 'DELETE',
//     ...(options || {}),
//   });
// }
