import axios from "axios";
interface InputLogin {
    username: string
    password: string
}
interface InputRegister{
    username :string
    email: string
    password:string
}

const url = 'http://127.0.0.1:3000/'
export const LoginApi = async ({ username, password }: InputLogin) => {
    try {
        const response = await axios.post(`${url}login`, {
            username: username,  
            password: password   
        },{
            withCredentials:true
        })
        return response.status;
    } catch (error) {
        console.error("Login error:", error);
        throw error;
    }
};
export const regisApi = async ({username,email,password}: InputRegister) => {
    try {
        const response = await axios.post(`${url}register`, {
            username: username,
            email: email,  
            password: password,
        });
        return response.status
    } catch (error) {
        return error
    }
}