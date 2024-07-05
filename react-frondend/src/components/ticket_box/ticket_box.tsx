
import './ticket_box.scss'

interface InputTicket {
    ticketId: string
    movieTitle: string
    theaterNumber: string
    seats: string[]
    time: string
}
export const Ticket_box = ({ ticketId, movieTitle, theaterNumber, seats, time }: InputTicket) => {
    const TransformTime = ()=>{
        const date = new Date(time)
        if (date.getMinutes() < 10) {
            return <p className='ticket-box-information'>{date.getHours() + ":0" + date.getMinutes()}</p>
        } else {
            return <p className='ticket-box-information'>{date.getHours() + ":" + date.getMinutes()}</p>
        }   
    }
    return (
        <>
            <div className="ticket-box">
                <p className='ticket-box-topic'>Ticket ID</p>
                <p className='ticket-box-information'>{ticketId}</p>

                <div className="ticket-information">
                    <div className='ticket-information-logo'></div>
                    <div className='ticket-information-describe'>
                        <div>
                            <p className='ticket-box-topic'>Movie Title</p>
                            <p className='ticket-box-information'>{movieTitle}</p>
                        </div>
                        <div>
                            <p className='ticket-box-topic'>Theater</p>
                            <p className='ticket-box-information'>{theaterNumber}</p></div>
                        <div>
                            <p className='ticket-box-topic'>Seats</p>
                            <p className='ticket-box-information'>
                                {seats.map((e)=>(e + " "))}
                            </p></div>
                        <div>
                            <p className='ticket-box-topic'>Show Time</p>
                            <TransformTime/></div>
                    </div>
                </div>
            </div>
        </>
    )
}