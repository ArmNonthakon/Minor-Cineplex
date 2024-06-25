import './moviebox2.scss'
import { ToUpperFirstLetter } from '../../utils/utils_string';

interface input {
    title : string
    genre : string
    duration: number
    poster : string
    theater: []
}


export const MovieBox2 = ({title,genre,duration,poster,theater}:input) => {
    const TheaterTime = ()=>{
        const theaterTime:any = []  
        theater.map((e:any)=>{
            const date = new Date(e.TimeStart)
            if (date.getMinutes() < 10){
                theaterTime.push(date.getHours() + ":0" + date.getMinutes())
            }else{
                theaterTime.push(date.getHours() + ":" + date.getMinutes())
            }     
        })
        theaterTime.sort()
        return theaterTime.map((e:any)=>(
            <div>{e}</div>
        ))
        
    }
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
                        <TheaterTime/>
                    </div>
                </div>

            </div>
        </>
    )
}