import { Ticket_box } from "../../components/ticket_box/ticket_box"
import './ticket.scss'

export const Ticket = ()=>{
    return (
        <>
            <h1>Ticket
            </h1>
            <div className="section-ticket">
                <Ticket_box/>
            </div>
        </>
    )
}