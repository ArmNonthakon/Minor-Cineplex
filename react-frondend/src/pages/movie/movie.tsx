import './movie.scss';
import { MovieBox } from '../../components/moviebox/moviebox';
import { GetMovie } from '../../service/movie_api';
import { useEffect, useState } from 'react';

interface MovieData {
    MovieTitle: string;
    Poster: string;
    TimeDuration: number;
    MovieGenre: string;
    ReleaseDate: string
}

export const Movie = () => {
    const [data, setData] = useState<MovieData[]>([]);

    const callApiMovie = async () => {
        try {
            const response = await GetMovie();
            setData(Array.isArray(response) ? response : []);
        } catch (error) {
            console.error('Error fetching movies:', error);
            setData([]);
        }
    };

    useEffect(() => {
        callApiMovie();
    }, []); // Empty array ensures the effect runs only once

    return (
        <>
        <div className='section-movie-topic'>
                <p>MOVIES</p>
            </div>
        <div className="section-movie">
            {data.map((item, index) => (
                <MovieBox
                    key={index}
                    title={item.MovieTitle}
                    poster={item.Poster}
                    duration={item.TimeDuration}
                    genre={item.MovieGenre}
                    date={item.ReleaseDate}
                />
            ))}
        </div>
        </>

    );
};
