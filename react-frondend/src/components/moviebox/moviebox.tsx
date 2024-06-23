

import { useState } from 'react'
import './moviebox.scss'
interface Input {
    title: string
    duration: number
    poster: string
    genre: string
    date: string
}
export const MovieBox = ({ title, duration, poster, genre, date }: Input) => {
    const [isHover, setIsHover] = useState(false)
    const TransformDate = () => {
        const tDate = new Date(date)
        const month = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"]  
        return <h4>{tDate.getDay()} {month[tDate.getMonth()]} {tDate.getFullYear()}</h4>
    }
    return (<>
        <div className='section-moviebox' onMouseOver={() => setIsHover(true)} onMouseLeave={() => setIsHover(false)}>
            <div className='movie-poster'
                style={{
                    backgroundImage: `url("/movie/${poster}")`,
                    opacity: isHover == true ? "0.5" : "1"
                }} >

            </div>
            <div className='movie-describe'>
                <TransformDate />
                <h3>{title.toUpperCase()}</h3>
                <div className='movie-describe-round'>
                    <div>{duration} minute</div>
                    <div>{genre}</div>
                </div>
            </div>

        </div>
    </>)
}