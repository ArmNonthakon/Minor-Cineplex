import './home.scss';
import { MovieBox } from '../../components/moviebox/moviebox';
import { GetMovie } from '../../service/movie_api';
import { useEffect, useState } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';

// Import Swiper styles
import 'swiper/css';
import 'swiper/css/pagination';
import { RemoveDuplicate } from '../../utils/filter';

interface MovieData {
    MovieTitle: string;
    Poster: string;
    TimeDuration: number;
    MovieGenre: string;
    ReleaseDate: string;
}

export const Home = () => {
    const [data, setData] = useState<MovieData[]>([]);
    const [genre, setGenre] = useState<string[]>([]);
    const [innerWidth,setInnerWidth] = useState(window.innerWidth)
    const callApiMovie = async () => {
        try {
            const response = await GetMovie();
            const movies = Array.isArray(response) ? response : [];
            setData(movies);

            const getGenre = movies.map((item) => item.MovieGenre);
            getGenre.sort()
            setGenre(RemoveDuplicate(getGenre));
        } catch (error) {
            console.error('Error fetching movies:', error);
            setData([]);
        }
    };
    window.addEventListener('resize', ()=>{
        setInnerWidth(window.innerWidth)
    });
   
    const getMoviesByGenre = (inputGenre: string) => {
        return data.filter((item) => item.MovieGenre === inputGenre);
    }

    useEffect(() => {
        callApiMovie();
        console.log(innerWidth)
    }, []);

    return (
        <>
            {genre.map((item, index) => (
                <div key={index}>
                    <p className='toppic-genre'>{item}</p>
                    <div className='section-home'>
                        <Swiper
                            slidesPerView={innerWidth <= 650 ? 2.1 : innerWidth <= 1000 ? 3.3 : innerWidth <= 1151 ? 4.4 : 5.4 }
                            spaceBetween={35}
                            className="mySwiper"
                        >
                            <div className='contain'>
                                {getMoviesByGenre(item).map((movie, idx) => (
                                    <SwiperSlide key={idx}>
                                        <MovieBox
                                            title={movie.MovieTitle}
                                            poster={movie.Poster}
                                            duration={movie.TimeDuration}
                                            genre={movie.MovieGenre}
                                            date={movie.ReleaseDate}
                                        />
                                    </SwiperSlide>
                                ))}
                            </div>
                        </Swiper>
                    </div>
                </div>
            ))}
        </>
    );
}
