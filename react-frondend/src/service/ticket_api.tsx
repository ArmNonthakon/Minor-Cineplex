import axios from "axios"
import { getCookie } from "typescript-cookie";
import parseJwt from '../utils/decode';


const url = '/api/'
export const ReserveSeat = async (MovieId: string, TheaterId: string, seatReserve: string[]) => {
    const jwt = getCookie("token") as string
    const decoded = parseJwt(jwt)
    try {
        const response = await axios.post(`${url}reserveSeat`, {
            MovieId: MovieId,
            UserName: decoded.name,
            TheaterId: TheaterId,
            SeatReserve: seatReserve
        }, {
            withCredentials: true
        })
        return response;
    } catch (error) {
        return error
    }
}
export const GetTicketByUser = async () => {
    const jwt = getCookie("token") as string
    const decoded = parseJwt(jwt)
    try {
        const response = await axios.post(`${url}getTicketByUserName`, {
            UserName: decoded.name,
        }, {
            withCredentials: true
        })
        return response;
    } catch (error) {
        return error
    }
}
export const GetTicketById = async (TicketId: string) => {
    try {
        const response = await axios.post(`${url}getTicketById`, {
            TicketId: TicketId,
        }, {
            withCredentials: true
        })
        return response;
    } catch (error) {
        return error
    }
}

