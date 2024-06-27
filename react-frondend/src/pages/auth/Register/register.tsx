import { useState } from 'react';
import './register.scss';
import { Link, useNavigate } from 'react-router-dom';
import { regisApi } from '../../../service/user_api';

const Register = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [response, setResponse] = useState('');
    const [loading, setLoading] = useState('');
    const [isSubmit, setIssubmit] = useState(false)

    const callRegisterApi = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setIssubmit(true)
        setResponse('');
        setLoading('Signing up...')
        if (username != "" && email != "" && password != "" && confirmPassword != "") {
            if (password === confirmPassword) {
                try {
                    const responseStatus = await regisApi({ username, email, password });
                    if (responseStatus === 201) {
                        setUsername('');
                        setEmail('');
                        setPassword('');
                        setConfirmPassword('');
                        setResponse('Sign up Success...');
                        setIssubmit(false)
                        setTimeout(() => {
                            setResponse('');
                            navigate('/login');
                        }, 3000);

                    } else {
                        setResponse('Registration failed. Please try again.');
                    }
                } catch (error) {
                    setResponse('Registration failed. Please try again.');
                } finally {
                    setLoading('')
                    setIssubmit(false)
                }
            } else {
                setResponse('Passwords do not match.');
                setLoading('')
                setIssubmit(false)
            }
        }else{
            setResponse('Please fill out the information completely.');
            setLoading('')
        }
    };

    return (
        <div className='register'>
            <h1>Sign up</h1>
            <form onSubmit={callRegisterApi}>
                <div>
                    <img src="/user.svg" alt="User icon" />
                    <input
                        type="text"
                        placeholder='Username'
                        id="username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        style={{ borderBottom: isSubmit == true && username == "" ? "3px solid rgb(245, 121, 121)" : "3px solid black" }}
                    />
                </div>
                <div>
                    <img src="/envelope.svg" alt="Email icon" />
                    <input
                        type="email"
                        placeholder='Email'
                        id='email'
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        style={{ borderBottom: isSubmit == true && email == "" ? "3px solid rgb(245, 121, 121)" : "3px solid black" }}
                    />
                </div>
                <div>
                    <img src="/lock.svg" alt="Lock icon" />
                    <input
                        type="password"
                        placeholder='Password'
                        id='Password'
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        style={{ borderBottom: isSubmit == true && password == "" ? "3px solid rgb(245, 121, 121)" : "3px solid black" }}
                    />
                </div>
                <div>
                    <img src="/lock.svg" alt="Lock icon" />
                    <input
                        type="password"
                        placeholder='Confirm Password'
                        id='confirmPassword'
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                        style={{ borderBottom: isSubmit == true && confirmPassword == "" ? "3px solid rgb(245, 121, 121)" : "3px solid black" }}
                    />
                </div>
                <div className='error' style={{ color: response == "Sign up Success..." ? "green" : "red" }}>{response && response}</div>
                <div className='register-button'>
                    <button type='submit' style={{ backgroundColor: loading && "rgb(247, 155, 155)" }}>
                        {loading ? loading : "Sign up"}
                    </button>
                    <p>Already have an account? <Link to="/login">Login</Link></p>
                </div>
            </form>
        </div>
    );
};

export default Register;
