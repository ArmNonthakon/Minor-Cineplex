import { Seats } from '../../../components/seats/seats';
import './movie_theater.scss';
import { Summary } from '../../../components/summary/summary';
import { MovieBox2 } from '../../../components/moviebox2/moviebox2';
import { GetMovieByTitle } from '../../../service/movie_api';
import { useCallback, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

interface MovieData {
    MovieId: string
    MovieTitle: string;
    Poster: string;
    TimeDuration: number;
    MovieGenre: string;
    Theaters: []
}

export const Movie_theater = () => {
    const [data, setData] = useState<MovieData | null>(null);
    const [currentTheater, setCurrentTheater] = useState("");
    const [currentTime, setCurrentTime] = useState("");
    const [reserveSeat, setReserveSeat] = useState<string[]>([]);
    const param = useParams();

    const callApiMovie = async () => {
        let title: string = param.id as string;
        try {
            const response = await GetMovieByTitle(title.replace(/_/g, " "));
            setData(response);
        } catch (error) {
            console.error('Error fetching movie:', error);
        }
    };

    const setTheaterFromMoviebox = useCallback((data: any, id: any) => {
        setCurrentTheater(id);
        setCurrentTime(data);
        setReserveSeat([])
    },[]);

    const setSeatsReserve = useCallback((seats: string[]) => {
        setReserveSeat(seats);
    }, []);

    const CurrentSeats = useCallback(() => {
        let seatCol = 0;
        let seatRow = 0;
        let seatReserve:any[] = []
        data?.Theaters.forEach((e: any) => {
            if (e.TheaterId === currentTheater) {
                seatCol = e.SeatCol;
                seatRow = e.SeatRow;
                seatReserve = e.SeatAlreadyReserve == null ? [] : e.SeatAlreadyReserve
            }
        });
        return <Seats row={seatRow} col={seatCol} seats={seatReserve}  sendSeat={setSeatsReserve} />;
    }, [data, currentTheater, setSeatsReserve]);

    useEffect(() => {
        callApiMovie();
        window.scrollTo(0, 0);
    }, [param.id]);

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
            <CurrentSeats />
            {data && (
                <Summary
                    MovieId={data.MovieId}
                    TheaterId={currentTheater}
                    title={data.MovieTitle}
                    poster={data.Poster}
                    duration={data.TimeDuration}
                    time={currentTime}
                    seatReserve={reserveSeat}
                />
            )}
        </div>
    );
};
