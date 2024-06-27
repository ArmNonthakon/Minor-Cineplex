import './login.scss';
import { Link, useNavigate } from 'react-router-dom';
import { LoginApi } from '../../../service/user_api';
import { useState } from 'react';
import { setCookie } from 'typescript-cookie';

const Login = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [response, setResponse] = useState('');
    const [loading, setLoading] = useState('');
    const [isSubmit, setIssubmit] = useState(false)

    const callLoginApi = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setResponse('');
        setLoading('Logging in...')
        setIssubmit(true)
        if (username != "" && password != "") {
            try {
                const response = await LoginApi({ username, password });
                if (response === 202) {
                    setUsername('');
                    setPassword('');
                    setResponse('Login success!!')
                    setIssubmit(true)
                    setCookie('userName',username)
                    setTimeout(() => {
                        navigate('/');
                    }, 3000);
                } else {
                    setResponse('Login failed. Please check your username and password.');
                }
            } catch (error) {
                setResponse('Login failed. Please check your username and password.');
            } finally {
                setLoading('')
                setIssubmit(false)
            }
        } else {
            setResponse('Please fill out the information completely.');
            setLoading('')
        }

    };

    return (
        <div className='login'>
            <h1>Login</h1>
            <form onSubmit={callLoginApi}>
                <div>
                    <img src="/user.svg" alt="User Icon" />
                    <input
                        type="text"
                        placeholder='Username'
                        id="Username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        style={{ borderBottom: isSubmit == true && username == "" ? "3px solid rgb(245, 121, 121)" : "3px solid black" }}
                    />
                </div>
                <div>
                    <img src="/lock.svg" alt="Lock Icon" />
                    <input
                        type="password"
                        placeholder='Password'
                        id='password'
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        style={{ borderBottom: isSubmit == true && password == "" ? "3px solid rgb(245, 121, 121)" : "3px solid black" }}
                    />
                </div>
                <div className='error' style={{ color: response == "Login success!!" ? "green" : "red" }}>{response && response}</div>
                <div className='login-button'>
                    <button type='submit' style={{ backgroundColor: loading && "rgb(147, 196, 239)" }}>
                        {loading ? loading : "Login"}
                    </button>
                    <p>Don't have an account? <Link to="/signup">Sign up</Link></p>
                </div>
            </form>
        </div>
    );
};

export default Login;
