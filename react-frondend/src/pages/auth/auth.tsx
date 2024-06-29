import { useEffect } from 'react'
import './auth.scss'
import Login from './Login/login'
import Register from './Register/register'
import { getCookie } from 'typescript-cookie'
import { useNavigate } from 'react-router-dom'
interface Auth {
    state: string
}
const Auth = ({ state }: Auth) => {
    const navigate = useNavigate()
    useEffect(() => {
        getCookie('token') && navigate('/')
    }, [])
    return (
        <>
            <div className='section-auth'>
                <div className='auth-picture'>
                    <img src="/logo.png" alt="" width="200px" />
                    <h1>Welcome to Minor Cineplex </h1>
                    <p>Please login account to continue.</p>

                </div>
                <div className="auth-login-regis">
                    <div id='auth-login' className='auth-login' style={{ display: state == 'login' ? 'inline' : 'none' }}>
                        <Login />

                    </div>
                    <div id='auth-regis' className='auth-regis' style={{ display: state == 'register' ? 'inline' : 'none' }}>
                        <Register />
                    </div>
                </div>
            </div>
        </>
    )
}
export default Auth