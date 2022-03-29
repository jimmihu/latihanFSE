import { message } from 'antd';
import jwt_decode from 'jwt-decode';

const users = () => {
    const token = localStorage.getItem('token');
        if (token) {
            const payload = jwt_decode(token)
            return {
                    islogin: true,
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