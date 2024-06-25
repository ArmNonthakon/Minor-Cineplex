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
    const param = useParams()
    const callApiMovie = async () => {
        let title:string = param.id as string
        try {
            const response = await GetMovieByTitle(title.replace(/_/g," "));
            setData(response);
        } catch (error) {
            console.error('Error fetching movie:', error);
        }
    };

    useEffect(() => {
        callApiMovie();
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
                />
            )}
            <Seats row={5} col={5}/>
            <Summary />
        </div>
    );
};
