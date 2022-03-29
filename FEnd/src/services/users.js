import { message } from 'antd';
import jwt_decode from 'jwt-decode';
import { GetUserDetail } from './ant-design-pro/api';

const users = async () => {
    const token = localStorage.getItem('token');
    
    if (token) {
        const payload = jwt_decode(token)
        return {
                isLogin: true,
                name: payload.name,
                payload: payload            
            };
    }

    return {
        islogin: false,
        name: null,
        payload: null,
    };
};

export default users