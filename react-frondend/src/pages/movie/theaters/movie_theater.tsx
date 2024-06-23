import { Seats } from '../../../components/seats/seats';
import './movie_theater.scss';
import { Summary } from '../../../components/summary/summary';
import { MovieBox2 } from '../../../components/moviebox2/moviebox2';
import { GetMovieByTitle } from '../../../service/movie_api';
import { useEffect, useState } from 'react';

interface MovieData {
    MovieTitle: string;
    Poster: string;
    TimeDuration: number;
    MovieGenre: string;
}

interface MovieTheaterProps {
    title: string;
}

export const Movie_theater = ({ title }: MovieTheaterProps) => {
    const [data, setData] = useState<MovieData | null>(null);

    const callApiMovie = async () => {
        try {
            const response = await GetMovieByTitle(title);
            setData(response);
        } catch (error) {
            console.error('Error fetching movie:', error);
        }
    };

    useEffect(() => {
        callApiMovie();
    }, [title]); // Add title as a dependency to refetch if it changes

    return (
        <div className='section-movie_theater'>
            {data && (
                <MovieBox2 
                    title={data.MovieTitle} 
                    poster={data.Poster} 
                    duration={data.TimeDuration} 
                    genre={data.MovieGenre} 
                />
            )}
            <Seats row={5} col={5}/>
            <Summary />
        </div>
    );
};
