import axios from "axios"
import { getCookie } from "typescript-cookie";
import parseJwt from '../utils/decode';


const url = 'http://127.0.0.1:3000/'
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

