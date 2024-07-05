
import { Ticket_box } from '../../../components/ticket_box/ticket_box'
import './ticket_id.scss'

export const Ticket_id = () => {
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
                <Ticket_box/>
                <Ticket_box/>
            </div>
        </>
    )
}