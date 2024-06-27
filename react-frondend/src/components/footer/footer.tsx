import './footer.scss'

const Footer = () => {
    return (
        <>
            <div className='section-footer'>
                <div className='section-footer-one'>
                    <div>
                        <h2>THEATER</h2>
                        <p>ALL THEATER</p>
                        <p>SOME THEATER</p>
                    </div>
                    <div>
                        <h2>MOVIES</h2>
                        <p>NOW SHOWING</p>
                        <p>COMMING SOON</p>
                    </div>
                    <div>
                        <h2>AUTH</h2>
                        <p>LOGIN</p>
                        <p>REGISTER</p>
                    </div>
                    <div>
                        <h2>CONTACT ME</h2>
                        <a href="https://www.facebook.com/people/Nonthakon-Tansamai/pfbid0zB9xYyGJNGkkPzwe2Wfsm6jaP5Yd3zLxwwU3C7hHbtNnHAqVhzWEfPJCjfHCKTSdl/">FACEBOOK</a>
                        <a href="https://www.instagram.com/3armm__9">INSTARGRAM</a>
                        <a href="https://www.linkedin.com/in/nonthakon-tansamai-09986430a">LINKEDIN</a>
                        <a href="https://github.com/ArmNonthakon">GITHUB</a>
                    </div>
                    

                </div>
                <div className='section-footer-copyright'>
                    <img src="/logo.png" alt="" />
                <p>
                    copyright ©2024; Designed by ArmNonthakon
                </p></div>
            </div>
        </>
    )
}

export default Footer