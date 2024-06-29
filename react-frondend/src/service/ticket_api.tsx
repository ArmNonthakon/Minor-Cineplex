import axios from "axios"
import { getCookie } from "typescript-cookie";
import parseJwt from '../utils/decode';


const url = '/api/'


export const ReserveSeat = async (MovieId: string, TheaterId: string, seatReserve: string[]) => {
    const jwt = getCookie("token")
    const decoded = parseJwt(jwt as string)
    console.log("MovieId: " + MovieId)
    console.log("TheaterId: " + TheaterId)
    console.log("UserName: " + decoded.name)
    console.log("Seat: "+ seatReserve)
    try {
        const response = await axios.post(`${url}reserveSeat`, {
            MovieId: MovieId,
            UserName: decoded.name,
            TheaterId: TheaterId,
            SeatReserve: seatReserve
        }, {
            withCredentials: true
        })
        return response.status;
    } catch (error) {
        return error
    }
}