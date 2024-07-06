import { useEffect, useState } from 'react';
import { Ticket_box } from '../../../components/ticket_box/ticket_box';
import './ticket_id.scss';
import { GetTicketById } from '../../../service/ticket_api';
import { useParams } from 'react-router-dom';

interface Ticket {
    TicketId: string;
    MovieTitle: string;
    TheaterNumber: string;
    Seats: string[];
    ShowTime: string;
}

export const Ticket_id = () => {
    const { id } = useParams<{ id: string }>();
    const [ticket, setTicket] = useState<Ticket | null>(null);

    const callTicketApiByUserId = async () => {
        try {
            const response: any = await GetTicketById(id as string);
            if (response.status === 202) {
                setTicket(response.data);
            }
            else{
                console.log(response.data)
            }
        } catch (error) {
            console.error("Error fetching ticket:", error);
        }
    };

    useEffect(() => {
        if (id) {
            callTicketApiByUserId();
        }
    }, [id]);

    return (
        <>
            <div className="section-ticket-topic">
                <div></div>
                <div>
                    <h1>Your Ticket</h1>
                </div>
                <div></div>
            </div>
            <div className="section-ticket-id">
                {ticket && (
                    <Ticket_box
                        ticketId={ticket.TicketId}
                        movieTitle={ticket.MovieTitle}
                        theaterNumber={ticket.TheaterNumber}
                        seats={ticket.Seats}
                        time={ticket.ShowTime}
                    />
                )}
            </div>
        </>
    );
};
