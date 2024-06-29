import axios from "axios"

const url = '/api/'
export const GetMovie = async () => {
    try {
        const response = await axios.get(`${url}getMovieWithTheater`)
        return response.data;
    } catch (error) {
        return error
    }
};

export const GetMovieByTitle = async (title:string) => {
    try {
        const response = await axios.post(`${url}getMovieByTitle`,{
            title: title
        })
        return response.data;
    } catch (error) {
        return error
    }
};




/*
const addMovie = async ()=>{
    try {
        const response = await axios.get("/addMovie")
        return response
    } catch (error) {
        return error
    }
    
}
*/