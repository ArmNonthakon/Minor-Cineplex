import { useEffect, useState } from "react"
import { Ticket_box } from '../../components/ticket_box/ticket_box';
import { GetTicketByUser } from "../../service/ticket_api"
import './ticket.scss'

export const Ticket = () => {
    const [tickets,setTickets] = useState([])
    const callTicketApiByUserName = async ()=>{
        try {
            const response:any = await GetTicketByUser();
            if (response.status == 202){
                setTickets(response.data)
            }
        } catch (error) {
            throw error
        }
    }
    useEffect(()=>{
        callTicketApiByUserName()
    })
    return (
        <>
            <div className="section-ticket-topic">
                <div></div>
                <div>
                    <h1>Your Ticket</h1>
                </div>
                
                <div>
                    <form action="">
                        <input type="text" placeholder="TICKET ID" name="" id="" />
                        <button><p>SEARCH</p></button>
                    </form>
                </div>
            </div>
            <div className="section-ticket">
                {tickets.map((e:any,i)=>(
                    <Ticket_box key={i} ticketId={e.TicketId} movieTitle={e.MovieTitle} theaterNumber={e.TheaterNumber} seats={e.Seats} time={e.ShowTime}/>
                ))}
            </div>
        </>
    )
}