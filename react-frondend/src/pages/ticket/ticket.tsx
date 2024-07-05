import { Ticket_box } from "../../components/ticket_box/ticket_box"
import './ticket.scss'

export const Ticket = () => {
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
                <Ticket_box />
                <Ticket_box />
                <Ticket_box />
                <Ticket_box />
            </div>
        </>
    )
}