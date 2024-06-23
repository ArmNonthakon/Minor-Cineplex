import './summary.scss'
import { Seats } from '../seats/seats';

export const Summary = () => {
    return (
        <>

            <div className="section-summary">
                <h2 style={{ color: "white", marginBottom: "20px" }}>Summary</h2>
                <div className='section-summary-box'>
                    <div className='section-summary-movie'>
                        <div className='section-summary-poster'>poster</div>
                        <div className='section-summary-describe'>
                            <h3>How to train your dragon</h3>
                            <div className='summary-describe'>
                                <img src="/clock.png" width="20px" alt="" />
                                <p>135 minute</p>
                            </div>

                        </div>
                    </div>
                    <div className='section-summary-round'>
                        <p></p>
                        <div>12.30</div>
                    </div>
                    <div className="section-summary-seats">
                        <h3>
                            Seats
                        </h3>
                        <h3>
                            Total Price
                        </h3>
                    </div>
                    <div className="section-summary-seats">
                        <h4>
                            C1,C2,C3
                        </h4>
                        <h4>
                            600
                        </h4>
                    </div>
                    <div className='section-summary-button'>
                         <button>BUY NOW</button>
                    </div>
                   
                </div>

            </div>
        </>
    )
}