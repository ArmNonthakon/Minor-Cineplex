import './ticket_box.scss'

export const Ticket_box = () => {
    return (
        <>
            <div className="ticket-box">
                <p className='ticket-box-topic'>Ticket ID</p>
                <p className='ticket-box-information'>017d14c8-0d49-472f-93a6-23814be9a9c9</p>

                <div className="ticket-information">
                    <div className='ticket-information-logo'></div>
                    <div className='ticket-information-describe'>
                        <div>
                            <p className='ticket-box-topic'>Movie Title</p>
                            <p className='ticket-box-information'>Bad boy for life</p>
                        </div>
                        <div>
                            <p className='ticket-box-topic'>Theater</p>
                            <p className='ticket-box-information'>T3</p></div>
                        <div>
                            <p className='ticket-box-topic'>Seats</p>
                            <p className='ticket-box-information'>C1 C2 C3 </p></div>
                        <div>
                            <p className='ticket-box-topic'>Show Time</p>
                            <p className='ticket-box-information'>12.30</p></div>
                    </div>
                </div>
            </div>
        </>
    )
}