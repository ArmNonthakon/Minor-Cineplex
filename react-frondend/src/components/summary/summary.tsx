import { useEffect } from 'react'
import { ReserveSeat } from '../../service/ticket_api'
import './summary.scss'

interface inputSummary{
    MovieId: string
    TheaterId: string
    title : string
    duration: number
    time : string
    poster:string
    seatReserve : string[]
}

export const Summary = ({MovieId,TheaterId,title,poster,duration,time,seatReserve}:inputSummary) => {
    const callApiReserve = async ()=>{
        try {
            const response = await ReserveSeat(MovieId,TheaterId,seatReserve);
            if (response == 202){
                console.log("Success")
                location.reload()
            }
            else{
                console.log("Err")
            }
        } catch (error) {
            throw error
        }
    }
    useEffect(()=>{
        
    },[])
    return (
        <>

            <div className="section-summary">
                <h2 style={{ color: "white", marginBottom: "20px" }}>Summary</h2>
                <div className='section-summary-box'>
                    <div className='section-summary-movie'>
                        <div className='section-summary-poster' style={{backgroundImage: `url("/movie/${poster}")`,}}></div>
                        <div className='section-summary-describe'>
                            <h3>{title}</h3>
                            <div className='summary-describe'>
                                <img src="/clock.png" width="20px" alt="" />
                                <p>{duration} minute</p>
                            </div>

                        </div>
                    </div>
                    <div className='section-summary-round'>
                        <p></p>
                        <div>{time ? time : "00:00"}</div>
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
                            {seatReserve.map((e)=>(
                                e + " "
                            ))} 
                        </h4>
                        <h4>
                            {seatReserve.length * 220}
                        </h4>
                    </div>
                    <div className='section-summary-button'>
                         <button onClick={callApiReserve}>BUY NOW</button>
                    </div>
                   
                </div>

            </div>
        </>
    )
}