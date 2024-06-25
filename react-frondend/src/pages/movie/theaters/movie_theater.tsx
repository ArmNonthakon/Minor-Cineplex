import { Seats } from '../../../components/seats/seats';
import './movie_theater.scss';
import { Summary } from '../../../components/summary/summary';
import { MovieBox2 } from '../../../components/moviebox2/moviebox2';
import { GetMovieByTitle } from '../../../service/movie_api';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

interface MovieData {
    MovieTitle: string;
    Poster: string;
    TimeDuration: number;
    MovieGenre: string;
    Theaters: []
    
}



export const Movie_theater = () => {
    const [data, setData] = useState<MovieData | null>(null);
    const [currentTheater,setCurrentTheater] = useState("")
    const [currentTime,setCurrentTime] = useState("")
    const param = useParams()
    const callApiMovie = async () => {
        let title: string = param.id as string
        try {
            const response = await GetMovieByTitle(title.replace(/_/g, " "));
            setData(response);
        } catch (error) {
            console.error('Error fetching movie:', error);
        }
    };
    const setTheaterFromMoviebox = (data:any,id:any)=>{
        setCurrentTheater(id)
        setCurrentTime(data)
    }
    const CurrentSeats = ()=>{
        let seatCol = 0;
        let seatRow = 0;
        data?.Theaters.map((e:any)=>{
            if(e.TheaterId == currentTheater){
                seatCol = e.SeatCol
                seatRow = e.SeatRow
            }
        })
        return <Seats row={seatCol} col={seatRow} />
    }
    useEffect(() => {
        callApiMovie();
        console.log(currentTheater)
    }, []); // Add title as a dependency to refetch if it changes

    return (
        <div className='section-movie_theater'>
            {data && (
                <MovieBox2
                    title={data.MovieTitle}
                    poster={data.Poster}
                    duration={data.TimeDuration}
                    genre={data.MovieGenre}
                    theater={data.Theaters}
                    ondata={setTheaterFromMoviebox}
                />
            )}
            <CurrentSeats/>
            {data && <Summary
                title={data.MovieTitle}
                poster={data.Poster}
                duration={data.TimeDuration}
                time={currentTime}
                
            />}
            
        </div>
    );
};
