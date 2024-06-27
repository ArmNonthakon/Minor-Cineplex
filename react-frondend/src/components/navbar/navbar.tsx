import { Link } from 'react-router-dom'
import './navbar.scss'

export const Navbar = () => {
    
    return (
        <>
            <div className='section-navbar'>
                <div className='section-navbar-topic'>
                    <img src="/logo.png" width="60px" alt="" />
                    <Link to={"/"}>HOME</Link>
                    <Link to={"/movie"}>MOVIES</Link>
                    <Link to={"/seat"}>THEATER</Link>
                </div>
                <div className='section-navbar-auth'>
                    <img src="/search.png" width="28px" alt="" />
                    <img src="/tally-1.png" width="28px" alt="" />
                    <Link to={"/login"}><img src="/user.svg" width="40px" alt="" /></Link>

                </div>

            </div>

        </>
    )
}