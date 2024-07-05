
import { useEffect, useState } from 'react'
import { Ticket_box } from '../../../components/ticket_box/ticket_box'
import './ticket_id.scss'
import { GetTicketById } from '../../../service/ticket_api'
import { useParams } from 'react-router-dom'

export const Ticket_id = () => {
    const param = useParams()
    const [ticket,setTicket] = useState<any>()
    const callTicketApiByUserName = async ()=>{
        try {
            const response:any = await GetTicketById(param.id as string);
            if (response.status == 202){
                setTicket(response.data)
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
                </div>
            </div>
            <div className="section-ticket-id">
            <Ticket_box  ticketId={ticket.TicketId} movieTitle={ticket.MovieTitle} theaterNumber={ticket.TheaterNumber} seats={ticket.Seats} time={ticket.ShowTime}/>
            </div>
        </>
    )
}