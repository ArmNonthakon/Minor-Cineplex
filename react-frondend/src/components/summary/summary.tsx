import { useEffect, useState } from 'react'
import { ReserveSeat } from '../../service/ticket_api'
import './summary.scss'
import { ToUpperFirstLetter } from '../../utils/utils_string';
import { useNavigate } from 'react-router-dom';

interface inputSummary {
    MovieId: string
    TheaterId: string
    title: string
    duration: number
    time: string
    poster: string
    seatReserve: string[]
}

export const Summary = ({ MovieId, TheaterId, title, poster, duration, time, seatReserve }: inputSummary) => {
    const [response, setResponse] = useState('');
    const [loading, setLoading] = useState('');
    const navigate = useNavigate()
    const callApiReserve = async () => {
        setLoading('Booking tickets...')
        try {
            const response:any = await ReserveSeat(MovieId, TheaterId, seatReserve);
            if (response.status == 202) {
                console.log(response.data)
                setResponse('Successfully booked tickets!!!')
                setTimeout(() => {
                    navigate('/ticket/'+ response.data.TicketId)
                }, 2000);
            }
            else {
                setResponse('Unsuccessfully booked tickets!!!')
            }
        } catch (error) {
            setResponse('Unsuccessfully booked tickets!!!')
            throw error
        } finally {
            setLoading('')
        }
    }
    useEffect(() => {

    }, [])
    return (
        <>

            <div className="section-summary">
                <h2 style={{ color: "white", marginBottom: "10px" }}>Summary</h2>
                <div className='section-summary-box'>
                    <div className='section-summary-movie'>
                        <div className='section-summary-poster' style={{ backgroundImage: `url("/movie/${poster}")`, }}></div>
                        <div className='section-summary-describe'>
                            <h3>{ToUpperFirstLetter(title)}</h3>
                            <div className='summary-describe'>
                                <img src="/clock.png" width="23px" alt="" />
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
                            {seatReserve.map((e) => (
                                e + " "
                            ))}
                        </h4>
                        <h4>
                            {seatReserve.length * 220}
                        </h4>
                    </div>
                    <div className='section-summary-button'>
                        <button style={{backgroundColor: loading && "rgb(241, 217, 157)"}} onClick={callApiReserve}>
                            {loading ? loading : "BUY NOW"}
                    </button>
                    </div>

                </div>
                <div className='section-summary-response'>
                    <p style={{color: response == "Unsuccessfully booked tickets!!!" ? "red" : "rgb(252, 200, 68)"}}>{response && response}</p>
                </div>

            </div>
        </>
    )
}