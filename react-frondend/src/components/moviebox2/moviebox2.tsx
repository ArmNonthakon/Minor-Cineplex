import './moviebox2.scss'
import { ToUpperFirstLetter } from '../../utils/utils_string';

interface input {
    title : string
    genre : string
    duration: number
    poster : string
}


export const MovieBox2 = ({title,genre,duration,poster}:input) => {
    return (
        <>
            <div className='section-movie_theater-movie'>
                <div className='section-movie_theater-movie-poster' 
                style={{backgroundImage: `url("/movie/${poster}")`,}}></div>
                <div className='section-movie_theater-movie-describe'>
                    <h1>{ToUpperFirstLetter(title)}</h1>
                    <div className='describe'>
                        <img src="/brand.png" width="23px" alt="" />
                        <h3>{genre}</h3>
                    </div>
                    <div className='describe'>
                        <img src="/clock.png" width="23px" alt="" />
                        <h3>{duration} minute</h3>
                    </div>
                    <div className='section-movie_theater-movie-round'>
                        <div>12.30</div>
                        <div>14.50</div>
                    </div>
                </div>

            </div>
        </>
    )
}